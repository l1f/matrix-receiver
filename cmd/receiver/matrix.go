package receiver

import (
	"matrix-alertmanager/internal/application"
	"matrix-alertmanager/internal/matrix"
)

func createClients(ctx application.Context) (map[string]matrix.Matrix, error) {
	clients := map[string]matrix.Matrix{}
	for name, user := range ctx.Config.Users {
		server := ctx.Config.HomeServers[user.HomeServer]

		if user.Token != "" {
			client, err := matrix.NewTokenClient(server.URL, user.Username, user.Token)
			if err != nil {
				return nil, err
			}
			clients[name] = *client
			continue
		}

		client, err := matrix.NewPasswordClient(server.URL, user.Username, user.Password)
		if err != nil {
			return nil, err
		}

		clients[name] = *client
	}

	return clients, nil
}
