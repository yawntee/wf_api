package internal

import (
	"encoding/base64"
	"github.com/vmihailenco/msgpack/v5"
	"io"
	"strings"
)

func EncodeHeader(header string) string {
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

func WrapMsgpack(data any) []byte {
	marshal, err := msgpack.Marshal(data)
	if err != nil {
		return nil
	}
	return []byte(base64.StdEncoding.EncodeToString(marshal))
}

func UnwrapMsgpack(reader *io.ReadCloser, v any) error {
	temp := base64.NewDecoder(base64.StdEncoding, *reader)

	err := msgpack.NewDecoder(temp).Decode(v)
	if err != nil {
		return err
	}
	return nil
}
