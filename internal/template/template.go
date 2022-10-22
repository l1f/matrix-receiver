package template

import (
	"fmt"
	"matrix-alertmanager/internal/alertmanager"
	"strings"
)

func getDescription(annotation alertmanager.KV) (string, bool) {
	description, ok := annotation["description"]
	if !ok {
		return "", false
	}

	return description, true
}

func Generate(status string, alerts []alertmanager.Alert) string {
	var icon string
	var statusStr string

	switch status {
	case "firing":
		icon = "🚨"
		statusStr = "FIRING"
	case "resolved":
		icon = "🆗"
		statusStr = "RESOLVED"
	default:
		icon = "❓"
		statusStr = "UNKNOWN"
	}

	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("%s %s: ", icon, statusStr))

	if !(len(alerts) > 1) {
		builder.WriteString(" ")
		description, ok := getDescription(alerts[0].Annotations)
		if !ok {
			// todo
			return ""
		}

		builder.WriteString(description)

		return builder.String()
	}

	for i, alert := range alerts {
		description, ok := getDescription(alert.Annotations)
		if !ok {
			// todo
			return ""
		}

		if i == 0 {
			builder.WriteString("\n")
		}
		builder.WriteString(fmt.Sprintf(" - %s\n", description))
	}

	return builder.String()
}
