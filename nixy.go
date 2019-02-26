package main

import (
	/* TODO: Look into alternative logging
	 * libraries, if needed. */
	"log"

	/* IRC client library (callback-based) */
	irc "github.com/thoj/go-ircevent"
)

const (
	/* `realname` IRC param */
	name = "Nixheads IRC bot"
	/* `nick` IRC param */
	nick = "Nixxy"
	/* Server to connect to. */
	server = "chat.freenode.net:6697"
)

func main() {
	/* Define list of channels to join */
	channels := [2]string{
		"##nixheads",
		"##nixheads-bots"}

	/* Create connection instance */
	conn := irc.IRC(nick, name)

	/* Enable TLS on connection. */
	conn.UseTLS = true

	/* Initialise, connect to IRC server */
	log.Println("nixy bot initialised.")
	log.Println("Connecting to IRC server..")
	err := conn.Connect(server)
	if err != nil {
		panic(err)
	}

	/* Define callbacks */

	/* Callback: on connect, join to channels */
	conn.AddCallback("001", func(e *irc.Event) {
		log.Println("Joining channels..")
		for _, c := range channels {
			conn.Join(c)
		}
	})

	/* Callback: on JOIN, log */
	conn.AddCallback("JOIN", func(e *irc.Event) {
		log.Println("Channel joined.")
	})

	/* Callback: on PRIVMSG, log message */
	conn.AddCallback("PRIVMSG", func(e *irc.Event) {
		log.Printf("Message received: %s\n",
			e.Message())
	})

	/* Callback: on INVITE, join, and log event */
	conn.AddCallback("INVITE", func(e *irc.Event) {
		log.Printf("Invited to channel %s, joining..",
			e.Arguments[1])

		conn.Join(e.Arguments[1])
		log.Println("Channel joined. Cause: invite.")
	})

	/* Callback: on KICK, log, and rejoin */
	/* See TODO entry - we should be polite when kicked */
	conn.AddCallback("KICK", func(e *irc.Event) {
		/* TODO: Rate-limiting, and option to disable
		*  rejoining. */
		log.Printf("Rejoining channel: %s\n",
			e.Arguments[0])
		log.Println("We were kicked, rejoining isn't advisable.")

		conn.Join(e.Arguments[0])
	})

	/* Begin connection! */
	conn.Loop()
}
