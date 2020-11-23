package transport

import (
	"context"

	"google.golang.org/grpc/metadata"
)

const (
	metadataKey = "x-request-id"
)

// fromContext returns a request ID from gRPC metadata if available in ctx.
func fromContext(ctx context.Context) (string, bool) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", false
	}

	ids, ok := md[metadataKey]
	if !ok || len(ids) == 0 {
		return "", false
	}

	return ids[0], true
}

// newMetadata constructs gRPC metadata with the request ID set.
func newMetadata(id string) metadata.MD {
	return metadata.Pairs(metadataKey, id)
}

// appendToOutgoingContext returns a context with the request-id added to the gRPC metadata.
func appendToOutgoingContext(ctx context.Context) context.Context {
	id, ok := fromContext(ctx)
	if !ok {
		return ctx
	}
	//generate id

	return metadata.AppendToOutgoingContext(ctx, metadataKey, id)
}
