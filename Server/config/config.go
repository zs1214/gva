package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Config struct {
	MysqlAdmin   MysqlAdmin   `json:"mysqlAdmin"`
	Qiniu        Qiniu        `json:"qiniu"`
	CasbinConfig CasbinConfig `json:"casbinConfig"`
	RedisAdmin   RedisAdmin   `json:"redisAdmin"`
	System       System       `json:"system"`
	JWT          JWT          `json:"jwt"`
	Captcha      Captcha      `json:"captcha"`
	Log          Log          `json:"log"`
}

type System struct {
	UseMultipoint bool   `json:"useMultipoint"`
	Env           string `json:"env"`
	Addr          int    `json:"addr"`
}

type JWT struct {
	SigningKey string `json:"signingKey"`
}

type CasbinConfig struct {
	ModelPath string `json:"modelPath"`
}

type MysqlAdmin struct { // mysql admin 数据库配置
	Username     string `json:"username"`
	Password     string `json:"password"`
	Path         string `json:"path"`
	Dbname       string `json:"dbname"`
	Config       string `json:"config"`
	MaxIdleConns int    `json:"maxIdleConns"`
	MaxOpenConns int    `json:"maxOpenConns"`
	LogMode      bool   `json:"maxOpenConns"`
}

type RedisAdmin struct { // Redis admin 数据库配置
	Addr     string `json:"addr"`
	Password string `json:"password"`
	DB       int    `json:"db"`
}
type Qiniu struct { // 七牛 密钥配置
	AccessKey string `json:"accessKey"`
	SecretKey string `json:"secretKey"`
}

type Captcha struct { // 验证码配置
	KeyLong   int `json:"keyLong"`
	ImgWidth  int `json:"imgWidth"`
	ImgHeight int `json:"imgHeight"`
}

type Log struct {
	//打印前缀
	Prefix string `json:"prefix"`
	//是否显示打印log文件具体路径
	LogFile bool `json:"logFile"`
	// 在控制台打印log的级别， []默认不打印
	Stdout []string `json:"stdout"`
	// 在文件中打印log的级别   []默认不打印
	File []string `json:"file"`
}

var GinVueAdminconfig Config
var VTool *viper.Viper

func init() {
	v := viper.New()
	v.SetConfigName("config")           //设置配置文件名
	v.AddConfigPath("./static/config/") //第一个搜索路径
	v.SetConfigType("json")
	err := v.ReadInConfig() //搜索路径，读取配置
	if err != nil {
		panic(fmt.Errorf("Fatal error config file : %s\n", err))
	}
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err := v.Unmarshal(&GinVueAdminconfig); err != nil {
			fmt.Println(err)
		}
	})
	if err := v.Unmarshal(&GinVueAdminconfig); err != nil {
		fmt.Println(err)
	}
	VTool = v
}
