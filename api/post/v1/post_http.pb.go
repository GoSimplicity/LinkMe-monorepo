// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.7.3
// - protoc             v5.27.1
// source: api/post/v1/post.proto

package post

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationPostCollectPost = "/api.post.Post/CollectPost"
const OperationPostCreatePost = "/api.post.Post/CreatePost"
const OperationPostDeletePost = "/api.post.Post/DeletePost"
const OperationPostDetailAdminPost = "/api.post.Post/DetailAdminPost"
const OperationPostDetailPost = "/api.post.Post/DetailPost"
const OperationPostDetailPubPost = "/api.post.Post/DetailPubPost"
const OperationPostGetPostStats = "/api.post.Post/GetPostStats"
const OperationPostLikePost = "/api.post.Post/LikePost"
const OperationPostListAdminPost = "/api.post.Post/ListAdminPost"
const OperationPostListPost = "/api.post.Post/ListPost"
const OperationPostListPubPost = "/api.post.Post/ListPubPost"
const OperationPostPublishPost = "/api.post.Post/PublishPost"
const OperationPostUpdatePost = "/api.post.Post/UpdatePost"
const OperationPostWithdrawPost = "/api.post.Post/WithdrawPost"

type PostHTTPServer interface {
	CollectPost(context.Context, *CollectPostRequest) (*CollectPostReply, error)
	CreatePost(context.Context, *CreatePostRequest) (*CreatePostReply, error)
	DeletePost(context.Context, *DeletePostRequest) (*DeletePostReply, error)
	DetailAdminPost(context.Context, *DetailAdminPostRequest) (*DetailAdminPostReply, error)
	DetailPost(context.Context, *DetailPostRequest) (*DetailPostReply, error)
	DetailPubPost(context.Context, *DetailPubPostRequest) (*DetailPubPostReply, error)
	GetPostStats(context.Context, *emptypb.Empty) (*GetPostStatsReply, error)
	LikePost(context.Context, *LikePostRequest) (*LikePostReply, error)
	ListAdminPost(context.Context, *ListAdminPostRequest) (*ListAdminPostReply, error)
	ListPost(context.Context, *ListPostRequest) (*ListPostReply, error)
	ListPubPost(context.Context, *ListPubPostRequest) (*ListPubPostReply, error)
	PublishPost(context.Context, *PublishPostRequest) (*PublishPostReply, error)
	UpdatePost(context.Context, *UpdatePostRequest) (*UpdatePostReply, error)
	WithdrawPost(context.Context, *WithdrawPostRequest) (*WithdrawPostReply, error)
}

func RegisterPostHTTPServer(s *http.Server, srv PostHTTPServer) {
	r := s.Route("/")
	r.POST("/create", _Post_CreatePost0_HTTP_Handler(srv))
	r.POST("/update", _Post_UpdatePost0_HTTP_Handler(srv))
	r.DELETE("/delete/{postId}", _Post_DeletePost0_HTTP_Handler(srv))
	r.POST("/publish", _Post_PublishPost0_HTTP_Handler(srv))
	r.POST("/withdraw", _Post_WithdrawPost0_HTTP_Handler(srv))
	r.POST("/list", _Post_ListPost0_HTTP_Handler(srv))
	r.POST("/list_pub", _Post_ListPubPost0_HTTP_Handler(srv))
	r.POST("/list_admin", _Post_ListAdminPost0_HTTP_Handler(srv))
	r.GET("/detail/{postId}", _Post_DetailPost0_HTTP_Handler(srv))
	r.GET("/detail_pub/{postId}", _Post_DetailPubPost0_HTTP_Handler(srv))
	r.GET("/detail_admin/{postId}", _Post_DetailAdminPost0_HTTP_Handler(srv))
	r.GET("/stats", _Post_GetPostStats0_HTTP_Handler(srv))
	r.POST("/like", _Post_LikePost0_HTTP_Handler(srv))
	r.POST("/collect", _Post_CollectPost0_HTTP_Handler(srv))
}

func _Post_CreatePost0_HTTP_Handler(srv PostHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreatePostRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationPostCreatePost)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreatePost(ctx, req.(*CreatePostRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreatePostReply)
		return ctx.Result(200, reply)
	}
}

func _Post_UpdatePost0_HTTP_Handler(srv PostHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdatePostRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationPostUpdatePost)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdatePost(ctx, req.(*UpdatePostRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdatePostReply)
		return ctx.Result(200, reply)
	}
}

func _Post_DeletePost0_HTTP_Handler(srv PostHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeletePostRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationPostDeletePost)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeletePost(ctx, req.(*DeletePostRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DeletePostReply)
		return ctx.Result(200, reply)
	}
}

func _Post_PublishPost0_HTTP_Handler(srv PostHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in PublishPostRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationPostPublishPost)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.PublishPost(ctx, req.(*PublishPostRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*PublishPostReply)
		return ctx.Result(200, reply)
	}
}

func _Post_WithdrawPost0_HTTP_Handler(srv PostHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in WithdrawPostRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationPostWithdrawPost)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.WithdrawPost(ctx, req.(*WithdrawPostRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*WithdrawPostReply)
		return ctx.Result(200, reply)
	}
}

func _Post_ListPost0_HTTP_Handler(srv PostHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListPostRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationPostListPost)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListPost(ctx, req.(*ListPostRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListPostReply)
		return ctx.Result(200, reply)
	}
}

func _Post_ListPubPost0_HTTP_Handler(srv PostHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListPubPostRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationPostListPubPost)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListPubPost(ctx, req.(*ListPubPostRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListPubPostReply)
		return ctx.Result(200, reply)
	}
}

func _Post_ListAdminPost0_HTTP_Handler(srv PostHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListAdminPostRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationPostListAdminPost)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListAdminPost(ctx, req.(*ListAdminPostRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListAdminPostReply)
		return ctx.Result(200, reply)
	}
}

func _Post_DetailPost0_HTTP_Handler(srv PostHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DetailPostRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationPostDetailPost)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DetailPost(ctx, req.(*DetailPostRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DetailPostReply)
		return ctx.Result(200, reply)
	}
}

func _Post_DetailPubPost0_HTTP_Handler(srv PostHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DetailPubPostRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationPostDetailPubPost)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DetailPubPost(ctx, req.(*DetailPubPostRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DetailPubPostReply)
		return ctx.Result(200, reply)
	}
}

func _Post_DetailAdminPost0_HTTP_Handler(srv PostHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DetailAdminPostRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationPostDetailAdminPost)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DetailAdminPost(ctx, req.(*DetailAdminPostRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DetailAdminPostReply)
		return ctx.Result(200, reply)
	}
}

func _Post_GetPostStats0_HTTP_Handler(srv PostHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in emptypb.Empty
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationPostGetPostStats)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetPostStats(ctx, req.(*emptypb.Empty))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetPostStatsReply)
		return ctx.Result(200, reply)
	}
}

func _Post_LikePost0_HTTP_Handler(srv PostHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in LikePostRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationPostLikePost)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.LikePost(ctx, req.(*LikePostRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*LikePostReply)
		return ctx.Result(200, reply)
	}
}

func _Post_CollectPost0_HTTP_Handler(srv PostHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CollectPostRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationPostCollectPost)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CollectPost(ctx, req.(*CollectPostRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CollectPostReply)
		return ctx.Result(200, reply)
	}
}

type PostHTTPClient interface {
	CollectPost(ctx context.Context, req *CollectPostRequest, opts ...http.CallOption) (rsp *CollectPostReply, err error)
	CreatePost(ctx context.Context, req *CreatePostRequest, opts ...http.CallOption) (rsp *CreatePostReply, err error)
	DeletePost(ctx context.Context, req *DeletePostRequest, opts ...http.CallOption) (rsp *DeletePostReply, err error)
	DetailAdminPost(ctx context.Context, req *DetailAdminPostRequest, opts ...http.CallOption) (rsp *DetailAdminPostReply, err error)
	DetailPost(ctx context.Context, req *DetailPostRequest, opts ...http.CallOption) (rsp *DetailPostReply, err error)
	DetailPubPost(ctx context.Context, req *DetailPubPostRequest, opts ...http.CallOption) (rsp *DetailPubPostReply, err error)
	GetPostStats(ctx context.Context, req *emptypb.Empty, opts ...http.CallOption) (rsp *GetPostStatsReply, err error)
	LikePost(ctx context.Context, req *LikePostRequest, opts ...http.CallOption) (rsp *LikePostReply, err error)
	ListAdminPost(ctx context.Context, req *ListAdminPostRequest, opts ...http.CallOption) (rsp *ListAdminPostReply, err error)
	ListPost(ctx context.Context, req *ListPostRequest, opts ...http.CallOption) (rsp *ListPostReply, err error)
	ListPubPost(ctx context.Context, req *ListPubPostRequest, opts ...http.CallOption) (rsp *ListPubPostReply, err error)
	PublishPost(ctx context.Context, req *PublishPostRequest, opts ...http.CallOption) (rsp *PublishPostReply, err error)
	UpdatePost(ctx context.Context, req *UpdatePostRequest, opts ...http.CallOption) (rsp *UpdatePostReply, err error)
	WithdrawPost(ctx context.Context, req *WithdrawPostRequest, opts ...http.CallOption) (rsp *WithdrawPostReply, err error)
}

type PostHTTPClientImpl struct {
	cc *http.Client
}

func NewPostHTTPClient(client *http.Client) PostHTTPClient {
	return &PostHTTPClientImpl{client}
}

func (c *PostHTTPClientImpl) CollectPost(ctx context.Context, in *CollectPostRequest, opts ...http.CallOption) (*CollectPostReply, error) {
	var out CollectPostReply
	pattern := "/collect"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationPostCollectPost))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *PostHTTPClientImpl) CreatePost(ctx context.Context, in *CreatePostRequest, opts ...http.CallOption) (*CreatePostReply, error) {
	var out CreatePostReply
	pattern := "/create"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationPostCreatePost))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *PostHTTPClientImpl) DeletePost(ctx context.Context, in *DeletePostRequest, opts ...http.CallOption) (*DeletePostReply, error) {
	var out DeletePostReply
	pattern := "/delete/{postId}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationPostDeletePost))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *PostHTTPClientImpl) DetailAdminPost(ctx context.Context, in *DetailAdminPostRequest, opts ...http.CallOption) (*DetailAdminPostReply, error) {
	var out DetailAdminPostReply
	pattern := "/detail_admin/{postId}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationPostDetailAdminPost))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *PostHTTPClientImpl) DetailPost(ctx context.Context, in *DetailPostRequest, opts ...http.CallOption) (*DetailPostReply, error) {
	var out DetailPostReply
	pattern := "/detail/{postId}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationPostDetailPost))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *PostHTTPClientImpl) DetailPubPost(ctx context.Context, in *DetailPubPostRequest, opts ...http.CallOption) (*DetailPubPostReply, error) {
	var out DetailPubPostReply
	pattern := "/detail_pub/{postId}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationPostDetailPubPost))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *PostHTTPClientImpl) GetPostStats(ctx context.Context, in *emptypb.Empty, opts ...http.CallOption) (*GetPostStatsReply, error) {
	var out GetPostStatsReply
	pattern := "/stats"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationPostGetPostStats))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *PostHTTPClientImpl) LikePost(ctx context.Context, in *LikePostRequest, opts ...http.CallOption) (*LikePostReply, error) {
	var out LikePostReply
	pattern := "/like"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationPostLikePost))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *PostHTTPClientImpl) ListAdminPost(ctx context.Context, in *ListAdminPostRequest, opts ...http.CallOption) (*ListAdminPostReply, error) {
	var out ListAdminPostReply
	pattern := "/list_admin"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationPostListAdminPost))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *PostHTTPClientImpl) ListPost(ctx context.Context, in *ListPostRequest, opts ...http.CallOption) (*ListPostReply, error) {
	var out ListPostReply
	pattern := "/list"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationPostListPost))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *PostHTTPClientImpl) ListPubPost(ctx context.Context, in *ListPubPostRequest, opts ...http.CallOption) (*ListPubPostReply, error) {
	var out ListPubPostReply
	pattern := "/list_pub"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationPostListPubPost))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *PostHTTPClientImpl) PublishPost(ctx context.Context, in *PublishPostRequest, opts ...http.CallOption) (*PublishPostReply, error) {
	var out PublishPostReply
	pattern := "/publish"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationPostPublishPost))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *PostHTTPClientImpl) UpdatePost(ctx context.Context, in *UpdatePostRequest, opts ...http.CallOption) (*UpdatePostReply, error) {
	var out UpdatePostReply
	pattern := "/update"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationPostUpdatePost))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *PostHTTPClientImpl) WithdrawPost(ctx context.Context, in *WithdrawPostRequest, opts ...http.CallOption) (*WithdrawPostReply, error) {
	var out WithdrawPostReply
	pattern := "/withdraw"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationPostWithdrawPost))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
