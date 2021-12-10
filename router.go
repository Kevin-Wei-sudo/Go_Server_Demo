package zinxNetwork

import "zinx/zinxFace"

type BaseRouter struct {


}

// Router 完全继承BaseRouter方法，不需要实现；因为有些Router并不想要Pre和Post操作

// PreHandle 在处理conn业务之前的钩子方法Hook
func (br*BaseRouter) PreHandle (request zinxFace.IRequest){

}

// Handle 在处理conn业务之中的钩子方法Hook
func (br*BaseRouter) Handle(request zinxFace.IRequest){

}

// PostHandle 在处理conn业务之后的钩子方法Hook
func (br*BaseRouter) PostHandle(request zinxFace.IRequest){

}

