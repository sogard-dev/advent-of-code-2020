package day19

import (
	"fmt"
	"regexp"
	"strings"
)

type rule interface {
	evaluate(rules *map[string]rule, depth int) string
}

type or struct {
	left, right []string
}

type single struct {
	s []string
}

type literal struct {
	v string
}

func (s single) evaluate(rules *map[string]rule, depth int) string {
	str := ""
	for _, l := range s.s {
		str += (*rules)[l].evaluate(rules, depth+1)
	}
	return str
}

func (o or) evaluate(rules *map[string]rule, depth int) string {
	if depth > 100 {
		return ""
	}

	leftSide := ""
	for _, l := range o.left {
		leftSide += (*rules)[l].evaluate(rules, depth+1)
	}

	rightSide := ""
	for _, r := range o.right {
		rightSide += (*rules)[r].evaluate(rules, depth+1)
	}

	return "(" + leftSide + "|" + rightSide + ")"
}

func (l literal) evaluate(_ *map[string]rule, _ int) string {
	return l.v
}

func part1(input string) int {
	return solve(input)
}

func part2(input string) int {
	input = strings.ReplaceAll(input, "8: 42", "8: 42 | 42 8")
	input = strings.ReplaceAll(input, "11: 42 31", "11: 42 31 | 42 11 31")
	return solve(input)
}

func solve(input string) int {
	blocks := strings.Split(input, "\n\n")
	rulesBlock := blocks[0]

	rules := map[string]rule{}
	for _, line := range strings.Split(rulesBlock, "\n") {
		spl := strings.Split(line, ": ")
		ruleNo := spl[0]
		ruleStr := spl[1]
		if strings.Contains(ruleStr, "|") {
			sides := strings.Split(ruleStr, " | ")
			rules[ruleNo] = or{
				left:  strings.Split(sides[0], " "),
				right: strings.Split(sides[1], " "),
			}
		} else if strings.Contains(ruleStr, `"`) {
			rules[ruleNo] = literal{
				v: ruleStr[len(ruleStr)-2 : len(ruleStr)-1],
			}
		} else {
			rules[ruleNo] = single{
				s: strings.Split(ruleStr, " "),
			}
		}

	}

	eval0 := rules["0"].evaluate(&rules, 0)
	fmt.Println("Rule 0: ", eval0)
	r, _ := regexp.Compile("^" + eval0 + "$")

	sum := 0
	messagesBlock := blocks[1]
	for _, line := range strings.Split(messagesBlock, "\n") {
		if r.MatchString(line) {
			fmt.Println(line)
			sum += 1
		}
	}

	return sum
}
