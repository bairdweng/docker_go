package main

import (
	"log"

	pb "com.miaoyou.server/myserver/grpc/model"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/consul/v2"

	"golang.org/x/net/context"
)

// 对象要和proto内定义的服务一致
type server struct {
}

// 实现rpc的 SayHi接口
func (s *server) SayHi(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReplay, error) {
	return &pb.HelloReplay{
		Message: "Hi " + in.Name,
	}, nil
}

// 实现rpc的 GetMsg接口
func (s *server) GetMsg(ctx context.Context, in *pb.HelloRequest) (*pb.HelloMessage, error) {
	return &pb.HelloMessage{
		Msg: "Server msg is coming...",
	}, nil
}

func registerConsul() *registry.Options {
	op := registry.Options{}
	op.Addrs = []string{
		"",
	}
	return &op
}

func main() {

	reg := consul.NewRegistry(registry.Addrs("127.0.0.1:8500"))

	service := micro.NewService(
		micro.Name("go.micro.service.helloworld"),
		micro.Version("latest"),
		micro.Registry(reg),
	)
	service.Init()

	micro.RegisterSubscriber("go.micro.service.helloworld", service.Server(), server{})
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
	println("服务注册成功")
	select {}
}
