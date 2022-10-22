package matrix

import (
	"maunium.net/go/mautrix"
	"maunium.net/go/mautrix/id"
)

type Matrix struct {
	client mautrix.Client
}

func NewPasswordClient(homeServer, username, password string) (*Matrix, error) {
	client, err := mautrix.NewClient(homeServer, "", "")
	if err != nil {
		return nil, err
	}

	_, err = client.Login(&mautrix.ReqLogin{
		Type: "m.login.password",
		Identifier: mautrix.UserIdentifier{
			Type: mautrix.IdentifierTypeUser, User: username,
		},
		Password:                 password,
		DeviceID:                 "Matrix-Alertmanager",
		InitialDeviceDisplayName: "Matrix-Alertmanager",
		StoreCredentials:         true,
	})
	if err != nil {
		return nil, err
	}

	return &Matrix{client: *client}, nil
}

func NewTokenClient(homeServer, username, token string) (*Matrix, error) {
	client, err := mautrix.NewClient(homeServer, "", "")
	if err != nil {
		return nil, err
	}

	_, err = client.Login(&mautrix.ReqLogin{
		Type: "m.login.password",
		Identifier: mautrix.UserIdentifier{
			Type: mautrix.IdentifierTypeUser, User: username,
		},
		Token:                    token,
		DeviceID:                 "Matrix-Alertmanager",
		InitialDeviceDisplayName: "Matrix-Alertmanager",
	})
	if err != nil {
		return nil, err
	}

	return &Matrix{client: *client}, nil
}

func (m Matrix) Send(roomID string, text string) error {
	_, err := m.client.SendText(id.RoomID(roomID), text)
	return err
}

func (m Matrix) JoinRoomByID(roomID string) error {
	_, err := m.client.JoinRoomByID(id.RoomID(roomID))
	return err
}

func (m Matrix) JoinRoomByAlias(roomAlias string) error {
	_, err := m.client.JoinRoom(roomAlias, "cyber-missile.io", nil)
	return err
}

func (m Matrix) UserIsJoined(roomID string) bool {
	room := mautrix.NewRoom(id.RoomID(roomID))
	membershipState := room.GetMembershipState(m.client.UserID)

	return !membershipState.IsLeaveOrBan()
}

func (m Matrix) GetRoomIdByAlias(alias string) (string, error) {
	resolveAlias, err := m.client.ResolveAlias(id.RoomAlias(alias))
	if err != nil {
		return "", err
	}

	return string(resolveAlias.RoomID), nil
}
