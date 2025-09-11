package db

import (
	"encoding/binary"
	"time"

	"github.com/boltdb/bolt"
)

// Bucket (table) name is byte slice in BoltDB.
var taskBucket = []byte("tasks")

// package-level BoltDB database, set in Init.
var db *bolt.DB

// How is task represented in memory.
type Task struct {
	Key int
	Value string
}

func Init(dbPath string) error {
	// global since db (below) is package level
	var err error

	// to avoid being stuck if DB is locked.
	db, err = bolt.Open(dbPath, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return err
	}

	// ensure "tasks" bucket is created
	return db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(taskBucket)
		return err
	}) 
}

func CreateTask(task string) (int, error) {
	var id int
	err := db.Update(func(tx *bolt.Tx) error {
		// open the "tasks" bucket
		b := tx.Bucket(taskBucket)

		// get the next automatically assigned id and do conversions
		id64, _ := b.NextSequence()
		id = int(id64) // for return
		key := itob(int(id)) // for BoltDB

		// add the task to the bucket
		return b.Put(key, []byte(task))
	})

	if err != nil {
		return -1, err
	}

	return id, nil
}

func AllTasks() ([]Task, error) {
	var tasks []Task
	err := db.View(func(tx *bolt.Tx) error {
		// open the "tasks" bucket
		b := tx.Bucket(taskBucket)

		// cursor to iterate over the enteries in the bucket
		c := b.Cursor()

		// append the enteries in tasks.
		for k, v := c.First(); k != nil; k, v = c.Next() {
			tasks = append(tasks, Task{
				Key: btoi(k),
				Value: string(v),
			})
		}
		
		return nil
	})

	if err != nil {
		return nil, err
	}

	return tasks, nil
}


func DeleteTask(key int) error { 
	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		return b.Delete(itob(key))
	})
}

func UpdateTask(key int, task string) error {
	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		k := itob(key)
		return b.Put(k, []byte(task))
	})
}

// helper: integer-to-byte (to convert the id64 to []byte to store in DB).
func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

// helper: byte-to-integer
func btoi(b []byte) int {
	return int(binary.BigEndian.Uint64(b))
}