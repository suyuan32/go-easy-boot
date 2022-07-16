package config

import (
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	DB      DBConfig `json:"DBConfig" yaml:"DBConfig"`
	LogConf LogConf  `json:"LogConf" yaml:"LogConf"`
}

type DBConfig struct {
	Type        string `json:"Type" yaml:"Type"`               // type of database: mysql, postpres
	Path        string `json:"Path" yaml:"Path"`               // address
	Port        string `json:"Port" yaml:"Port"`               // port
	Config      string `json:"Config" yaml:"Config"`           // extra config such as charset=utf8mb4&parseTime=True
	Dbname      string `json:"DBName" yaml:"DBName"`           // database name
	Username    string `json:"Username" yaml:"Username"`       // username
	Password    string `json:"Password" yaml:"Password"`       // password
	MaxIdleConn int    `json:"MaxIdleConn" yaml:"MaxIdleConn"` // the maximum number of connections in the idle connection pool
	MaxOpenConn int    `json:"MaxOpenConn" yaml:"MaxOpenConn"` // the maximum number of open connections to the database
	LogMode     string `json:"LogMode" yaml:"LogMode"`         // open gorm's global logger
	LogZap      bool   `json:"LogZap" yaml:"LogZap"`
}

func (d *DBConfig) MysqlDSN() string {
	return d.Username + ":" + d.Password + "@tcp(" + d.Path + ":" + d.Port + ")/" + d.Dbname + "?" + d.Config
}

func (d *DBConfig) PostgresDSN() string {
	return "host=" + d.Path + " user=" + d.Username + " password=" + d.Password + " dbname=" + d.Dbname + " port=" + d.Port + " " + d.Config
}

type LogConf struct {
	ServiceName         string `json:",optional"`                                    // service name
	Mode                string `json:",default=console,options=console|file|volume"` // mode: console-output to console，file-output to file，，volume-output to the docker volume
	Path                string `json:",default=logs"`                                // store path
	Level               string `json:",default=info,options=info|error|severe"`      // the level to be shown
	Compress            bool   `json:",optional"`                                    // gzip compress
	KeepDays            int    `json:",optional"`                                    // the period to be stored
	StackCooldownMillis int    `json:",default=100"`                                 // the period between two writing (ms)
}
