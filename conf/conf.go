package conf

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"path/filepath"
	"sync"
)

type tomlConfig struct {
	Port     string
	Database database
	Redis    redis
	Quote    quote
}

type database struct {
	DbHost string
	DbPort string
	DbUser string
	DbPass string
	DbName string
	DbType string
}

type redis struct {
	RdbHost string
	RdbPort string
	RdbPass string
}

type quote struct {
	QUName string
	QPsswd string
	QPort  string
	QHost  string
}

//读取环境变量
var conf *tomlConfig
var once sync.Once
var confLock = new(sync.RWMutex)

/** 平滑的替换掉配置，不需要重启项目,起一个协程，监视我们定义好的信号，
如果接收到信号就重新加载配置。
在Config方法中获取Conf的时候加了读锁，防止在读的时候，也在写入，导致配置错乱。
*/
func Config() *tomlConfig {
	once.Do(ReloadConfig)
	confLock.RLock()
	defer confLock.RUnlock()
	return conf
}

func ReloadConfig() {
	filePath, err := filepath.Abs("./conf/conf.toml")
	if err != nil {
		panic(err)
	}

	fmt.Printf("解析文件, 路径:%s\n", filePath)

	config := new(tomlConfig)

	if _, err := toml.DecodeFile(filePath, config); err != nil {
		panic(err)
	}

	confLock.Lock()
	defer confLock.Unlock()
	conf = config
}
