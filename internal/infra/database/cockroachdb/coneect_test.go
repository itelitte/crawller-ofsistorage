package cockroachdb

import (
	"testing"
)

func TestConnect(t *testing.T) {

	t.Run("should connect to database", func(t *testing.T) {
		db := NewCockroachDB()
		if db == nil {
			t.Errorf("expected db to be not nil")
		}
	})

}
