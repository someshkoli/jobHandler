package db

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"runtime"
	"sync"
	"time"
)

// DB - Database connector
type DB struct {
	sync.Mutex
	FileName string `json:"FileName"`
	FilePath string `json:"FilePath"`
}

// NewDB - Creates new instance for db connector
func NewDB(id string) *DB {
	_, fileName, _, _ := runtime.Caller(0)
	dirname := path.Dir(fileName)
	return &DB{
		FileName: id,
		FilePath: path.Join(dirname, "data", id),
	}
}

// Insert - data for a this key
func (D *DB) Insert(data string) error {
	d := []byte(data)
	err := ioutil.WriteFile(D.FilePath, d, 0644)
	if err != nil {
		return err
	}
	fmt.Printf("[DB-WRITE][%s][%s]\n", D.FilePath, time.Now().Format("2 Jan 2006 15:04:05"))
	return nil
}

// Lookup checks for availability of data and if available then returns data
func Lookup(id string) (string, error) {
	_, fileName, _, _ := runtime.Caller(0)
	dirname := path.Dir(fileName)
	filepath := path.Join(dirname, "data", id)
	_, err := os.Stat(filepath)
	if err != nil {
		return "", err
	}

	buf, err := ioutil.ReadFile(filepath)
	data := string(buf)
	return data, err
}

// Delete - Removes a key from storage
func (D *DB) Delete() error {
	return os.Remove(D.FilePath)
}

// Update - Update a record in the db
func (D *DB) Update(data string) error {
	bytes := []byte(data)
	err := os.Truncate(D.FilePath, 0)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(D.FilePath, bytes, 0644)
	return err
}
