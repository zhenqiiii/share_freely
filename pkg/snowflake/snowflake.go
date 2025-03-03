package snowflake

// 常量定义
const (
	timestampBits  = 41                         // 时间戳位数
	machineIDBits  = 10                         // 机器ID位数
	sequenceBits   = 12                         // 序列号位数，以上共63位
	maxMachineID   = -1 ^ (-1 << machineIDBits) // 最大机器ID； << ：左边的数左移n位，n为右边的数
	maxSequenceNum = -1 ^ (-1 << sequenceBits)  // 最大序列号
)

// 结构体定义：该结构体用于存储雪花算法生成ID所需的各个参数
// 1.时间戳 2.机器ID 3.序列号
type snowflake struct {
	timestamp   int64 // 时间戳
	machineID   int64 // 机器ID
	sequenceNum int64 // 序列号
}
