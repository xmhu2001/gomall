package service

import (
	"context"
	"errors"

	"github.com/xmhu2001/gomall/app/user/biz/dal/mysql"
	"github.com/xmhu2001/gomall/app/user/biz/model"
	user "github.com/xmhu2001/gomall/rpc_gen/kitex_gen/user"
	"golang.org/x/crypto/bcrypt"
)

type RegisterService struct {
	ctx context.Context
} // NewRegisterService new RegisterService
func NewRegisterService(ctx context.Context) *RegisterService {
	return &RegisterService{ctx: ctx}
}

// Run create note info
func (s *RegisterService) Run(req *user.RegisterReq) (resp *user.RegisterResp, err error) {
	// Finish your business logic.
	// 参数校验
	if req.Password != req.PasswordConfirm {
		return nil, errors.New("password and password_confirm not match")
	}
	bcryptPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	newUser := &model.User{
		Email:          req.Email,
		PasswordHashed: string(bcryptPassword),
	}
	err = model.Create(mysql.DB, newUser)
	if err != nil {
		return nil, err
	}
	return &user.RegisterResp{
		UserId: int32(newUser.ID),
	}, nil
}
