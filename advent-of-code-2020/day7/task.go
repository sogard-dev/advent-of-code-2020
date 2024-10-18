package day7

import (
	"regexp"
	"slices"
	"strings"

	"github.com/sogard-dev/advent-of-code-2020/utils"
)

type luggage struct {
	amt  int
	name string
}

type rules map[string][]luggage

func parse(input string) rules {
	ret := rules{}
	regName, _ := regexp.Compile(`(\w+) (\w+)`)
	regContains, _ := regexp.Compile(`(\d+) (\w+) (\w+)`)

	for _, line := range strings.Split(input, "\n") {
		bag := regName.FindString(line)
		ret[bag] = []luggage(nil)

		for _, spl := range regContains.FindAllStringSubmatch(line, -1) {
			name := strings.Join(spl[2:], " ")
			num := utils.GetAllNumbers(spl[1])[0]

			ret[bag] = append(ret[bag], luggage{
				name: name,
				amt:  num,
			})
		}
	}

	return ret
}

func part1(input string) int {
	rules := parse(input)
	bag := "shiny gold"

	bags := []string{bag}
	for {
		before := len(bags)

		for k, arr := range rules {
			for _, l := range arr {
				canWePick := slices.Contains(bags, l.name)
				wePickedAlready := !slices.Contains(bags, k)
				if canWePick && wePickedAlready {
					bags = append(bags, k)
				}
			}
		}

		if before == len(bags) {
			break
		}
	}

	return len(bags) - 1
}

func part2(input string) int {
	rules := parse(input)

	var count func(b string) int
	count = func(b string) int {
		cnt := 1
		for _, l := range rules[b] {
			cnt += count(l.name) * l.amt
		}
		return cnt
	}

	return count("shiny gold") - 1
}
