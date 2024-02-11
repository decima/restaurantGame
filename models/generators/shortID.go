package generators

import (
	"math/rand"
	"strings"
	"time"
)

type shortIDGenerator struct{}

func (s *shortIDGenerator) Generate(isValid func(probe string) bool, withPrefix ...string) string {

	i := 3
	for {
		id := s.make(i)
		if len(withPrefix) > 0 {
			id = strings.Join(withPrefix, "") + id
		}
		if isValid(id) {
			return id
		}
		i++
	}
}

func (s *shortIDGenerator) make(iteration int) string {
	//generate a short string of length iteration
	randomizer := rand.New(rand.NewSource(time.Now().UnixNano()))
	//characters to be used, expect easily confusable ones
	chars := "0123456789abcdefghjkmpqrtuvwxyzABCDEFGHJKMPQRTUVWXYZ_"
	length := len(chars)
	id := ""
	for len(id) < iteration {
		id += string(chars[randomizer.Intn(length)])
	}

	return id

}

var ShortID = shortIDGenerator{}
