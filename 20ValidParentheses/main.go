package main

import "fmt"

type Stack []string

func (s *Stack) Push(value string) {
    *s = append(*s, value)
}

func (s *Stack) Pop() (string, error) {
    if len(*s) == 0 {
        return "", fmt.Errorf("Stack is empty")
    }
    index := len(*s) - 1
    value := (*s)[index]
    *s = (*s)[:index]
    return value, nil
}

func isValid(s string) bool {
    stack := Stack{}
    for _, char := range s {
        char_str := string(char) 
        if char_str == "(" || char_str == "[" || char_str == "{" {
            stack.Push(char_str)
        } else {
            if len(stack) == 0 {
                return false
            }
            last_char, _ := stack.Pop()
            if (char_str == ")" && last_char != "(") || (char_str == "]" && last_char != "[") || (char_str == "}" && last_char != "{") {
                return false
            }
        }
    }
    return len(stack) == 0
}