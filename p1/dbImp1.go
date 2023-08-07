package main

import (
	"errors"
	"fmt"
	"time"
)

type memRecord struct {
	Text      string
	CreatedAt time.Time
}

type InMemoryDb struct {
	// The int ID to make the simulation
	// of a PostGRE database feel more
	// accurate.
	storage map[int]memRecord
}

// NewInMemory is a test database
// that will log the updated database
// each time an operation is called.
func NewInMemoryDb() *InMemoryDb {
	return &InMemoryDb{
		map[int]memRecord{},
	}
}

func logDB(db InMemoryDb) {
	fmt.Println("\n\nCurrent database:")
	for key, data := range db.storage {
		fmt.Println("Row", key, "Data:", data)
	}
}

func (db *InMemoryDb) Create(r any) error {

	nextId := len(db.storage)
	entry, ok := r.(memRecord)
	if !ok {
		return errors.New("Wrong concrete type passed")
	}
	entry.CreatedAt = time.Now()
	db.storage[nextId] = entry
	logDB(*db)
	return nil
}

func (db *InMemoryDb) Read(q int, r any) error {

	data, ok := db.storage[q]

	// return early if item not present.
	if !ok {
		return errors.New("Record not found")
	}

	applyDataToPointer(data, r)
	return nil
}

func (db *InMemoryDb) Update(q int, r any) error {
	data, ok := db.storage[q]

	// return early if item not present.
	if !ok {
		return errors.New("Record not found")
	}

	data.Text = r.(memRecord).Text
	db.storage[q] = data
	logDB(*db)
	return nil
}

func (db *InMemoryDb) Delete(q int) error {
	// This delete function will only work
	// if you're removing the last element.
	// This is due to the fact that I'm guessing
	// the next ID with this statement:
	//	`nextId := len(db.storage)`
	delete(db.storage, q)
	logDB(*db)
	return nil
}
