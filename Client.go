package main

import (
	"fmt"
	"net"
	"time"
)

/*
  simulate client
 */
func main()  {
	fmt.Println("client start...")

	time.Sleep(1 *time.Second)
	// step1 connect to remote server, and then get the connection link
	conn, err := net.Dial("tcp", "127.0.0.1:8999")
	if err != nil {
		fmt.Println("client start err, exit")
		return
	}

	for {
		//step 2 use write function to input data
		_, err := conn.Write([]byte("Hello Zinx V0"))
		if err != nil {
			fmt.Println("write conn err")
			return
		}

		buf := make([]byte, 512)
		cnt, err := conn.Read(buf)
		if err != nil {
			fmt.Println("read buf error")
			return
		}

		fmt.Printf("server call back: %s, cnt = %d\n", buf, cnt)

		// cpu阻塞
		time.Sleep(1*time.Second)
	}


}
