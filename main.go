package main

import (
	"os"

	"github.com/andersfylling/disgord"
	"github.com/sirupsen/logrus"
	"github.com/zackartz/cmdlr2"
)

var log = &logrus.Logger{
	Out:       os.Stderr,
	Formatter: new(logrus.TextFormatter),
	Hooks:     make(logrus.LevelHooks),
	Level:     logrus.DebugLevel,
}

func main() {
	client := disgord.New(disgord.Config{
		ProjectName: "julesbot",
		Logger:      log,
		BotToken:    "ODM4MTU1MDkzNDUwNDI0MzYy.YI2-wg.C8w5H2ua6KS8zO3KtFeoQp7L34o",
		DMIntents:   disgord.IntentDirectMessages | disgord.IntentDirectMessageReactions | disgord.IntentDirectMessageTyping,
	})
	defer client.Gateway().StayConnectedUntilInterrupted()

	router := cmdlr2.Create(&cmdlr2.Router{
		Prefixes:         []string{"$"},
		Client:           client,
		BotsAllowed:      false,
		IgnorePrefixCase: true,
	})

	router.RegisterCMD(&cmdlr2.Command{
		Name:        "ping",
		Description: "It pings.",
		Usage:       "ping",
		Example:     "ping",
		Handler: func(ctx *cmdlr2.Ctx) {
			ctx.ResponseText("pong")
		},
	})

	router.RegisterDefaultHelpCommand(client)

	router.Initialize(client)
}
