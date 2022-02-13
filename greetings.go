package greetings

import (
    "fmt"
    "errors"
)


type N struct {
    m int
    n int
}

type P struct {
    width int
    high int
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

func (n N) Max() int {
    x := n.n
    y := n.m
    if x > y {
        return x
    }
    return y
}

func (p P) Area() int {
    return p.width * p.high
}

func (p P) Circumference() int {
    return (p.width + p.high) * 2
}
