package main

import "fmt"

// ----- 抽象层 -----

// AbstractDouYin 抽象接口
type AbstractDouYin interface {
	RunDouYin()
}

// AbstractWeiXin 抽象接口
type AbstractWeiXin interface {
	RunWeiXin()
}

// AbstractTaoBao 抽象接口
type AbstractTaoBao interface {
	RunTaoBao()
}

// AbstractFactory 抽象工厂接口
type AbstractFactory interface {
	CreateDouYin(version string) AbstractDouYin
	CreateWeiXin(version string) AbstractWeiXin
	CreateTaoBao(version string) AbstractTaoBao
}

// ----- 实体层 -----

// CN App & Factory 实现

type DouYinCN struct {
	version string
}

func (dcn *DouYinCN) RunDouYin() {
	fmt.Printf("DouYinCN-%s is running\n", dcn.version)
}

type WeiXinCN struct {
	version string
}

func (wcn *WeiXinCN) RunWeiXin() {
	fmt.Printf("WeiXinCN-%s is running\n", wcn.version)
}

type TaoBaoCN struct {
	version string
}

func (tcn *TaoBaoCN) RunTaoBao() {
	fmt.Printf("TaoBaoCN-%s is running\n", tcn.version)
}

type CNFactory struct{}

func (cn *CNFactory) CreateDouYin(version string) AbstractDouYin {
	return &DouYinCN{version: version}
}

func (cn *CNFactory) CreateWeiXin(version string) AbstractWeiXin {
	return &WeiXinCN{version: version}
}

func (cn *CNFactory) CreateTaoBao(version string) AbstractTaoBao {
	return &TaoBaoCN{version: version}
}

// US App & Factory 实现

type DouYinUS struct {
	version string
}

func (dus *DouYinUS) RunDouYin() {
	fmt.Printf("DouYinUS-%s is running\n", dus.version)
}

type WeiXinUS struct {
	version string
}

func (wus *WeiXinUS) RunWeiXin() {
	fmt.Printf("WeiXinUS-%s is running\n", wus.version)
}

type TaoBaoUS struct {
	version string
}

func (tus *TaoBaoUS) RunTaoBao() {
	fmt.Printf("TaoBaoUS-%s is running\n", tus.version)
}

type USFactory struct{}

func (us *USFactory) CreateDouYin(version string) AbstractDouYin {
	return &DouYinUS{version: version}
}

func (us *USFactory) CreateWeiXin(version string) AbstractWeiXin {
	return &WeiXinUS{version: version}
}

func (us *USFactory) CreateTaoBao(version string) AbstractTaoBao {
	return &TaoBaoUS{version: version}
}

// 主函数 - 业务逻辑
func main() {
	cnFactory := new(CNFactory)
	cnFactory.CreateDouYin("1.0").RunDouYin()
	cnFactory.CreateDouYin("2.0").RunDouYin()
	cnFactory.CreateWeiXin("1.0").RunWeiXin()
	cnFactory.CreateTaoBao("1.0").RunTaoBao()

	usFactory := new(USFactory)
	usFactory.CreateDouYin("1.0").RunDouYin()
	usFactory.CreateWeiXin("1.0").RunWeiXin()
	usFactory.CreateTaoBao("1.0").RunTaoBao()
	usFactory.CreateTaoBao("2.0").RunTaoBao()
}
