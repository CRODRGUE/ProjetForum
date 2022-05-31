package main

import (
	"net/http"
	"regexp"
	"strconv"
)

func CheckInputUpdateTopic(input []checkInputs) bool {
	letter := regexp.MustCompile(`^[a-zA-Z]+$`)

	for i := 0; i < len(input); i++ {
		println("TEST POUR LA VALEUR : ", input[i].Value)
		switch input[i].InputName {
		case "Description":
			if !(1 <= len(input[i].Value) && len(input[i].Value) <= 500) {
				return false
			}
		case "Title-Subjet":
			if !(1 <= len(input[i].Value) && len(input[i].Value) <= 255) && !letter.MatchString(input[i].Value) {
				return false
			}
		case "Id-picture":
			value, _ := strconv.Atoi(input[i].Value)
			if !(1 <= value && value <= 8) && !(input[i].Value == "0") {
				return false
			}
		case "Id-category":
			value, _ := strconv.Atoi(input[i].Value)
			if !(10 <= value && value <= 16) && !(input[i].Value == "0") {

				return false
			}
		}
	}
	return true
}

func validateDataTopic(tab []bool, updateTopic *DataCreateTopic, id int) {
	topic := getTopicUp(id)
	for _, element := range tab {
		println(element)
		switch element {
		case strconv.Itoa(updateTopic.Id_category) == "0":
			if true {
				updateTopic.Id_category = topic.Id_category
			}
		case strconv.Itoa(updateTopic.Id_topic_picture) == "0":
			if true {
				updateTopic.Id_topic_picture = topic.Id_topic_picture
			}
		}
	}
}

func updateTopicUser(topic DataCreateTopic, id int) bool {
	db := DbCon()
	req, _ := db.Prepare("UPDATE `topics` SET `title`= ?,`subject`= ?,`description`= ?,`id_topic_picture`= ?,`id_category`= ? WHERE `id_topic` = ?")
	res, err := req.Exec(&topic.Title, &topic.Subjet, &topic.Description, &topic.Id_topic_picture, &topic.Id_category, id)
	println(res)
	if err != nil {
		println("erreur update topic : ", err.Error())
		return false
	}
	return true
}

func updateTopic(w http.ResponseWriter, r *http.Request) {
	ruler, _ := regexp.MatchString("/favicon.ico", r.URL.Path)

	if !ruler {
		queries := r.URL.Query()
		id, _ := strconv.Atoi(queries.Get("topic"))
		println("updateTopic : ", id, r.URL.Path)
		id_cate, _ := strconv.Atoi(queries.Get("topic_category"))
		id_pic, _ := strconv.Atoi(queries.Get("topic_picture"))

		updateTopic := DataCreateTopic{
			id,
			queries.Get("topic_titre"),
			queries.Get("topic_sujet"),
			"",
			0,
			queries.Get("topic_description"),
			id_pic,
			id_cate,
			0,
		}

		InputsCheck := []checkInputs{
			{Value: updateTopic.Description, InputName: "Description"},
			{Value: updateTopic.Subjet, InputName: "Title-Subjet"},
			{Value: updateTopic.Title, InputName: "Title-Subjet"},
			{Value: strconv.Itoa(updateTopic.Id_topic_picture), InputName: "Id-picture"},
			{Value: strconv.Itoa(updateTopic.Id_category), InputName: "Id-category"},
		}

		println("avant check")
		test := CheckInputUpdateTopic(InputsCheck)
		println("valeur du test : ", test)
		if test {
			println("Avant id : ", updateTopic.Id_category, updateTopic.Id_topic_picture)

			validateDataTopic([]bool{strconv.Itoa(updateTopic.Id_category) == "0", strconv.Itoa(updateTopic.Id_topic_picture) == "0"}, &updateTopic, updateTopic.Id)

			println("Apres id : ", updateTopic.Id_category, updateTopic.Id_topic_picture, " val id topic : ", updateTopic.Id)
			req := updateTopicUser(updateTopic, updateTopic.Id)
			if req {
				session, _ := store.Get(r, "cookies-ses")
				id, _ := session.Values["id"].(int)
				path := "/profil/" + strconv.Itoa(id)
				println("redirection vers : ", path)
				http.Redirect(w, r, path, http.StatusMovedPermanently)
			} else {
				erreur.Message = "Serveur injoignable... Veuillez recommencer plus tard !"
				http.Redirect(w, r, "/erreur", http.StatusMovedPermanently)
			}
		} else {
			erreur.Message = "Un probleme avec les données indiquées... Veuillez recommencer plus tard !"
			http.Redirect(w, r, "/erreur", http.StatusMovedPermanently)
		}
	}
}
