package faker

import (
	"github.com/brianvoe/gofakeit/v6"
	"time"
)

func init() {
	gofakeit.Seed(time.Now().UnixNano())

}
func PersonName() string {
	return gofakeit.Name()
}
