package main

import (
	"fmt"
	"math/rand"
	"time"
)

func GenerateID() string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("TXN-%06d", rand.Intn(999999))
}
