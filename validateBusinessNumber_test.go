package validate

import (
	"github.com/stretchr/testify/assert"
	"testing"
)


func TestValidateBusinessNumber_10자리아닌경우(t *testing.T) {
	//given
	a := "123456789"
	//when
	actual :=  BusinessNumber(a)
	//then
	assert.Equal(t, false, actual)
	assert.Equal(t, len(a), 9)
}


func TestValidateBusinessNumber_유효한번호아닌경우(t *testing.T) {
	//given
	a := "1234567890"
	//when
	actual :=  BusinessNumber(a)
	//then
	assert.Equal(t, false, actual)
}

func TestValidateBusinessNumber_Ok(t *testing.T) {
	//given
	a := "2018121515"
	//when
	actual :=  BusinessNumber(a)
	//then
	assert.Equal(t, true, actual)
}

