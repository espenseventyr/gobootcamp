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
	vowels       = [10]rune{'A', 'E', 'I', 'O', 'U', 'a', 'e', 'i', 'o', 'u'}
	distribution = make(map[string]int, len(users))
)

func main() {
	for index, user := range users {
		for _, char := range users[index] {
			switch char {
			case 'A', 'a':
				distribution[user]++
			case 'E', 'e':
				distribution[user]++
			case 'I', 'i':
				distribution[user] += 2
			case 'O', 'o':
				distribution[user] += 3
			case 'U', 'u':
				distribution[user] += 4

				/*			for _, vowel := range vowels {
							//				fmt.Printf("index:%v\nuser:%v\nchar:%v\nvowel:%v\n", index, user, char, vowel)
							// For each vowel equals 1 coin
								if char == vowel {
								distribution[user]++
								coins--
							}
						}*/
			}
		}
	}
	fmt.Println(distribution)
	fmt.Println("Coins left:", coins)

}
