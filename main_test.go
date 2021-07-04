package main

import (
	"testing"

	"gopkg.in/go-playground/assert.v1"
)

func TestYoda(t *testing.T) {
	t.Run("Can do odd word count", func(t *testing.T) {
		k := ToYodaSpeak("The man ate a creamcake")
		assert.Equal(t, k, "A creamcake the man ate")
	})

	t.Run("Can do even word count", func(t *testing.T) {
		k := ToYodaSpeak("I am a cat")
		assert.Equal(t, k, "A cat i am")
	})
}
