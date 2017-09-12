package yapool

import (
	"github.com/boltdb/bolt"
	"github.com/sirupsen/logrus"
)

type DB interface {
	Add(agentName, staus string)
	Delete(agentName string)
	Read(agentName string) string
}

type db struct {
	boltDB *bolt.DB
}

func GetDB() DB {
	boltDb, err := bolt.Open("yapool.db", 0600, nil)
	if err != nil {
		logrus.Fatalf("yapool.db  create  failed  :  ", err)
	}
	return &db{
		boltDB: boltDb,
	}
}

func (d *db) Add(agentName, status string) {
	d.boltDB.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte("Agent"))
		if err != nil {
			return err
		}
		err = bucket.Put([]byte(agentName), []byte(status))
		if err != nil {
			return err
		}
		return nil
	})
}

func (d *db) Delete(agentName string) {
	d.boltDB.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("Agent"))
		err := bucket.Delete([]byte(agentName))
		if err != nil {
			return err
		}
		return nil
	})
}

func (d *db) Read(agentName string) (status string) {
	d.boltDB.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("Agent"))
		if bucket != nil {
			statusByt := bucket.Get([]byte(agentName))
			status = string(statusByt)
		}
		return nil
	})
	return
}
