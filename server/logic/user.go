package logic

import (
	"encoding/json"
	"errors"
	"golang.org/x/net/context"
	"log"
	"net/http"
	"zg5/Homework01/server/proto/server"
	service2 "zg5/Homework01/service"
)

type RpcServers struct {
	server.UnimplementedServerServer
}

func (c *RpcServers) Product(ctx context.Context, in *server.ProductRequest) (*server.Response, error) {
	res, err := service2.PreheatTheProduct(in.Name)
	log.Println(res, "dlsdksak")
	if err != nil {
		return &server.Response{}, errors.New("product request failed")
	}

	var productList []*server.ProductList

	for _, p := range *res {
		product := &server.ProductList{
			Name:        p.Name,
			Description: p.Description,
			Price:       float32(p.Price),
			Stock:       int64(p.Stock),
		}
		productList = append(productList, product)
	}

	return &server.Response{
		ProductList: productList,
	}, nil

}

func (c *RpcServers) ProductAdd(ctx context.Context, in *server.ProductAddToRequest) (*server.ProductAddToResponse, error) {
	product := service2.Product{
		Name:        in.Name,
		Description: in.Description,
		Price:       float64(in.Price),
		Stock:       int(in.Stock),
		Category:    in.Category,
	}

	_, err := service2.ProductAddTo(product)
	if err != nil {
		return &server.ProductAddToResponse{
			Code: http.StatusAccepted,
			Msg:  "mysql" + err.Error(),
		}, nil
	}
	p, _ := json.Marshal(product)

	err = service2.RedisDBHSet("商品列表:", p)
	if err != nil {
		return &server.ProductAddToResponse{
			Code: http.StatusAccepted,
			Msg:  "redis" + err.Error(),
		}, nil
	}

	return &server.ProductAddToResponse{
		Code: http.StatusOK,
		Msg:  "购物车添加成功",
	}, nil
}

func (c *RpcServers) Register(ctx context.Context, in *server.RegisterRequest) (*server.RegisterResponse, error) {
	//fmt.Println(in.Username)
	//user, err := utils.QueryTheUser(in.Username)
	//if err != nil {
	//	fmt.Print("22222", err.Error())
	//	return nil, err
	//}
	//fmt.Print("1111111111", user)
	//if user.Id != 0 {
	//	return &server.RegisterResponse{Data: "账号已存在"}, nil
	//}
	//pwd, err := utils.EncryptPasswords([]byte(in.Password))
	//if err != nil {
	//	return &server.RegisterResponse{Data: "密码加密出错"}, nil
	//}
	//_, err = utils.UserRegistration(utils.User{
	//	Username: in.Username,
	//	Password: string(pwd),
	//	Mobile:   in.Mobile,
	//})
	//if err != nil {
	//	return &server.RegisterResponse{Data: "注册失败"}, nil
	//}
	////token, _ := common.SetJwtToken(common.GAOWEIMING, time.Now().Unix(), 3600, strconv.Itoa(user.Id))
	//
	return &server.RegisterResponse{Data: "注册成功"}, nil
	//return nil, nil
}
