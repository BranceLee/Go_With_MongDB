package api

// 建立mongoDB 数据连接

import (
	"log"
	"github.com/Plee/Mongolang/models"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

/*
** 后端侧数据抽象,自定义它的行为
*/
type MoviesAPI struct{
	Server string
	Database string
}

/*
** 为了代表，真实数据库的实例
*/ 
var db *mgo.Database

/*
** 代表哪张表
*/
const (
	COLLECTION ="movies"
)

/* 
** 连接 MongoDB 数据库
** server被实例的DB,与数据库建立联系
*/
func (m *MoviesAPI) Connect() {
	session, err :=mgo.Dial(m.Server)
	if err !=nil{
		log.Fatal("Connect",err)
	}
	db = session.DB(m.Database)
	print("Connect Success")
}

/*
** db.C(COLLECTION) 这个对象中的方法，能操作改变数据库
*/
func (m *MoviesAPI) FindAll()([]models.Movies,error){
	var movies []models.Movies
	err:=db.C(COLLECTION).Find(bson.M{}).All(&movies)
	return movies, err
}

func (m *MoviesAPI) Create(movie models.Movies) error{
	err := db.C(COLLECTION).Insert(&movie)
	return err
}

func (m *MoviesAPI) FindById(id string) (models.Movies, error){
	var movie models.Movies
	err :=db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&movie)
	return movie, err
}

func (m *MoviesAPI) Delete(movie models.Movies) error{
	err := db.C(COLLECTION).Remove(&movie)
	return err
}