package cfg

import (
	"ecomb-go-demo/web/pkg/localtime"
	"fmt"
	"net/url"
	"time"
)

const dns = "%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=true&interpolateParams=true&loc=%s"

type MySqlConfig struct {
	Host            string        `mapstructure:"host"`
	Port            int           `mapstructure:"port"`
	Database        string        `mapstructure:"database"`
	Username        string        `mapstructure:"username"`
	Password        string        `mapstructure:"password"`
	MaxIdleConns    int           `mapstructure:"max_idle_conns"`
	MaxOpenConns    int           `mapstructure:"max_open_conns"`
	ConnMaxLifetime time.Duration `mapstructure:"conn_max_life_time"`
	Timezone        string        `mapstructure:"timezone"`
}

func (mc *MySqlConfig) BuildDNS() string {
	if mc.Timezone == "" {
		mc.Timezone = localtime.Timezone
	}
	return fmt.Sprintf(
		dns,
		mc.Username,
		mc.Password,
		mc.Host,
		mc.Port,
		mc.Database,
		url.QueryEscape(mc.Timezone),
		)
}