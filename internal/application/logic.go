package application

import "matrix-alertmanager/internal/alertmanager"

type Logic interface {
	SendMessage(webhook alertmanager.Webhook) error
}
