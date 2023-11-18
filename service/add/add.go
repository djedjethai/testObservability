package add

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	obs "gitlab.com/grpasr/common/observability"
	logs "gitlab.com/grpasr/common/observability/logging"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	// "time"
	// // added
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/metric/global"
	"go.opentelemetry.io/otel/metric/instrument"
	"go.opentelemetry.io/otel/trace"
)

const (
	jaegerEndpoint string = "http://127.0.0.1:14268/api/traces"
	serviceName    string = "add"
	environment    string = "development"
	id                    = 1
)

func Start() {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	tls, err := getTls()
	if err != nil {
		fmt.Println("err from tls: ", err)
	}

	{
		tp, err := obs.SetupTracing(ctx, serviceName, tls)
		if err != nil {
			panic(err)
		}
		defer tp.Shutdown(ctx)

		mp, err := obs.SetupMetrics(ctx, serviceName, tls)
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

	mux := http.NewServeMux()
	mux.HandleFunc("/", addHandler)
	handler := otelhttp.NewHandler(mux, "add-server")
	server := &http.Server{Addr: ":4001", Handler: handler}
	log.Println("api start, port 4001 ...")
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}

func addHandler(w http.ResponseWriter, req *http.Request) {
	ctx, span := otel.Tracer("go.opentelemetry.io").Start(req.Context(), "add-server_addHandler")
	defer span.End()

	values := strings.Split(req.URL.Query()["o"][0], ",")
	fmt.Println("In add - addHandler: ", values)

	res, err := addRunner(ctx, values)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "%d", res)
}

func addRunner(ctx context.Context, values []string) (res int, err error) {
	ctx, span := otel.Tracer("go.opentelemetry.io").Start(
		ctx,
		"add-server_addRunner",
		// add labels/tags/resources(if any) that are specific to this scope.
		trace.WithAttributes(attribute.String("component", "addition")),
		trace.WithAttributes(attribute.String("someKey", "someValue")),
	)
	defer span.End()

	for _, n := range values {
		i, errE := strconv.Atoi(n)
		if err != nil {
			// http.Error(w, err.Error(), http.StatusBadRequest)
			err = errE
			return
		}
		res += i
	}

	// setted on the same span as is trace.WithAttributes(attribute.String(...))
	span.SetAttributes(attribute.Int("add-res", res))

	counter, _ := global.MeterProvider().
		Meter(
			"go.opentelemetry.io",
			metric.WithInstrumentationVersion("0.0.1"),
		).Int64Counter(
		"add_counter",
		instrument.WithDescription("how many times add function has been called."),
	)

	counter.Add(
		ctx,
		1,
		// labels/tag
		attribute.String("component", "addition"),
		attribute.Int("add-res", res),
	)

	log := logs.NewZerolog(ctx)
	log.Info().Msg("add_called")

	// return res
	return
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
