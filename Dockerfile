# === builder: ビルド用 ===
FROM golang:1.26 AS builder
WORKDIR /app

# 依存関係のみ先にダウンロードし、レイヤーキャッシュを効かせる
COPY go.mod go.sum ./
RUN go mod download

COPY . /app

# lint/vetはCloud Buildの別ステップ（cloudbuild_push.yamlのlint）で実行する
# go-buildキャッシュはローカルのBuildKitビルドで増分コンパイルに効く
RUN --mount=type=cache,target=/root/.cache/go-build \
    GOOS=linux GOARCH=amd64 CGO_ENABLED=0 \
    go build -ldflags "-s -w" -o /app/bin/main ./main.go

# === runner: 本番イメージ ===
# 静的バイナリのためdistroless staticで動作する（ca-certificates同梱・非root実行）
FROM gcr.io/distroless/static-debian12:nonroot AS runner

COPY --from=builder /app/bin/main /main
EXPOSE 8081
ENTRYPOINT ["/main"]
