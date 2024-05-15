package global

import (
	"github.com/gin-gonic/gin"
	"github.com/hashicorp/consul/api"
	"zg5/Homework01/server/proto/server"
)

var (
	ApiALLConf   *ApiAllClient
	NacosConf    *NacosConfig
	ConsulClient *api.Client
	GinClient    *gin.Engine
	ServerClient server.ServerClient
)

type ApiAllClient struct {
	ApiIp   *ApiIp         `yaml:"ApiIp"`
	Mysql   *MysqlConfig   `yaml:"Mysql"`
	Consul  *ConsulConfig  `yaml:"Consul"`
	Elastic *ElasticConfig `yaml:"Elastic"`
	Redis   *RedisConfig   `yaml:"Redis"`
}

type ApiIp struct {
	Host string `yaml:"Host"`
	Port int    `yaml:"Port"`
}

type NacosConfig struct {
	NamespaceId string `yaml:"NamespaceId"`
	DataId      string `yaml:"DataId"`
	Group       string `yaml:"Group"`
	Host        string `yaml:"Host"`
	Port        int    `yaml:"Port"`
}

type ConsulConfig struct {
	Id            string   `yaml:"Id"`
	Name          string   `yaml:"Name"`
	Tags          []string `yaml:"Tags"`
	ConsulRpcName string   `yaml:"ConsulRpcName"`
}

type MysqlConfig struct {
	Username string `yaml:"Username"`
	Password string `yaml:"Password"`
	Host     string `yaml:"Host"`
	Port     int    `yaml:"Port"`
	Dbname   string `yaml:"Dbname"`
}

type ElasticConfig struct {
	Host string `yaml:"Host"`
	Port int    `yaml:"Port"`
}

type RedisConfig struct {
	Host string `yaml:"Host"`
	Port int    `yaml:"Port"`
}
