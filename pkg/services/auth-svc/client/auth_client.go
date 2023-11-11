package client

import (
	"fmt"
	"log"

	"github.com/ashiqsabith123/api-gateway/pkg/config"
	intefaces "github.com/ashiqsabith123/api-gateway/pkg/services/auth-svc/client/interface"

	"google.golang.org/grpc"
)

type AuthClient struct {
	config config.Config
}

var Conn *grpc.ClientConn
var err error

func NewAuthClient(config config.Config) intefaces.AuthClient {
	client := &AuthClient{config: config}
	client.InitAuthClient()
	return &AuthClient{}
}

func (A *AuthClient) InitAuthClient() {
	Conn, err = grpc.Dial(A.config.PORTS.AuthSvcPort, grpc.WithInsecure())
	if err != nil {
		log.Fatal("Could not connect the auth server:", err)
	}

	fmt.Println("Auth service connected at port ", A.config.PORTS.AuthSvcPort)

}

func (A *AuthClient) GetClient() pb.AuthServiceClient {
	return pb.NewAuthServiceClient(Conn)
}