package main

import "fmt"

// Handler 抽象类 包含处理程序的步骤
type Handler interface {
	Request()
	Check()
	Save()
	Response()

	NeedBackup() bool
}

// Template 封装流程模版
type Template struct {
	h Handler
}

func (t *Template) NetHandler() {
	t.h.Request()
	t.h.Check()
	t.h.Save()
	t.h.Response()

	if t.h.NeedBackup() {
		fmt.Println("备份数据")
	}
}

// HTTPHandler 模版实现类
type HTTPHandler struct {
	Template
}

func NewHTTPHandler() *HTTPHandler {
	httpHandler := new(HTTPHandler)
	httpHandler.h = httpHandler
	return httpHandler
}

func (h *HTTPHandler) Request() {
	fmt.Println("处理HTTP请求")
}

func (h *HTTPHandler) Check() {
	fmt.Println("校验数据")
}

func (h *HTTPHandler) Save() {
	fmt.Println("保存数据")
}

func (h *HTTPHandler) Response() {
	fmt.Println("响应请求")
}

func (h *HTTPHandler) NeedBackup() bool {
	return true
}

// TCPHandler 模版实现类
type TCPHandler struct {
	Template
}

func NewTCPHandler() *TCPHandler {
	tcpHandler := new(TCPHandler)
	tcpHandler.h = tcpHandler
	return tcpHandler
}

func (h *TCPHandler) Request() {
	fmt.Println("处理TCP请求")
}

func (h *TCPHandler) Check() {
	fmt.Println("校验数据")
}

func (h *TCPHandler) Save() {
	fmt.Println("保存数据")
}

func (h *TCPHandler) Response() {
	fmt.Println("响应请求")
}

func (h *TCPHandler) NeedBackup() bool {
	return false
}

// 主函数 - 业务逻辑
func main() {
	httpHandler := NewHTTPHandler()
	httpHandler.NetHandler()

	tcpHandler := NewTCPHandler()
	tcpHandler.NetHandler()
}
