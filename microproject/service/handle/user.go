package handle

import(
	"context"
	"strconv"

	"github.com/micro/go-micro/v2/logger"
	user "microproject/proto/user"
)

type User struct{

}
// 实现了user.pb.micro.go中的UserHandler接口
func (e *User) QueryUserByName(ctx context.Context, req *user.Request, rsp *user.Response) error {
	logger.Info("Received QueryUserByName request:", req.GetUserName())
	ID64,_:=strconv.ParseInt(req.UserID,10,64)
	rsp.User = &user.User{
		Id: ID64,
		Name: req.UserName,
		Pwd: req.UserPwd,
	}
	rsp.Success = true
	return nil
}