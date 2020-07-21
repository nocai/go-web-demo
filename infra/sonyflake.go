package infra

import (
	"github.com/golang/glog"
	"github.com/sony/sonyflake"
	"strconv"
)

var sf = sonyflake.NewSonyflake(sonyflake.Settings{})

func NextUint64() (uint64, error) {
	id, err := sf.NextID()
	if err != nil {
		glog.Error(err)
	}
	return id, err
}

func NextString() (string, error) {
	id, err := NextUint64()
	if err != nil {
		return "", err
	}
	return strconv.FormatUint(id, 10), nil
}


