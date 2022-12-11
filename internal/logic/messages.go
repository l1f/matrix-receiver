package logic

import (
	"errors"
	"matrix-alertmanager/internal/alertmanager"
	"matrix-alertmanager/internal/config"
	"matrix-alertmanager/internal/matrix"
	"matrix-alertmanager/internal/template"
)

var (
	ErrorReceiverNotFound = errors.New("receiver not found")
	ErrorUserNotFound     = errors.New("user not found")
	ErrorRoomNotFound     = errors.New("room not found")
)

var roomCache = map[string]string{}

func getRoomId(room config.Room, client matrix.Matrix) (string, error) {
	roomID, ok := roomCache[room.Address]
	if ok {
		return roomID, nil
	}

	roomIDResult, err := client.GetRoomIdByAlias(room.Address)
	if err != nil {
		// TODO
		return "", err
	}

	roomCache[room.Address] = roomIDResult
	return roomIDResult, nil
}

func joinRoom(client matrix.Matrix, roomID string) error {
	ok := client.UserIsJoined(roomID)
	if ok {
		return nil
	}

	return client.JoinRoomByAlias(roomID)
}

func (l logic) SendMessage(webhook alertmanager.Webhook) error {
	receiver, ok := l.ctx.Config.Receiver[webhook.Receiver]
	if !ok {
		return ErrorReceiverNotFound
	}

	client, ok := l.ctx.Clients[receiver.User]
	if !ok {
		return ErrorUserNotFound
	}

	room, ok := l.ctx.Config.Rooms[receiver.Room]
	if !ok {
		return ErrorRoomNotFound
	}

	roomID := room.Id
	if room.Id == "" {
		id, err := getRoomId(room, client)
		if err != nil {
			return err
		}

		roomID = id
	}

	err := joinRoom(client, roomID)
	if err != nil {
		return err
	}

	message := template.Generate(webhook.Status, webhook.Alerts)
	err = client.Send(roomID, message)
	if err != nil {
		l.ctx.Logger.Error().Err(err)
		return err
	}

	return nil
}
