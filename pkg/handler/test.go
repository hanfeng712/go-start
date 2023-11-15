package handler

import (
	"context"
	"gtstart/design/gen/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TestLogic struct {
	logx.Logger
}

func NewGreetLogic(logs logx.Logger) *TestLogic {
	return &TestLogic{
		Logger: logs,
	}
}

func (l *TestLogic) Greet(svcCtx context.Context, req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line
	return
}
