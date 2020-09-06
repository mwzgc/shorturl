package logic

import (
	"context"
	"net"

	"shorturl/rpc/transform/internal/svc"
	"shorturl/rpc/transform/model"
	transform "shorturl/rpc/transform/pb"

	"github.com/tal-tech/go-zero/core/hash"
	"github.com/tal-tech/go-zero/core/logx"
)

type ShortenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewShortenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShortenLogic {
	return &ShortenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ShortenLogic) Shorten(in *transform.ShortenReq) (*transform.ShortenResp, error) {
	// todo: add your logic here and delete this line
	key := hash.Md5Hex([]byte(in.Url))[:6]

	row, _ := l.svcCtx.Model.FindOne(key)

	if row == nil {
		_, err := l.svcCtx.Model.Insert(model.Shorturl{
			Shorten: key,
			Url:     in.Url,
		})
		if err != nil {
			return nil, err
		}
	}

	return &transform.ShortenResp{
		Shorten: key + ", ip:" + getIP(),
	}, nil
	// return &transform.ShortenResp{}, nil
}

func getIP() string {
	addrs, err := net.InterfaceAddrs()

	if err != nil {
		return ""
	}

	for _, address := range addrs {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}

		}
	}

	return ""
}
