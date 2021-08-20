package bindb

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBINDB(t *testing.T) {
	db, error := LoadDB("./fixtures/test_data.txt", FixLine)
	assert.Equal(t, nil, error)
	assert.Equal(t, "VISA", db.Map["477938"].Brand)
}
