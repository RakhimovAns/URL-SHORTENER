package initializers

import (
	"crypto/md5"
	"encoding/base64"
)

func HashString(link string) string {
	data := []byte(link)
	hash := md5.Sum(data)
	hashBase64 := base64.StdEncoding.EncodeToString(hash[:])
	return hashBase64[:7]
}

func Parser(link string) string {
	if (link[0:min(len(link), 8)] == "https://") == false {
		link = "https://" + link
	}
	return link
}
