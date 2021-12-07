package main

import "zinx/zinxNetwork"

/*
	基于Zinx框架来开发的服务器端应用程序
 */

func main() {
	// step1:创建一个Server句柄，使用Zinx的Api
	s := zinxNetwork.NewServer("[zinxV0]")
	// step2:启动Server
	s.Run()
}
