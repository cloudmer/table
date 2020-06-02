package service

import (
	"fmt"
	"github.com/Shopify/sarama"
	"table/library/logger"
	"table/share"
)

// 执行命令
func execute(session sarama.ConsumerGroupSession, message *sarama.ConsumerMessage)  {
	//log.Printf("Message claimed: offset = %d, key = %s, value = %s, timestamp = %v, topic = %s", message.Offset, string(message.Key), string(message.Value), message.Timestamp, message.Topic)
	logger.Info(fmt.Sprintf("需要消费的数据: offset = %d, key = %s, value = %s, timestamp = %v, topic = %s", message.Offset, string(message.Key), string(message.Value), message.Timestamp, message.Topic))
	logger.Info(fmt.Sprintf("需要消费的数据 value: value = %s", string(message.Value)))

	// 开启事物
	conn, err := share.ShareDb.Begin()
	if err != nil {
		logger.Warning(err.Error())
		return
	}

	fmt.Println("sql: ", string(message.Value))
	_, err = share.ShareDb.Exec(string(message.Value))
	if err != nil {
		logger.Warning(err.Error())
		// 事物回滚
		conn.Rollback()
		logger.Info("事物回滚了")
	}else {
		// 事物提交
		conn.Commit()
		logger.Info("本次SQL执行成功")
	}

	// 消费本次消息
	session.MarkMessage(message, "")
}
