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
					URL: "https://github.com/pieza/zephex-bot/blob/master/assets/images/zephex.jpeg?raw=true",
				},
		}
		s.ChannelMessageSendEmbed(m.ChannelID, embed)
	}
}
