package serverlogic

import (
	"log"
	"net"
	"time"
)

func RouteLogic(conn net.Conn, cmd, version byte, content string)  {
	log.Printf("handle client login - cmd: %d - version: %d - content: %s", cmd, version, content)
	switch cmd {
		case 1:  // tcp 心跳
			log.Println("收到心跳，刷新 tcp 保活时间")
			conn.SetDeadline(time.Now().Add(time.Minute * 1))
		case 2:
			log.Println("客户端登陆")
		case 3:
			log.Println("客户端登出")
		default:
			log.Println("unhandle logic")
			return
	}
}
