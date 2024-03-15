package utils

import (
	"context"
	"fmt"

	"log/slog"
	"runtime"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/google/uuid"
)

func NewTrace(ctx context.Context) context.Context {
	if ctx.Value("requestID") == nil {
		requestID := middleware.GetReqID(ctx)
		if requestID == "" {
			requestID = uuid.New().String()
		}

		ctx = context.WithValue(ctx, "requestID", requestID)
		ctx = context.WithValue(ctx, "parentID", ctx.Value("requestID"))
		ctx = context.WithValue(ctx, "spanID", uuid.New().String())
	} else {
		ctx = context.WithValue(ctx, "parentID", ctx.Value("spanID"))
		ctx = context.WithValue(ctx, "spanID", uuid.New().String())
	}

	pc, file, _, ok := runtime.Caller(1)
	if !ok {
		fmt.Println("No caller information")
	}

	fn := runtime.FuncForPC(pc)

	ctx = context.WithValue(ctx, "file", file)
	ctx = context.WithValue(ctx, "function", fn.Name())

	return ctx
}

func InfoTrace(ctx context.Context, message string) {
	slog.Info(
		message,
		"requestID", ctx.Value("requestID"),
		"parentID", ctx.Value("parentID"),
		"spanID", ctx.Value("spanID"),
		"file", ctx.Value("file"),
		"function", ctx.Value("function"),
	)
}

func DebugTrace(ctx context.Context, message string) {
	slog.Debug(
		message,
		"requestID", ctx.Value("requestID"),
		"parentID", ctx.Value("parentID"),
		"spanID", ctx.Value("spanID"),
		"file", ctx.Value("file"),
		"function", ctx.Value("function"),
	)
}

func ErrorTrace(ctx context.Context, message string, err error, args ...interface{}) {
	slog.Error(
		message,
		"requestID", ctx.Value("requestID"),
		"parentID", ctx.Value("parentID"),
		"spanID", ctx.Value("spanID"),
		"file", ctx.Value("file"),
		"function", ctx.Value("function"),
		"args", args,
		"error", err,
	)
}
