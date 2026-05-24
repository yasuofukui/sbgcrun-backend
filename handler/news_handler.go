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
			Title:     "子犬・子猫の新入荷情報🐶🐱",
			Content:   "今週はトイプードル、柴犬、スコティッシュフォールド、マンチカンの子たちが新しく仲間入りしました！ぜひ会いに来てくださいね。",
			Category:  "new_arrival",
			Author:    "店長 田中",
			CreatedAt: "2026-05-01T10:00:00+09:00",
		},
		{
			ID:        "2",
			Title:     "ペットフード20%OFFセール開催中🎉",
			Content:   "ロイヤルカナン、ヒルズ、シーバなど人気ブランドのフードが全品20%OFF！6月末までの期間限定セールです。",
			Category:  "campaign",
			Author:    "販売スタッフ",
			CreatedAt: "2026-05-15T12:30:00+09:00",
		},
		{
			ID:        "3",
			Title:     "無料しつけ相談会のご案内",
			Content:   "プロのドッグトレーナーによる無料しつけ相談会を毎週日曜日に開催中。お困りごとはお気軽にご相談ください。",
			Category:  "event",
			Author:    "トレーナー 山田",
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

		// 80%の確率で500エラーを返す
		if rand.N(10) < 8 {
			span.SetAttributes(attribute.Bool("simulated_error", true))
			return echo.NewHTTPError(http.StatusInternalServerError, "Error")
		}

		span.SetAttributes(attribute.Int("news.count", len(news)))

		return c.JSON(http.StatusOK, model.APIResponse{
			Data: news,
		})
	}
}
