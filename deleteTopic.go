package main

import (
	"net/http"
	"strconv"
)

//DELETE FROM `topics` WHERE `id_topic` = ?
func delTopic(id int) bool {
	db := DbCon()
	req, _ := db.Prepare("DELETE FROM `topics` WHERE `id_topic` = ?")
	res, err := req.Exec(id)
	println(res)
	if err != nil {
		println("erreur delTopic : ", err.Error())
		return false
	}
	return true
}

func delMessage(id int) bool {
	db := DbCon()
	req, _ := db.Prepare("DELETE FROM `messages` WHERE `id_topic` = ?")
	res, err := req.Exec(id)
	println(res)
	if err != nil {
		println("erreur delMessage : ", err.Error())
		return false
	}
	return true
}

func delLike(id int) bool {
	db := DbCon()
	req, _ := db.Prepare("DELETE FROM `likes` WHERE `id_topic` = ?")
	res, err := req.Exec(id)
	println(res)
	if err != nil {
		println("erreur delLike : ", err.Error())
		return false
	}
	return true
}

func DeleteTopic(w http.ResponseWriter, r *http.Request) {
	querise := r.URL.Query()
	id_topic, err := strconv.Atoi(querise.Get("id"))
	id_p, err0 := strconv.Atoi(querise.Get("q"))

	session, _ := store.Get(r, "cookies-ses")
	auth, checkAuth := session.Values["auth"].(bool)
	id_ses, checkId := session.Values["id"].(int)
	println("valeur bool (auth puis test valeur) : ", auth)
	println("valeur id_user (id puis test valeur) : ", id_ses)
	if (!checkAuth || !auth) || (!checkId) || !(id_ses == id_p) {
		http.Redirect(w, r, "/login", http.StatusMovedPermanently)
	}

	if delTopic(id_topic) && delMessage(id_topic) && delLike(id_topic) {
		http.Redirect(w, r, r.Header.Get("Referer"), http.StatusMovedPermanently)
	} else {
		erreur.Message = "Serveur injoignable... Veuillez recommencer plus tard !"
		http.Redirect(w, r, "/erreur", http.StatusMovedPermanently)
	}

	println(id_topic, err, err0)
}
