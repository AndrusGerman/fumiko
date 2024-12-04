package discord

import (
	"context"
	"fmt"

	"github.com/AndrusGerman/fumiko/internal/core/ports"
	"github.com/bwmarrin/discordgo"

	"go.uber.org/fx"
)

type discord struct {
	config         ports.Config
	socialHandlers []ports.SocialHandler
	s              *discordgo.Session
}

// AddHandlers implements ports.Social.
func (t *discord) AddHandlers(handlers ...ports.SocialHandler) {
	t.socialHandlers = handlers

}

// Register implements ports.Social.
func (t *discord) Register() error {

	return nil
}

func (t *discord) defaulHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Content == "" {
		return
	}
	if m.Author.ID == s.State.User.ID {
		return
	}

	var socialMessage = newSocialMessage(s, m)
	fmt.Println("Received a discord message!", socialMessage.GetText())

	for i := range t.socialHandlers {
		if t.socialHandlers[i].IsValid(socialMessage) {
			go t.socialHandlers[i].Message(socialMessage)
		}
	}

}

func (t *discord) Start(c context.Context) error {
	var err error

	if t.s, err = discordgo.New("Bot " + t.config.GetDiscordToken()); err != nil {
		return err
	}

	t.s.AddHandler(t.defaulHandler)

	return t.s.Open()
}

func (t *discord) Close(c context.Context) error {

	if t.s == nil {
		return nil
	}

	return t.s.Close()
}

func New(lc fx.Lifecycle, config ports.Config) ports.Social {
	var discord = new(discord)
	discord.config = config

	lc.Append(fx.StartHook(discord.Start))
	lc.Append(fx.StopHook(discord.Close))
	return discord
}
