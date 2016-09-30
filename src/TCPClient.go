package main

import (
	"net"
	"log"
	"os"
	"encode"
	"encoding/binary"
	"bytes"
	"time"
)

var heartBeatTicker *time.Ticker = time.NewTicker(61 * time.Second)
var heartBeatChan = make(chan bool)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Println("connect failed")
		os.Exit(1)
	}
	go handleRead(conn)
	go HeartBeat(conn)
	//sendHeartBeatData(conn)
	time.Sleep(time.Duration(time.Minute*2))
	//sendLoginData(conn)
	//sendLogoutData(conn)
	//sendUnkownData(conn)
}

func handleRead(c net.Conn)  {
	defer c.Close()
	buf := make([]byte, 1024)
	for {
		len, err := c.Read(buf)
		if err != nil {
			log.Println("error to read message: "+err.Error())
			c.Close()
			heartBeatChan <- true
			return
		}
		log.Println(string(buf[:len]))
	}
}

func HeartBeat(c net.Conn){
	if <- heartBeatChan {
		close(heartBeatChan)
		return
	}
	for t := range heartBeatTicker.C {
		t.Day()
		sendHeartBeatData(c)
	}
}

func sendHeartBeatData(c net.Conn)  {
	con := []byte("客户端心跳包")
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, con)
	data := encode.GetPackage(1, 2, buf.Bytes())
	c.Write(data)
}


func sendLoginData(c net.Conn)  {
	con := []byte("客户端登陆包")
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, con)
	data := encode.GetPackage(2, 2, buf.Bytes())
	c.Write(data)
}

func sendLogoutData(c net.Conn)  {
	con := []byte("客户端登出包")
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, con)
	data := encode.GetPackage(3, 2, buf.Bytes())
	c.Write(data)
}

func sendUnkownData(c net.Conn)  {
	con := []byte("未知逻辑包")
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, con)
	data := encode.GetPackage(4, 2, buf.Bytes())
	c.Write(data)
}