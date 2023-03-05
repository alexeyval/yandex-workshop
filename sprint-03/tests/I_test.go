package tests

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
	"time"
)

func TestI_Performance(t *testing.T) {
	// Подготовка данных
	assert := assert.New(t)
	a := time.Now()

	const N = 15000
	const K = 10000
	ints := make([]int, N)

	rand.Seed(time.Now().UnixNano())

	for i := range ints {
		ints[i] = rand.Intn(K)
	}

	// Тестируемая функция
	_ = I(ints, K, 3)

	// Укладываемся в лимит?
	b := time.Now()
	assert.WithinDuration(a, b, 80*time.Millisecond)
}
