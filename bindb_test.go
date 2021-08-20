package bindb

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBINDB(t *testing.T) {
	db, error := LoadDB("/Users/hegong/Downloads/iindb3/main.txt", FixLine)
	assert.Equal(t, nil, error)
	assert.Equal(t, "KHAN BANK", db.Map["499999"].Bank)
}
