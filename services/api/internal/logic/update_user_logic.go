package logic

import (
	"context"

	"gorm-zero-example/services/api/internal/svc"
	"gorm-zero-example/services/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) UpdateUserLogic {
	return UpdateUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateUserLogic) UpdateUser(req types.UpdateUserReq) error {
	// todo: add your logic here and delete this line

	u, err := l.svcCtx.UserModel.FindOne(int64(req.Id))
	if err != nil {
		return err
	}
	u.NickName = req.NickName

	err = l.svcCtx.UserModel.Update(u)

	return err
}
