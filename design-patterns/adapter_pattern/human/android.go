package human

import (
	"fmt"
	"math/rand"
)

var words = []string{"gopher", "golang", "compiler", "concurrency", "channel", "pointer"}

type androidSound string

type Android struct {
	ID string
}

func (a *Android) MoveLeft() {
	fmt.Printf("%s moved left\n", a.ID)
}

func (a *Android) MoveRight() {
	fmt.Printf("%s moved right\n", a.ID)
}

func (a *Android) MoveForward() {
	fmt.Printf("%s moved forward\n", a.ID)
}

func (a *Android) MoveBack() {
	fmt.Printf("%s moved back\n", a.ID)
}

func (a *Android) ThinkWord() string {
	word := words[rand.Intn(len(words))]
	fmt.Printf("%s thought: %s\n", a.ID, word)
	return word
}

func (a *Android) TextToSound(txt string) androidSound {
	fmt.Printf("%s converted text to sound: %s\n", a.ID, txt)
	return androidSound(txt)
}

func (a *Android) ReproduceSound(sound androidSound) {
	fmt.Printf("%s reproduced sound: %s\n", a.ID, sound)
}
