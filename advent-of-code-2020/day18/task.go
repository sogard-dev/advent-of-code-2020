package day18

import (
	"strconv"
	"strings"

	"github.com/sogard-dev/advent-of-code-2020/utils"
)

func strToTokens(line string) []string {
	line = strings.ReplaceAll(line, "(", "( ")
	line = strings.ReplaceAll(line, ")", " )")
	tokens := strings.Split(line, " ")
	return tokens
}

func part1(input string) int {
	sum := 0
	for _, line := range strings.Split(input, "\n") {
		lineSum, _ := evaluateExpressionPart1(strToTokens(line), 0)
		sum += lineSum
	}
	return sum
}

func part2(input string) int {
	sum := 0
	for _, line := range strings.Split(input, "\n") {
		lineSum, _ := evaluateExpressionPart2(strToTokens(line), 0)
		sum += lineSum
	}
	return sum
}

func evaluateExpressionPart1(tokens []string, from int) (int, int) {
	sum := 0

	add := func(i int) {
		sum += i
	}
	mult := func(i int) {
		sum *= i
	}

	op := add

	for i := from; i < len(tokens); i++ {
		t := tokens[i]
		switch t {
		case "(":
			subSum, subIndex := evaluateExpressionPart1(tokens, i+1)
			op(subSum)
			i = subIndex
		case ")":
			return sum, i
		case "+":
			op = add
		case "*":
			op = mult
		default:
			op(utils.GetAllNumbers(t)[0])
		}
	}

	return sum, -1
}

func evaluateExpressionPart2(tokens []string, from int) (int, int) {
	evalutatedTokens := []string{}

	i := from
	cont := true
	for ; i < len(tokens) && cont; i++ {
		t := tokens[i]
		switch t {
		case "(":
			subSum, subIndex := evaluateExpressionPart2(tokens, i+1)
			evalutatedTokens = append(evalutatedTokens, strconv.Itoa(subSum))
			i = subIndex - 1
		case ")":
			cont = false
		default:
			evalutatedTokens = append(evalutatedTokens, t)
		}
	}

	sum := 0
	for len(evalutatedTokens) > 0 {
		if len(evalutatedTokens) == 1 {
			sum = utils.GetAllNumbers(evalutatedTokens[0])[0]
			break
		} else {
			reduce := func(k int, op func(int, int) int) {
				s := op(utils.GetAllNumbers(evalutatedTokens[k-1])[0], utils.GetAllNumbers(evalutatedTokens[k+1])[0])
				newEvalutatedTokens := []string{}
				newEvalutatedTokens = append(newEvalutatedTokens, evalutatedTokens[:k-1]...)
				newEvalutatedTokens = append(newEvalutatedTokens, strconv.Itoa(s))
				newEvalutatedTokens = append(newEvalutatedTokens, evalutatedTokens[k+2:]...)
				evalutatedTokens = newEvalutatedTokens
			}

			for k := 1; k < len(evalutatedTokens); k += 2 {
				if evalutatedTokens[k] == "+" {
					reduce(k, func(a, b int) int {
						return a + b
					})
					k -= 2
				}
			}
			for k := 1; k < len(evalutatedTokens); k += 2 {
				if evalutatedTokens[k] == "*" {
					reduce(k, func(a, b int) int {
						return a * b
					})
					k -= 2
				}
			}
		}
	}

	return sum, i
}
