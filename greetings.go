package greetings

import (
    "fmt"
    "errors"
)

func Sayhi(name string) (string, error) {
    if name == "" {
        return "", errors.New("empty name")
    }

    message := fmt.Sprintf("Hi, %v. Welcome.", name)
    return message, nil
}

type Num struct {
    M,N int
}

func (n Num) Max() int {
    m := n.M
    n := n.N

    if m > n {
        return m
    }

    return n
}
