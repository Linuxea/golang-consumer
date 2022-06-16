package main

import (
	"fmt"
	"testing"
)

type orderConsumer struct {
}

func (*orderConsumer) Consume(data string) {
	fmt.Println("老子订单处理器", data)
}

type order2Consumer struct {
}

func (*order2Consumer) Consume(data string) {
	fmt.Println("老子订单处理器2", data)
}

type nothingConsumer struct {
}

func (*nothingConsumer) Consume(data string) {
	fmt.Println("老子无所事事", data)
}

type gift struct {
	id   int
	name string
	num  int
}

func (g *gift) send() {
	s := fmt.Sprintf("将礼物%d送给%s%d份", g.id, g.name, g.num)
	fmt.Println(s)
}

func (g *gift) Consume(data gift) {
	data.send()
}

func TestConsumer(t *testing.T) {

	cm := NewconsumerManager()
	cm.registerConsumer("order", &orderConsumer{})
	cm.registerConsumer("order", &order2Consumer{})
	cm.registerConsumer("order2", &nothingConsumer{})
	// cm.registerConsumer("gift", &gift{})

	cm.data <- "abcdefg"
	// cm.data <- gift{id: 1, name: "flower", num: 3}

	// never return
	<-make(chan int)
}
