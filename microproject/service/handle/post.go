package handle

import (
	"context"
	post "microproject/proto/post"
	"strconv"

	"github.com/micro/go-micro/v2/logger"
)
type Post struct{

}

//实现post.pb.micro.go中的PostHandler接口

func (e *Post) QueryUserPosts(ctx context.Context,req *post.Request, rsp *post.Response) error  {
	logger.Info("Received QueryUserPosts request:",req.GetPostID())
	ID64,_:=strconv.ParseInt(req.PostID,10,64)
	rsp.Post = &post.Post{
		Id: ID64,
		Title: "the title get by id",
	}
	rsp.Success=true
	rsp.Error = nil
	return nil
}