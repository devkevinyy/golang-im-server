package serverlogic

import (
	"log"
	"net"
	"encoding/binary"
	"bytes"
)

var client = make(map[string]net.Conn)


func HandleConn(c net.Conn) {
	client_ip := c.RemoteAddr().String()
	log.Println("client: "+client_ip+" connect success!")
	client[client_ip] = c
	defer c.Close()

	buffer := make([]byte, 1)
	body_length := make([]byte, 4)
	for {
		n, err := c.Read(buffer)
		if err != nil {
			log.Println("error to read: "+err.Error())
			return
		}
		cmd := parseByteArrayToByte(buffer[:n])
		log.Printf("收到包,cmd 为: %d", cmd)

		len, err := c.Read(buffer)
		if err != nil {
			log.Println("error to read: "+err.Error())
			return
		}
		version := parseByteArrayToByte(buffer[:len])
		log.Printf("收到包,version 为: %d", version)

		le, err := c.Read(body_length)
		if err != nil {
			log.Println("error to read: "+err.Error())
			return
		}
		bodyLen := parseByteArrayToUint32(body_length[:le])
		log.Printf("收到包,body length 为: %d", bodyLen)

		dataBuf := make([]byte, bodyLen)
		length, err := c.Read(dataBuf)
		if err != nil {
			log.Println("error to read: "+err.Error())
			return
		}
		data := string(dataBuf[:length])
		log.Println("收到包,data 为: "+data)

		go RouteLogic(c, cmd, version, data)
	}
}

func parseByteArrayToByte(array []byte) byte {
	var value byte
	buf := bytes.NewBuffer(array)
	binary.Read(buf, binary.LittleEndian, &value)
	return value
}

func parseByteArrayToUint32(array []byte) uint32 {
	var value uint32
	buf := bytes.NewBuffer(array)
	binary.Read(buf, binary.LittleEndian, &value)
	return value
}
