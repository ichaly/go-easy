package base

import (
	"github.com/ichaly/go-easy/base/utils"
	"github.com/sony/sonyflake"
	"strings"
)

var sf *sonyflake.Sonyflake

func init() {
	st := sonyflake.Settings{}
	// machineID是个回调函数
	st.MachineID = getMachineID
	sf = sonyflake.NewSonyflake(st)
}

// 模拟获取本机的机器ID
func getMachineID() (mID uint16, err error) {
	result := strings.Join(utils.GetMAC(), ",")
	mID = uint16(utils.HashCode(result) % 10)
	return
}

func GenerateID() (uint64, error) {
	return sf.NextID()
}
