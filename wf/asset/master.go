package asset

import (
	"bytes"
	"encoding/binary"
	"io"
	"os"
	"strconv"
	"strings"
)

func parseIntMap(input io.Reader) map[int][]string {
	strMap := parseStrMap(input)
	intMap := make(map[int][]string)
	for k, v := range strMap {
		num, err := strconv.Atoi(k)
		if err != nil {
			panic(err)
		}
		intMap[num] = v
	}
	return intMap
}

func parseStrMap(input io.Reader) map[string][]string {
	anyMap := parseAnyMap(input, func(data []byte) []string {
		uncompressed := uncompress(data)
		return strings.Split(string(uncompressed), ",")
	})
	return anyMap
}

func parseAnyMap[T any](input io.Reader, handler func(data []byte) T) map[string]T {
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
	uncompressed := uncompress(header)
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
	ks := make([]string, count)
	vs := make([]T, count)
	//read keys
	var prev int32 = 0
	for i, ko := range kos {
		buf := make([]byte, ko-prev)
		_, err = headerReader.Read(buf)
		if err != nil {
			panic(err)
		}
		ks[i] = string(buf)
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
		vs[i] = handler(buf)
		if err != nil {
			panic(err)
		}
		prev = vo
	}
	table := make(map[string]T)
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
