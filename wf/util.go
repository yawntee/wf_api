package wf

import (
	"encoding/base64"
	"fmt"
	"github.com/vmihailenco/msgpack/v5"
	"io"
	"strings"
)

const (
	separator = "----------------"
)

func DebugMsg(title, msg string) {
	fmt.Printf("%s%s%s\n", separator, title, separator)
	fmt.Println(msg)
}

func DebugMsgf(title, format string, a ...any) {
	fmt.Printf("%s%s%s\n", separator, title, separator)
	fmt.Printf(format+"\n", a...)
}

func encodeHeader(header string) string {
	var builder strings.Builder
	for _, ch := range header {
		switch ch {
		case '%':
			builder.WriteString("%25")
		case ' ':
			builder.WriteString("%20")
		case '+':
			builder.WriteString("%2B")
		case '/':
			builder.WriteString("%2F")
		case '?':
			builder.WriteString("%3F")
		case '#':
			builder.WriteString("%23")
		case '&':
			builder.WriteString("%26")
		default:
			builder.WriteByte(byte(ch))
		}

	}
	return builder.String()
}

func wrapMsgpack(data any) []byte {
	marshal, err := msgpack.Marshal(data)
	if err != nil {
		return nil
	}
	return []byte(base64.StdEncoding.EncodeToString(marshal))
}

func unwrapMsgpack(reader *io.ReadCloser, v any) error {
	temp := base64.NewDecoder(base64.StdEncoding, *reader)

	err := msgpack.NewDecoder(temp).Decode(v)
	if err != nil {
		return err
	}
	return nil
}
