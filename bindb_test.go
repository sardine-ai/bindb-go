package bindb

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBINDB(t *testing.T) {
	db, error := LoadDB("./fixtures/", nil)
	assert.Equal(t, nil, error)
	bindbRecord, error := Find(db, "238006")
	assert.Equal(t, nil, error)
	assert.Equal(t, bindbRecord.Brand, "MASTERCARD")
	assert.Equal(t, bindbRecord.Level, "")
	assert.Equal(t, nil, error)
	bindbRecord, err := Find(db, "477938")
	assert.Equal(t, nil, err)
	assert.Equal(t, "VISA", bindbRecord.Brand)
	bindbRecord, err = Find(db, "000000")
	assert.Equal(t, errors.New("Couldn't find this BIN number 000000 in the BINDB"), err)
}
