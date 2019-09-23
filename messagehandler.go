package main

import (
	"regexp"

	"github.com/bwmarrin/discordgo"
	"github.com/kayteh/midori/chatops"
)

func findInUsers(ul []*discordgo.User, pred func(*discordgo.User) bool) *discordgo.User {
	for _, u := range ul {
		if pred(u) {
			return u
		}
	}

	return nil
}

type messageHandler struct {
	matcher *regexp.Regexp
	chatops *chatops.ChatOpsProvider
}

func (mh *messageHandler) createRegexpFromUser(u *discordgo.User) {
	mh.matcher = regexp.MustCompile(`^<@!?` + u.ID + `>`)
}

func (mh *messageHandler) Handle(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.Bot {
		return
	}

	if mh.matcher == nil {
		return
	}

	if m.GuildID == "" && !mh.matcher.MatchString(m.Content) {
		return
	}

}
