package logic

import (
	"context"
	"math"
	"math/rand"
	"time"

	"greet/internal/svc"
	"greet/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoadLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoadLogic(ctx context.Context, svcCtx *svc.ServiceContext) LoadLogic {
	return LoadLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoadLogic) Load(req types.LoadRequest) (*types.LoadResponse, error) {
	rand.Seed(time.Now().UnixNano())
	count := 1000000

	for i:=0; i<count; i++ {
		num := rand.Intn(1000) + 1000000
		math.Sqrt(float64(num))
	}

	return &types.LoadResponse{
		Message: "ok",
	}, nil
}
