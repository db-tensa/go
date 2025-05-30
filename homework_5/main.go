//coded at night, feels boring
package main

import (
	"bufio"   // for grabbin user input
	"fmt"     // for printin stuff
	"os"      // for readin stdin and crashin
	"sort"    // for sortin players
	"strings" // for messin with strings
	"unicode" // for checkin nicknames
)

// Global vars, all player stuff, uga buga
var players []string              // Slice of nicknames
var scores map[string]int         // Map: nick -> rating
var matchesPlayed map[string]int  // Map: nick -> matches
var wins map[string]int           // Map: nick -> wins
var losses map[string]int         // Map: nick -> losses

// Registers player, starts at 1000
func registerPlayer(nickname string) bool {
	// Empty or taken nick? Go touch grass
	if strings.TrimSpace(nickname) == "" || playerExists(nickname) {
		return false
	}
	players = append(players, nickname)
	scores[nickname] = 1000
	matchesPlayed[nickname] = 0
	wins[nickname] = 0
	losses[nickname] = 0
	return true // You in!
}

// Yeets player from system
func removePlayer(nickname string) bool {
	// No player? Read line 12
	if !playerExists(nickname) {
		return false
	}
	// Find and yeet from slice
	for i, name := range players {
		if name == nickname {
			players = append(players[:i], players[i+1:]...)
			break
		}
	}
	// Clean maps, no trash
	delete(scores, nickname)
	delete(matchesPlayed, nickname)
	delete(wins, nickname)
	delete(losses, nickname)
	return true // Poof, gone
}

// Finds player index, -1
func findPlayerIndex(nickname string) int {
	for i, name := range players {
		// Case-insensitive, no dramaaaa-a-a
		if strings.ToLower(name) == strings.ToLower(nickname) {
			return i
		}
	}
	return -1 // What did you forget there ? 
}

// Checks if player real
func playerExists(nickname string) bool {
	return findPlayerIndex(nickname) != -1 // If not -1, they exist
}

// Shows all players
func displayAllPlayers() {
	// No players? Touch grass
	if len(players) == 0 {
		fmt.Println("\r\x1b[31mNo players, what’s that?\x1b[0m")
		return
	}
	fmt.Println("\r\x1b[33mAll players:\x1b[0m")
	fmt.Println("\r", "--------")
	for _, nickname := range players {
		displayPlayerStats(nickname) // Stats 
	}
}

// Updates rating, win or lose
func updateRating(nickname string, won bool, pointsChange int) {
	// No player? Who this?
	if !playerExists(nickname) {
		fmt.Println("\r\x1b[31mError: Player not found\x1b[0m")
		return
	}
	matchesPlayed[nickname]++ // Count match
	if won {
		scores[nickname] += pointsChange // Win points
		wins[nickname]++                 // Add win
	} else {
		scores[nickname] = int(float64(scores[nickname]-pointsChange) * 0.8) // Lose points, 80%, I did in this way because, because this is my program ! 
		losses[nickname]++                                                  // Add loss
	}
	// No negative ratings, sad :(
	if scores[nickname] < 0 {
		scores[nickname] = 0
	}
}

// Gets top N players
func getTopPlayers(count int) []string {
	// No players? No top
	if len(players) == 0 {
		return []string{}
	}
	// Sort by rating
	sorted := sortPlayersByRating()
	// Less than count? Take what we got
	if len(sorted) < count {
		count = len(sorted)
	}
	return sorted[:count] // Top slice
}

// Sorts players by rating
func sortPlayersByRating() []string {
	// Copy slice, don’t mess original
	sortedPlayers := make([]string, len(players))
	copy(sortedPlayers, players)
	// Sort high to low
	sort.Slice(sortedPlayers, func(i, j int) bool {
		return scores[sortedPlayers[i]] > scores[sortedPlayers[j]]
	})
	return sortedPlayers
}

// Gets best player
func getBestPlayer() string {
	// No players? No king
	if len(players) == 0 {
		return "No players, no champ"
	}
	// Sort, grab top
	sorted := sortPlayersByRating()
	return sorted[0] // Top dog
}

// Gets worst player
func getWorstPlayer() string {
	// No players? No loser
	if len(players) == 0 {
		return "No players, no flop"
	}
	// Sort, grab last
	sorted := sortPlayersByRating()
	return sorted[len(sorted)-1] // Sad times
}

// Counts win rate, %
func calculateWinRate(nickname string) float64 {
	// No player or matches? Zero, read line 12
	if !playerExists(nickname) || matchesPlayed[nickname] == 0 {
		return 0
	}
	return (float64(wins[nickname]) / float64(matchesPlayed[nickname])) * 100 // Win %
}

// Shows player stats
func displayPlayerStats(nickname string) {
	// No player? Who this?
	if !playerExists(nickname) {
		fmt.Println("\r\x1b[31mError: Player not found\x1b[0m")
		return
	}
	// outputrs which are generetad by grok because i am lazy
	fmt.Printf("\r\x1b[36mStats for %s:\x1b[0m\n", nickname)
	fmt.Printf("\rRating: %d\n", scores[nickname])
	fmt.Printf("\rMatches: %d\n", matchesPlayed[nickname])
	fmt.Printf("\rWins: %d\n", wins[nickname])
	fmt.Printf("\rLosses: %d\n", losses[nickname])
	fmt.Printf("\rWin rate: %.2f%%\n", calculateWinRate(nickname))
	fmt.Println("\r", "--------")
}

// Shows system stats
func displaySystemStats() {
	// No players? System empty!
	if len(players) == 0 {
		fmt.Println("\r\x1b[31mError: No players, paste broken?\x1b[0m")
		return
	}
	fmt.Println("\r\x1b[33mSystem stats:\x1b[0m")
	fmt.Printf("\rTotal players: %d\n", len(players))
	// Average rating
	total := 0
	for _, nickname := range players {
		total += scores[nickname]
	}
	// also generetaed
	avg := float64(total) / float64(len(players))
	fmt.Printf("\rAverage rating: %.2f\n", avg)
	fmt.Printf("\rBest player: %s (Rating: %d)\n", getBestPlayer(), scores[getBestPlayer()])
	fmt.Printf("\rWorst player: %s (Rating: %d)\n", getWorstPlayer(), scores[getWorstPlayer()])
	totalMatches := 0
	totalWins := 0
	totalLosses := 0
	for _, nickname := range players {
		totalMatches += matchesPlayed[nickname]
		totalWins += wins[nickname]
		totalLosses += losses[nickname]
	}
	// also
	fmt.Printf("\rTotal matches: %d\n", totalMatches)
	fmt.Printf("\rTotal wins: %d\n", totalWins)
	fmt.Printf("\rTotal losses: %d\n", totalLosses)
	fmt.Println("\r", "--------")
}

// Shows menu, keep it cool
func displayMenu() {
	// AGAIN GROK
	fmt.Println("\r\x1b[32m=== Rating System ===\x1b[0m")
	fmt.Println("\r\x1b[34mPick option:\x1b[0m")
	fmt.Println("\r1. Add player")
	fmt.Println("\r2. Yeet player")
	fmt.Println("\r3. Update rating")
	fmt.Println("\r4. Find player")
	fmt.Println("\r5. Show all players")
	fmt.Println("\r6. Top-10 players")
	fmt.Println("\r7. Rating range search")
	fmt.Println("\r8. Player stats")
	fmt.Println("\r9. System stats")
	fmt.Println("\r10. Exit")
}

// Grabs string input
func getStringInput(prompt string) string {
	fmt.Print("\r\x1b[33m", prompt, "\x1b[0m")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return strings.TrimSpace(scanner.Text())
}

// Grabs number input
func getIntInput(prompt string) int {
	var input int
	fmt.Print("\r\x1b[33m", prompt, "\x1b[0m")
	fmt.Scanln(&input)
	return input
}

func main() {
	// Init maps, 
	scores = make(map[string]int)
	matchesPlayed = make(map[string]int)
	wins = make(map[string]int)
	losses = make(map[string]int)

	// Clear screen, make it cool
	fmt.Print("\033[2J\033[H")
	for {
		displayMenu()
		choice := getIntInput("> ") // Grab choice

		// Wanna exit? touch grass 
		if choice == 10 {
			fmt.Println("\r\x1b[31mError! Joke, shutting down...\x1b[0m")
			break
		}

		// yeah yeah i know that switch case better, but i don't like it 
		if choice == 1 {
			// Add new player
			nickname := getStringInput("Enter nick: ")
			// Check nick is letters/digits/spaces
			valid := true
			for _, r := range nickname {
				if !unicode.IsLetter(r) && !unicode.IsDigit(r) && !unicode.IsSpace(r) {
					valid = false
					break
				}
			}
			if !valid || strings.TrimSpace(nickname) == "" {
				fmt.Println("\r\x1b[31mBad nick, crashin!\x1b[0m")
				os.Exit(1) // Babah
			}
			if registerPlayer(nickname) {
				fmt.Printf("\r\x1b[36mPlayer %s added!\x1b[0m\n", nickname)
			} else {
				fmt.Println("\r\x1b[31mNick taken, crashin!\x1b[0m")
				os.Exit(1) // Yeet program
			}
		} else if choice == 2 {
			// Yeet player
			nickname := getStringInput("Enter nick to yeet: ")
			if removePlayer(nickname) {
				fmt.Printf("\r\x1b[36mPlayer %s yeeted!\x1b[0m\n", nickname)
			} else {
				fmt.Println("\r\x1b[31mError: Player not found!\x1b[0m")
			}
		} else if choice == 3 {
			// Update rating
			nickname := getStringInput("Enter nick: ")
			if !playerExists(nickname) {
				fmt.Println("\r\x1b[31mError: Player not found!\x1b[0m")
				continue
			}
			result := getStringInput("Win? (yes/no): ")
			points := getIntInput("Points change: ")
			if points < 0 {
				fmt.Println("\r\x1b[31mPoints bad, crashin!\x1b[0m")
				os.Exit(1) // No retry, just die
			}
			if strings.ToLower(result) == "yes" {
				updateRating(nickname, true, points)
				fmt.Printf("\r\x1b[36mRating updated for %s, nice!\x1b[0m\n", nickname)
			} else if strings.ToLower(result) == "no" {
				updateRating(nickname, false, points)
				fmt.Printf("\r\x1b[36mRating updated for %s, oof!\x1b[0m\n", nickname)
			} else {
				fmt.Println("\r\x1b[31mSay ‘yes’ or ‘no’, crashin!\x1b[0m")
				os.Exit(1) // Crash on bad input
			}
		} else if choice == 4 {
			// Find player
			nickname := getStringInput("Enter nick to find: ")
			if playerExists(nickname) {
				fmt.Printf("\r\x1b[36mPlayer %s found!\x1b[0m\n", nickname)
				displayPlayerStats(nickname)
			} else {
				fmt.Println("\r\x1b[31mError: Player not found!\x1b[0m")
			}
		} else if choice == 5 {
			// Show all players
			displayAllPlayers()
		} else if choice == 6 {
			// Show top-10
			fmt.Println("\r\x1b[33mTop-10 Players:\x1b[0m")
			fmt.Println("\r", "--------")
			top := getTopPlayers(10)
			if len(top) == 0 {
				fmt.Println("\r\x1b[31mNo players, no champs!\x1b[0m")
			} else {
				for i, nickname := range top {
					fmt.Printf("\r%d. %s (Rating: %d)\n", i+1, nickname, scores[nickname])
				}
			}
		} else if choice == 7 {
			// Rating range search, messy in main
			minRating := getIntInput("Min rating: ")
			maxRating := getIntInput("Max rating: ")
			if minRating > maxRating {
				fmt.Println("\r\x1b[31mMin bigger than max, crashin!\x1b[0m")
				os.Exit(1) // No retry, just die
			}
			var result []string
			for _, nickname := range players {
				// Check rating range, no func
				if scores[nickname] >= minRating && scores[nickname] <= maxRating {
					result = append(result, nickname)
				}
			}

			// there is a warning maybe because of my lsp server but idk
			fmt.Println("\r\x1b[33mPlayers in range %d-%d:\x1b[0m", minRating, maxRating)
			if len(result) == 0 {
				fmt.Println("\r\x1b[31mNo players found!\x1b[0m")
			} else {
				for _, nickname := range result {
					fmt.Printf("\r%s (Rating: %d)\n", nickname, scores[nickname])
				}
			}
		} else if choice == 8 {
			// Show player stats
			nickname := getStringInput("Enter nick for stats: ")
			displayPlayerStats(nickname)
		} else if choice == 9 {
			// System stats
			displaySystemStats()
		} else {
			fmt.Println("\r\x1b[31mBad option, crashin!\x1b[0m")
			os.Exit(1) // Crash on invalid choice
		}

		fmt.Println("\r", "--------")
	}
	
	fmt.Println("\r\x1b[32m=== Thanks for  Rating System! ===\x1b[0m")
}
