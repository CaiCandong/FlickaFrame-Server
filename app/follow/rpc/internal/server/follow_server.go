// Code generated by goctl. DO NOT EDIT.
// Source: follow.proto

package server

import (
	"context"

	"github.com/FlickaFrame/FlickaFrame-Server/app/follow/rpc/internal/logic"
	"github.com/FlickaFrame/FlickaFrame-Server/app/follow/rpc/internal/svc"
	"github.com/FlickaFrame/FlickaFrame-Server/app/follow/rpc/pb/pb"
)

type FollowServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedFollowServer
}

func NewFollowServer(svcCtx *svc.ServiceContext) *FollowServer {
	return &FollowServer{
		svcCtx: svcCtx,
	}
}

func (s *FollowServer) Follow(ctx context.Context, in *pb.FollowRequest) (*pb.FollowResponse, error) {
	l := logic.NewFollowLogic(ctx, s.svcCtx)
	return l.Follow(in)
}

func (s *FollowServer) UnFollow(ctx context.Context, in *pb.UnFollowRequest) (*pb.UnFollowResponse, error) {
	l := logic.NewUnFollowLogic(ctx, s.svcCtx)
	return l.UnFollow(in)
}

func (s *FollowServer) FollowList(ctx context.Context, in *pb.FollowListRequest) (*pb.FollowListResponse, error) {
	l := logic.NewFollowListLogic(ctx, s.svcCtx)
	return l.FollowList(in)
}

func (s *FollowServer) FansList(ctx context.Context, in *pb.FansListRequest) (*pb.FansListResponse, error) {
	l := logic.NewFansListLogic(ctx, s.svcCtx)
	return l.FansList(in)
}

func (s *FollowServer) IsFollow(ctx context.Context, in *pb.IsFollowReq) (*pb.IsFollowResp, error) {
	l := logic.NewIsFollowLogic(ctx, s.svcCtx)
	return l.IsFollow(in)
}