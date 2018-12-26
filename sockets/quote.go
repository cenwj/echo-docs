package sockets
import (
	"bytes"

	"fmt"
	"github.com/google/uuid"
	"net"
	"time"
	"io"
	"unsafe"
	"reflect"
	"github.com/syyongx/php2go"
	"log"
	"echo-docs/conf"
)

// ByteString converts []byte to string without memory allocation by block magic
func ByteString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

// StringByte converts string to string without memory allocation by block magic
func StringByte(s string) []byte {
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := reflect.SliceHeader{
		Data: sh.Data,
		Len:  sh.Len,
		Cap:  sh.Len,
	}
	return *(*[]byte)(unsafe.Pointer(&bh))
}

// 获取登录信息
func GetLoginData() string {
	loginTime := fmt.Sprintf("%v", time.Now().Format("2006-01-02 15:04:05"))
	DevId := GetUid()
	data := make(map[string]string)
	data["UserName"] = conf.Config().Quote.QUName
	data["LoginTime"] = loginTime
	data["DevId"] = DevId
	data["Token"] = php2go.Md5(conf.Config().Quote.QPsswd + loginTime + conf.Config().Quote.QUName + DevId)
	BytesData, _ := php2go.JsonEncode(data)
	bodyExcept := len(BytesData) / 256
	bodyFmod := len(BytesData) % 256
	byteslist := []byte{1, 2, 3, byte(bodyExcept), byte(bodyFmod)}
	var buffer bytes.Buffer
	buffer.Write(byteslist)
	buffer.Write(BytesData)
	loginStr := ByteString(buffer.Bytes())
	fmt.Println(loginStr)
	return loginStr
}
func GetResult(symbols []string) string {
	data := make(map[string]interface{})
	data["Params"] = symbols
	data["RequestNo"] = "0"
	data["ServiceCode"] = "00001"
	BytesData, _ := php2go.JsonEncode(data)
	bodyExcept := len(BytesData) / 256
	bodyFmod := len(BytesData) % 256
	byteslist := []byte{1, 3, 3, byte(bodyExcept), byte(bodyFmod)}
	var buffer bytes.Buffer
	buffer.Write(byteslist)
	buffer.Write(BytesData)
	QuoteStr := ByteString(buffer.Bytes())
	return QuoteStr
}
func GetUid() string {
	UUid, _ := uuid.NewRandom()
	return fmt.Sprintf("%s", UUid)
}

// 定时器
func HeartTimer(conn net.Conn) {
	ticker := time.NewTicker(5 * time.Second)
	for {
		select {
		case <-ticker.C:
			byteslist := []byte{1, 1, 3, 0, 0}
			conn.Write(byteslist)
			log.Printf("heart at %v\r\n", time.Now().Format("2006-01-02 15:04:05"))
		}
	}
}

//检测是否登录成功
func CheckSuclogin(login []byte) bool {
	//2 2 3 0 66
	loginHeader := []byte{2, 2, 3, 0, 66}
	res := bytes.Compare(login, loginHeader)
	if res == 0 {
		return true
	} else {
		return false
	}
}

// 检测心跳
func CheckIsHeart(heart []byte) bool {
	heart1 := []byte{2, 1, 3, 0, 0}
	heart2 := []byte{1, 1, 3, 0, 0}
	heartRes1 := bytes.Compare(heart, heart1)
	heartRes2 := bytes.Compare(heart, heart2)
	if heartRes1 == 0 {
		return true
	}
	if heartRes2 == 0 {
		return true
	}
	return false
}

func RetDataHandle(Buffers *bytes.Buffer, readerChannel chan string) {
	//读取头信息
	DataHeaders := make([]byte, 5)
	Buffers.Read(DataHeaders)
	// header：1-3-3-bodysize/256-bodySize%256 body
	// 被除数 = 除数 x 商 + 余数
	bodySize := (int(DataHeaders[3]))*int(256) + int(DataHeaders[4])
	//log.Println("bodySize==>", bodySize)
	//log.Println("Buffers==>", Buffers.Bytes())
	// 如果超过Buffer的长度
	if bodySize > len(Buffers.Bytes()) {

	} else {
		BodyDatas := make([]byte, bodySize)
		Buffers.Read(BodyDatas)
		readerChannel <- ByteString(BodyDatas)
	}
}

func reader(readerChannel chan string) {
	for {
		select {
		case data := <-readerChannel:
			log.Println(data)
		}
	}
}

func main() {
	fmt.Println(GetUid())
	fmt.Printf("Client Start at %v \r\n", time.Now().Format("2006-01-02 15:04:05"))
	conn, err := net.Dial("tcp", "quote.tigerwit.com:7777")
	if err != nil {
		log.Fatal("Err => %v", err)
	}
	conn.Write(StringByte(GetLoginData()))
	go HeartTimer(conn)

	buf := make([]byte, 1024)
	Buffers := bytes.NewBuffer([]byte{})
	defer Buffers.Reset()
	//声明一个管道用于接收解包的数据
	readerChannel := make(chan string, 1)
	go reader(readerChannel)
	for {
		n, err := conn.Read(buf)
		if CheckSuclogin(buf[:5]) {
			log.Println("LOGIN IS SUC")
			symbols := []string{"XAUUSD", "AUDUSD", "GBPUSD", "EURJPY", "GBPJPY", "GBPNZD", "XAGUSD", "USDJPY"}
			conn.Write(StringByte(GetResult(symbols)))
		}
		if CheckIsHeart(buf[:5]) {
			log.Println("====>>> Send Heart =====")
			byteslist := []byte{2, 1, 3, 0, 0}
			conn.Write(byteslist)
		}
		if n > 5 {
			Buffers.Write(buf[:n])
			go RetDataHandle(Buffers, readerChannel)
		}
		if err != nil && err == io.EOF { //io.EOF在网络编程中表示对端把链接关闭了。
			log.Fatal(err)
		}
	}

}
