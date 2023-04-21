package main

import (
	"fmt"
)

type Person struct {
	Name   string
	Level  int
	Exp    int
	Health int
	Attack int
}

type Attacker interface {
	Attack(target *Person)
}

type Warrior struct {
	Person
}

func (w *Warrior) Attack(target *Person) {
	fmt.Printf("%s 攻击了 %s\n", w.Name, target.Name)
	target.Health -= w.Person.Attack
	fmt.Printf("%s 剩余血量: %d\n", target.Name, target.Health)
}

func main() {
	warrior1 := &Warrior{
		Person{
			Name:   "warrior1",
			Level:  10,
			Exp:    100,
			Health: 1000,
			Attack: 200,
		},
	}

	warrior2 := &Warrior{
		Person{
			Name:   "warrior2",
			Level:  8,
			Exp:    80,
			Health: 800,
			Attack: 150,
		},
	}

	// 让 warrior1 攻击 warrior2
	warrior1.Attack(&warrior2.Person)
	fmt.Println("开始反击")
	// 让 warrior2 攻击 warrior1
	warrior2.Attack(&warrior1.Person)
}
