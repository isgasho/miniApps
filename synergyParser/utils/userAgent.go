package utils

import (
	"bufio"
	"math/rand"
	"os"
	"time"
)

type UA struct {
	random     *rand.Rand
	userAgents []string
}

func NewRandomUA() *UA {
	src := rand.NewSource(time.Now().Unix())
	rnd := rand.New(src)
	ua := &UA{}
	ua.random = rnd
	return ua
}

func (ua *UA) LoadUserAgents(src string) {
	ua.userAgents = make([]string, 0)
	fd, err := os.Open(src)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(fd)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		ua.userAgents = append(ua.userAgents, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
}

func (ua *UA) GetAllUserAgents() []string {
	return ua.userAgents
}

func (ua *UA) GetRndUserAgent() string {
	return ua.userAgents[ua.random.Intn(len(ua.userAgents))]
}
