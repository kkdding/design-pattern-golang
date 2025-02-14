package main

import "fmt"

// ----- 抽象层 -----

// AI 抽象接口
type AI interface {
	Dialogue()
}

type Decorator struct {
	ai AI
}

func (d *Decorator) Dialogue() {}

// ----- 实现层 -----

// GPT 实现类
type GPT struct{}

func (gpt *GPT) Dialogue() {
	fmt.Println("dialogue with GPT")
}

// DouBao 实现类
type DouBao struct{}

func (db *DouBao) Dialogue() {
	fmt.Println("dialogue with DouBao")
}

// ----- 装饰器层 -----

// GraphDecorator 装饰器实现类
type GraphDecorator struct {
	Decorator
}

func (gd *GraphDecorator) Dialogue() {
	gd.ai.Dialogue()
	fmt.Println("generate graph")
}

func NewGraphDecorator(ai AI) AI {
	return &GraphDecorator{Decorator{ai: ai}}
}

// VideoDecorator 装饰器实现类
type VideoDecorator struct {
	Decorator
}

func (vd *VideoDecorator) Dialogue() {
	vd.ai.Dialogue()
	fmt.Println("generate video")
}

func NewVideoDecorator(ai AI) AI {
	return &VideoDecorator{Decorator{ai: ai}}
}

// 主函数 - 业务逻辑
func main() {
	gpt := new(GPT)
	gpt.Dialogue()

	gptGenerateGraph := NewGraphDecorator(gpt)
	gptGenerateGraph.Dialogue()

	douBao := new(DouBao)
	douBao.Dialogue()

	douBaoGenerateVideo := NewVideoDecorator(douBao)
	douBaoGenerateVideo.Dialogue()
}
