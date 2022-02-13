package greetings

import (
    "fmt"
    "errors"
)


type N struct {
    m, n int
}

type P struct {
    width, high int
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
    x := v.n
    y := v.m
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
