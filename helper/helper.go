package helper

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
)

func GetMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}
func GetToken(text string) string {
	token := sha256.Sum256([]byte(text))
	strToken := hex.EncodeToString(token[:])
	return strToken
}
