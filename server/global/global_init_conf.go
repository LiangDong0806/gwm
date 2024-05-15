package global

import "github.com/hashicorp/consul/api"

var (
	RpcALLConf   *RpcAllClient
	NacosConf    *NacosConfig
	ConsulClient *api.Client
)

type RpcAllClient struct {
	GrpcIp *GrpcConfig   `yaml:"GrpcIp"`
	Consul *ConsulConfig `yaml:"Consul"`
	Mysql  *MysqlConfig  `yaml:"Mysql"`
	Redis  *RedisConfig  `yaml:"Redis"`
}
type RedisConfig struct {
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

type GrpcConfig struct {
	Host string `yaml:"Host"`
	Port int    `yaml:"Port"`
}

type ConsulConfig struct {
	Id   string   `yaml:"Id"`
	Name string   `yaml:"Name"`
	Tags []string `yaml:"Tags"`
}

type MysqlConfig struct {
	Username string `yaml:"Username"`
	Password string `yaml:"Password"`
	Host     string `yaml:"Host"`
	Port     int    `yaml:"Port"`
	Dbname   string `yaml:"Dbname"`
}
