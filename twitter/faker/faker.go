package faker

import (
	"fmt"
	"hash/maphash"
	"math/rand"

	"github.com/syedwshah/twitter/uuid"
)

func init() {
	rand.New(rand.NewSource(int64(new(maphash.Hash).Sum64())))
}

// Password is hashed password
var Password = "$2a$04$fVXQNwA4cxoLiQN23lnOEucoBaU9h0zfKbzbSYuqFTKovmZsvujlu"

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// private:

func randStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}

	return string(b)
}

func randStringLowerRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes)/2)]
	}

	return string(b)
}

// public:

func RandInt(min, max int) int {
	return rand.Intn(max-min+1) + min
}

func Username() string {
	return randStringRunes(RandInt(2, 10))
}

func ID() string {
	return fmt.Sprintf("%s-%s-%s-%s", randStringRunes(4), randStringRunes(4), randStringRunes(4), randStringRunes(4))
}

func UUID() string {
	return uuid.Generate()
}

func Email() string {
	return fmt.Sprintf("%s@example.com", randStringLowerRunes(RandInt(5, 10)))
}

func RandStr(n int) string {
	return randStringRunes(n)
}
