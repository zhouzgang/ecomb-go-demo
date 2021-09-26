package di

import (
	"ecomb-go-demo/web/cfg"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
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
	if err := di.OpenDBs(); err != nil {
		log.Fatal().Msg("fail to init DI, due to DB error")
	}
}

func (di *Container) OpenDBs() error  {
	for k, v := range cfg.C.DB {
		db, err := OpenDB(v)
		if err != nil {
			return err
		}
		di.Bind(KeyDB + k, db)
	}
	return nil
}

func (di *Container) DB(dbname string) *gorm.DB  {
	db := di.Get(KeyDB + dbname)
	if db == nil {
		log.Error().Msgf("db %s id nil", dbname)
		return nil
	}
	// todo 这里的写法没搞明白
	return db.(*gorm.DB)
}