package text

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/roleypoly/discord/responders"
)

var textResponder = &responders.Responder{
	Commands: []responders.Command{{
		Match: regexp.MustCompile(`do ([a-zA-Z0-9-]{1,38}/[a-zA-Z0-9-_\.]+) (\w+)`),
		Handler: func(matches [][]string, s *discordgo.Session, m *discordgo.MessageCreate) string {
			// err := chatops.Do(matches[0][0], matches[0][1])
			// if err != nil {
			// 	if err == chatops.ErrNoPermissions {
			// 		return "Please invite `midorichatops` to have contributor access to your repository, then send `join your-name/repo`."
			// 	}
			// 	return "An error occurred. Please yell."
			// }
		},
	}},
}

func TextResponder(s *discordgo.Session, m *discordgo.MessageCreate) {
	myID := s.State.User.ID
	msg := strings.Replace(m.Content, fmt.Sprintf(`<@%s>`, myID), "", 1)
	msg = strings.Replace(m.Content, fmt.Sprintf(`<@!%s>`, myID), "", 1)

	if !textResponder.FindAndExecute(msg, s, m) {
		SendDefaultResponse(s, m)
	}
}

func SendDefaultResponse(s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSend(m.ChannelID, `I'm not sure about that.`)
}
