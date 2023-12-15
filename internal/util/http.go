package util

import (
	"context"
	"time"

	"github.com/qaultsabit/bengkel/domain"
)

func ResponseInterceptor(ctx context.Context, res *domain.ApiResponse) {
	traceIdInf := ctx.Value("requestid")
	traceId := ""
	if traceIdInf != nil {
		traceId = traceIdInf.(string)
	}

	res.Timestamp = time.Now()
	res.TraceID = traceId
}
