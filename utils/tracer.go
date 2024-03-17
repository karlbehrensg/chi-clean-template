package utils

import (
	"context"
	"fmt"

	"log/slog"
	"runtime"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/google/uuid"
)

type contextKey string

const (
	requestIDKey contextKey = "requestID"
	parentIDKey  contextKey = "parentID"
	spanIDKey    contextKey = "spanID"
	fileKey      contextKey = "file"
	functionKey  contextKey = "function"
)

func NewTrace(ctx context.Context) context.Context {
	if ctx.Value(requestIDKey) == nil {
		requestID := middleware.GetReqID(ctx)
		if requestID == "" {
			requestID = uuid.New().String()
		}

		ctx = context.WithValue(ctx, requestIDKey, requestID)
		ctx = context.WithValue(ctx, parentIDKey, ctx.Value(requestIDKey))
		ctx = context.WithValue(ctx, spanIDKey, uuid.New().String())
	} else {
		ctx = context.WithValue(ctx, parentIDKey, ctx.Value(spanIDKey))
		ctx = context.WithValue(ctx, spanIDKey, uuid.New().String())
	}

	pc, file, _, ok := runtime.Caller(1)
	if !ok {
		fmt.Println("No caller information")
	}

	fn := runtime.FuncForPC(pc)

	ctx = context.WithValue(ctx, fileKey, file)
	ctx = context.WithValue(ctx, functionKey, fn.Name())

	return ctx
}

func InfoTrace(ctx context.Context, message string) {
	slog.Info(
		message,
		"requestID", ctx.Value(requestIDKey),
		"parentID", ctx.Value(parentIDKey),
		"spanID", ctx.Value(spanIDKey),
		"file", ctx.Value(fileKey),
		"function", ctx.Value(functionKey),
	)
}

func DebugTrace(ctx context.Context, message string) {
	slog.Debug(
		message,
		"requestID", ctx.Value(requestIDKey),
		"parentID", ctx.Value(parentIDKey),
		"spanID", ctx.Value(spanIDKey),
		"file", ctx.Value(fileKey),
		"function", ctx.Value(functionKey),
	)
}

func ErrorTrace(ctx context.Context, message string, err error, args ...interface{}) {
	slog.Error(
		message,
		"requestID", ctx.Value(requestIDKey),
		"parentID", ctx.Value(parentIDKey),
		"spanID", ctx.Value(spanIDKey),
		"file", ctx.Value(fileKey),
		"function", ctx.Value(functionKey),
		"args", args,
		"error", err,
	)
}
