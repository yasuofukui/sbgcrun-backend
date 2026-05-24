package handlers

import (
	"math/rand/v2"
	"net/http"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"

	"github.com/labstack/echo/v4"
	"github.com/uma-arai/sbcntr-backend/domain/model"
)

// NewsHandler ...
type NewsHandler struct{}

// NewNewsHandler ...
func NewNewsHandler() *NewsHandler {
	return &NewsHandler{}
}

// GetNews ...
func (handler *NewsHandler) GetNews() echo.HandlerFunc {
	news := []model.News{
		{
			ID:        "1",
			Title:     "新サービス『SBGCRun』提供開始のお知らせ",
			Content:   "本日より新サービス『SBGCRun』の提供を開始いたしました。皆様のご利用を心よりお待ちしております。",
			Category:  "announcement",
			Author:    "運営チーム",
			CreatedAt: "2026-05-01T10:00:00+09:00",
		},
		{
			ID:        "2",
			Title:     "メンテナンスのお知らせ",
			Content:   "2026年6月1日 深夜2:00〜4:00にシステムメンテナンスを実施いたします。",
			Category:  "maintenance",
			Author:    "インフラチーム",
			CreatedAt: "2026-05-15T12:30:00+09:00",
		},
		{
			ID:        "3",
			Title:     "新機能リリースのお知らせ",
			Content:   "ペット予約機能をリリースしました。ぜひご利用ください。",
			Category:  "release",
			Author:    "開発チーム",
			CreatedAt: "2026-05-20T09:15:00+09:00",
		},
	}

	return func(c echo.Context) error {
		ctx := c.Request().Context()

		tracer := otel.Tracer("news-handler")
		_, span := tracer.Start(ctx, "GetNews",
			trace.WithSpanKind(trace.SpanKindInternal),
		)
		defer span.End()

		// 30%の確率で500エラーを返す
		if rand.N(10) < 3 {
			span.SetAttributes(attribute.Bool("simulated_error", true))
			return echo.NewHTTPError(http.StatusInternalServerError, "Error")
		}

		span.SetAttributes(attribute.Int("news.count", len(news)))

		return c.JSON(http.StatusOK, model.APIResponse{
			Data: news,
		})
	}
}
