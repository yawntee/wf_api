package util

import "time"

var loc *time.Location

func init() {
	_loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		panic(err)
	}
	loc = _loc
}

func ParseIso(str string) time.Time {
	t, err := time.ParseInLocation("2006-01-02 15:04:05", str, loc)
	if err != nil {
		panic(err)
	}
	return t
}
