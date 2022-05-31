package main

import (
	"net/http"
	"strconv"
	"time"
)

func insertMessage(message *Message) bool {
	db := DbCon()
	println(message.Message)
	req, _ := db.Prepare(`INSERT INTO messages (message, date_message, id_user, id_topic) VALUES (?,?,?,?);`)
	res, err0 := req.Exec(message.Message, message.Date_message, message.Id_user_message, message.Id_Topic)
	println(res)
	if err0 != nil {
		println("Probleme avec la requete pour creer un topic : ", err0.Error())
		return false
	}
	return true
}

func sendMessage(w http.ResponseWriter, r *http.Request) {

	queries := r.URL.Query()
	id_topic, _ := strconv.Atoi(queries.Get("topic"))
	message := r.FormValue("new_message")

	session, _ := store.Get(r, "cookies-ses")
	auth, _ := session.Values["auth"].(bool)
	id, _ := session.Values["id"].(int)

	println("valeur bool (auth puis test valeur) : ", auth)
	println("valeur id_user (id puis test valeur) : ", id)

	dt := time.Now()
	date := editeDate(dt.Format("01-02-2006")) + " " + dt.Format("15:04:05")

	newMessage := Message{0, message, date, "", id, id_topic}
	checkMessage := insertMessage(&newMessage)
	if checkMessage {
		path := "/topic?topic=" + strconv.Itoa(id_topic) + "&page=0"
		http.Redirect(w, r, path, 301)
	} else {
		erreur.Message = "Impossible de joindre le server... Veuillez recommencer !"
		http.Redirect(w, r, "/erreur", 301)
	}

}
