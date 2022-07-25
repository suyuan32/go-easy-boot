package user

import (
	"context"
	"fmt"
	"net/http"
	"system/api/common/errorx"

	"system/api/internal/svc"
	"system/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserPermCodeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserPermCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserPermCodeLogic {
	return &GetUserPermCodeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserPermCodeLogic) GetUserPermCode() (resp *types.PermCodeResp, err error) {
	roleId := l.ctx.Value("roleId")
	fmt.Println(roleId)
	if roleId == nil {
		return nil, &errorx.ApiError{
			Code: http.StatusUnauthorized,
			Msg:  "sys.login.requireLogin",
		}
	}
	return &types.PermCodeResp{Data: []string{fmt.Sprintf("%v", roleId)}}, nil
}
