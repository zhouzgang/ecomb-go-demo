package di

import (
	"ecomb-go-demo/web/cfg"
	"sync"
)

const (
	KeyDB = "db"
)

type DependencyInjecter interface {
	Bind(k interface{}, v interface{}) DependencyInjecter
	Unbind(k interface{}) DependencyInjecter
	Get(key interface{}) interface{}
}

type Container struct {
	store sync.Map
}

var _ DependencyInjecter = new(Container)

func New() *Container {
	return &Container{
		store: sync.Map{},
	}
}

func (di *Container) Bind(k interface{}, v interface{}) DependencyInjecter {
	di.store.Store(k, v)
	return di
}

func (di *Container) Unbind(k interface{}) DependencyInjecter {
	di.store.Delete(k)
	return di
}

func (di *Container) Get(key interface{}) interface{} {
	v, _ := di.store.Load(key)
	return v
}

func (di *Container) Init() {
	if err := di.
}

func (di *Container) OpenDBs() error  {
	for k, v := range cfg.C.DB {
		db, err := OpenDB(v)
	}
}