package crypto

import (
	"encoding/hex"
	"github.com/go-park-mail-ru/2020_1_k-on/pkg/constants"
	"golang.org/x/crypto/argon2"
)

func CreateToken(sessionId string) string {
	hash := argon2.IDKey([]byte(constants.CSRFKey), []byte(sessionId), 1, 64*1024, 4, 32)
	return hex.EncodeToString(hash[:])
}

func CheckToken(sessionId, token string) bool {
	return token == CreateToken(sessionId)
}
