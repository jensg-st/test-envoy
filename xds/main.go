package main

import (
	"context"
	"fmt"
	"os"

	"github.com/envoyproxy/go-control-plane/pkg/cache/types"
	"github.com/envoyproxy/go-control-plane/pkg/cache/v3"

	"go.uber.org/zap"
)

var (
	endpoints     []types.Resource
	version       int
	snapshotCache cache.SnapshotCache
	// endpointInformers []k8scache.SharedIndexInformer
)

func main() {
	fmt.Println("starting")

	nodeID := "sss"

	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any
	sugar := logger.Sugar()

	cache := cache.NewSnapshotCache(false, cache.IDHash{}, sugar)

	snapshot := GenerateSnapshot()

	// Create the snapshot that we'll serve to Envoy
	// snapshot := example.GenerateSnapshot()
	if err := snapshot.Consistent(); err != nil {
		sugar.Errorf("snapshot inconsistency: %+v\n%+v", snapshot, err)
		os.Exit(1)
	}
	sugar.Debugf("will serve snapshot %+v", snapshot)

	// Add the snapshot to the cache
	if err := cache.SetSnapshot(context.Background(), nodeID, snapshot); err != nil {
		sugar.Errorf("snapshot error %q for %+v", err, snapshot)
		os.Exit(1)
	}

	ctx := context.Background()

	srv := NewServer(ctx, cache, cb)
	RunServer(srv, 12000)

	// cb := &test.Callbacks{Debug: l.Debug}
	// srv := server.NewServer(ctx, cache, cb)
	// example.RunServer(srv, port)

	// version = 0
	// snapshotCache = cache.NewSnapshotCache(false, cache.IDHash{}, nil)
	// server := xds.NewServer(context.Background(), snapshotCache, nil)
	// grpcServer := grpc.NewServer()
	// lis, _ := net.Listen("tcp4", ":12000")

	// discoverygrpc.RegisterAggregatedDiscoveryServiceServer(grpcServer, server)
	// endpointservice.RegisterEndpointDiscoveryServiceServer(grpcServer, server)
	// clusterservice.RegisterClusterDiscoveryServiceServer(grpcServer, server)
	// routeservice.RegisterRouteDiscoveryServiceServer(grpcServer, server)
	// listenerservice.RegisterListenerDiscoveryServiceServer(grpcServer, server)

	// // clusters, _ := CreateBootstrapClients()

	// // for _, cluster := range clusters {
	// // 	stop := make(chan struct{})
	// // 	defer close(stop)

	// // 	factory := informers.NewSharedInformerFactoryWithOptions(cluster, time.Second*10, informers.WithNamespace("demo"))
	// // 	informer := factory.Core().V1().Endpoints().Informer()
	// // 	endpointInformers = append(endpointInformers, informer)

	// // 	// informer.AddEventHandler(k8scache.ResourceEventHandlerFuncs{
	// // 	// 	UpdateFunc: HandleEndpointsUpdate,
	// // 	// })

	// // 	go func() {
	// // 		informer.Run(stop)
	// // 	}()
	// // }

	// if err := grpcServer.Serve(lis); err != nil {
	// 	fmt.Printf("%v", err)
	// }

}

func doit() {

}
