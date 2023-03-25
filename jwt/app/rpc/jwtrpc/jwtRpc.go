// Code generated by goctl. DO NOT EDIT.
// Source: jwt.proto

package jwtrpc

import (
	Jwt2 "Mini-Tiktok/jwt/app/rpc/Jwt"
	"context"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	CreateTokenReq   = Jwt2.CreateTokenReq
	CreateTokenResp  = Jwt2.CreateTokenResp
	IsValidTokenReq  = Jwt2.IsValidTokenReq
	IsValidTokenResp = Jwt2.IsValidTokenResp
	ParseTokenReq    = Jwt2.ParseTokenReq
	ParseTokenResp   = Jwt2.ParseTokenResp

	JwtRpc interface {
		CreateToken(ctx context.Context, in *CreateTokenReq, opts ...grpc.CallOption) (*CreateTokenResp, error)
		ParseToken(ctx context.Context, in *ParseTokenReq, opts ...grpc.CallOption) (*ParseTokenResp, error)
		IsValidToken(ctx context.Context, in *IsValidTokenReq, opts ...grpc.CallOption) (*IsValidTokenResp, error)
	}

	defaultJwtRpc struct {
		cli zrpc.Client
	}
)

func NewJwtRpc(cli zrpc.Client) JwtRpc {
	return &defaultJwtRpc{
		cli: cli,
	}
}

func (m *defaultJwtRpc) CreateToken(ctx context.Context, in *CreateTokenReq, opts ...grpc.CallOption) (*CreateTokenResp, error) {
	client := Jwt2.NewJwtRpcClient(m.cli.Conn())
	return client.CreateToken(ctx, in, opts...)
}

func (m *defaultJwtRpc) ParseToken(ctx context.Context, in *ParseTokenReq, opts ...grpc.CallOption) (*ParseTokenResp, error) {
	client := Jwt2.NewJwtRpcClient(m.cli.Conn())
	return client.ParseToken(ctx, in, opts...)
}

func (m *defaultJwtRpc) IsValidToken(ctx context.Context, in *IsValidTokenReq, opts ...grpc.CallOption) (*IsValidTokenResp, error) {
	client := Jwt2.NewJwtRpcClient(m.cli.Conn())
	return client.IsValidToken(ctx, in, opts...)
}
