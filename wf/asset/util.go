package asset

import (
	"bytes"
	"compress/zlib"
	"wf_api/util"
)

func uncompress(b []byte) []byte {
	reader, err := zlib.NewReader(bytes.NewReader(b))
	if err != nil {
		panic(err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(reader)
	if err != nil {
		panic(err)
	}
	return buf.Bytes()
}

func hashPath(path string) string {
	return util.Sha1([]byte(path + "K6R9T9Hz22OpeIGEWB0ui6c6PYFQnJGy"))
}
