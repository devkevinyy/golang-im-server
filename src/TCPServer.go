package main

import(
	"net"
	"log"
	"flag"
	"util"
	"fmt"
	"serverlogic"
	"time"
)

const (
	Name    string = "Golang-IM-Server"
	Version string = "1.0"
	TIME_OUT int = 1
)


func startServer(config *util.IMConfig)  {
	log.Println("*********************************************")
	log.Printf("       系统:[%s] 版本:[%s]    ", Name, Version)
	log.Println("*********************************************")
	address := fmt.Sprintf("%s:%d", config.IMHost, config.IMPort)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Println(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			return
		}
		conn.SetDeadline(time.Now().Add(time.Minute * 1))
		go serverlogic.HandleConn(conn)
	}
}

func main() {
	configPath := flag.String("config", "/Users/yangchujie/Go_Projects/HKApp/src/util/config.json", "Configuration file to use")
	flag.Parse()
	config, err := util.ReadConfig(*configPath)
	if err != nil {
		log.Fatalf("读取配置文件错误: %s", err.Error())
	}
	startServer(config)
}
