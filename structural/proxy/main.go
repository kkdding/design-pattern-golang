package main

import "fmt"

type Resource struct {
	kind   string
	id     int
	exists bool
}

// ----- 抽象层 -----

// System 抽象接口
type System interface {
	Get(resource *Resource)
}

// ----- 实现层 -----

// PictureSystem 实体类
type PictureSystem struct{}

func (p *PictureSystem) Get(resource *Resource) {
	fmt.Println("Get [Picture] Resource")
}

// VideoSystem 实体类
type VideoSystem struct{}

func (v *VideoSystem) Get(resource *Resource) {
	fmt.Println("Get [Video] Resource")
}

// TextSystem 实体类
type TextSystem struct{}

func (t *TextSystem) Get(resource *Resource) {
	fmt.Println("Get [Text] Resource")
}

// ----- 代理层 -----

// ResourceProxy 代理类
type ResourceProxy struct {
	system System
}

func (rp *ResourceProxy) Get(resource *Resource) {
	if rp.Exists(resource) == false {
		fmt.Println("Resource -", resource.kind, "-", resource.id, " not exists")
		return
	}
	rp.system.Get(resource)
	rp.Cache(resource)
}

func (rp *ResourceProxy) Exists(resource *Resource) bool {
	return resource.exists
}

func (rp *ResourceProxy) Cache(resource *Resource) {
	fmt.Println("Cache Resource -", resource.kind, "-", resource.id)
}

func NewProxy(system System) System {
	return &ResourceProxy{system: system}
}

// 主函数 - 业务逻辑
func main() {
	r1 := &Resource{
		id:     1,
		kind:   "picture",
		exists: true,
	}

	r2 := &Resource{
		id:     2,
		kind:   "picture",
		exists: false,
	}

	system := &PictureSystem{}
	proxy := NewProxy(system)
	proxy.Get(r1)
	proxy.Get(r2)
}
