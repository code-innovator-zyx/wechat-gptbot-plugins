package utils

import (
	"bytes"
	"fmt"
	"io"
	"math/rand"
	"net/url"
	"strings"
	"time"
)

var r *rand.Rand

func init() {
	source := rand.NewSource(time.Now().UnixNano())
	r = rand.New(source)
}

type Element interface {
	~int | ~string | ~float64 | ~float32
}

func EleInArray[T Element](slice []T, elem T) bool {
	for _, e := range slice {
		if e == elem {
			return true
		}
	}
	return false
}

func HasPrefixes(s string, profiles []string) bool {
	for _, profile := range profiles {
		if strings.HasPrefix(s, profile) {
			return true
		}
	}
	return false
}

type compress func(io.Reader, *bytes.Buffer) error

//func compressPng(reader io.Reader, buffer *bytes.Buffer) error {
//	// Png 无损压缩
//	io.Copy(buffer, reader) //
//	return nil
//}

func BuildPersonalMessage(userName, content string) string {
	builder := strings.Builder{}
	builder.WriteString("【")
	builder.WriteString(userName)
	builder.WriteString("】:")
	builder.WriteString(content)
	return builder.String()
}

func BuildResponseMessage(userName, content, reply string) string {
	builder := strings.Builder{}
	//builder.WriteString("[")
	//builder.WriteString(userName)
	//builder.WriteString("]:")
	//builder.WriteString(content)
	//builder.WriteString("\n---------------------------------------------\n")
	builder.WriteString(reply)
	return builder.String()
}

func FakeIP() string {
	// 随便找的国内IP段：223.64.0.0 - 223.117.255.255
	return fmt.Sprintf("223.%d.%d.%d", r.Intn(54)+64, r.Intn(254), r.Intn(254))
}

func GetFromData(data map[string]string) io.Reader {
	formData := url.Values{}
	for k, v := range data {
		formData.Add(k, v)
	}
	return strings.NewReader(formData.Encode())
}

func GetRandInt64(n int64) int64 {
	return r.Int63n(n)

}
