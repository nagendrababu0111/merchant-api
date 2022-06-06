package mongo

import (
	"fmt"
	"log"
	"net/url"

	"gopkg.in/mgo.v2"

	"merchant-api/utils/config"
	"merchant-api/utils/types"
)

var sessions map[string]*mgo.Session

// pull session
func GetSession(db string) *mgo.Session {
	initSessions()
	uri := getConnection(db)
	if sessions[db] == nil {
		mgoSession, err := connect(uri)
		if err != nil {
			panic(err)
		}
		mgoSession.SetSafe(&mgo.Safe{WMode: "majority"})
		sessions[db] = mgoSession
	}
	return sessions[db].Clone()
}

func connect(uri string) (*mgo.Session, error) {
	var err error
	var mgoSession *mgo.Session

	for index := 1; index < 5; index++ {
		mgoSession, err = mgo.Dial(uri)
		if err == nil {
			return mgoSession, err
		}
	}
	return mgoSession, err
}

func getConnection(db string) string {
	conn := config.GetDBCredentials(db)
	password := url.QueryEscape(conn.Password)
	uri := fmt.Sprintf("mongodb://%s:%s@%s:%d/%s", conn.User, password, conn.Host, conn.Port, db)
	return uri
}

func initSessions() {
	if sessions == nil {
		sessions = make(map[string]*mgo.Session)
	}
}

//FindOne - Get documents from a table based on given query
func FindOne(db string, table string, query types.Map) (types.Map, error) {
	result := make(types.Map, 0)
	session := GetSession(db)
	defer session.Close()
	coll := session.DB(db).C(table)
	q := coll.Find(query).One(&result)
	if q != nil && q.Error() != "" {
		log.Println(q.Error())
	}
	return result, q
}

//Update - Update multiple documents
func Update(db string, table string, query types.Map, data interface{}) error {
	obj := make(types.Map, 0)
	obj["$set"] = data
	log.Println("data-->", obj, query)
	session := GetSession(db)
	defer session.Close()
	prog := session.DB(db).C(table)
	_, err := prog.UpdateAll(query, data)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func UpdateOne(db string, table string, id string, data interface{}) error {
	obj := make(types.Map, 0)
	obj["$set"] = data
	session := GetSession(db)
	defer session.Close()
	prog := session.DB(db).C(table)
	err := prog.UpdateId(id, data)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

//Delete
func Delete(db string, table string, query types.Map) error {
	session := GetSession(db)
	defer session.Close()
	tab := session.DB(db).C(table)
	_, err := tab.RemoveAll(query)
	return err
}

// Create -- Insert a doc and return document with if any error.
func Create(db string, table string, data interface{}) (interface{}, error) {

	session := GetSession(db)
	defer session.Close()
	tab := session.DB(db).C(table)
	err := tab.Insert(data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

//FindRecords - produce result bassed on paging
func FindRecords(db string, table string, query types.Map, paging types.Page, result interface{}) (int, error) {
	// result := make([]types.Map, 0)
	session := GetSession(db)
	defer session.Close()
	coll := session.DB(db).C(table)
	totalRows, err := coll.Find(query).Count()
	if err != nil {
		log.Println("error while reading data in find ", query)
		return 0, err
	}
	coll = session.DB(db).C(table)
	err = coll.Find(query).Skip(paging.Skip).Limit(paging.Limit).All(result)
	if err != nil {
		return 0, err
	}
	return totalRows, err
}

func CreateUniqueIndex(db string, table string, key []string) error {
	session := GetSession(db)
	defer session.Close()
	coll := session.DB(db).C(table)

	index := mgo.Index{
		Key:    key,
		Unique: true,
	}
	return coll.EnsureIndex(index)

}

//GetTableRowsCount - Get documents count from a table based on given query
func GetTableRowsCount(db string, table string, query types.Map) int {
	session := GetSession(db)
	defer session.Close()
	coll := session.DB(db).C(table)
	totalRecords, err := coll.Find(query).Count()
	if err != nil {
		log.Println(err)
	}
	log.Println("totalRecords", totalRecords)
	return totalRecords
}
