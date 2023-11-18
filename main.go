// package main

// import (
// 	"context"
// 	"crypto/tls"
// 	"crypto/x509"
// 	"fmt"
// 	obs "gitlab.com/grpasr/common/observability"
// 	logs "gitlab.com/grpasr/common/observability/logging"
// 	"net/http"
// 	"os"
// 	// added
// 	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
// 	"go.opentelemetry.io/otel"
// 	"go.opentelemetry.io/otel/attribute"
// 	"go.opentelemetry.io/otel/metric"
// 	"go.opentelemetry.io/otel/metric/global"
// 	"go.opentelemetry.io/otel/metric/instrument"
// 	"go.opentelemetry.io/otel/trace"
// )
//
// // curl -I http://127.0.0.1:8081/serviceA
//
// // package to add
// // go get go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp@v0.39.0
// // go get go.opentelemetry.io/otel@v1.13.0
// // go get go.opentelemetry.io/otel/metric@v0.36.0
// // go get go.opentelemetry.io/otel/trace@v1.13.0
//
// const serviceName = "svcName"
//
// func main() {
// 	ctx := context.Background()
//
// 	tls, err := getTls()
// 	if err != nil {
// 		fmt.Println("err from tls: ", err)
// 	}
//
// 	{
// 		tp, err := obs.SetupTracing(ctx, serviceName, tls)
// 		if err != nil {
// 			panic(err)
// 		}
// 		defer tp.Shutdown(ctx)
//
// 		mp, err := obs.SetupMetrics(ctx, serviceName, tls)
// 		if err != nil {
// 			panic(err)
// 		}
// 		defer mp.Shutdown(ctx)
// 	}
//
// 	go serviceA(ctx, 8081)
// 	serviceB(ctx, 8082)
// }
//
// // curl -vkL http://127.0.0.1:8081/serviceA
// func serviceA(ctx context.Context, port int) {
// 	mux := http.NewServeMux()
// 	mux.HandleFunc("/serviceA", serviceA_HttpHandler)
// 	handler := otelhttp.NewHandler(mux, "server.http")
// 	serverPort := fmt.Sprintf(":%d", port)
// 	server := &http.Server{Addr: serverPort, Handler: handler}
//
// 	fmt.Println("serviceA listening on", server.Addr)
// 	if err := server.ListenAndServe(); err != nil {
// 		panic(err)
// 	}
// }
//
// func serviceA_HttpHandler(w http.ResponseWriter, r *http.Request) {
// 	ctx, span := otel.Tracer("myTracer").Start(r.Context(), "serviceA_HttpHandler")
// 	defer span.End()
//
// 	cli := &http.Client{
// 		Transport: otelhttp.NewTransport(http.DefaultTransport),
// 	}
// 	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost:8082/serviceB", nil)
// 	if err != nil {
// 		panic(err)
// 	}
// 	resp, err := cli.Do(req)
// 	if err != nil {
// 		panic(err)
// 	}
//
// 	w.Header().Add("SVC-RESPONSE", resp.Header.Get("SVC-RESPONSE"))
// }
//
// func serviceB(ctx context.Context, port int) {
// 	mux := http.NewServeMux()
// 	mux.HandleFunc("/serviceB", serviceB_HttpHandler)
// 	handler := otelhttp.NewHandler(mux, "server.http")
// 	serverPort := fmt.Sprintf(":%d", port)
// 	server := &http.Server{Addr: serverPort, Handler: handler}
//
// 	fmt.Println("serviceB listening on", server.Addr)
// 	if err := server.ListenAndServe(); err != nil {
// 		panic(err)
// 	}
// }
//
// func serviceB_HttpHandler(w http.ResponseWriter, r *http.Request) {
// 	ctx, span := otel.Tracer("myTracer").Start(r.Context(), "serviceB_HttpHandler")
// 	defer span.End()
//
// 	answer := add(ctx, 42, 1813)
// 	w.Header().Add("SVC-RESPONSE", fmt.Sprint(answer))
// 	fmt.Fprintf(w, "hello from serviceB: Answer is: %d", answer)
// }
//
// func add(ctx context.Context, x, y int64) int64 {
// 	ctx, span := otel.Tracer("myTracer").Start(
// 		ctx,
// 		"add",
// 		// add labels/tags/resources(if any) that are specific to this scope.
// 		trace.WithAttributes(attribute.String("component", "addition")),
// 		trace.WithAttributes(attribute.String("someKey", "someValue")),
// 		trace.WithAttributes(attribute.Int("age", 89)),
// 	)
// 	defer span.End()
//
// 	counter, _ := global.MeterProvider().
// 		Meter(
// 			"instrumentation/package/name",
// 			metric.WithInstrumentationVersion("0.0.1"),
// 		).
// 		Int64Counter(
// 			"add_counter",
// 			instrument.WithDescription("how many times add function has been called."),
// 		)
// 	counter.Add(
// 		ctx,
// 		1,
// 		// labels/tags
// 		attribute.String("component", "addition"),
// 		attribute.Int("age", 89),
// 	)
//
// 	log := logs.NewZerolog(ctx)
// 	log.Info().Msg("add_called")
//
// 	return x + y
//
// }
//
// // getTls returns a configuration that enables the use of mutual TLS.
// func getTls() (*tls.Config, error) {
// 	clientAuth, err := tls.LoadX509KeyPair("./confs/client.crt", "./confs/client.key")
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	caCert, err := os.ReadFile("./confs/rootCA.crt")
// 	if err != nil {
// 		return nil, err
// 	}
// 	caCertPool := x509.NewCertPool()
// 	caCertPool.AppendCertsFromPEM(caCert)
//
// 	c := &tls.Config{
// 		RootCAs:      caCertPool,
// 		Certificates: []tls.Certificate{clientAuth},
// 	}
//
// 	return c, nil
// }

// ------------------------------------------

// func serviceA(ctx context.Context, port int) {
// 	mux := http.NewServeMux()
// 	mux.HandleFunc("/serviceA", serviceA_HttpHandler)
// 	serverPort := fmt.Sprintf(":%d", port)
// 	server := &http.Server{Addr: serverPort, Handler: mux}
//
// 	fmt.Println("serviceA listening on", server.Addr)
// 	if err := server.ListenAndServe(); err != nil {
// 		panic(err)
// 	}
// }
//
// func serviceA_HttpHandler(w http.ResponseWriter, r *http.Request) {
// 	cli := &http.Client{}
// 	req, err := http.NewRequestWithContext(r.Context(), http.MethodGet, "http://localhost:8082/serviceB", nil)
// 	if err != nil {
// 		panic(err)
// 	}
// 	resp, err := cli.Do(req)
// 	if err != nil {
// 		panic(err)
// 	}
//
// 	w.Header().Add("SVC-RESPONSE", resp.Header.Get("SVC-RESPONSE"))
// }
//
// func serviceB(ctx context.Context, port int) {
// 	mux := http.NewServeMux()
// 	mux.HandleFunc("/serviceB", serviceB_HttpHandler)
// 	serverPort := fmt.Sprintf(":%d", port)
// 	server := &http.Server{Addr: serverPort, Handler: mux}
//
// 	fmt.Println("serviceB listening on", server.Addr)
// 	if err := server.ListenAndServe(); err != nil {
// 		panic(err)
// 	}
// }
//
// func serviceB_HttpHandler(w http.ResponseWriter, r *http.Request) {
// 	answer := add(r.Context(), 42, 1813)
// 	w.Header().Add("SVC-RESPONSE", fmt.Sprint(answer))
// 	fmt.Fprintf(w, "hello from serviceB: Answer is: %d", answer)
// }
//
// func add(ctx context.Context, x, y int64) int64 { return x + y }
