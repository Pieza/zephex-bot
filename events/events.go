package events

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
)

func Ready(s *discordgo.Session, event *discordgo.Ready) {
	s.UpdateStatus(0, "Manda huevo")
	fmt.Println("logged in as user " + string(s.State.User.ID))
}

func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Content == "manda huevo" {
		embed := &discordgo.MessageEmbed{
		    Image: &discordgo.MessageEmbedImage {
					URL: "https://cdn.discordapp.com/avatars/119249192806776836/cc32c5c3ee602e1fe252f9f595f9010e.jpg?size=2048",
				},
		}
		s.ChannelMessageSendEmbed(m.ChannelID, embed)
	}
}
