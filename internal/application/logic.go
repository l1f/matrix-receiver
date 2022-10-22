package application

import "matrix-alertmanager/internal/alertmanager"

type Logic interface {
	ScheduleMessage(webhook alertmanager.Webhook) error
}
