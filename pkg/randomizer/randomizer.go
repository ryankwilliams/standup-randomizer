package randomizer

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

// Collection representing a team
type Team struct {
	Members []string
}

// Collection representing a team member
type Member struct {
	Name string
}

// Load the team members from a file on the filesystem
func loadTeamMembers() Team {
	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	// TODO: Make this config filename configurable from the CLI
	content, err := os.ReadFile(currentDir + "/teamMembers.json")
	if err != nil {
		log.Fatal(err)
	}

	team := []Member{}
	finalTeam := Team{}

	_ = json.Unmarshal(content, &team)
	for _, v := range team {
		finalTeam.Members = append(
			finalTeam.Members,
			v.Name,
		)
	}

	return finalTeam
}

// Randomizes the team member list for the daily stand up
func randomizeGuests(team *Team) {
	rand.Seed(time.Now().Unix())
	rand.Shuffle(len(team.Members), func(i, j int) {
		team.Members[i], team.Members[j] =
			team.Members[j], team.Members[i]
	})
}

// Prints the team members in the order for the daily stand up
func printOrder(team *Team) {
	for _, v := range team.Members {
		fmt.Println(v)
	}
}

// Generates the stand up order for the team
// Standup order gets printed to the console output
func GenerateOrder() {
	team := loadTeamMembers()
	randomizeGuests(&team)
	printOrder(&team)
}
