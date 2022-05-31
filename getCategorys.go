package main

import (
	_ "github.com/go-sql-driver/mysql"
)

type category struct {
	Id          int
	Category    string
	Picto       string
	SubCategory string
	Path        string
}

type topicImage struct {
	Id   int
	Path string
}

type pageNewTopic struct {
	Category   []category
	TopicImage []topicImage
	Id_ses     int
}

func getCategory() []category {
	db := DbCon()
	getCat, err := db.Query("SELECT * FROM `categorys`")
	if err != nil {
		println("get")
		panic(err.Error())
	}

	var allCate []category

	for getCat.Next() {
		var cate category
		err = getCat.Scan(&cate.Id, &cate.Category, &cate.Picto, &cate.SubCategory, &cate.Path)
		if err != nil {
			println("3")
			panic(err.Error())
		}
		allCate = append(allCate, cate)
	}
	return allCate
}

func getTopicPicture() []topicImage {
	db := DbCon()
	getImg, err := db.Query("SELECT * FROM `topics_pictures`")
	if err != nil {
		panic(err.Error())
	}

	var allImg []topicImage

	for getImg.Next() {
		var Img topicImage
		err = getImg.Scan(&Img.Id, &Img.Path)
		if err != nil {
			panic(err.Error())
		}
		allImg = append(allImg, Img)
	}
	return allImg
}
