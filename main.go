package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	//fmt.Println("执行项目")
	//config := sarama.NewConfig()
	//config.Producer.RequiredAcks = sarama.WaitForAll          // 让leader 和 follower都确认
	//config.Producer.Partitioner = sarama.NewRandomPartitioner // 新选出一个partition
	//config.Producer.Return.Successes = true                   // 成功交付的消息将在success channel返回
	//
	//msg := &sarama.ProducerMessage{}
	//msg.Topic = "bai"
	//msg.Value = sarama.StringEncoder("love")
	//
	//producer, err := sarama.NewSyncProducer([]string{"121.196.163.8:9092"}, config)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println("连接成功")
	//
	//defer producer.Close()
	//pid, offset, err := producer.SendMessage(msg)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Printf("pid : %v,offset: %v\n\n", pid, offset)
	//fmt.Println("执行完成")
	r := gin.Default()
	// 2.绑定路由规则，执行的函数
	// gin.Context，封装了request和response
	r.GET("/he", func(c *gin.Context) {
		c.String(http.StatusOK, "hello World!")
	})
	// 3.监听端口，默认在8080
	// Run("里面不指定端口号默认为8080")
	r.Run(":8000")

}
