package service

import (
	"context"
	"errors"

	"github.com/xmhu2001/gomall/app/user/biz/dal/mysql"
	"github.com/xmhu2001/gomall/app/user/biz/model"
	user "github.com/xmhu2001/gomall/rpc_gen/kitex_gen/user"
	"golang.org/x/crypto/bcrypt"
)

type LoginService struct {
	ctx context.Context
} // NewLoginService new LoginService
func NewLoginService(ctx context.Context) *LoginService {
	return &LoginService{ctx: ctx}
}

// Run create note info
func (s *LoginService) Run(req *user.LoginReq) (resp *user.LoginResp, err error) {
	// Finish your business logic.
	if req.Email == "" || req.Password == "" {
		return nil, errors.New("email or password is empty")
	}
	u, err := model.GetByEmail(s.ctx, mysql.DB, req.Email)
	if err != nil {
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(u.PasswordHashed), []byte(req.Password))
	if err != nil {
		return nil, err
	}
	resp = &user.LoginResp{
		UserId: int32(u.ID),
	}
	return resp, nil
}
