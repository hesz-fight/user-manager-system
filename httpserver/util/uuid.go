package util

import (
	"crypto/rand"
	"fmt"
)

func CreateUUID() (string, error) {
	u := new([16]byte)
	if _, err := rand.Read(u[:]); err != nil {
		return "", err
	}
	u[8] = (u[8] | 0x40) & 0x7F
	u[6] = (u[6] & 0xF) | (0x4 << 4)
	uuid := fmt.Sprintf("%x-%x-%x-%x-%x", u[0:4], u[4:6], u[6:8], u[8:10], u[10:])

	return uuid, nil
}
