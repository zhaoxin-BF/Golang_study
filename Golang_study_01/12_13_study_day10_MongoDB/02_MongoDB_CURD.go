package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
	//"gopkg.in/mgo.v2/bson"
)

type Person struct {
	Name string
	Phone string
}

func main(){
	session, err := mgo.Dial("172.18.183.132:27017")
	if err != nil{
		panic(err)
	}
	defer session.Clone()

	session.SetMode(mgo.Monotonic, true)
	c := session.DB("test").C("student")
	for i := 0;i<100000;i++{
		err = c.Insert(&Person{"boreas", "666666666"})
		if err != nil{
			panic(err)
		}
	}

	fmt.Println("插入完毕！")
}