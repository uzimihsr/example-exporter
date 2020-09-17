package main

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// メトリクスの準備
var (
	// Counter
	exampleCounter = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "example_total",
			Help: "Example Counter",
		},
		[]string{"hoge"},
	)
	// Gauge
	exampleGauge = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "example_number",
			Help: "Example Gauge",
		},
		[]string{"fuga"},
	)
)

// 30秒ごとにCounterの値を1つ増やす関数
func count() {
	for {
		exampleCounter.With(prometheus.Labels{"hoge": "hogehoge"}).Inc()
		time.Sleep(30 * time.Second)
	}
}

// 10秒ごとにGaugeに乱数をセットする関数
func setRandomValue() {
	for {
		rand.Seed(time.Now().UnixNano())
		n := -1 + rand.Float64()*2
		exampleGauge.With(prometheus.Labels{"fuga": "fugafuga"}).Set(n)
		time.Sleep(10 * time.Second)
	}
}

func main() {

	// 各メトリクスの値を更新する関数をgo routineで実行
	go count()
	go setRandomValue()

	// /metricsへのリクエストを処理
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":2112", nil)
}
