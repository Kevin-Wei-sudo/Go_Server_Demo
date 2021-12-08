package zinxFace

import "net"

type IConnection interface {
	// Start 启动链接 让当前的链接准备开始工作
	Start()

	// Stop 停止链接 结束当前链接的工作
	Stop()

	// GetTCPConnection 获取当前链接的绑定socket conn
	GetTCPConnection() *net.TCPConn

	// GetConnID 获取当前链接模块的链接ID
	GetConnID() uint32

	// RemoteAddr 获取远程客户端的TCP状态 IP Port
	RemoteAddr() net.Addr

	// Send 发送数据，将数据发送给远程的客户端
	Send(data []byte) error

}

// HandleFunc 定义处理业务的方式
type HandleFunc func(*net.TCPConn, []byte, int) error