package service

import (
	"context"
	"errors"

	pb "github.com/sunjin110/gRPC/server"
)

// MyCatService is server instance?
type MyCatService struct {
}

// GetMyCat is 猫の名前と種類を返す
func (s *MyCatService) GetMyCat(ctx context.Context, message *pb.GetMyCatMessage) (*pb.MyCatResponse, error) {
	switch message.TargetCat {
	case "tama":
		return &pb.MyCatResponse{
			Name: "tama",
			Kind: "mainecoon",
		}, nil
	case "mike":
		return &pb.MyCatResponse{
			Name: "mike",
			Kind: "Norwegian Forest Cat",
		}, nil
	default:
		return nil, errors.New("Not Found Your Cat")
	}
}
