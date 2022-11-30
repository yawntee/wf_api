package asset

import (
	"bytes"
	"encoding/binary"
	"io"
	"os"
	"strconv"
	"strings"
)

func intKeyParser(key string) int {
	num, err := strconv.Atoi(key)
	if err != nil {
		panic(err)
	}
	return num
}

func parseIntMap(input io.Reader) map[int][]string {
	return parseAnyMap(input, intKeyParser, func(data []byte) []string {
		uncompressed := UncompressZlib(data)
		return strings.Split(string(uncompressed), ",")
	})
}

func parseStrMap(input io.Reader) map[string][]string {
	return parseAnyMap(input, func(key string) string {
		return key
	}, func(data []byte) []string {
		uncompressed := UncompressZlib(data)
		return strings.Split(string(uncompressed), ",")
	})
}

func parseAnyMap[K int | string, V any](input io.Reader, keyHandler func(key string) K, valueHandler func(data []byte) V) map[K]V {
	//read header's length
	var headerLen int32
	err := binary.Read(input, binary.LittleEndian, &headerLen)
	if err != nil {
		panic(err)
	}
	//read header
	header := make([]byte, headerLen)
	_, err = input.Read(header)
	if err != nil {
		panic(err)
	}
	uncompressed := UncompressZlib(header)
	if err != nil {
		panic(err)
	}
	headerReader := bytes.NewReader(uncompressed)
	//read offsets
	var count int32
	err = binary.Read(headerReader, binary.LittleEndian, &count)
	if err != nil {
		panic(err)
	}
	kos := make([]int32, count)
	vos := make([]int32, count)
	for i := int32(0); i < count; i++ {
		err = binary.Read(headerReader, binary.LittleEndian, &kos[i])
		if err != nil {
			panic(err)
		}
		err = binary.Read(headerReader, binary.LittleEndian, &vos[i])
		if err != nil {
			panic(err)
		}
	}
	ks := make([]K, count)
	vs := make([]V, count)
	//read keys
	var prev int32 = 0
	for i, ko := range kos {
		buf := make([]byte, ko-prev)
		_, err = headerReader.Read(buf)
		if err != nil {
			panic(err)
		}
		ks[i] = keyHandler(string(buf))
		prev = ko
	}
	//read values
	prev = 0
	for i, vo := range vos {
		buf := make([]byte, vo-prev)
		_, err = input.Read(buf)
		if err != nil {
			panic(err)
		}
		vs[i] = valueHandler(buf)
		if err != nil {
			panic(err)
		}
		prev = vo
	}
	table := make(map[K]V)
	for i := 0; i < int(count); i++ {
		table[ks[i]] = vs[i]
	}
	return table
}

func parseStr(path string) []string {
	file, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return strings.Split(string(file), ",")
}
