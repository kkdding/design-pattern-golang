package main

import "fmt"

// Server 服务器-命令接收者
type Server struct{}

func (s *Server) functionA() {
	fmt.Println("调用功能A")
}

func (s *Server) functionB() {
	fmt.Println("调用功能B")
}

// Command 抽象的命令
type Command interface {
	Request()
}

// CommandFunctionA 功能A命令队列
type CommandFunctionA struct {
	server *Server
}

func (cmdA *CommandFunctionA) Request() {
	cmdA.server.functionA()
}

// CommandFunctionB 功能B命令队列
type CommandFunctionB struct {
	server *Server
}

func (cmdB *CommandFunctionB) Request() {
	cmdB.server.functionB()
}

// Scheduler 调用命令者
type Scheduler struct {
	CmdList []Command
}

// Notify 发送命令的方法
func (s *Scheduler) Notify() {
	if s.CmdList == nil {
		return
	}

	for _, cmd := range s.CmdList {
		cmd.Request()
	}
}

func main() {
	server := new(Server)
	cmdA := CommandFunctionA{server: server}
	cmdB := CommandFunctionB{server: server}

	scheduler := new(Scheduler)
	//收集命令
	scheduler.CmdList = append(scheduler.CmdList, &cmdA)
	scheduler.CmdList = append(scheduler.CmdList, &cmdB)

	//执行指令
	scheduler.Notify()
}
