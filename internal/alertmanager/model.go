package alertmanager

import "time"

type KV map[string]string

type Webhook struct {
	Version           string  `json:"version"`
	GroupKey          string  `json:"groupKey"`
	TruncatedAlerts   int     `json:"truncatedAlerts"`
	Status            string  `json:"status"`
	Receiver          string  `json:"receiver"`
	ExternalURL       string  `json:"externalURL"`
	GroupLabels       KV      `json:"groupLabels"`
	CommonLabels      KV      `json:"commonLabels"`
	CommonAnnotations KV      `json:"commonAnnotations"`
	Alerts            []Alert `json:"alerts"`
}

type Alert struct {
	Status       string    `json:"status"`
	Labels       KV        `json:"labels"`
	Annotations  KV        `json:"annotations"`
	StartsAt     time.Time `json:"startsAt"`
	EndsAt       time.Time `json:"endsAt"`
	GeneratorURL string    `json:"generatorURL"`
}
