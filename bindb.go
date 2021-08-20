package bindb

import (
	"bufio"
	"fmt"
	"github.com/iancoleman/strcase"
	"os"
	"strconv"
	"strings"
)

type BINDBrecord struct {
	BIN     string
	Brand   string
	Bank    string
	Type    string
	Level   string
	Info    string
	Country string
	WWW     string
	Phone   string
	Address string
	City    string
	Zip     string
}

type DB struct {
	Map map[string]*BINDBrecord
}

func LoadDB(dbpath string, autofix func(string) string) (*DB, error) {
	var f *os.File
	var db = &DB{}
	db.Map = make(map[string]*BINDBrecord)
	var err error
	f, err = os.Open(dbpath)
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	var line string
	var fields []string
	for scanner.Scan() {
		line = scanner.Text()
		if autofix != nil {
			line = autofix(line)
		}
		fields = strings.Split(line, "\t")
		_, err = strconv.ParseInt(fields[0], 10, 32)
		if err != nil || len(fields) < 12 {
			fmt.Printf("BINDB row is not valid: %s\n", line)
			continue
		}
		db.Map[fields[0]] = &BINDBrecord{
			BIN:     fields[0],
			Brand:   strings.ToUpper(fields[1]),
			Bank:    strings.ToUpper(fields[2]),
			Type:    strings.ToUpper(fields[3]),
			Level:   strings.ToUpper(fields[4]),
			Info:    fields[5],
			Country: strings.ToUpper(fields[6]),
			WWW:     fields[7],
			Phone:   fields[8],
			Address: fields[9],
			City:    strcase.ToCamel(fields[10]),
			Zip:     fields[11],
		}
	}
	_ = f.Close()
	return db, err
}

func FixLine(line string) string {
	line = strings.Replace(line, "EG\twww.banqueducaire.com\t16990\t\\", "EG\twww.banqueducaire.com\t16990\t\t\t", 1)
	return line
}

func Printrecord(x BINDBrecord) {
	fmt.Printf("BIN: %s\n", x.BIN)
	fmt.Printf("Brand: %s\n", x.Brand)
	fmt.Printf("Bank: %s\n", x.Bank)
	fmt.Printf("Type: %s\n", x.Type)
	fmt.Printf("Level: %s\n", x.Level)
	fmt.Printf("Info: %s\n", x.Info)
	fmt.Printf("Country: %s\n", x.Country)
	fmt.Printf("WWW: %s\n", x.WWW)
	fmt.Printf("Phone: %s\n", x.Phone)
	fmt.Printf("Address: %s\n", x.Address)
	fmt.Printf("City: %s\n", x.City)
	fmt.Printf("Zip: %s\n", x.Zip)
}
