package cryptos

import (
	"crypto/md5"
	"encoding/hex"
)

//GetMd5 : To Hash the Password
func GetMd5(input string) string {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(input))
	return hex.EncodeToString(hash.Sum(nil))
}
