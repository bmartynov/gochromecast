package utils

import (
	"io"
	"fmt"
	"crypto/md5"
)

func CalcMd5(src string) string {
	h := md5.New()
	io.WriteString(h, src)
	return fmt.Sprintf("%x", h.Sum(nil))
}