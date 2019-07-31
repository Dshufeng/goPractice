package mongo

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Info struct {
	Id    bson.ObjectId `bson:"_id"`
	Src   string        `bson:"src"`
	Url   string        `bson:"url"`
	Title string        `bson:"title"`
}

const URL = "localhost:27017"

var (
	mgoSession *mgo.Session
	database   = "zhanku"
)

func getSession() *mgo.Session {
	if mgoSession == nil {
		var err error
		mgoSession, err = mgo.Dial(URL)
		if err != nil {
			panic(err) //直接终止程序运行
		}
	}
	return mgoSession.Clone()
}

func witchCollection(collection string, s func(c *mgo.Collection) error) error {
	session := getSession()
	//session.SetMode(mgo.Monotonic, true)

	defer session.Close()

	c := session.DB(database).C(collection)
	return s(c)
}

func AddList(l Info) string {
	l.Id = bson.NewObjectId()
	query := func(c *mgo.Collection) error {
		return c.Insert(l)
	}
	err := witchCollection("list", query)
	if err != nil {
		return "false"
	}
	return l.Id.Hex()
}

func GetListById(id string) Info {
	objId := bson.ObjectIdHex(id)

	result := Info{}
	query := func(c *mgo.Collection) error {
		return c.FindId(objId).One(&result)
	}
	witchCollection("list", query)
	return result
}

func AllList() []Info {
	var infoSlice []Info
	query := func(c *mgo.Collection) error {
		return c.Find(nil).All(&infoSlice)
	}
	witchCollection("list", query)
	return infoSlice
}
func UpdateList(selector bson.M, change bson.M) string {
	query := func(c *mgo.Collection) error {
		_, err := c.UpdateAll(selector, change)
		return err
	}
	err := witchCollection("list", query)
	if err == nil {
		return "true"
	}
	return "false"
}

func UpdateListById(id string, change bson.M) string {
	objId := bson.ObjectIdHex(id)
	query := func(c *mgo.Collection) error {
		return c.UpdateId(objId, change)
	}
	err := witchCollection("list", query)
	if err != nil {
		return "false"
	}
	return "true"
}
