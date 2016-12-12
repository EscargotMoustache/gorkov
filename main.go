package main

import (
	"crypto/tls"
	"log"
	"strings"

	"github.com/thoj/go-ircevent"
	"./conf"
	"./markov"
)

func main() {
	var err error

	if err = conf.Load("conf.yml"); err != nil {
		log.Fatal(err)
	}

	markov.Init()
	markov.MainChain.Load("bjf.txt")

	ib := irc.IRC(conf.C.BotName, conf.C.BotName)

	if conf.C.TLS {
		ib.UseTLS = true
		if conf.C.InsecureTLS {
			ib.TLSConfig = &tls.Config{InsecureSkipVerify: true}
		}
	}

	if err = ib.Connect(conf.C.Server); err != nil {
		log.Fatal(err)
	}

	ib.AddCallback("001", func (e *irc.Event) {
		log.Printf(conf.C.Channel)
		ib.Join(conf.C.Channel)
	})

	ib.AddCallback("366", func (e *irc.Event) { })

	ib.AddCallback("PRIVMSG", func (e *irc.Event) {
		m := e.Message()
		if strings.HasPrefix(m, "!") {
			if strings.HasPrefix(m, "!bible") {
				ib.Privmsg(conf.C.Channel, markov.MainChain.Generate())
			}
		} else {
			markov.MainChain.Build(m)
		}
	})

	ib.Loop()
}
