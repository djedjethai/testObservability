package api

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	obs "gitlab.com/grpasr/common/observability"
	tracing "gitlab.com/grpasr/common/observability/tracing/http"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	// "time"

	// "google.golang.org/grpc/codes"
	"gopkg.in/yaml.v2"
	// // added
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	// "go.opentelemetry.io/otel/metric"
	// "go.opentelemetry.io/otel/metric/global"
	// "go.opentelemetry.io/otel/metric/instrument"
	"go.opentelemetry.io/otel/trace"
)

// go get go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp

const (
	jaegerEndpoint string = "http://127.0.0.1:14268/api/traces"
	serviceName    string = "api"
	environment    string = "development"
	id                    = 0
)

var services Config

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

	// NOTE add tracingMiddleware to the handlers
	tr := otel.Tracer("go.opentelemetry.io")
	rootHandlerWithMiddleware := tracing.TracingMiddleware(tr, rootHandler, "rootHandler_http_req_res")
	calcHandlerWithMiddleware := tracing.TracingMiddleware(tr, calcHandler, "calcHandler_http_req_res")

	mux := http.NewServeMux()
	mux.HandleFunc("/", rootHandlerWithMiddleware)
	mux.HandleFunc("/calculate", calcHandlerWithMiddleware)
	services = GetServices()
	fmt.Println("see the services: ", services)
	handler := otelhttp.NewHandler(mux, "api-server")
	server := &http.Server{Addr: ":4000", Handler: handler}
	log.Println("api start, port 4000 ...")
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}

func rootHandler(w http.ResponseWriter, req *http.Request) {

	fmt.Println("hittttt the api /...................")
	fmt.Fprintf(w, "%v", services)
}

func enableCors(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-B3-SpanId, X-B3-TraceId, X-B3-Sampled, traceparent")
}

func calcHandler(w http.ResponseWriter, req *http.Request) {

	enableCors(&w, req)
	if (*req).Method == "OPTIONS" {
		return
	}

	// get a span and its context
	ctx, span := otel.Tracer("go.opentelemetry.io").Start(req.Context(), "calcHandler_HttpHandler")
	defer span.End()

	// calcRequest, err := ParseCalcRequest(req.Body, span)
	calcRequest, err := ParseCalcRequest(req.Body, span)
	if err != nil {
		fmt.Println("errr parse request: ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Println("hittttt the api /calculate...................: ", calcRequest)

	var url string

	for _, n := range services.Services {
		if strings.ToLower(calcRequest.Method) == strings.ToLower(n.Name) {
			j, _ := json.Marshal(calcRequest.Operands)
			url = fmt.Sprintf("http://%s:%d/%s?o=%s", n.Host, n.Port, strings.ToLower(n.Name), strings.Trim(string(j), "[]"))
		}
	}

	if url == "" {
		http.Error(w, "could not find requested calculation method", http.StatusBadRequest)
	}

	client := &http.Client{
		Transport: otelhttp.NewTransport(http.DefaultTransport),
	}
	request, _ := http.NewRequestWithContext(ctx, "GET", url, nil)

	res, err := client.Do(request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp, err := strconv.Atoi(string(body))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println("In api see response before sending: ", resp)
	fmt.Fprintf(w, "%d", resp)
}

type CalcRequest struct {
	Method   string `json:"method"`
	Operands []int  `json:"operands"`
}

// func ParseCalcRequest(body io.Reader, span trace.Span) (CalcRequest, error) {
func ParseCalcRequest(body io.Reader, span trace.Span) (CalcRequest, error) {
	var parsedRequest CalcRequest

	// Add event: attempting to parse body
	span.AddEvent("Attempting to parse body", trace.WithAttributes(attribute.String("event_key", "event_value")))
	span.AddEvent(fmt.Sprintf("%s", body))
	err := json.NewDecoder(body).Decode(&parsedRequest)
	if err != nil {
		// span.SetStatus(codes.InvalidArgument)
		// 500 is the http.Code
		span.SetStatus(500, "the description")
		span.AddEvent(err.Error())
		span.End()
		return parsedRequest, err
	}
	// NOTE this would have cut the span recording at this point.
	// removing it let the caller func close it and get this added events
	// span.End()

	return parsedRequest, nil
}

type Config struct {
	Services []struct {
		Name string `yaml:"name"`
		Host string `yaml:"host"`
		Port int    `yaml:"port"`
	} `yaml:"services"`
}

func GetServices() Config {
	wd, err := os.Getwd()
	if err != nil {
		log.Println("Error:", err)
	}

	relativePath := "../../services.yaml"

	absolutePath := filepath.Join(wd, relativePath)

	f, err := os.Open(absolutePath)
	if err != nil {
		log.Fatal("could not open config")
	}
	defer f.Close()

	var cfg Config
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		log.Fatal("could not process config")
	}
	return cfg
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

// func GetServices() Config {
// 	f, err := os.Open("services.yaml")
// 	if err != nil {
// 		log.Fatal("could not open config")
// 	}
// 	defer f.Close()
//
// 	var cfg Config
// 	decoder := yaml.NewDecoder(f)
// 	err = decoder.Decode(&cfg)
// 	if err != nil {
// 		log.Fatal("could not process config")
// 	}
// 	return cfg
// }
