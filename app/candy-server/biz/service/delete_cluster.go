package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	v1 "github.com/mokaz111/candy-server/hertz_gen/api/v1"
)

type DeleteClusterService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewDeleteClusterService(Context context.Context, RequestContext *app.RequestContext) *DeleteClusterService {
	return &DeleteClusterService{RequestContext: RequestContext, Context: Context}
}

func (h *DeleteClusterService) Run(req *v1.DeleteClusterRequest) (resp *v1.DeleteClusterResponse, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}
