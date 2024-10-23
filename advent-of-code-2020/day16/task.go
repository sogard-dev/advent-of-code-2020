package day16

import (
	"strings"

	"github.com/sogard-dev/advent-of-code-2020/utils"
)

type rule struct {
	name      string
	intervals []interval
}

type interval struct {
	from int
	to   int
}

func parseRules(rulesString string) []rule {
	rulesBlock := strings.Split(rulesString, "\n")

	rules := []rule{}
	for _, rulesLine := range rulesBlock {
		validTicketRanges := []interval{}

		nums := utils.GetAllNumbers(strings.ReplaceAll(rulesLine, "-", "_"))
		for i := 0; i < len(nums); i += 2 {
			validTicketRanges = append(validTicketRanges, interval{
				from: nums[i],
				to:   nums[i+1],
			})
		}

		rules = append(rules, rule{
			name:      rulesLine[:strings.IndexRune(rulesLine, ':')],
			intervals: validTicketRanges,
		})
	}
	return rules
}

func getErrorRate(rules []rule, ticket []int) int {
	errorRate := 0
	for _, n := range ticket {
		isValid := false
		for _, rule := range rules {
			for _, valid := range rule.intervals {
				if valid.from <= n && valid.to >= n {
					isValid = true
					break
				}
			}
		}

		if !isValid {
			errorRate += n
		}
	}
	return errorRate
}

func part1(input string) int {
	blocks := strings.Split(input, "\n\n")
	//yourTicketLine := strings.Split(blocks[1], "\n")[1]
	nearbyTicketsBlock := strings.Split(blocks[2], "\n")[1:]
	rules := parseRules(blocks[0])

	errorRate := 0
	for _, nearbyTicket := range nearbyTicketsBlock {
		nums := utils.GetAllNumbers(nearbyTicket)
		errorRate += getErrorRate(rules, nums)
	}

	return errorRate
}

func part2(input string) int {
	blocks := strings.Split(input, "\n\n")
	yourTicketLine := utils.GetAllNumbers(strings.Split(blocks[1], "\n")[1])
	nearbyTicketsBlock := strings.Split(blocks[2], "\n")[1:]
	rules := parseRules(blocks[0])

	options := map[int]map[int]bool{}
	//rule => list of potential columns
	for r := range rules {
		c := map[int]bool{}
		for i := range yourTicketLine {
			c[i] = true
		}
		options[r] = c
	}

	for _, nearbyTicket := range nearbyTicketsBlock {
		ticket := utils.GetAllNumbers(nearbyTicket)
		if getErrorRate(rules, ticket) == 0 {
			for c, n := range ticket {
				for r, rule := range rules {

					isValid := false
					for _, valid := range rule.intervals {
						if valid.from <= n && valid.to >= n {
							isValid = true
							break
						}
					}

					if !isValid {
						delete(options[r], c)
					}
				}
			}
		}
	}

	changed := true
	for changed {
		changed = false
		for _, opt := range options {
			if len(opt) == 1 {
				for c := range opt {
					for _, cleanOpt := range options {
						if _, exist := cleanOpt[c]; exist && len(cleanOpt) > 1 {
							delete(cleanOpt, c)
							changed = true
						}
					}
				}
			}
		}
	}

	mult := 1
	for r, c := range options {
		if strings.Contains(rules[r].name, "departure") {
			for v := range c {
				mult *= yourTicketLine[v]
			}
		}
	}
	return mult
}
