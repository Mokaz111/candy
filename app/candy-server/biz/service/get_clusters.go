package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	v1 "github.com/mokaz111/candy-server/hertz_gen/api/v1"
)

type GetClustersService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGetClustersService(Context context.Context, RequestContext *app.RequestContext) *GetClustersService {
	return &GetClustersService{RequestContext: RequestContext, Context: Context}
}

func (h *GetClustersService) Run(req *v1.GetClustersRequest) (resp *v1.GetClustersResponse, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}
