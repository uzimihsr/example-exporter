# ビルド用イメージ
FROM golang:1.14

# mainパッケージがあるディレクトリ(.)をまるごとコピー
COPY . ./goapp
WORKDIR ./goapp

# goapp内のgo.mod, go.sumで依存関係を管理している場合に使用
RUN go mod download

# クロスコンパイル
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /app .

# バイナリを載せるイメージ
FROM scratch
WORKDIR goapp

# ビルド済みのバイナリをコピー
COPY --from=0 /app ./

# httpsで通信を行う場合に使用
COPY --from=0 /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

ENTRYPOINT ["./app"]
