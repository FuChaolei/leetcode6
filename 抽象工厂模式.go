//go:build ignore

package main

import (
	"fmt"
)

type abstractApple interface {
	ShowName()
}
type abstractBanana interface {
	ShowName()
}

type ChinaApple struct{}
type ChinaBanana struct{}
type USAApple struct{}
type USABanana struct{}

func (a *ChinaApple) ShowName() {
	fmt.Println("i'm china apple")
}

func (a *ChinaBanana) ShowName() {
	fmt.Println("i'm china banana")
}

func (a *USAApple) ShowName() {
	fmt.Println("i'm usa apple")
}

func (a *USABanana) ShowName() {
	fmt.Println("i'm usa banana")
}

type AbstractFactory interface {
	CreateApple() abstractApple
	CreateBanana() abstractBanana
}

type ChinaFactory struct{}
type USAFactory struct{}

func (c *ChinaFactory) CreateApple() abstractApple {
	return new(ChinaApple)
}
func (c *ChinaFactory) CreateBanana() abstractBanana {
	return new(ChinaBanana)
}
func (c *USAFactory) CreateApple() abstractApple {
	return new(USAApple)
}
func (c *USAFactory) CreateBanana() abstractBanana {
	return new(USABanana)
}

func main() {
	var factory AbstractFactory
	var fruit abstractApple
	factory = new(ChinaFactory)
	fruit = factory.CreateApple()
	fruit.ShowName()
}
