package api

import "github.com/go-chi/chi/v5"

func InitRoute(r chi.Router) {
	InitMonitorHandler(r)
	initDbExcelRouter(r)
	InitMetricsQueryHandler(r)
}
