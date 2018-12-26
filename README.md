## vendor依赖包

1. 下载vendor包 `go get -u github.com/kardianos/govendor`
2. 进入到依赖包目录编译二进制文件,`go build` 生成`govendor`文件
3. 可以执行文件拷贝到`bin`目录或添加到`GOPATH`
4. 进入到工程项目
* `govendor init` 生成`vendor`目录
* 执行` govendor add +external `导入依赖包
5. 已经导入依赖的编译包，此时执行`go build `使用的包全在vendor下

## 交叉编译

    build.sh
    
## 配置文件
不提交本地配置文件/conf/conf.toml，线上单独配置
## 部署上线
1. supervisor
2. 执行sh build.sh，生成编译文件和复制相关前端依赖目录\文件
3. 停止supervisor服务，线上拉dist目录，启动supervisor服务

## Supervisor

1. 安装
2. 创建配置文件 `echo_supervisord_conf > /etc/supervisord.conf`
3. 启动 `systemctl start supervisord.service && systemctl enable supervisord.service`或者加载的时候`supervisord -c /etc/supervisord.conf`