package actions

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
)

func Execute(s *discordgo.Session, event *discordgo.Ready) {
	s.UpdateStatus(0, "with Go")
	fmt.Println("logged in as user " + string(s.State.User.ID))
}
