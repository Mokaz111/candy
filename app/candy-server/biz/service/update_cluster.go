package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	v1 "github.com/mokaz111/candy-server/hertz_gen/api/v1"
)

type UpdateClusterService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewUpdateClusterService(Context context.Context, RequestContext *app.RequestContext) *UpdateClusterService {
	return &UpdateClusterService{RequestContext: RequestContext, Context: Context}
}

func (h *UpdateClusterService) Run(req *v1.UpdateClusterRequest) (resp *v1.UpdateClusterResponse, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}
