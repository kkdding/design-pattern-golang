package main

import "fmt"

// ----- 抽象层 -----

// App 抽象接口
type App interface {
	Run()
}

// ----- 实体层 -----

// DouYin 实体
type DouYin struct {
	version string
}

func (d *DouYin) Run() {
	fmt.Printf("DouYin-%s is running\n", d.version)
}

// WeiXin 实体
type WeiXin struct {
	version string
}

func (w *WeiXin) Run() {
	fmt.Printf("WeiXin-%s is running\n", w.version)
}

// TaoBao 实体
type TaoBao struct {
	version string
}

func (t *TaoBao) Run() {
	fmt.Printf("TaoBao-%s is running\n", t.version)
}

// ----- 工厂模块 -----

// AppFactory 工厂类
type AppFactory struct{}

func (fac *AppFactory) Create(appName string, version string) App {
	switch appName {
	case "DouYin":
		return &DouYin{version: version}
	case "WeiXin":
		return &WeiXin{version: version}
	case "TaoBao":
		return &TaoBao{version: version}
	default:
		fmt.Println("不支持的app")
		return nil
	}
}

// 主函数 - 业务逻辑
func main() {
	factory := AppFactory{}
	fmt.Println("app工厂已创建...")

	douyin := factory.Create("DouYin", "1.0.0")
	douyin.Run()

	weixin := factory.Create("WeiXin", "1.0.0")
	weixin.Run()

	taobao := factory.Create("TaoBao", "1.0.0")
	taobao.Run()
}
