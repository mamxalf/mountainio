package helper

import (
	"github.com/google/uuid"
	"mountainio/app/constant"
)

func CheckNilDataFromUUID() uuid.UUID {
	return uuid.MustParse(constant.NilUUID)
}

func ConvertUUID(id string) uuid.UUID {
	return uuid.MustParse(id)
}
