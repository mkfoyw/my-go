package pool

import (
	"errors"
	"math/rand"
	"time"
)

var ErrAlreadyPopAll = errors.New("all has already been poped")

type Pool struct {
	totalCount      int
	alreadyPopCount int

	pool []int
}

func NewPool(totalCount int) *Pool {
	return &Pool{
		totalCount:      totalCount,
		alreadyPopCount: 0,
		pool:            make([]int, totalCount),
	}
}

// 从奖池获中随机的抽取一个物品（编号）
func (p *Pool) GetIndex() (int, error) {
	if p.alreadyPopCount >= p.totalCount {
		return 0, ErrAlreadyPopAll
	}
	rd := p.getRandomInt()
	randomIndex := rd%(p.totalCount-p.alreadyPopCount) + p.alreadyPopCount
	number := p.getNumber(randomIndex)
	number2 := p.getNumber(p.alreadyPopCount)

	p.pool[randomIndex] = number2
	p.alreadyPopCount++
	return number, nil
}

// 返回该编号指向的物品编号
func (p *Pool) getNumber(index int) int {
	// 物品还没有被抽取。
	if p.pool[index] == 0 {
		return index + 1
	}
	return p.pool[index]
}

// 随机的生成一个整数
func (p *Pool) getRandomInt() int {
	rand.Seed(time.Now().Unix())
	return rand.Int()
}
