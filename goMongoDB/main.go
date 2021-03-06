package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var SERVER_CONNECT_STRING = "mongodb://heroku:DsRBxVA1CIFLJQB2kSZZnsP0L2XosvaVRlZbxZlcU4in8pokyFeZ5CMSVKV3d9nsyiyX-ZcjA3mSPciqeTCSiw@candidate.14.mongolayer.com:10799,candidate.15.mongolayer.com:10745/app27707026"

type msg struct {
	// Msg   string        `bson:"msg"`
	Count int `bson:"count"`
}

// Note: attribute name must be upper-case start. Or it will not save to DB. (could not identical with document (at least one upper-case))
type Book struct {
	Isbn  string `bson:"isbn"`
	Title string `bson:"title"`
	Price int    `bson:"price"`
}

type Person struct {
	Name  string
	Phone string
}

func main() {
	fmt.Println("Starting connect mongoDB....")
	session, err := mgo.Dial(SERVER_CONNECT_STRING)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	fmt.Println("Connection works....")
	// Optional. Switch the session to a monotonic behavior.
	session.SetSafe(&mgo.Safe{})
	c := session.DB("app27707026").C("test")

	fmt.Println("Getting  Database count....")
	count, err2 := c.Find(bson.M{}).Count()
	if err2 != nil {
		panic(err)
	}
	fmt.Printf("total book count = %d\n", count)

	if count == 0 {
		err = c.Insert(&Book{"Ale1", "Book1", 35})
		err = c.Insert(&Book{"Ale2", "Book2", 20})
		err = c.Insert(&Book{"Ale3", "Book3", 40})
		err = c.Insert(&Book{"Ale4", "Book4", 15})
		err = c.Insert(&Book{"Ale5", "Book5", 55})
		err = c.Insert(&Book{"Ale6", "Book6", 45})
	} else {
		err = c.Insert(&Book{"Ale7", "Book7", 45})

	}

	result := Book{}
	fmt.Println("Getting  data....")
	// Find book which prices is greater than(gt) 40
	//iter := c.Find(bson.M{"price": bson.M{"$gt": 40}}).Iter()
	//Order By: bson.M{"$query": bson.M{}, "$orderby": bson.M{"price": -1}}
	//AND: bson.M{"$and": []bson.M{bson.M{"title": "Book7"}, bson.M{"price": 35}}}
	iter := c.Find(bson.M{"$and": []bson.M{bson.M{"title": "Book7"}, bson.M{"price": 35}}}).Iter()
	var index = 1

	for iter.Next(&result) {
		fmt.Printf("current result is [%d] result =%+v\n", index, result)
		index++
	}
	//when search the DB it must all lower-case to avoid any error.
	if err2 := iter.Close(); err2 != nil {
		fmt.Printf("No data\n")
	} else {
		fmt.Printf("result =%+v\n", result)
	}
}
