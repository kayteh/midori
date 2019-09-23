package main

import (
	"regexp"

	"github.com/bwmarrin/discordgo"
)

func findInUsers(ul []*discordgo.User, pred func(*discordgo.User) bool) *discordgo.User {
	for _, u := range ul {
		if pred(u) {
			return u
		}
	}

	return nil
}

type msgHandler struct {
	matcher *regexp.Regexp
}

func (mh *msgHandler) createRegexpFromUser(u *discordgo.User) {
	mh.matcher = regexp.MustCompile(`^<@!?` + u.ID + `> (.*)`)
}

func (mh *msgHandler) messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.Bot {
		return
	}

	if !mh.matcher.MatchString(m.Message.Content) {
		return
	}

}
