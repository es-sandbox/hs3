package storage

import (
	"encoding/binary"
	"github.com/boltdb/bolt"
	"github.com/es-sandbox/hs3/bolt_db"

	"github.com/es-sandbox/hs3/message"
)

var (
	environmentInfoBucketName = []byte("env")
	humanHeartInfoBucketName  = []byte("human/heart")
	humanCommonInfoBucketName = []byte("human/common")
	flowerpotInfoBucketName   = []byte("flowerpot")
	metaInfoBucketName        = []byte("meta")
	headInfoBucketName        = []byte("head")

	robotModeKey = []byte("mode")
)

type storage struct {
	boltDb *bolt.DB
}

func New(boltDb *bolt.DB) bolt_db.Store {
	return &storage{
		boltDb: boltDb,
	}
}

// ------------------------------ GET ALL METHODS ------------------------------

func (db *storage) GetAllEnvironmentInfoRecords() ([]*message.EnvironmentInfo, error) {
	msgs := make([]*message.EnvironmentInfo, 0)
	err := db.boltDb.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(environmentInfoBucketName)
		if bucket == nil {
			return nil
		}

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

func (db *storage) GetAllHumanHeartInfoRecords() ([]*message.HumanHeartInfo, error) {
	msgs := make([]*message.HumanHeartInfo, 0)
	err := db.boltDb.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(humanHeartInfoBucketName)
		if bucket == nil {
			return nil
		}

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

func (db *storage) GetAllHumanCommonInfoRecords() ([]*message.HumanCommonInfo, error) {
	msgs := make([]*message.HumanCommonInfo, 0)
	err := db.boltDb.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(humanCommonInfoBucketName)
		if bucket == nil {
			return nil
		}

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

func (db *storage) GetAllFlowerpotInfoRecords() ([]*message.FlowerpotInfo, error) {
	msgs := make([]*message.FlowerpotInfo, 0)
	err := db.boltDb.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(flowerpotInfoBucketName)
		if bucket == nil {
			return nil
		}

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

func (db *storage) GetAllHeadInfoRecords() ([]*message.Head, error) {
	msgs := make([]*message.Head, 0)
	err := db.boltDb.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(headInfoBucketName)
		if bucket == nil {
			return nil
		}

		return bucket.ForEach(func(key, value []byte) error {
			head, err := message.NewHeadFromBytes(value)
			if err != nil {
				return err
			}
			msgs = append(msgs, head)
			return nil
		})
	})
	if err != nil {
		return nil, err
	}
	return msgs, nil
}

// -----------------------------------------------------------------------------

// ------------------------------ GET METHODS ------------------------------

func (db *storage) GetEnvironmentInfoRecord() (*message.EnvironmentInfo, error) {
	var msg *message.EnvironmentInfo
	err := db.boltDb.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(environmentInfoBucketName)
		if bucket == nil {
			return nil
		}

		key, value := bucket.Cursor().Last()
		if key == nil || value == nil {
			return nil
		}

		var err error
		msg, err = message.NewEnvironmentInfoFromBytes(value)
		return err
	})
	if err != nil {
		return nil, err
	}
	return msg, nil
}

func (db *storage) GetHumanHeartInfoRecord() (*message.HumanHeartInfo, error) {
	var msg *message.HumanHeartInfo
	err := db.boltDb.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(humanHeartInfoBucketName)
		if bucket == nil {
			return nil
		}

		key, value := bucket.Cursor().Last()
		if key == nil || value == nil {
			return nil
		}

		var err error
		msg, err = message.NewHumanHeartInfoFromBytes(value)
		return err
	})
	if err != nil {
		return nil, err
	}
	return msg, nil
}

func (db *storage) GetHumanCommonInfoRecord() (*message.HumanCommonInfo, error) {
	var msg *message.HumanCommonInfo
	err := db.boltDb.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(humanCommonInfoBucketName)
		if bucket == nil {
			return nil
		}

		key, value := bucket.Cursor().Last()
		if key == nil || value == nil {
			return nil
		}

		var err error
		msg, err = message.NewHumanCommonInfoFromBytes(value)
		return err
	})
	if err != nil {
		return nil, err
	}
	return msg, nil
}

func (db *storage) GetFlowerpotInfoRecord() (*message.FlowerpotInfo, error) {
	var msg *message.FlowerpotInfo
	err := db.boltDb.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(flowerpotInfoBucketName)
		if bucket == nil {
			return nil
		}

		key, value := bucket.Cursor().Last()
		if key == nil || value == nil {
			return nil
		}

		var err error
		msg, err = message.NewFlowerpotInfoFromBytes(value)
		return err
	})
	if err != nil {
		return nil, err
	}
	return msg, nil
}

func (db *storage) GetRobotMode() (*message.RobotMode, error) {
	var mode *message.RobotMode
	err := db.boltDb.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(metaInfoBucketName)
		if bucket == nil {
			return nil
		}

		var (
			err error
			raw = bucket.Get(robotModeKey)
		)
		mode, err = message.NewRobotModeFromBytes(raw)
		return err
	})
	if err != nil {
		return nil, err
	}
	return mode, nil
}

func (db *storage) GetHeadInfoRecord() (*message.Head, error) {
	var msg *message.Head
	err := db.boltDb.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(headInfoBucketName)
		if bucket == nil {
			return nil
		}

		key, value := bucket.Cursor().Last()
		if key == nil || value == nil {
			return nil
		}

		var err error
		msg, err = message.NewHeadFromBytes(value)
		return err
	})
	if err != nil {
		return nil, err
	}
	return msg, nil
}

// -------------------------------------------------------------------------

// ------------------------------ PUT ALL METHODS ------------------------------

func (db *storage) PutEnvironmentInfoRecord(envInfo *message.EnvironmentInfo) error {
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

func (db *storage) PutHumanHeartInfo(humanHeartInfo *message.HumanHeartInfo) error {
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

func (db *storage) PutHumanCommonInfo(hcInfo *message.HumanCommonInfo) error {
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

func (db *storage) PutFlowerpotInfo(flowerpotInfo *message.FlowerpotInfo) error {
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

func (db *storage) PutRobotMode(mode *message.RobotMode) error {
	return db.boltDb.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists(metaInfoBucketName)
		if err != nil {
			return err
		}

		rawToDatabase, err := mode.Encode()
		if err != nil {
			return err
		}

		return bucket.Put(robotModeKey, rawToDatabase)
	})
}

func (db *storage) PutHeadInfoRecord(head *message.Head) error {
	return db.boltDb.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists(headInfoBucketName)
		if err != nil {
			return err
		}

		id, _ := bucket.NextSequence()
		head.Id = id

		rawToDatabase, err := head.Encode()
		if err != nil {
			return err
		}

		return bucket.Put(itob(head.Id), rawToDatabase)
	})
}

// -----------------------------------------------------------------------------

// itob returns an 8-byte big endian representation of v.
func itob(v uint64) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}