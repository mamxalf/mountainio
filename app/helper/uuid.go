package helper

import (
	"github.com/google/uuid"
	"mountainio/app/constant"
	"mountainio/app/exception"
)

func CheckNilDataFromUUID() uuid.UUID {
	nilUUID, err := uuid.FromBytes([]byte(constant.NilUUID))
	exception.PanicIfNeeded(err)

	return nilUUID
}

func ConvertUUID(id string) uuid.UUID {
	nilUUID, err := uuid.FromBytes([]byte(id))
	exception.PanicIfNeeded(err)

	return nilUUID
}
