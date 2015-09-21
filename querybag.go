package querybag

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
)

// Bag contains all the queries loaded from a specific directory
type Bag map[string]string

// Get allows you to retrieve a specific query from the bag
func (b Bag) Get(key string) string {
	if query, ok := b[key]; ok {
		return query
	}

	panic(fmt.Sprintf("Query not found for '%s'", key))
}

// New constructs a new Bag based on the to be loaded directory
func New(path string) (*Bag, error) {

	dir, err := ioutil.ReadDir(path)

	if err != nil {
		return nil, err
	}

	b := Bag{}

	for _, f := range dir {
		if !f.IsDir() && isSQL(f.Name()) {

			key := sanitizeName(f.Name())
			data, err := ioutil.ReadFile(filepath.Join(path, f.Name()))

			if err != nil {
				return nil, err
			}

			b[key] = string(data)

		}
	}

	return &b, nil
}

func isSQL(name string) bool {

	length := len(name)

	if length < 4 {
		return false
	}

	return name[length-4:] == ".sql"
}

func sanitizeName(name string) string {
	return strings.Replace(name, ".sql", "", -1)
}
