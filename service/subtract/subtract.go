package subtract

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	obs "gitlab.com/grpasr/common/observability"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"log"
	"net"
	"os"
	"strconv"
	pb "testobservability/service/proto/v1/operand"
)

const (
	jaegerEndpoint    string  = "http://127.0.0.1:14268/api/traces"
	serviceName       string  = "substract"
	environment       string  = "development"
	id                        = 2
	collectorEndpoint         = "localhost:4317"
	samplingRatio     float64 = 0.6
	scratchDelay      int     = 30
)

func Start() {

	obs.SetObservabilityFacade(serviceName)

	obs.Logging.SetLoggingEnvToDevelopment()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	tls, err := getTls()
	if err != nil {
		fmt.Println("err from tls: ", err)
	}

	{
		tp, err := obs.Tracing.SetupTracing(
			ctx,
			tls,
			samplingRatio,
			serviceName,
			collectorEndpoint,
			environment)
		if err != nil {
			panic(err)
		}
		defer tp.Shutdown(ctx)

		mp, err := obs.Metrics.SetupMetrics(
			ctx,
			tls,
			scratchDelay,
			serviceName,
			collectorEndpoint,
			environment)
		if err != nil {
			panic(err)
		}
		defer mp.Shutdown(ctx)
	}

	// defer func(ctx context.Context) {
	// 	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	// 	defer cancel()
	// 	if err := tp.Shutdown(ctx); err != nil {
	// 		log.Fatal(err)
	// 	}
	// }(ctx)

	// set the interceptor
	serverOptions := []grpc.ServerOption{
		grpc.UnaryInterceptor(traceInterceptor),
	}

	osvc, err := NewOperandServer(serverOptions...)
	if err != nil {
		log.Fatal(err)
	}

	var grpcPort = ":50001"
	lis, err := net.Listen("tcp", grpcPort)
	if err != nil {
		log.Fatal("err creating the listener: ", err)
	}

	log.Println("GRPC Server is listening on port: ", grpcPort)
	err = osvc.Serve(lis)
	if err != nil {
		log.Println("err server listen: ", err)
	}
}

type OperandServer struct {
	pb.UnimplementedOperandManagementServer
}

// func NewOperandServer(opt ...grpc.ServerOption) (*grpc.Server, error) {
func NewOperandServer(opt ...grpc.ServerOption) (*grpc.Server, error) {
	// gsrv := grpc.NewServer(opt...)
	gsrv := grpc.NewServer(opt...)

	osrv := &OperandServer{}

	pb.RegisterOperandManagementServer(gsrv, osrv)

	return gsrv, nil
}

var res = float32(0)

func (o *OperandServer) SendOperand(ctx context.Context, dt *pb.Data) (*wrapperspb.StringValue, error) {
	// Access the operand using request.Operand
	operandValue := dt.Operand

	ctx, span := obs.Tracing.SPNGetFromCTX(
		ctx,
		"subtract-server_SendOperand",
		obs.Tracing.TAString("component", "subtraction"),
		obs.Tracing.TAString("somekey", "somevalue"),
	)
	defer span.End()

	for _, n := range operandValue {
		// fmt.Println("see res0: ", res)
		// if i == 0 {
		// 	fmt.Println("see n: ", n)
		// 	res += n
		// }
		// fmt.Println("see n2: ", n)
		// fmt.Println("see res: ", res)
		res -= n
	}

	// setted on the same span as is trace.WithAttributes(attribute.String(...))
	obs.Tracing.SPNSetAttributes(
		span,
		obs.Tracing.TAFloat64("subtract-res", float64(res)))

	strValue := strconv.FormatFloat(float64(res), 'f', -1, 32)

	return &wrapperspb.StringValue{Value: strValue}, nil

}

// grpc jaegger staff --------------------------------
func traceInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	// from client to server

	method := info.FullMethod
	log.Println("Received gRPC request for method:", method)

	gh := obs.Tracing.NewGrpcTracingHandler()
	// Extract the span context from the gRPC metadata
	ok := gh.MetadataExtractor(ctx)
	// md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		// propagator := propagation.TraceContext{}
		ctx, _ = gh.ContextExtractor(ctx)
		// ctx = propagator.Extract(ctx, metadataCarrier(md))
		// ctx = propagator.Extract(ctx, propagation.NewCarrier(md))
	}

	// Call the gRPC handler with the modified context
	resp, err = handler(ctx, req)

	// response to the client

	// // Invoke the gRPC method
	// ctxSp := trace.ContextWithSpan(ctx, span)
	// ctx = metadata.NewOutgoingContext(ctxSp, md)

	return resp, err
}

// getTls returns a configuration that enables the use of mutual TLS.
func getTls() (*tls.Config, error) {
	clientAuth, err := tls.LoadX509KeyPair("../../confs/client.crt", "../../confs/client.key")
	if err != nil {
		return nil, err
	}

	caCert, err := os.ReadFile("../../confs/rootCA.crt")
	if err != nil {
		return nil, err
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	c := &tls.Config{
		RootCAs:      caCertPool,
		Certificates: []tls.Certificate{clientAuth},
	}

	return c, nil
}
