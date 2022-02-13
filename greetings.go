package greetings

import (
    "fmt"
    "errors"
)


type N struct {
    A int
    B int
}

type P struct {
    Width, High int
}

type Rectangle interface {
    Circumference() int
    Area() int
}

func Sayhi(name string) (string, error) {
    if name == "" {
        return "", errors.New("empty name")
    }

    message := fmt.Sprintf("Hi, %v. Welcome.", name)
    return message, nil
}

func (v N) Max() int {
    if v.A > v.B {
        return v.A
    }
    return v.B
}

func (p P) Area() int {
    return p.Width * P.High
}

func (p P) Circumference() int {
    return (p.Width + p.High) * 2
}
