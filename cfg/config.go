package cfg

import (
	"encoding/json"
	"io/ioutil"
)

// 定义配置信息结构体，从配置文件读入
//mysql config
type Connect struct {
	Database string `json:"database"` // MySQL数据库名称
	Host     string `json:"host"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
}

type Config struct {
	Host       string `json:"host"` // 监听的http地址
	Port       string `json:"port"` // 监听的http端口
	Connection Connect
}

// 读入配置文件
func LoadConfig(file string) (c *Config, err error) {
	// 将文件读到内存中，为一个切片类型
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	// json解析切片数据，反序列化到结构体中
	c = &Config{}
	err = json.Unmarshal(data, c)
	return c, err
}
