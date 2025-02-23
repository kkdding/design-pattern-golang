package main

import "fmt"

type Handler interface {
	execute(*Product)
	nextHandler(Handler)
}

type HandlerA struct {
	next Handler
}

func (a *HandlerA) execute(p *Product) {
	if p.ADone {
		fmt.Println("A is Done")
		a.next.execute(p)
		return
	}
	fmt.Println("A Done")
	p.ADone = true
	a.next.execute(p)
}

func (a *HandlerA) nextHandler(h Handler) {
	a.next = h
}

type HandlerB struct {
	next Handler
}

func (b *HandlerB) execute(p *Product) {
	if p.BDone {
		fmt.Println("B is Done")
		b.next.execute(p)
		return
	}
	fmt.Println("B Done")
	p.BDone = true
	b.next.execute(p)
}

func (b *HandlerB) nextHandler(h Handler) {
	b.next = h
}

type HandlerC struct {
	next Handler
}

func (c *HandlerC) execute(p *Product) {
	if p.CDone {
		fmt.Println("C is Done")
		return
	}
	fmt.Println("C Done")
	p.CDone = true
}

func (c *HandlerC) nextHandler(h Handler) {
	c.next = h
}

type Product struct {
	name  string
	ADone bool
	BDone bool
	CDone bool
}

func main() {
	a := &HandlerA{}
	b := &HandlerB{}
	c := &HandlerC{}

	a.nextHandler(b)
	b.nextHandler(c)

	p := &Product{name: "Product"}
	a.execute(p)

	fmt.Println(p)
}
