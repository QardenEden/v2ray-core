// +build json

package protocol

import (
	"encoding/json"

	"github.com/v2ray/v2ray-core/common/uuid"
)

func (u *User) UnmarshalJSON(data []byte) error {
	type rawUser struct {
		IdString     string `json:"id"`
		EmailString  string `json:"email"`
		LevelByte    byte   `json:"level"`
		AlterIdCount uint16 `json:"alterId"`
	}
	var rawUserValue rawUser
	if err := json.Unmarshal(data, &rawUserValue); err != nil {
		return err
	}
	id, err := uuid.ParseString(rawUserValue.IdString)
	if err != nil {
		return err
	}
	primaryID := NewID(id)
	alterIDs := NewAlterIDs(primaryID, rawUserValue.AlterIdCount)
	*u = *NewUser(primaryID, alterIDs, UserLevel(rawUserValue.LevelByte), rawUserValue.EmailString)

	return nil
}
