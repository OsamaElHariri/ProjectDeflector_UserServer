package users

import (
	"crypto/md5"
	"encoding/binary"
	"math/rand"
	"time"

	"github.com/lucasb-eyer/go-colorful"
)

func getPlayerColors(playerId string, colorCount int) []string {
	hashGen := md5.New()
	hashGen.Write([]byte(playerId + time.Now().Format("Jan 2 2006")))
	var seed uint64 = binary.BigEndian.Uint64(hashGen.Sum(nil))
	rand.Seed(int64(seed))

	colors := make([]string, colorCount)
	for i := 0; i < colorCount; i++ {
		color := getValidColor()
		colors[i] = color.Hex()
	}
	return colors
}

func getValidColor() (c colorful.Color) {
	for c = getColor(); !c.IsValid(); c = getColor() {
	}
	return c
}

func getColor() colorful.Color {
	return colorful.Hcl(rand.Float64()*360.0,
		0.4+rand.Float64()*0.4,
		0.4+rand.Float64()*0.2)
}
