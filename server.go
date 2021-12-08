package zinxNetwork

import (
	"errors"
	"fmt"
	"net"
	"zinx/zinxFace"
)

// iServer api implement, define

type Server struct {
	Name string
	IPVersion string
	Ip string
	Port int
}

// CallBackToClient 定义当前客户端链接的所绑定Handle api（目前这个Handle是写死的）
func CallBackToClient( conn *net.TCPConn, data []byte, cnt int) error  {
	// 回显的业务
	fmt.Println("[Conn Handle] CallBackToClient...")
	if _, err := conn.Write(data[:cnt]); err != nil {
		fmt.Println("write back buf err", err)
		return errors.New("CallBackToClient error")
	}
	return nil
}

func (s * Server) Start()  {
	fmt.Printf("[Start] Server Listenner at IP: %s, Port %d, is Starting\n", s.Ip, s.Port)

	go func() {
		// step1 : acquire a TCP address
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.Ip, s.Port))

		if err != nil {
			fmt.Println("resolve the tcp address error", err)
			return
		}
		// step2 : try to listen TCP address

		listener, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			fmt.Println("listen", s.IPVersion, "err", err)
			return
		}

		fmt.Println("start Zinx server succeed", s.Name, " is listening")
		var cid uint32
		cid = 0

		// step3 : if succeed, wait until client connect to the server, then deal with the action(read & write)
		for {
			// if client join in
			conn, err := listener.AcceptTCP()
			if err != nil {
				fmt.Println("Accept err", err)
				continue
			}

			// 将处理新连接的业务方法和conn进行绑定，得到我们的链接模块
			dealConn := NewConnection(conn, cid,CallBackToClient )
			cid ++

			// 启动当前的链接业务处理
			go dealConn.Start()

			// if server connects to the client, do the basic thing such as returning 512 byte

			go func() {
				for {
					buf := make([]byte, 512)
					cnt, err := conn.Read(buf)
					if err != nil {
						fmt.Println("receive buf err", err)
						continue
					}

					fmt.Printf("recv client buf %s, cnt %d\n", buf, cnt)
					// 回显功能
					if _, err := conn.Write(buf[:cnt]); err != nil {
						fmt.Println("write back buf err", err)
						continue
					}
				}
			}()

		}
	}()

}

func (s * Server) Stop()  {
	// TODO 将服务器的资源、状态或者一些已经开辟的链接信息进行回收

}

func (s * Server) Run()  {
	// Start server function
	s.Start()

	// TODO 做一些启动服务器之外额外的业务

	// 阻塞状态
	select {

	}
}

func NewServer(name string) zinxFace.IServer  {
	s := &Server{
		Name : name,
		IPVersion : "tcp4",
		Ip : "0.0.0.0",
		Port: 8999,
	}
	return s
}