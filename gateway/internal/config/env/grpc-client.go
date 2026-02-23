package env

import (
	"net"

	"github.com/caarlos0/env/v11"
)

type grpcClientBoardEnvConfig struct {
	Host string `env:"BOARD_GRPC_HOST,required"`
	Port string `env:"BOARD_GRPC_PORT,required"`
}

type grpcClientAuthEnvConfig struct {
	Host string `env:"AUTH_GRPC_HOST,required"`
	Port string `env:"AUTH_GRPC_PORT,required"`
}

type GrpcClient struct {
	rawBoard grpcClientBoardEnvConfig
	rawAuth  grpcClientAuthEnvConfig
}

func NewGrpcClient() (*GrpcClient, error) {
	var rawBoard grpcClientBoardEnvConfig

	var rawAuth grpcClientAuthEnvConfig

	if err := env.Parse(&rawBoard); err != nil {
		return nil, err
	}

	if err := env.Parse(&rawAuth); err != nil {
		return nil, err
	}

	return &GrpcClient{
		rawBoard: rawBoard,
		rawAuth:  rawAuth,
	}, nil
}

func (cfg *GrpcClient) BoardAddress() string {
	return net.JoinHostPort(cfg.rawBoard.Host, cfg.rawBoard.Port)
}

func (cfg *GrpcClient) AuthAddress() string {
	return net.JoinHostPort(cfg.rawAuth.Host, cfg.rawAuth.Port)
}
