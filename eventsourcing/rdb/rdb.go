package rdb

import (
	"github.com/jinzhu/gorm"

	"github.com/xozrc/cqrs/eventsourcing"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type RdbEventStore struct {
	db *gorm.DB
}

func (res *RdbEventStore) Load(partitionKey string, version int64) (es []*eventsourcing.EventEntity, err error) {
	es = make([]*eventsourcing.EventEntity, 0)
	tdb := res.db.Where("partition_key =? && version>?", partitionKey, version).Find(&es)

	if tdb.Error != nil {
		return nil, tdb.Error
	}
	return
}

func (res *RdbEventStore) Save(partitionKey string, events []*eventsourcing.EventEntity) error {

	for _, e := range events {
		e.PartitionKey = partitionKey

		tdb := res.db.Create(e)
		if tdb.Error != nil {
			return tdb.Error
		}
	}
	return nil
}

func NewStore(db *gorm.DB) (res *RdbEventStore, err error) {
	res = &RdbEventStore{}
	res.db = db

	tmpDb := res.db.AutoMigrate(&eventsourcing.EventEntity{})
	if tmpDb.Error != nil {
		return nil, tmpDb.Error
	}

	tmpDb = res.db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&eventsourcing.EventEntity{})
	if tmpDb.Error != nil {
		return nil, tmpDb.Error
	}
	return
}
