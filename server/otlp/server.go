package otlp

import (
	"context"
	"fmt"
	"net"
	"strings"

	"github.com/kubeshop/tracetest/server/model"
	"github.com/kubeshop/tracetest/server/traces"
	"go.opentelemetry.io/otel/trace"
	pb "go.opentelemetry.io/proto/otlp/collector/trace/v1"
	v1 "go.opentelemetry.io/proto/otlp/trace/v1"
	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedTraceServiceServer

	addr string
	db   model.Repository

	gServer *grpc.Server
}

func NewServer(addr string, db model.Repository) *Server {
	return &Server{
		addr: addr,
		db:   db,
	}
}

func (s *Server) Start() error {
	s.gServer = grpc.NewServer()
	listener, err := net.Listen("tcp", s.addr)
	if err != nil {
		return fmt.Errorf("cannot listen on address %s: %w", s.addr, err)
	}
	pb.RegisterTraceServiceServer(s.gServer, s)
	return s.gServer.Serve(listener)
}

func (s *Server) Stop() {
	s.gServer.Stop()
}

func (s Server) Export(ctx context.Context, request *pb.ExportTraceServiceRequest) (*pb.ExportTraceServiceResponse, error) {
	ds, err := s.db.DefaultDataStore(ctx)

	if err != nil || !ds.IsOTLPBasedProvider() {
		fmt.Println("OTLP server is not enabled. Ignoring request")
		return &pb.ExportTraceServiceResponse{}, nil
	}

	if len(request.ResourceSpans) == 0 {
		return &pb.ExportTraceServiceResponse{}, nil
	}

	spansByTrace := s.getSpansByTrace(request)

	for traceID, spans := range spansByTrace {
		s.saveSpansIntoTest(ctx, traceID, spans)
	}

	return &pb.ExportTraceServiceResponse{
		PartialSuccess: &pb.ExportTracePartialSuccess{
			RejectedSpans: 0,
		},
	}, nil
}

func (s Server) getSpansByTrace(request *pb.ExportTraceServiceRequest) map[trace.TraceID][]model.Span {
	otelSpans := make([]*v1.Span, 0)
	for _, resourceSpan := range request.ResourceSpans {
		for _, spans := range resourceSpan.ScopeSpans {
			otelSpans = append(otelSpans, spans.Spans...)
		}
	}

	spansByTrace := make(map[trace.TraceID][]model.Span)

	for _, span := range otelSpans {
		traceID := traces.CreateTraceID(span.TraceId)
		var existingArray []model.Span
		if spansArray, ok := spansByTrace[traceID]; ok {
			existingArray = spansArray
		} else {
			existingArray = make([]model.Span, 0)
		}

		existingArray = append(existingArray, *traces.ConvertOtelSpanIntoSpan(span))
		spansByTrace[traceID] = existingArray
	}

	return spansByTrace
}

func (s Server) saveSpansIntoTest(ctx context.Context, traceID trace.TraceID, spans []model.Span) error {
	run, err := s.db.GetRunByTraceID(ctx, traceID)
	if err != nil && strings.Contains(err.Error(), "record not found") {
		// span is not part of any known test run. So it will be ignored
		return nil
	}

	if err != nil {
		return fmt.Errorf("could not find test run with traceID %s: %w", traceID.String(), err)
	}

	if run.State != model.RunStateAwaitingTrace {
		// test is not waiting for trace, so we can completely ignore those as they might
		// mess up with the test integrity.
		//
		// For example:
		// Imagine that a test failed because Span A wasn't available in the trace and one minute
		// later, the span is received and added to the span. When investigating the issue,
		// one might be confused and maybe think it's a bug in our assertion engine
		// because the assertion failed, but the span is there. However, it wasn't at
		// the moment the assertion ran.
		//
		// So, to reduce friction and prevent long debugging hours, we can just disable this.

		return nil
	}

	existingSpans := run.Trace.Spans()
	newSpans := append(existingSpans, spans...)
	newTrace := model.NewTrace(traceID.String(), newSpans)

	run.Trace = &newTrace

	err = s.db.UpdateRun(ctx, run)
	if err != nil {
		return fmt.Errorf("could not update run: %w", err)
	}

	return nil
}
