package model

import (
	"errors"
	"sync"
)

var PullUsers = Pull{P: map[string]int64{}}

type Pull struct {
	sync.Mutex
	P     map[string]int64
	Stage int
}

type SendTo struct {
	ChatId int64
	Msg    string
}

func (p *Pull) AddUser(chatId int64, name string, stage int) {
	p.Lock()
	p.P[name] = chatId
	p.Stage = stage
	p.Unlock()
}

func (p *Pull) IncStage(stage int) {
	p.Lock()
	p.Stage = stage
	p.Unlock()
}

func (p *Pull) GetUser(name string) (int64, error) {
	p.Lock()
	defer p.Unlock()
	usId, ok := p.P[name]
	if !ok {
		return 0, errors.New("Нету юзера")
	}
	return usId, nil
}

func (p *Pull) GetAllUserIds() []int64 {

	p.Lock()
	defer p.Unlock()

	var userIds []int64
	for _, u := range p.P {
		userIds = append(userIds, u)
	}
	return userIds
}
