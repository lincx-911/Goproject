package handler

import (
	"context"
	"encoding/json"
	post "microproject/proto/post"
	"strconv"
	"strings"

	api "github.com/micro/go-micro/api/proto"
	"github.com/micro/go-micro/errors"
	"github.com/micro/go-micro/v2/logger"
)

type Article struct {
	Client post.PostService
}

func (s *Article) GetArtivle(ctx context.Context, req *api.Request, rsp *api.Response) error {
	logger.Info("Received Article.GetArticle API request")
	ID, ok := req.Get["id"]
	if !ok || len(ID.Values) == 0 {
		return errors.BadRequest("go.micro.api.article", "id cannot be blank")
	}
	response, err := s.Client.QueryUserPosts(ctx, &post.Request{
		PostID: strings.Join(ID.Values, " "),
	})
	if err != nil {
		return err
	}
	rsp.StatusCode = 200
	b, _ := json.Marshal(map[string]string{
		"id":      strconv.FormatInt(response.Post.Id, 10),
		"title":   response.Post.Title,
		"content": response.Post.Content,
	})
	rsp.Body = string(b)
	return nil
}
