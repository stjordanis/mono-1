package svc

import (
	"context"

	v0proto "github.com/owncloud/mono/thumbnails/pkg/proto/v0"
	"go.opencensus.io/trace"
)

// NewTracing returns a service that instruments traces.
func NewTracing(next v0proto.ThumbnailServiceHandler) v0proto.ThumbnailServiceHandler {
	return tracing{
		next: next,
	}
}

type tracing struct {
	next v0proto.ThumbnailServiceHandler
}

// GetThumbnail implements the ThumbnailServiceHandler interface.
func (t tracing) GetThumbnail(ctx context.Context, req *v0proto.GetRequest, rsp *v0proto.GetResponse) error {
	ctx, span := trace.StartSpan(ctx, "Thumbnails.GetThumbnail")
	defer span.End()

	span.Annotate([]trace.Attribute{
		trace.StringAttribute("filepath", req.Filepath),
		trace.StringAttribute("filetype", req.Filetype.String()),
		trace.StringAttribute("etag", req.Etag),
		trace.Int64Attribute("width", int64(req.Width)),
		trace.Int64Attribute("height", int64(req.Height)),
	}, "Execute Thumbnails.GetThumbnail handler")

	return t.next.GetThumbnail(ctx, req, rsp)
}
