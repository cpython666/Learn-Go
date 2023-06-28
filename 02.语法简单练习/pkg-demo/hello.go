package main

import "gitee.com/aurora-engine/aurora"

// Server 嵌套Engine定义一个服务 实例
type Server struct {
	*aurora.Engine
}
