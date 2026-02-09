package main

import (
	"fmt"
)

type Speaker interface {
	Say() (string, error)
	WhoIM() (string, error)
}

type Cat struct {
	Name  string
	Voice string
}

func (c Cat) Say() (string, error) {
	if c.Voice == "" {
		return "", fmt.Errorf("voice is missing")
	}

	return "meow", nil
}

func (c Cat) WhoIM() (string, error) {
	if c.Name == "" {
		return "", fmt.Errorf("name is required for pet")
	}

	return "I am a cat named " + c.Name, nil
}

type Dog struct {
	Name  string
	Voice string
}

func (d Dog) Say() (string, error) {
	if d.Voice == "" {
		return "", fmt.Errorf("voice is missing")
	}

	return "woof", nil
}

func (d Dog) WhoIM() (string, error) {
	if d.Name == "" {
		return "", fmt.Errorf("name is required for pet")
	}

	return "I am a dog named " + d.Name, nil
}

type Cow struct {
	Name  string
	Voice string
}

func (c Cow) Say() (string, error) {
	if c.Voice == "" {
		return "", fmt.Errorf("voice is missing")
	}

	return "moo", nil
}

func (c Cow) WhoIM() (string, error) {
	if c.Name == "" {
		return "", fmt.Errorf("name is required for cow")
	}
	return "I am a cow named " + c.Name, nil
}

func MakeSomeNoises(s Speaker) {
	sound, err := s.Say()
	if err != nil {
		fmt.Println("Silence because of error:", err)
		return
	}

	fmt.Println(sound)
}

func main() {
	cat := Cat{Name: "Barsik", Voice: "meow"}
	dog := Dog{Name: "Rex", Voice: "woof"}
	cow := Cow{Name: "Burenka", Voice: "moo"}

	zoo := []Speaker{cat, dog, cow}

	fmt.Println("Animals are speaking:")

	for _, z := range zoo {
		name, errName := z.WhoIM()
		voice, errVoice := z.Say()

		if errName != nil {
			fmt.Println("Data error:", errName)
			continue
		} else if errVoice != nil {
			fmt.Println("Voice error:", errVoice)
		} else {
			fmt.Printf("%s says %s\n", name, voice)
		}
	}
}