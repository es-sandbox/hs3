package main

import (
	"github.com/boltdb/bolt"

	"github.com/es-sandbox/hs3/message"
)

var (
	environmentInfoBucketName = []byte("env")
	humanHeartInfoBucketName  = []byte("human/heart")
	humanCommonInfoBucketName = []byte("human/common")
	flowerpotInfoBucketName   = []byte("flowerpot")
)

type db struct {
	boltDb *bolt.DB
}

func newDb(boltDb *bolt.DB) *db {
	return &db{
		boltDb: boltDb,
	}
}

func (db *db) getAllEnvironmentInfoRecords() ([]*message.EnvironmentInfo, error) {
	msgs := make([]*message.EnvironmentInfo, 0)
	err := db.boltDb.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(environmentInfoBucketName)

		return bucket.ForEach(func(key, value []byte) error {
			envInvo, err := message.NewEnvironmentInfoFromBytes(value)
			if err != nil {
				return err
			}
			msgs = append(msgs, envInvo)
			return nil
		})
	})
	if err != nil {
		return nil, err
	}
	return msgs, nil
}

func (db *db) putEnvironmentInfoRecord(envInfo *message.EnvironmentInfo) error {
	return db.boltDb.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists(environmentInfoBucketName)
		if err != nil {
			return err
		}

		id, _ := bucket.NextSequence()
		envInfo.Id = id

		rawToDatabase, err := envInfo.Encode()
		if err != nil {
			return err
		}

		return bucket.Put(itob(envInfo.Id), rawToDatabase)
	})
}

func (db *db) getAllHumanHeartInfoRecords() ([]*message.HumanHeartInfo, error) {
	msgs := make([]*message.HumanHeartInfo, 0)
	err := db.boltDb.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(humanHeartInfoBucketName)

		return bucket.ForEach(func(key, value []byte) error {
			hhInfo, err := message.NewHumanHeartInfoFromBytes(value)
			if err != nil {
				return err
			}
			msgs = append(msgs, hhInfo)
			return nil
		})
	})
	if err != nil {
		return nil, err
	}
	return msgs, nil
}

func (db *db) putHumanHeartInfo(humanHeartInfo *message.HumanHeartInfo) error {
	return db.boltDb.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists(humanHeartInfoBucketName)
		if err != nil {
			return err
		}

		id, _ := bucket.NextSequence()
		humanHeartInfo.Id = id

		rawToDatabase, err := humanHeartInfo.Encode()
		if err != nil {
			return err
		}

		return bucket.Put(itob(humanHeartInfo.Id), rawToDatabase)
	})
}

func (db *db) getAllHumanCommonInfoRecords() ([]*message.HumanCommonInfo, error) {
	msgs := make([]*message.HumanCommonInfo, 0)
	err := db.boltDb.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(humanCommonInfoBucketName)

		return bucket.ForEach(func(key, value []byte) error {
			hcInfo, err := message.NewHumanCommonInfoFromBytes(value)
			if err != nil {
				return err
			}
			msgs = append(msgs, hcInfo)
			return nil
		})
	})
	if err != nil {
		return nil, err
	}
	return msgs, nil
}

func (db *db) putHumanCommonInfo(hcInfo *message.HumanCommonInfo) error {
	return db.boltDb.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists(humanCommonInfoBucketName)
		if err != nil {
			return err
		}

		id, _ := bucket.NextSequence()
		hcInfo.Id = id

		rawToDatabase, err := hcInfo.Encode()
		if err != nil {
			return err
		}

		return bucket.Put(itob(hcInfo.Id), rawToDatabase)
	})
}

func (db *db) getAllFlowerpotInfoRecords() ([]*message.FlowerpotInfo, error) {
	msgs := make([]*message.FlowerpotInfo, 0)
	err := db.boltDb.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(flowerpotInfoBucketName)

		return bucket.ForEach(func(key, value []byte) error {
			fpInfo, err := message.NewFlowerpotInfoFromBytes(value)
			if err != nil {
				return err
			}
			msgs = append(msgs, fpInfo)
			return nil
		})
	})
	if err != nil {
		return nil, err
	}
	return msgs, nil
}

func (db *db) putFlowerpotInfo(flowerpotInfo *message.FlowerpotInfo) error {
	return db.boltDb.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists(flowerpotInfoBucketName)
		if err != nil {
			return err
		}

		id, _ := bucket.NextSequence()
		flowerpotInfo.Id = id

		rawToDatabase, err := flowerpotInfo.Encode()
		if err != nil {
			return err
		}

		return bucket.Put(itob(flowerpotInfo.Id), rawToDatabase)
	})
}
