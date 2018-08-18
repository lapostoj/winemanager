package service_test

import (
	"testing"

	"github.com/lapostoj/winemanager/service/application/service"
	"github.com/stretchr/testify/assert"
)

func TestStringToIntConvertsString(t *testing.T) {
	assert.Equal(t, service.StringToInt("1"), 1)
	assert.Equal(t, service.StringToInt(""), 0)
}

func TestStringToIntPanicsIfInvalidString(t *testing.T) {
	assert.Panics(t, func() { service.StringToInt("a") })
}

func TestEqualsStringSlicesReturnsTrue(t *testing.T) {
	stringSlice1 := []string{"string1", "string2"}
	stringSlice2 := []string{"string1", "string2"}

	assert.True(t, service.EqualsStringSlices(stringSlice1, stringSlice2))
}

func TestEqualsStringSlicesReturnsFalse(t *testing.T) {
	stringSlice1 := []string{"string1", "string2"}
	stringSlice2 := []string{"string1", "string3"}

	assert.False(t, service.EqualsStringSlices(stringSlice1, stringSlice2))
}

func TestEqualsStringSlicesSupportsDifferentLength(t *testing.T) {
	stringSlice1 := []string{"string1", "string2"}
	stringSlice2 := []string{"string1"}

	assert.False(t, service.EqualsStringSlices(stringSlice1, stringSlice2))
}
