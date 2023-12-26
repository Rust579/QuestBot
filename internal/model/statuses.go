package model

import "sync"

var pull = Pull{P: map[int64]User{}}

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

func (p *Pull) AddUser(chatId int64, name string) {
	p.Lock()
	p.P[chatId] = User{ChatId: chatId, Name: name}
	p.Unlock()
}

func (p *Pull) IncStage(chatId int64) {
	p.Lock()
	u := p.P[chatId]
	u.Stage++
	p.P[chatId] = u
	p.Unlock()
}

func (p *Pull) GetUser(chatId int64) User {
	p.Lock()
	defer p.Unlock()
	return p.P[chatId]
}
