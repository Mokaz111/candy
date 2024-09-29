package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	v1 "github.com/mokaz111/candy-server/hertz_gen/api/v1"
)

type CreateClusterService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewCreateClusterService(Context context.Context, RequestContext *app.RequestContext) *CreateClusterService {
	return &CreateClusterService{RequestContext: RequestContext, Context: Context}
}

func (h *CreateClusterService) Run(req *v1.CreateClusterRequest) (resp *v1.CreateClusterResponse, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}
