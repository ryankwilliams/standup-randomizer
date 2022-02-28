package randomizer

import (
	"fmt"
	"math/rand"
	"time"
)

type TeamMembers struct {
	Members []string
}

func loadTeamMembers() TeamMembers {
	// TODO: Load from filesystem
	teamMembers := TeamMembers{
		Members: []string{"User1", "User2", "User3", "User4", "User5"},
	}
	return teamMembers
}

func randomizeGuests(teamMembers *TeamMembers) {
	rand.Seed(time.Now().Unix())
	rand.Shuffle(len(teamMembers.Members), func(i, j int) {
		teamMembers.Members[i], teamMembers.Members[j] =
			teamMembers.Members[j], teamMembers.Members[i]
	})
}

func printOrder(teamMembers *TeamMembers) {
	for _, v := range teamMembers.Members {
		fmt.Println(v)
	}
}

func GenerateOrder() {
	teamMembers := loadTeamMembers()
	randomizeGuests(&teamMembers)
	printOrder(&teamMembers)
}
