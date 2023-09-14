// Code generated by "weaver generate". DO NOT EDIT.
//go:build !ignoreWeaverGen

package interview

import (
	"context"
	"errors"
	"github.com/ServiceWeaver/weaver"
	"github.com/ServiceWeaver/weaver/runtime/codegen"
	"github.com/timliudream/go-test/serviceWeaverDemo/model"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"reflect"
)

func init() {
	codegen.Register(codegen.Registration{
		Name:  "github.com/timliudream/go-test/serviceWeaverDemo/interview/Service",
		Iface: reflect.TypeOf((*Service)(nil)).Elem(),
		Impl:  reflect.TypeOf(implementation{}),
		LocalStubFn: func(impl any, caller string, tracer trace.Tracer) any {
			return service_local_stub{impl: impl.(Service), tracer: tracer, makeInterviewMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "github.com/timliudream/go-test/serviceWeaverDemo/interview/Service", Method: "MakeInterview", Remote: false})}
		},
		ClientStubFn: func(stub codegen.Stub, caller string) any {
			return service_client_stub{stub: stub, makeInterviewMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "github.com/timliudream/go-test/serviceWeaverDemo/interview/Service", Method: "MakeInterview", Remote: true})}
		},
		ServerStubFn: func(impl any, addLoad func(uint64, float64)) codegen.Server {
			return service_server_stub{impl: impl.(Service), addLoad: addLoad}
		},
		ReflectStubFn: func(caller func(string, context.Context, []any, []any) error) any {
			return service_reflect_stub{caller: caller}
		},
		RefData: "",
	})
}

// weaver.InstanceOf checks.
var _ weaver.InstanceOf[Service] = (*implementation)(nil)

// weaver.Router checks.
var _ weaver.Unrouted = (*implementation)(nil)

// Local stub implementations.

type service_local_stub struct {
	impl                 Service
	tracer               trace.Tracer
	makeInterviewMetrics *codegen.MethodMetrics
}

// Check that service_local_stub implements the Service interface.
var _ Service = (*service_local_stub)(nil)

func (s service_local_stub) MakeInterview(ctx context.Context, a0 model.Golang) (err error) {
	// Update metrics.
	begin := s.makeInterviewMetrics.Begin()
	defer func() { s.makeInterviewMetrics.End(begin, err != nil, 0, 0) }()
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.tracer.Start(ctx, "interview.Service.MakeInterview", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.MakeInterview(ctx, a0)
}

// Client stub implementations.

type service_client_stub struct {
	stub                 codegen.Stub
	makeInterviewMetrics *codegen.MethodMetrics
}

// Check that service_client_stub implements the Service interface.
var _ Service = (*service_client_stub)(nil)

func (s service_client_stub) MakeInterview(ctx context.Context, a0 model.Golang) (err error) {
	// Update metrics.
	var requestBytes, replyBytes int
	begin := s.makeInterviewMetrics.Begin()
	defer func() { s.makeInterviewMetrics.End(begin, err != nil, requestBytes, replyBytes) }()

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.stub.Tracer().Start(ctx, "interview.Service.MakeInterview", trace.WithSpanKind(trace.SpanKindClient))
	}

	defer func() {
		// Catch and return any panics detected during encoding/decoding/rpc.
		if err == nil {
			err = codegen.CatchPanics(recover())
			if err != nil {
				err = errors.Join(weaver.RemoteCallError, err)
			}
		}

		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
		}
		span.End()

	}()

	// Encode arguments.
	enc := codegen.NewEncoder()
	(a0).WeaverMarshal(enc)
	var shardKey uint64

	// Call the remote method.
	requestBytes = len(enc.Data())
	var results []byte
	results, err = s.stub.Run(ctx, 0, enc.Data(), shardKey)
	replyBytes = len(results)
	if err != nil {
		err = errors.Join(weaver.RemoteCallError, err)
		return
	}

	// Decode the results.
	dec := codegen.NewDecoder(results)
	err = dec.Error()
	return
}

// Note that "weaver generate" will always generate the error message below.
// Everything is okay. The error message is only relevant if you see it when
// you run "go build" or "go run".
var _ codegen.LatestVersion = codegen.Version[[0][20]struct{}](`

ERROR: You generated this file with 'weaver generate' v0.21.2 (codegen
version v0.20.0). The generated code is incompatible with the version of the
github.com/ServiceWeaver/weaver module that you're using. The weaver module
version can be found in your go.mod file or by running the following command.

    go list -m github.com/ServiceWeaver/weaver

We recommend updating the weaver module and the 'weaver generate' command by
running the following.

    go get github.com/ServiceWeaver/weaver@latest
    go install github.com/ServiceWeaver/weaver/cmd/weaver@latest

Then, re-run 'weaver generate' and re-build your code. If the problem persists,
please file an issue at https://github.com/ServiceWeaver/weaver/issues.

`)

// Server stub implementations.

type service_server_stub struct {
	impl    Service
	addLoad func(key uint64, load float64)
}

// Check that service_server_stub implements the codegen.Server interface.
var _ codegen.Server = (*service_server_stub)(nil)

// GetStubFn implements the codegen.Server interface.
func (s service_server_stub) GetStubFn(method string) func(ctx context.Context, args []byte) ([]byte, error) {
	switch method {
	case "MakeInterview":
		return s.makeInterview
	default:
		return nil
	}
}

func (s service_server_stub) makeInterview(ctx context.Context, args []byte) (res []byte, err error) {
	// Catch and return any panics detected during encoding/decoding/rpc.
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	// Decode arguments.
	dec := codegen.NewDecoder(args)
	var a0 model.Golang
	(&a0).WeaverUnmarshal(dec)

	// TODO(rgrandl): The deferred function above will recover from panics in the
	// user code: fix this.
	// Call the local method.
	appErr := s.impl.MakeInterview(ctx, a0)

	// Encode the results.
	enc := codegen.NewEncoder()
	enc.Error(appErr)
	return enc.Data(), nil
}

// Reflect stub implementations.

type service_reflect_stub struct {
	caller func(string, context.Context, []any, []any) error
}

// Check that service_reflect_stub implements the Service interface.
var _ Service = (*service_reflect_stub)(nil)

func (s service_reflect_stub) MakeInterview(ctx context.Context, a0 model.Golang) (err error) {
	err = s.caller("MakeInterview", ctx, []any{a0}, []any{})
	return
}
