# gorkov
Gorkov is a Go IRC bot that learns from people chatting, and generates Markov chains when asked to. (Still pretty stupid right now, but heh, works anyway.)

Gorkov depends on `yaml.v2` and `go-ircevent`.

The Markov chain generation is shamelessly stolen from [this page](https://golang.org/doc/codewalk/markov/).
