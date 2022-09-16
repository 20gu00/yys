package MD5

import (
	"crypto/md5"
	"fmt"
	"io"
)

func MD5(s string) string {
	h := md5.New()
	io.WriteString(h, s)
	return fmt.Sprintf("%x", h.Sum(nil))
}
