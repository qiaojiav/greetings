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

func (v Num) Max() int {
    if v.M > v.N {
        return v.M
    }
    return v.N
}
