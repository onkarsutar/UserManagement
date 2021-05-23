package middleware

import (
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/labstack/echo"
)

type Stats struct {
	Uptime       time.Time      `json:"uptime"`
	RequestCount int            `json:"requestCount"`
	Status       map[string]int `json:"status"`
	mutex        sync.RWMutex
}

var ServerStats *Stats

func Init(e *echo.Echo, o, r *echo.Group) {
	ServerStats = NewStats()
	e.Use(ServerStats.Process)
	o.GET("/serverstats", ServerStatsRoute)
}

func NewStats() *Stats {
	return &Stats{
		Uptime: time.Now(),
		Status: map[string]int{},
	}
}

func (s *Stats) Process(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if err := next(c); err != nil {
			c.Error(err)
		}
		s.mutex.Lock()
		defer s.mutex.Unlock()
		s.RequestCount++
		status := strconv.Itoa(c.Response().Status)
		s.Status[status]++
		return nil
	}
}

// ServerStatsRoute : API to check server stats.
func ServerStatsRoute(c echo.Context) error {
	ServerStats.mutex.RLock()
	defer ServerStats.mutex.RUnlock()
	return c.JSON(http.StatusOK, ServerStats)
}
