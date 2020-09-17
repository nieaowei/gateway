package lib

import (
	"github.com/spf13/viper"
)

type Config interface {
	ConfName() string
}

type BaseConf struct {
	Base struct {
		DebugMode    string `mapstructure:"debug_mode"`
		TimeLocation string `mapstructure:"time_location"`
		StaticPath   string `mapstructure:"static_path"`
		Domain       string `mapstructure:"domain"`
	} `mapstructure:"base"`
	Http struct {
		Addr           string   `mapstructure:"addr"`
		ReadTimeout    int      `mapstructure:"read_timeout"`
		WriteTimeout   int      `mapstructure:"write_timeout"`
		MaxHeaderBytes int      `mapstructure:"max_header_bytes"`
		AllowIP        []string `mapstructure:"allow_ip"`
	} `mapstructure:"http"`
	Cluster struct {
		Ip      string `mapstructure:"ip"`
		Port    string `mapstructure:"port"`
		SslPort string `mapstructure:"ssl_port"`
	} `mapstructure:"cluster"`
}

func (p *BaseConf) ConfName() string {
	return "base"
}

type LogConfFileWriter struct {
	On              bool   `mapstructure:"on"`
	LogPath         string `mapstructure:"log_path"`
	RotateLogPath   string `mapstructure:"rotate_log_path"`
	WfLogPath       string `mapstructure:"wf_log_path"`
	RotateWfLogPath string `mapstructure:"rotate_wf_log_path"`
}

type LogConfConsoleWriter struct {
	On    bool `mapstructure:"on"`
	Color bool `mapstructure:"color"`
}

type LogConfig struct {
	Level string               `mapstructure:"log_level"`
	FW    LogConfFileWriter    `mapstructure:"file_writer"`
	CW    LogConfConsoleWriter `mapstructure:"console_writer"`
}

type MysqlConf struct {
	List map[string]*MySQLConf `mapstructure:"list"`
}

func (p *MysqlConf) ConfName() string {
	return "mysql"
}

type MySQLConf struct {
	DriverName      string `mapstructure:"driver_name"`
	DataSourceName  string `mapstructure:"data_source_name"`
	MaxOpenConn     int    `mapstructure:"max_open_conn"`
	MaxIdleConn     int    `mapstructure:"max_idle_conn"`
	MaxConnLifeTime int    `mapstructure:"max_conn_life_time"`
}

type RedisMapConf struct {
	List map[string]*RedisConf `mapstructure:"list"`
}

func (p *RedisMapConf) ConfName() string {
	return "redis"
}

type RedisConf struct {
	ProxyList []string `mapstructure:"proxy_list"`
	MaxActive int      `mapstructure:"max_active"`
	MaxIdle   int      `mapstructure:"max_idle"`
	DownGrade bool     `mapstructure:"down_grade"`
}

type ProxyConf struct {
	Base struct {
		DebugMode    string `mapstructure:"debug_mode"`
		TimeLocation string `mapstructure:"time_location"`
		StaticPath   string `mapstructure:"static_path"`
		Domain       string `mapstructure:"domain"`
	} `mapstructure:"base"`
	Http struct {
		Addr           string `mapstructure:"addr"`
		ReadTimeout    int    `mapstructure:"read_timeout"`
		WriteTimeout   int    `mapstructure:"write_timeout"`
		MaxHeaderBytes int    `mapstructure:"max_header_bytes"`
	} `mapstructure:"http"`
	Https struct {
		Addr           string `mapstructure:"addr"`
		ReadTimeout    int    `mapstructure:"read_timeout"`
		WriteTimeout   int    `mapstructure:"write_timeout"`
		MaxHeaderBytes int    `mapstructure:"max_header_bytes"`
	} `mapstructure:"https"`
}

func (p *ProxyConf) ConfName() string {
	return "proxy"
}

//全局变量
var ConfBase *BaseConf
var ConfRedis *RedisMapConf
var ConfMysql *MysqlConf
var ConfProxy *ProxyConf

var ViperConfMap map[string]*viper.Viper = map[string]*viper.Viper{}

func InitConf(path string, config Config) (err error) {
	v := viper.New()
	v.SetConfigName(config.ConfName())
	if path == "" {
		v.AddConfigPath("../conf/pro")
	}
	v.AddConfigPath(path)
	err = v.ReadInConfig()
	if err != nil { // Handle errors reading the config file
		return
	}
	err = v.Unmarshal(config)
	if err != nil { // Handle errors reading the config file
		return
	}
	return
}

func InitBaseConf(path string) {
	conf := &BaseConf{}
	err := InitConf(path, conf)
	if err != nil {
		panic("init conf base")
	}
	ConfBase = conf
	return
}

func InitRedisConf(path string) {
	conf := &RedisMapConf{}
	err := InitConf(path, conf)
	if err != nil {
		panic("init conf redis")
	}
	ConfRedis = conf
	return
}

func InitMysqlConf(path string) {
	conf := &MysqlConf{}
	err := InitConf(path, conf)
	if err != nil {
		panic("init conf mysql")
	}
	ConfMysql = conf
	return
}

func InitProxyConf(path string) {
	conf := &ProxyConf{}
	err := InitConf(path, conf)
	if err != nil {
		panic("init conf proxy " + err.Error())
	}
	ConfProxy = conf
	return
}

//初始化配置文件
func InitViperConf() error {

	return nil
}

func GetDefaultConfMysql() *MySQLConf {
	if ConfMysql == nil {
		InitMysqlConf("")
	}
	return ConfMysql.List["default"]
}

func GetDefaultConfRedis() *RedisConf {
	if ConfRedis == nil {
		InitRedisConf("")
	}
	return ConfRedis.List["default"]
}

func GetDefaultConfBase() *BaseConf {
	if ConfBase == nil {
		InitBaseConf("")
	}
	return ConfBase
}

func GetDefaultConfProxy() *ProxyConf {
	if ConfProxy == nil {
		InitProxyConf("")
	}
	return ConfProxy
}
