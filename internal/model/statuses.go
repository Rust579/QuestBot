package model

import (
	"sync"
)

var PullUsers = Pull{P: map[int64]User{}}

type Pull struct {
	sync.Mutex
	P map[int64]User
}

type User struct {
	ChatId int64
	Name   string
	Stage  int
	//Team   int
}

func (p *Pull) AddUser(chatId int64, name string, stage int) {
	p.Lock()
	p.P[chatId] = User{ChatId: chatId, Name: name, Stage: stage}
	p.Unlock()
}

func (p *Pull) IncStage(chatId int64, stage int) {
	p.Lock()
	u, ok := p.P[chatId]
	if ok {
		u.Stage = stage
		p.P[chatId] = u
	}
	p.Unlock()
}

func (p *Pull) GetUser(chatId int64) User {
	p.Lock()
	defer p.Unlock()
	return p.P[chatId]
}
