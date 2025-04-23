package metrics

import "github.com/prometheus/client_golang/prometheus"

var (
	TaskCreated = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "task_created_total",
			Help: "Total number of tasks created, grouped by type.",
		},
		[]string{"type"},
	)
	TaskCompleted = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "task_completed_total",
			Help: "Total number of tasks completed, grouped by type.",
		},
		[]string{"type"},
	)
	TaskFailed = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "task_failed_total",
			Help: "Total number of tasks failed, grouped by type.",
		},
		[]string{"type"},
	)
	TaskDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "task_duration_seconds",
			Help:    "Duration of tasks in seconds, grouped by type.",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"type"},
	)
)

func Register() {
	prometheus.MustRegister(TaskCreated, TaskCompleted, TaskFailed, TaskDuration)
}
