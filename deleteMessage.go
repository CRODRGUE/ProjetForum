package main

import (
	"net/http"
	"strconv"
)

func delMessageUser(id int) bool {
	db := DbCon()
	req, _ := db.Prepare("DELETE FROM `messages` WHERE `id_message` = ?")
	res, err := req.Exec(id)
	println(res)
	if err != nil {
		println("erreur delMessage : ", err.Error())
		return false
	}
	return true
}

func DeleteMessage(w http.ResponseWriter, r *http.Request) {
	querise := r.URL.Query()
	id_mess, _ := strconv.Atoi(querise.Get("id"))
	id_p, _ := strconv.Atoi(querise.Get("q"))

	session, _ := store.Get(r, "cookies-ses")
	auth, checkAuth := session.Values["auth"].(bool)
	id_ses, checkId := session.Values["id"].(int)
	println("valeur bool (auth puis test valeur) : ", auth)
	println("valeur id_user (id puis test valeur) : ", id_ses)
	if (!checkAuth || !auth) || (!checkId) || !(id_ses == id_p) {
		http.Redirect(w, r, "/login", http.StatusMovedPermanently)
	}

	if delMessageUser(id_mess) {
		http.Redirect(w, r, r.Header.Get("Referer"), http.StatusMovedPermanently)
	} else {
		erreur.Message = "Serveur injoignable... Veuillez recommencer plus tard !"
		http.Redirect(w, r, "/erreur", http.StatusMovedPermanently)
	}

}
