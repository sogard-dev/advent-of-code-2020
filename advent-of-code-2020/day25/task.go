package day25

import (
	"github.com/sogard-dev/advent-of-code-2020/utils"
)

func part1(input string) int {
	nums := utils.GetAllNumbers(input)
	cardPublicKey := nums[0]
	doorPublicKey := nums[1]

	cardSecret := getSecret(7, cardPublicKey)

	return getEncryptionKey(doorPublicKey, cardSecret)
}

func getSecret(subjectNumber int, cardPublicKey int) int {
	subject := 1
	for i := 1; ; i++ {
		if subject == cardPublicKey {
			return i
		}
		subject = (subject * subjectNumber) % 20201227
	}
}

func getEncryptionKey(subjectNumber int, transformations int) int {
	subject := 1
	for i := 1; i < transformations; i++ {
		subject = (subject * subjectNumber) % 20201227
	}
	return subject
}
