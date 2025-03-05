package snowflake

import "github.com/bwmarrin/snowflake"

// 节点
var node *snowflake.Node

func GenUID() (uid int64) {
	// 设置起始时间和机器ID
	// Init("2025-2-23", 1)

	// 先使用默认的起始时间，节点ID设为1
	node, _ = snowflake.NewNode(1)
	// 返回生成ID
	return node.Generate().Int64()
}
