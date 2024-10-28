package day21

import (
	"maps"
	"slices"
	"strings"
)

func part1(input string) int {
	allIngredients, allAllergens := solve(input)

	for _, i := range allAllergens {
		delete(allIngredients, i)
	}

	sum := 0
	for _, v := range allIngredients {
		sum += v
	}

	return sum
}

func part2(input string) string {
	_, allAllergens := solve(input)

	allergensSorted := slices.Collect(maps.Keys(allAllergens))
	slices.Sort(allergensSorted)

	canon := []string{}
	for _, a := range allergensSorted {
		canon = append(canon, allAllergens[a])
	}

	return strings.Join(canon, ",")
}

func solve(input string) (map[string]int, map[string]string) {
	input = strings.ReplaceAll(input, ")", "")

	allIngredients := map[string]int{}
	allAllergens := map[string]map[string]bool{}

	for _, line := range strings.Split(input, "\n") {
		spl := strings.Split(line, " (contains ")

		ingredients := strings.Split(spl[0], " ")
		for _, v := range ingredients {
			allIngredients[v] += 1
		}

		allergens := strings.Split(spl[1], ", ")
		for _, a := range allergens {
			if prevIngredients, exist := allAllergens[a]; !exist {
				allAllergens[a] = map[string]bool{}

				for _, v := range ingredients {
					allAllergens[a][v] = true
				}
			} else {
				for prev := range prevIngredients {
					if !slices.Contains(ingredients, prev) {
						delete(allAllergens[a], prev)
					}
				}
			}
		}
	}

	changed := true
	for changed {
		changed = false
		for _, ingredients := range allAllergens {
			if len(ingredients) == 1 {
				for i := range ingredients {
					for _, otherIngredients := range allAllergens {
						if _, exist := otherIngredients[i]; exist && len(otherIngredients) > 1 {
							delete(otherIngredients, i)
							changed = true
						}
					}
				}
			}
		}
	}

	allergenMapping := map[string]string{}

	for a, ingredients := range allAllergens {
		if len(ingredients) != 1 {
			panic("ohh noes")
		}
		for i := range ingredients {
			allergenMapping[a] = i
		}
	}
	return allIngredients, allergenMapping
}
