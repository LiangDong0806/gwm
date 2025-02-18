package initialized

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/hashicorp/consul/api"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"gopkg.in/yaml.v2"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"zg5/Homework01/client/consul"
	"zg5/Homework01/client/global"
	"zg5/Homework01/client/routers"
	"zg5/Homework01/server/proto/server"
)

func init() {
	initViper()
	initNacos()
	initZapLog()
	initRouter()
	//utils.InitRedis()
}

func initViper() {
	v := viper.New()
	v.SetConfigFile("./conf/conf.yaml")
	err := v.ReadInConfig()
	if err != nil {
	}
	v.Unmarshal(&global.NacosConf)
}

func initZapLog() {
	log := zap.NewDevelopmentConfig()
	log.OutputPaths = []string{
		"./logger/message.log",
		"stderr",
		"stdout",
	}
	logs, _ := log.Build()
	zap.ReplaceGlobals(logs)
}

func initNacos() {
	// 创建clientConfig
	clientConfig := constant.ClientConfig{ //TODO: 11111
		NamespaceId:         global.NacosConf.NamespaceId, // 如果需要支持多namespace，我们可以场景多个client,它们有不同的NamespaceId。当namespace是public时，此处填空字符串。
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "/tmp/nacos/log",
		CacheDir:            "/tmp/nacos/cache",
		LogLevel:            "debug",
	}
	// 至少一个ServerConfig
	serverConfigs := []constant.ServerConfig{ //TODO: 22222
		{
			IpAddr:      global.NacosConf.Host,
			ContextPath: "/nacos",
			Port:        uint64(global.NacosConf.Port),
			Scheme:      "http",
		},
	}
	// 创建动态配置客户端的另一种方式 (推荐)
	configClient, err := clients.NewConfigClient( //TODO: 33333
		vo.NacosClientParam{
			ClientConfig:  &clientConfig,
			ServerConfigs: serverConfigs,
		},
	)
	if err != nil {
		panic(err)
	}
	content, err := configClient.GetConfig(vo.ConfigParam{ //TODO:44444
		DataId: global.NacosConf.DataId,
		Group:  global.NacosConf.Group,
	})
	yaml.Unmarshal([]byte(content), &global.ApiALLConf)
	zap.S().Info("client configuration：", &content)
}

func initConsul() consul.RegistryClient {
	global.ApiALLConf.Consul.Id = uuid.NewString()
	var err error
	global.ConsulClient, err = api.NewClient(api.DefaultConfig())
	if err != nil {
		panic(err)
	}
	consulClient := consul.NewConsulClient(global.ApiALLConf.ApiIp.Host, global.ApiALLConf.ApiIp.Port)
	car := consulClient.AgentHealthServiceByName(global.ApiALLConf.Consul.ConsulRpcName)
	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", car[0].Service.Address, car[0].Service.Port),
		grpc.WithBlock(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	global.ServerClient = server.NewServerClient(conn)
	err = consulClient.RegisterConsul(global.ApiALLConf.Consul.Id, global.ApiALLConf.Consul.Name, global.ApiALLConf.Consul.Tags)
	if err != nil {
		panic(err)
	}
	return consulClient
}

func InitGrpc() {
	var consulClient = initConsul()
	client := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", global.ApiALLConf.ApiIp.Host, global.ApiALLConf.ApiIp.Port),
		Handler: global.GinClient,
	}
	go func() {
		if err := client.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
			panic(err)
		}
	}()
	zap.S().Info("服务正在监听...", global.ApiALLConf.ApiIp.Host+":", global.ApiALLConf.ApiIp.Port)
	SignChan := make(chan os.Signal, 1)
	signal.Notify(SignChan, syscall.SIGINT, syscall.SIGTERM)
	<-SignChan
	err := consulClient.DeregisterConsulByID(global.ApiALLConf.Consul.Id)
	if err != nil {
		panic(err)
	}
}

func initRouter() {
	global.GinClient = gin.New()
	Group := global.GinClient.Group("vv2")
	routers.GinRouter(Group)
}
