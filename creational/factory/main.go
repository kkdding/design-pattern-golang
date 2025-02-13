package main

import "fmt"

// ----- 抽象层 -----

// App 抽象接口
type App interface {
	Run()
}

// AbstractFactory 抽象工厂接口
type AbstractFactory interface {
	CreateApp(version string) App
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

// DouYinFactory douyin工厂类
type DouYinFactory struct{}

func (df *DouYinFactory) CreateApp(version string) App {
	return &DouYin{version: version}
}

// WeiXinFactory weixin工厂类
type WeiXinFactory struct{}

func (wf *WeiXinFactory) CreateApp(version string) App {
	return &WeiXin{version: version}
}

// TaoBaoFactory taobao工厂类
type TaoBaoFactory struct{}

func (tf *TaoBaoFactory) CreateApp(version string) App {
	return &TaoBao{version: version}
}

// 主函数 - 业务逻辑
func main() {
	douyinFactory := new(DouYinFactory)
	douyinFactory.CreateApp("v1.0").Run()
	douyinFactory.CreateApp("v2.0").Run()

	weixinFactory := new(WeiXinFactory)
	weixinFactory.CreateApp("v1.0").Run()
	weixinFactory.CreateApp("v2.0").Run()

	taobaoFactory := new(TaoBaoFactory)
	taobaoFactory.CreateApp("v1.0").Run()
	taobaoFactory.CreateApp("v2.0").Run()
}
