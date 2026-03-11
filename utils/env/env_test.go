package env

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetEnv(t *testing.T){
	t.Run("Should Return enviroment default", func(t *testing.T){
		defaultValue := "GOLANG"
		enviroment := "PROGRAM"
		assert.Equal(t, GetEnv(enviroment, defaultValue), defaultValue)
	})

	t.Run("Should Return enviroment default", func(t *testing.T) {
		defaultValue := ""
		environment := "HOME"
		assert.NotEmpty(t, GetEnv(environment, defaultValue))
	})
}
