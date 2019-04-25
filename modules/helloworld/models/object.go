package models

import (
	"errors"
	"strconv"
	"time"
)

var (
	// Objects ...
	Objects map[string]*Object	
)

// Object ...
type Object struct {
	ObjectID   string
	Score      int64
	PlayerName string
}

func init() {
	Objects = make(map[string]*Object)
	Objects["test001"] = &Object{"test001", 100, "hopher"}
	Objects["test002"] = &Object{"test002", 101, "xulq"}
}

// AddOne ...
func AddOne(object Object) (ObjectID string) {
	object.ObjectID = "astaxie" + strconv.FormatInt(time.Now().UnixNano(), 10)
	Objects[object.ObjectID] = &object
	return object.ObjectID
}

// GetOne ...
func GetOne(ObjectID string) (object *Object, err error) {
	if v, ok := Objects[ObjectID]; ok {
		return v, nil
	}
	return nil, errors.New("ObjectID Not Exist")
}

// GetAll ...
func GetAll() map[string]*Object {
	return Objects
}

// Update ...
func Update(ObjectID string, Score int64) (err error) {
	if v, ok := Objects[ObjectID]; ok {
		v.Score = Score
		return nil
	}
	return errors.New("ObjectID Not Exist")
}

// Delete ...
func Delete(ObjectID string) {
	delete(Objects, ObjectID)
}
