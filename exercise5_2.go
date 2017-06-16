// exercise5_2.go
// Distribute coins according to how many vowels their name contains.
// a: 1 coin e: 2 i: 3 o: 4 u: 5 y: semi-vowel...
package main

import "fmt"

var (
	coins = 50
	users = []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie",
		"Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
	}
	distribution = make(map[string]int, len(users))
)

func main() {
	for index, user := range users {
		for _, char := range users[index] {
			switch char {
			case 'A', 'a':
				distribution[user]++
				coins--
			case 'E', 'e':
				distribution[user]++
				coins--
			case 'I', 'i':
				distribution[user] += 2
				coins -= 2
			case 'O', 'o':
				distribution[user] += 3
				coins -= 3
			case 'U', 'u':
				distribution[user] += 4
				coins -= 4
			}
		}
	}

	for _, user := range users {
		if distribution[user] > 10 {
			coins += (distribution[user] - 10)
			distribution[user] = 10
		}
	}

	fmt.Println(distribution)
	fmt.Println("Coins left:", coins)

}
