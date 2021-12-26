package utils

import (
	"crypto/md5"
	"encoding/hex"
)

//GenMd5 生成md5加密串
func GenMd5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}
