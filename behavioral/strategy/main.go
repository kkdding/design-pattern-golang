package main

import "fmt"

// PayStrategy 支付策略抽象类
type PayStrategy interface {
	Pay() //使用武器
}

// WeiXinPay 具体策略
type WeiXinPay struct{}

func (wx *WeiXinPay) Pay() {
	fmt.Println("使用微信支付")
}

// AliPay 具体策略
type AliPay struct{}

func (ali *AliPay) Pay() {
	fmt.Println("使用支付宝支付")
}

// Buy 支付场景
type Buy struct {
	strategy PayStrategy //拥有一个抽象的策略
}

// SetPayStrategy 设置一个支付策略
func (b *Buy) SetPayStrategy(p PayStrategy) {
	b.strategy = p
}

func (b *Buy) UsePay() {
	b.strategy.Pay() //调用支付策略
}

func main() {
	customer := &Buy{strategy: &WeiXinPay{}}

	// 微信支付
	customer.SetPayStrategy(new(WeiXinPay))
	customer.UsePay()

	// 支付宝支付
	customer.SetPayStrategy(new(AliPay))
	customer.UsePay()
}
