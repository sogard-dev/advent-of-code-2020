package day22

import (
	"strconv"
	"strings"

	"github.com/sogard-dev/advent-of-code-2020/utils"
)

func part1(input string) int {
	players := parse(input)

	for {
		a := players[0].Pop()
		b := players[1].Pop()

		winner := 0
		if a < b {
			winner = 1
		}

		players[winner].Push(max(a, b))
		players[winner].Push(min(a, b))

		if players[1-winner].Len() == 0 {
			return scoreWinner(players, winner)
		}
	}
}

func playersToString(players []utils.Queue[int]) string {
	o := []string{}
	for _, q := range players {
		t := []string{}

		for _, n := range q.Slice() {
			t = append(t, strconv.Itoa(n))
		}
		o = append(o, strings.Join(t, ", "))
	}
	return strings.Join(o, " # ")
}

func part2(input string) int {
	var playGame func([]utils.Queue[int]) (int, int)
	playGame = func(players []utils.Queue[int]) (int, int) {
		seen := map[string]bool{}

		for r := 1; ; r++ {
			seenKey := playersToString(players)
			if _, exist := seen[seenKey]; exist {
				return 0, 0
			}
			seen[seenKey] = true

			for i, v := range players {
				if v.Len() == 0 {
					winner := 1 - i
					return winner, scoreWinner(players, winner)
				}
			}

			a := players[0].Pop()
			b := players[1].Pop()

			winner := 0

			if players[0].Len() >= a && players[1].Len() >= b {
				subPlayers := []utils.Queue[int]{}

				q0 := utils.NewUnboundedQueue[int]()
				for _, n := range players[0].Slice()[0:a] {
					q0.Push(n)
				}
				q1 := utils.NewUnboundedQueue[int]()
				for _, n := range players[1].Slice()[0:b] {
					q1.Push(n)
				}
				subPlayers = append(subPlayers, q0, q1)
				winner, _ = playGame(subPlayers)
				if winner == 0 {
					players[winner].Push(a)
					players[winner].Push(b)
				} else {
					players[winner].Push(b)
					players[winner].Push(a)
				}
			} else {
				if a < b {
					winner = 1
				}

				players[winner].Push(max(a, b))
				players[winner].Push(min(a, b))
			}
		}
	}

	initialPlayers := parse(input)
	_, score := playGame(initialPlayers)
	return score
}

func parse(input string) []utils.Queue[int] {
	blocks := strings.Split(input, "\n\n")

	players := []utils.Queue[int]{}
	for _, block := range blocks {
		q := utils.NewUnboundedQueue[int]()
		for _, num := range utils.GetAllNumbers(block)[1:] {
			q.Push(num)
		}
		players = append(players, q)
	}
	return players
}

func scoreWinner(players []utils.Queue[int], winner int) int {
	sum := 0
	mult := players[winner].Len()
	it := players[winner].Iterator()
	for it.HasNext() {
		sum += mult * it.Next()
		mult--
	}
	return sum
}
