package httpclient

import (
	"crypto/md5"
	"fmt"
)

type MD5Digest struct {
	Digest   string `json:"md5"`
	Original string `json:"original"`
}

func Example_basic() {
	var md5Digest MD5Digest
	message := "hello world from httpclient"
	digest := fmt.Sprintf("%x", md5.Sum([]byte(message)))

	if err := New().WithEndpoint("http://md5.jsontest.com/").Get("/").Query("text", message).JsonDecode(&md5Digest); err != nil {
		fmt.Println(err)
		return
	}

	if md5Digest.Digest != digest {
		fmt.Println("Failed to calcualte MD5 digest")
	}

}
