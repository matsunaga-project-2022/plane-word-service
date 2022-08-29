package controllers

import (
	"context"
	"github.com/matsunaga-project-2022/shanks/internal/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

type WordController struct {
	// TODO: add service of word
	// service *service.Word
}

func (w *WordController) CreateWord(ctx context.Context, request *proto.CreateWordRequest) (*emptypb.Empty, error) {

	//TODO implement me
	panic("implement me")
}

func (w *WordController) ListWord(ctx context.Context, request *proto.ListWordRequest) (*proto.ListWordResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (w *WordController) UpdateWord(ctx context.Context, request *proto.Request) (*emptypb.Empty, error) {
	//TODO implement me
	panic("implement me")
}

func (w *WordController) DeleteWord(ctx context.Context, request *proto.Request) (*emptypb.Empty, error) {
	//TODO implement me
	panic("implement me")
}
