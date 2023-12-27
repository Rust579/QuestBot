package model

import (
	"sync"
)

var PullUsers = Pull{P: map[int64]User{}}

type Pull struct {
	sync.Mutex
	P     map[int64]User
	Stage int
}

type User struct {
	ChatId int64
	Name   string
	//Team   int
}

func (p *Pull) AddUser(chatId int64, name string, stage int) {
	p.Lock()
	p.P[chatId] = User{ChatId: chatId, Name: name}
	p.Stage = stage
	p.Unlock()
}

func (p *Pull) IncStage(stage int) {
	p.Lock()
	p.Stage = stage
	p.Unlock()
}

func (p *Pull) GetUser(chatId int64) User {
	p.Lock()
	defer p.Unlock()
	return p.P[chatId]
}
