package youtube

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func randUint(max uint) (uint, error) {
	var min uint
	diff := max - min + 1

	n, err := rand.Int(rand.Reader, big.NewInt(int64(diff)))
	if err != nil {
		return 0, fmt.Errorf("failed to get random int: %w", err)
	}

	return min + uint(n.Int64()), nil
}

func randFloat64(min, max float64) (float64, error) {
	var maxRandomValue uint = 1_000_000

	ratioScale := float64(maxRandomValue)

	randUint, err := randUint(maxRandomValue)
	if err != nil {
		return 0, err
	}

	ratio := float64(randUint) / ratioScale

	return min + ratio*(max-min), nil
}

func randCoordinatesStr() (string, error) {
	var (
		minLat float64 = -90
		maxLat float64 = 90

		minLong float64 = -180
		maxLong float64 = 180
	)

	latitude, err := randFloat64(minLat, maxLat)
	if err != nil {
		return "", err
	}

	longitude, err := randFloat64(minLong, maxLong)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%.4f, %.4f", latitude, longitude), nil
}

func randEnStringRunes(n uint) (string, error) {
	enLetterRunes := []rune(
		"abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	runes := make([]rune, n)

	randRuneIdx, err := randUint(uint(len(enLetterRunes)))
	if err != nil {
		return "", err
	}

	for i := range runes {
		runes[i] = enLetterRunes[randRuneIdx]
	}

	return string(runes), nil
}

func randUaStringRunes(n uint) (string, error) {
	uaLetterRunes := []rune(
		"абвгґдеєжзиіїйклмнопрстуфхцчшщьюяАБВГҐДЕЄЖЗИІЇЙКЛМНОПРСТУФХЦЧШЩЬЮЯ")
	runes := make([]rune, n)

	randRuneIdx, err := randUint(uint(len(uaLetterRunes)))
	if err != nil {
		return "", err
	}

	for i := range runes {
		runes[i] = uaLetterRunes[randRuneIdx]
	}

	return string(runes), nil
}

func randString(strLen uint) (string, error) {
	var fns []func(n uint) (string, error)
	fns = append(fns, randEnStringRunes, randUaStringRunes)

	randLangFn, err := randUint(uint(len(fns)))
	if err != nil {
		return "", err
	}

	return fns[randLangFn](strLen)
}

func randOrder() (string, error) {
	orders := [6]string{
		"date", "rating", "relevance", "title", "videoCount", "viewCount",
	}

	randomIndex, err := randUint(uint(len(orders)))
	if err != nil {
		return "", err
	}

	return orders[randomIndex], nil
}
