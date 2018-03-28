package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	opentracing "github.com/opentracing/opentracing-go"
	zipkin "github.com/openzipkin/zipkin-go-opentracing"
)

const (
	// Our service name.
	serviceName = "svc1"

	// Host + port of our service.
	hostPort = "127.0.0.1:61001"

	// Endpoint to send Zipkin spans to.
	zipkinHTTPEndpoint = "http://127.0.0.1:9411/api/v1/spans"

	// Debug mode.
	debug = false

	// same span can be set to true for RPC style spans (Zipkin V1) vs Node style (OpenTracing)
	sameSpan = true

	// make Tracer generate 128 bit traceID's for root spans.
	traceID128Bit = true
)

func main() {
	initGlobalTracer()

	mux := http.NewServeMux()
	mux.HandleFunc("/", Tracing)
	http.ListenAndServe(hostPort, mux)
}

// Tracing handle all request
func Tracing(w http.ResponseWriter, req *http.Request) {
	tracer := opentracing.GlobalTracer()
	span := tracer.StartSpan("tracing")
	defer span.Finish()

	carrier := opentracing.HTTPHeadersCarrier(req.Header)
	if err := tracer.Inject(span.Context(), opentracing.HTTPHeaders, carrier); err != nil {
		w.WriteHeader(400)
		w.Write([]byte("err:" + err.Error()))
		return
	}

	logEvent(span, 1)
	logEvent(span, 2)
	logEvent(span, 3)

	// Inject the Span context into the outgoing HTTP Request.
	if err := tracer.Inject(
		span.Context(),
		opentracing.TextMap,
		opentracing.HTTPHeadersCarrier(req.Header),
	); err != nil {
		fmt.Printf("error encountered while trying to inject span: %+v\n", err)
	}

	w.Write([]byte("Hello World"))
}

func logEvent(span opentracing.Span, step int) {
	span.LogEvent(fmt.Sprintf("event:%d", step))
	time.Sleep(time.Second * 5)
}

func initGlobalTracer() {
	collector, err := zipkin.NewHTTPCollector(zipkinHTTPEndpoint)
	if err != nil {
		fmt.Printf("unable to create Zipkin HTTP collector: %+v\n", err)
		os.Exit(-1)
	}

	recorder := zipkin.NewRecorder(collector, debug, hostPort, serviceName)

	tracer, err := zipkin.NewTracer(recorder, zipkin.ClientServerSameSpan(sameSpan))
	if err != nil {
		fmt.Printf("unable to create Zipkin tracer: %+v\n", err)
		os.Exit(-1)
	}

	opentracing.InitGlobalTracer(tracer)
}
