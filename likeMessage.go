package main

import (
	"net/http"
	"strconv"
)

type LikeMessage struct {
	Id_user    int
	id_message int
}

func CheckLike(id_message, id_user int) bool {
	db := DbCon()
	getTopic := db.QueryRow("SELECT id_user, id_message FROM likes where id_user=? AND id_message=?;", id_user, id_message)

	var like LikeMessage

	err := getTopic.Scan(&like.Id_user, &like.id_message)
	if err != nil {
		return false
	}
	return true
}

func AddLike(id_message, id_user, id_topic int) {
	db := DbCon()
	req, _ := db.Prepare("INSERT INTO `likes`(`id_user`, `id_message`,`id_topic`) VALUES (?,?,?)")
	res, err0 := req.Exec(id_user, id_message, id_topic)
	println(res)
	if err0 != nil {
		println("Probleme avec l'ajout de like au message : ", err0.Error())
	}
}

func DelLike(id_message, id_user int) {
	db := DbCon()
	req, _ := db.Prepare("DELETE FROM `likes` WHERE id_user=? AND id_message=?;")
	res, err0 := req.Exec(id_user, id_message)
	println(res)
	if err0 != nil {
		println("Probleme avec la suppression du like au message : ", err0.Error())
	}
}

func getLikeUserTopic(id_user, id_topic int) []LikeMessage {
	db := DbCon()
	getLike, err := db.Query("SELECT id_user, id_message FROM likes where id_user=? AND id_topic=?;", id_user, id_topic)
	if err != nil {
		panic(err.Error())
	}

	var allLikeMessage []LikeMessage
	for getLike.Next() {
		var Like LikeMessage
		err = getLike.Scan(&Like.Id_user, &Like.id_message)
		if err != nil {
			panic(err.Error())
		}
		allLikeMessage = append(allLikeMessage, Like)
	}
	return allLikeMessage
}

func LikeMessages(w http.ResponseWriter, r *http.Request) {
	querise := r.URL.Query()
	id_user, errU := strconv.Atoi(querise.Get("u"))
	id_message, errM := strconv.Atoi(querise.Get("m"))
	id_topic, errT := strconv.Atoi(querise.Get("t"))
	println(errU, errM, errT)
	println(id_user, id_message)
	check := CheckLike(id_message, id_user)
	println(check)
	if !check {
		AddLike(id_message, id_user, id_topic)
	} else {
		DelLike(id_message, id_user)
	}
	http.Redirect(w, r, r.Header.Get("Referer"), 301)
}
