package main

import "gitee.com/aurora-engine/aurora"

// Server 嵌套Engine定义一个服务 实例
type Server struct {
	*aurora.Engine
}

func (server *Server) Server() {
	// 进行一下初始化操作，比如 控制器实例，全局中间件，全局变量，第三方依赖库等操作
	// 请不要在 Server()函数中进行路由注册，Server()函数 仅负责加载组件依赖
}

func (server *Server) Router() {
	// 添加 web 路由
	// Router() 函数内可以做任何处理，路由分组，路由中间件等

	server.Get("/", func() string {
		return "hello world"
	})
}
func main() {
	err := aurora.Run(&Server{aurora.New()})
	if err != nil {
		panic(err)
	}
}
