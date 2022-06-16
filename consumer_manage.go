package main

import (
	"math/rand"
)

func NewconsumerManager() *consumerManager {
	cm := &consumerManager{
		consumers: make(map[string][]Consume),
		data:      make(chan interface{}, 1),
	}

	go func() {
		cm.loopWork()
	}()

	return cm
}

type consumerManager struct {
	consumers map[string][]Consume
	data      chan interface{}
}

func (cm *consumerManager) registerConsumer(queue string, consume Consume) {
	cm.consumers[queue] = append(cm.consumers[queue], consume)
}

func (cm *consumerManager) loopWork() {

	for d := range cm.data {
		for _, v := range cm.consumers {
			// 同组随机消费者
			v[rand.Intn(len(v))].consume(d)
		}

	}

}
