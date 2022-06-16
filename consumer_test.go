package main

import (
	"fmt"
	"testing"
)

type orderConsumer struct {
}

func (*orderConsumer) consume(data interface{}) {
	if s, ok := data.(string); ok {
		fmt.Println("老子订单处理器", s)
	}
}

type order2Consumer struct {
}

func (*order2Consumer) consume(data interface{}) {
	if s, ok := data.(string); ok {
		fmt.Println("老子订单处理器2", s)
	}
}

type nothingConsumer struct {
}

func (*nothingConsumer) consume(data interface{}) {
	if s, ok := data.(string); ok {
		fmt.Println("老子无所事事", s)
	}
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

func (g *gift) consume(data interface{}) {

	d, ok := data.(gift)
	if ok {
		d.send()
	}
}

func TestConsumer(t *testing.T) {

	cm := NewconsumerManager()
	cm.registerConsumer("order", &orderConsumer{})
	cm.registerConsumer("order", &order2Consumer{})
	cm.registerConsumer("order2", &nothingConsumer{})
	cm.registerConsumer("gift", &gift{})

	cm.data <- "abcdefg"
	cm.data <- gift{id: 1, name: "flower", num: 3}

	// never return
	<-make(chan int)
}
