// THIS FILE IS AUTO GENERATED BY GK-CLI DO NOT EDIT!!
package grpc

import (
	grpc "github.com/go-kit/kit/transport/grpc"
	endpoint "github.com/moecasts/microcasts/novels/pkg/endpoint"
	pb "github.com/moecasts/microcasts/novels/pkg/grpc/pb"
)

// NewGRPCServer makes a set of endpoints available as a gRPC AddServer
type grpcServer struct {
	browse       grpc.Handler
	read         grpc.Handler
	add          grpc.Handler
	edit         grpc.Handler
	trash        grpc.Handler
	restore      grpc.Handler
	destroy      grpc.Handler
	batchAdd     grpc.Handler
	batchEdit    grpc.Handler
	batchTrash   grpc.Handler
	batchRestore grpc.Handler
	batchDestroy grpc.Handler
}

func NewGRPCServer(endpoints endpoint.Endpoints, options map[string][]grpc.ServerOption) pb.NovelsServer {
	return &grpcServer{
		add:          makeAddHandler(endpoints, options["Add"]),
		batchAdd:     makeBatchAddHandler(endpoints, options["BatchAdd"]),
		batchDestroy: makeBatchDestroyHandler(endpoints, options["BatchDestroy"]),
		batchEdit:    makeBatchEditHandler(endpoints, options["BatchEdit"]),
		batchRestore: makeBatchRestoreHandler(endpoints, options["BatchRestore"]),
		batchTrash:   makeBatchTrashHandler(endpoints, options["BatchTrash"]),
		browse:       makeBrowseHandler(endpoints, options["Browse"]),
		destroy:      makeDestroyHandler(endpoints, options["Destroy"]),
		edit:         makeEditHandler(endpoints, options["Edit"]),
		read:         makeReadHandler(endpoints, options["Read"]),
		restore:      makeRestoreHandler(endpoints, options["Restore"]),
		trash:        makeTrashHandler(endpoints, options["Trash"]),
	}
}