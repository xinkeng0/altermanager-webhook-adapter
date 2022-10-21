package adapter

import "time"

// alert manager webhook mesage struc
type AlertManagerMessage struct {
	Receiver string `json:"receiver"`
	Status   string `json:"status"`
	Alerts   []struct {
		Status string `json:"status"`
		Labels struct {
			Alertname string `json:"alertname"`
			Severity  string `json:"severity"`
			Instance  string `json:"instance"`
		}
		Annotations struct {
			Description string `json:"description"`
			Summary     string `json:"summary"`
			Value       string `json:"value"`
		}
		StartsAt time.Time `json:"startsAt"`
		EndsAt   time.Time `json:"endsAt"`
	}
}
