package main

import "fmt"

// 适配目标

// JsonData 数据结构接口
type JsonData interface {
	JsonFormat()
}

// BusinessData 实体类依赖 JsonData 接口
type BusinessData struct {
	data JsonData
}

func NewBusinessData(data JsonData) *BusinessData {
	return &BusinessData{data: data}
}

func (b *BusinessData) showData() {
	b.data.JsonFormat()
	fmt.Printf(" json data\n json data\n json data\n")
}

// 适配者

type XmlData struct{}

func (x XmlData) XmlFormat() {}

// 适配器

// DataAdapter 数据适配器
type DataAdapter struct {
	xmlData *XmlData
}

func (d *DataAdapter) JsonFormat() {
	d.xmlData.XmlFormat()
}

func NewDataAdapter(xmlData *XmlData) *DataAdapter {
	return &DataAdapter{xmlData: xmlData}
}

// 主函数 - 业务逻辑
func main() {
	configXmlData := NewBusinessData(NewDataAdapter(new(XmlData)))
	configXmlData.showData()
}
