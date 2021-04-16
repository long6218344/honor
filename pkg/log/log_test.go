package log

import "testing"

func TestErrorf(t *testing.T) {
	t.Run("errorf", func(t *testing.T) {
		Errorf("lebron has %d nba champions", 4)
	})
}
