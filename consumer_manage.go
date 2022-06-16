package main

import (
	"math/rand"
	"reflect"
)

func NewconsumerManager() *consumerManager {
	cm := &consumerManager{
		consumers: make(map[string][]reflect.Value),
		data:      make(chan interface{}, 1),
	}

	go func() {
		cm.loopWork()
	}()

	return cm
}

type consumerManager struct {
	consumers map[string][]reflect.Value
	data      chan interface{}
}

func (cm *consumerManager) registerConsumer(queue string, any interface{}) {
	cm.consumers[queue] = append(cm.consumers[queue], reflect.ValueOf(any))
}

func (cm *consumerManager) loopWork() {

	for d := range cm.data {
		for _, v := range cm.consumers {
			// 同组随机消费者
			cs := v[rand.Intn(len(v))]
			consumerMethod := cs.MethodByName("Consume")
			consumerMethod.Call([]reflect.Value{reflect.ValueOf(d)})
		}

	}

}
