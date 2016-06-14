package rdb_test

import (
	"fmt"
	"sync"
	"testing"
)
import (
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
	"github.com/xozrc/cqrs/eventsourcing"
	. "github.com/xozrc/cqrs/eventsourcing/rdb"
)

var _ = fmt.Print

const (
	dialect  = "mysql"
	user     = "root"
	password = "root"
	host     = "127.0.0.1:3306"
	dbName   = "event_store_test"
	charset  = "utf8"
)

const (
	partitionKey = "test"
)

var (
	version int64 = 0
)

var (
	once sync.Once
	st   *RdbEventStore
)

func setup() {
	once.Do(
		func() {
			dbArgs := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s", user, password, host, dbName, charset)
			s, err := gorm.Open(dialect, dbArgs)
			if err != nil {
				panic(err)
			}
			s.LogMode(true)
			st, err = NewStore(s)
			if err != nil {
				panic(err)
			}
		},
	)

}

func save() error {
	ed := &eventsourcing.EventEntity{}
	version += 1
	ed.Version = version
	eds := make([]*eventsourcing.EventEntity, 0)
	eds = append(eds, ed)
	return st.Save(partitionKey, eds)
}

func load() error {
	setup()
	eds, err := st.Load(partitionKey, 0)
	if err != nil {
		return err
	}
	if len(eds) == 0 {
		return nil
	}

	version = eds[len(eds)-1].Version
	return nil
}

func TestSave(t *testing.T) {
	err := load()

	if !assert.NoError(t, err, "save error") {
		return
	}

	err = save()
	if !assert.NoError(t, err, "save error") {

		return
	}

}
