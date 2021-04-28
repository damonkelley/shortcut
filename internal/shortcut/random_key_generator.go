package shortcut

import (
	"math/rand"
	"time"
)

type RandomKeyConfig struct {
	Length int
	Seed   int64
}

type randomKeyGenerator struct {
	random *rand.Rand
	length int
}

func (this *randomKeyGenerator) Generate(key string) string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	accumulator := make([]rune, this.length)

	for i := range accumulator {
		accumulator[i] = letters[this.random.Intn(len(letters))]
	}

	return string(accumulator)
}

func DefaultRandomKeyConfig() *RandomKeyConfig {
	return &RandomKeyConfig{Length: 8, Seed: time.Now().UnixNano()}
}

func NewRandomKeyGenerator(config *RandomKeyConfig) KeyGenerator {
	return &randomKeyGenerator{
		length: config.Length,
		random: rand.New(rand.NewSource(config.Seed)),
	}
}
