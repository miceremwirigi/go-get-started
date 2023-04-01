package random

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckName(t *testing.T) {
	result := CheckName("Dan")
	assert.True(t, result)

	result = CheckName("Jeff")
	assert.True(t, result)

	result = CheckName("John")
	assert.False(t, result)
}

func TestGetFullName(t *testing.T){
	result, err := GetFullName("","")
	msg := "missing firstName and lastName"
	assert.Equal(t, result, "")
	assert.Equal(t, err.Error(), msg)

	result, err = GetFullName("","Kennedy")
	msg = "missing firstName"
	assert.Equal(t, result, "")
	assert.Equal(t, err.Error(), msg)

	result, err = GetFullName("Jeff","")
	msg = "missing lastName"
	assert.Equal(t, result, "")
	assert.Equal(t, err.Error(), msg)

	result, err = GetFullName("Jeff","Kennedy")
	assert.Equal(t, result, "Jeff Kennedy")
	// assert.Equal(t, err, nil)
	assert.Nil(t, err)
}
