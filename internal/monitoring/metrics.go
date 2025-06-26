package monitoring

import (
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	startTime = time.Now()

	ConnectionsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "connections_total",
			Help: "Total number of connections",
		},
		[]string{"service"},
	)

	ConnectionsCurrent = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "connections_current",
			Help: "Current active connections",
		},
		[]string{"service"},
	)

	PlayersCurrent = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "players_current",
			Help: "Current active players",
		},
		[]string{"service"},
	)

	// Database metrics
	DatabaseFastSaveQueueSize = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "database_fast_save_queue_size",
			Help: "Number of items in the fast save queue",
		},
	)

	DatabaseSlowSaveQueueSize = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "database_slow_save_queue_size",
			Help: "Number of items in the slow save queue",
		},
	)

	DatabaseCommitTotal = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "database_commit_total",
			Help: "Total number of database commit operations",
		},
	)

	DatabaseModifyTotal = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "database_modify_total",
			Help: "Total number of database modify operations",
		},
	)

	DatabaseWriteTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "database_write_total",
			Help: "Total number of database write operations",
		},
		[]string{"type"}, // "write" or "delete"
	)

	// Company metrics
	CompanyTimerTickTotal = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "company_timer_tick_total",
			Help: "Total number of company timer ticks",
		},
	)

	// Duchy metrics
	TransportationTimerTickTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "transportation_timer_tick_total",
			Help: "Total number of transportation timer ticks by duchy",
		},
		[]string{"duchy"},
	)

	// Planet metrics
	ExchangeTimerTickTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "exchange_timer_tick_total",
			Help: "Total number of exchange timer ticks by system and planet",
		},
		[]string{"system", "planet"},
	)

	// System metrics
	CleanTimerTickTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "clean_timer_tick_total",
			Help: "Total number of clean timer ticks by system",
		},
		[]string{"system"},
	)

	FightTimerTickTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "fight_timer_tick_total",
			Help: "Total number of fight timer ticks by system",
		},
		[]string{"system"},
	)

	MoveTimerTickTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "move_timer_tick_total",
			Help: "Total number of move timer ticks by system",
		},
		[]string{"system"},
	)

	PublicAddressTimerTickTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "public_address_timer_tick_total",
			Help: "Total number of public address timer ticks by system",
		},
		[]string{"system"},
	)

	RecycleTimerTickTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "recycle_timer_tick_total",
			Help: "Total number of recycle timer ticks by system",
		},
		[]string{"system"},
	)

	ShuttleTimerTickTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "shuttle_timer_tick_total",
			Help: "Total number of shuttle timer ticks by system",
		},
		[]string{"system"},
	)
)

func init() {
	prometheus.MustRegister(CompanyTimerTickTotal)
	prometheus.MustRegister(ConnectionsCurrent)
	prometheus.MustRegister(ConnectionsTotal)
	prometheus.MustRegister(DatabaseCommitTotal)
	prometheus.MustRegister(DatabaseFastSaveQueueSize)
	prometheus.MustRegister(DatabaseModifyTotal)
	prometheus.MustRegister(DatabaseSlowSaveQueueSize)
	prometheus.MustRegister(DatabaseWriteTotal)
	prometheus.MustRegister(ExchangeTimerTickTotal)
	prometheus.MustRegister(CleanTimerTickTotal)
	prometheus.MustRegister(FightTimerTickTotal)
	prometheus.MustRegister(MoveTimerTickTotal)
	prometheus.MustRegister(PlayersCurrent)
	prometheus.MustRegister(PublicAddressTimerTickTotal)
	prometheus.MustRegister(RecycleTimerTickTotal)
	prometheus.MustRegister(ShuttleTimerTickTotal)
	prometheus.MustRegister(TransportationTimerTickTotal)

	prometheus.MustRegister(prometheus.NewGaugeFunc(
		prometheus.GaugeOpts{
			Name: "uptime_seconds",
			Help: "Process uptime in seconds",
		},
		func() float64 { return time.Since(startTime).Seconds() },
	))
}

func Handler() http.Handler {
	return promhttp.Handler()
}
