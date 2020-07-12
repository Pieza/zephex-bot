package events

import (
	"fmt"
	"main/db"
	"regexp"
	"strings"

	"github.com/bwmarrin/discordgo"
)

// Ready sets the bot as connected in discord
func Ready(s *discordgo.Session, event *discordgo.Ready) {
	s.UpdateStatus(0, "Manda huevo")
	fmt.Println("logged in as user " + string(s.State.User.ID))
}

// MessageCreate handles the messages received by the bot
func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if strings.Contains(strings.ToLower(m.Content), "manda huevo") {
		embed := &discordgo.MessageEmbed{
			Image: &discordgo.MessageEmbedImage{
				URL: "https://github.com/pieza/zephex-bot/blob/master/assets/images/zephex.jpeg?raw=true",
			},
		}
		s.ChannelMessageSendEmbed(m.ChannelID, embed)
	}

	if strings.Contains(strings.ToLower(m.Content), "gordo") || strings.Contains(strings.ToLower(m.Content), "empresario") {
		embed := &discordgo.MessageEmbed{
			Image: &discordgo.MessageEmbedImage{
				URL: "https://github.com/pieza/zephex-bot/blob/master/assets/images/choche.png?raw=true",
			},
		}
		s.ChannelMessageSendEmbed(m.ChannelID, embed)
	}

	var re = regexp.MustCompile(`(?m)^\$zephex when \w+ then .+$`)
	if re.MatchString(m.Content) {
		words := strings.Fields(m.Content)
		command := words[2]
		url := words[4]
		db.InsertCommand(command, url)
	}

	var memes = db.SelectAllCommands()
	for _, meme := range memes {

		if strings.Contains(strings.ToLower(m.Content), meme.Command) {
			embed := &discordgo.MessageEmbed{
				Image: &discordgo.MessageEmbedImage{
					URL: meme.URL,
				},
			}
			s.ChannelMessageSendEmbed(m.ChannelID, embed)
		}
	}

}
