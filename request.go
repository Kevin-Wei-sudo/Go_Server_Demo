package zinxNetwork

import "zinx/zinxFace"

type Request struct {
	// 已经和客户端建立好的连接
	conn zinxFace.IConnection

	// 客户端请求的数据
	data []byte
}

// GetConnection 得到当前链接
func (r*Request) GetConnection() zinxFace.IConnection  {
	return r.conn
}

// GetData 得到请求的消息数据
func (r*Request) GetData()[] byte  {
	return r.data
}
