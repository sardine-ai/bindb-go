package bindb

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/iancoleman/strcase"
	"os"
	"strconv"
	"strings"
)

type Record struct {
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
	Map map[string]*Record
}

func LoadDB(path string, autofix func(string) string) (*DB, error) {
	var err error

	if path[len(path)-1] != '/' {
		path = path + "/"
	}
	var db = &DB{}
	db.Map = make(map[string]*Record)
	err = LoadMain(db, path+"main.txt", autofix)
	if err != nil {
		return nil, err
	}
	err = LoadMulti(db, path+"multi.txt", autofix)
	if err != nil {
		return nil, err
	}
	return db, err
}

func LoadMain(db *DB, dbpath string, autofix func(string) string) error {
	var f *os.File
	var err error
	f, err = os.Open(dbpath)
	if err != nil {
		return err
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
		db.Map[fields[0]] = BuildRecord(fields)
	}
	_ = f.Close()
	return nil
}

func LoadMulti(db *DB, path string, autofix func(string) string) (err error) {
	var f *os.File
	f, err = os.Open(path)
	if err != nil {
		return err
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
			fmt.Printf("BINDB row is not valid: %s, fields number is %d\n", line, len(fields))
			continue
		}
		newRecord := BuildRecord(fields)
		if item, exist := db.Map[fields[0]]; exist {
			if item.Brand != newRecord.Brand {
				item.Brand = ""
			}
			if item.Bank != newRecord.Bank {
				item.Bank = ""
			}
			if item.Type != newRecord.Type {
				item.Type = ""
			}
			if item.Level != newRecord.Level {
				item.Level = ""
			}
			if item.Info != newRecord.Info {
				item.Info = ""
			}
			if item.Country != newRecord.Country {
				item.Country = ""
			}
			if item.WWW != newRecord.WWW {
				item.WWW = ""
			}
			if item.Phone != newRecord.Phone {
				item.Phone = ""
			}
			if item.Address != newRecord.Address {
				item.Address = ""
			}
			if item.City != newRecord.City {
				item.City = ""
			}
			if item.Zip != newRecord.Zip {
				item.Zip = ""
			}
		} else {
			db.Map[fields[0]] = newRecord
		}
	}
	_ = f.Close()
	return nil
}

func Find(db *DB, bin string) (bindbRecord *Record, err error) {
	var ok bool
	if bindbRecord, ok = db.Map[bin]; ok {
		return db.Map[bin], nil
	} else {
		return nil, errors.New(fmt.Sprintf("Couldn't find this BIN number %s in the BINDB", bin))
	}
}

func Printrecord(x Record) {
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

func BuildRecord(fields []string) *Record {
	return &Record{
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
