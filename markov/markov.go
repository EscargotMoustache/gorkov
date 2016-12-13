package markov

import (
	"math/rand"
	"strings"
	"time"
	"os"
	"log"
	"bufio"
)

var PrefixLen int
var MainChain *Chain

type Prefix []string

func (p Prefix) String() string {
	return strings.Join(p, " ")
}

func (p Prefix) Shift(word string) {
	copy(p, p[1:])
	p[len(p) - 1] = word
}

type Chain struct {
	Chain map[string][]string
}

func (c *Chain) Load(f string) {
	file, err := os.Open(f)
	if err != nil {
		log.Println(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		c.Build(scanner.Text())
	}
}

func (c *Chain) Build(s string) {
	p := make(Prefix, PrefixLen)
	for _, v := range strings.Split(s, " ") {
		key := p.String()
		c.Chain[key] = append(c.Chain[key], v)
		p.Shift(v)
	}
}

func (c *Chain) Generate(n int) string {
	p := make(Prefix, PrefixLen)
	var words []string
	for i := 0; i < n; i++ {
		choices := c.Chain[p.String()]
		if len(choices) == 0 {
			break
		}
		next := choices[rand.Intn(len(choices))]
		words = append(words, next)
		p.Shift(next)
	}
	return strings.Join(words, " ")
}

func NewChain() *Chain {
	return &Chain{
		make(map[string][]string),
	}
}

func Init(pLen int) {
	rand.Seed(time.Now().UnixNano())
	MainChain = NewChain()
	PrefixLen = pLen
}
