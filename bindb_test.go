package bindb

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestBINDB(t *testing.T) {
	db, error := LoadDB("./fixtures/test_data.txt", nil)
	assert.Equal(t, nil, error)
	bindbRecord, err := Find(db, "477938")
	assert.Equal(t, nil, err)
	assert.Equal(t, "VISA", bindbRecord.Brand)
	bindbRecord, err = Find(db, "000000")
	assert.Equal(t, errors.New("Couldn't find this BIN number 000000 in the BINDB"), err)
}

func TestListDB(t *testing.T) {
	var dbInfo *DBInfo
	var err error
	apiKey := os.Getenv("BINDB_APIKEY")
	dbInfo, err = ListDB(apiKey)
	assert.Equal(t, nil, err)
	fmt.Printf("%v", dbInfo)
	fmt.Printf("%v", dbInfo.time)
}
