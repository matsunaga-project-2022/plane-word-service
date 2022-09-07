package controllers

import (
	"context"
	"github.com/matsunaga-project-2022/shanks/internal/domain/service"
	"github.com/matsunaga-project-2022/shanks/internal/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

type WordController struct {
	service *service.Word
}

func NewWordController(service *service.Word) *WordController {
	return &WordController{service: service}
}

func (w *WordController) CreateWord(ctx context.Context, request *proto.CreateWordRequest) (*emptypb.Empty, error) {
	err := w.service.CreateWord(ctx, request.GetWord(), request.GetMeaning(), request.GetUserID())
	if err != nil {
		return nil, err
	}

	res := &emptypb.Empty{}
	return res, nil
}

func (w *WordController) GetWord(ctx context.Context, request *proto.GetWordRequest) (*proto.GetWordResponse, error) {
	word, err := w.service.GetWordByID(ctx, request.GetId())
	if err != nil {
		return nil, err
	}
	resWord := &proto.Word{
		Id:      word.ID,
		UserID:  request.GetUserID(),
		Word:    word.WordName,
		Meaning: word.Meaning,
	}
	res := &proto.GetWordResponse{Word: resWord}

	return res, nil
}

func (w *WordController) ListWord(ctx context.Context, request *proto.ListWordRequest) (*proto.ListWordResponse, error) {
	words, err := w.service.ListWordByUserID(ctx, request.GetUserID())
	if err != nil {
		return nil, err
	}

	var resWords []*proto.Word
	for _, word := range words {
		resWord := &proto.Word{
			Id:      word.ID,
			UserID:  request.GetUserID(),
			Word:    word.WordName,
			Meaning: word.Meaning,
		}
		resWords = append(resWords, resWord)
	}

	res := &proto.ListWordResponse{
		Words: resWords,
	}

	return res, nil
}

func (w *WordController) UpdateWord(ctx context.Context, request *proto.UpdateRequest) (*emptypb.Empty, error) {
	err := w.service.UpdateWord(ctx, request.GetId(), request.GetMeaning())
	if err != nil {
		return nil, err
	}

	res := &emptypb.Empty{}
	return res, nil
}

func (w *WordController) DeleteWord(ctx context.Context, request *proto.DeleteRequest) (*emptypb.Empty, error) {
	err := w.service.DeleteWord(ctx, request.GetId())
	if err != nil {
		return nil, err
	}

	res := &emptypb.Empty{}
	return res, nil
}
