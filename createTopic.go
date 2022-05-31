package main

import (
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type DataCreateTopic struct {
	Id               int
	Title            string
	Subjet           string
	Start_date       string
	Polity           int
	Description      string
	Id_topic_picture int
	Id_category      int
	Id_user          int
}

func CheckInputCreate(input []checkInputs) bool {
	letter := regexp.MustCompile(`^[a-zA-Z_-~\/'"]+$`)

	for i := 0; i < len(input); i++ {
		println("TEST POUR LA VALEUR : ", input[i].Value)
		switch input[i].InputName {
		case "Description":
			if !(1 <= len(input[i].Value) && len(input[i].Value) <= 500) {
				println("description non")
				return false
			}
			println("description ok")
		case "Title-Subjet":
			println("taille du title-subjet : ", len(input[i].Value))
			if !(1 <= len(input[i].Value) && len(input[i].Value) <= 255) && !letter.MatchString(input[i].Value) {
				println("title-subjet non")
				return false
			}
			println("title-subjet ok")
		case "Id-picture":
			value, err := strconv.Atoi(input[i].Value)
			if !(1 <= value && value <= 8) || err != nil {
				println("id-picture non")
				return false
			}
			println("id-picture ok")
		case "Id-category":
			value, err := strconv.Atoi(input[i].Value)
			if !(10 <= value && value <= 16) || err != nil {
				println("id-cate non")
				return false
			}
			println("id-cate ok")
		}

	}
	return true
}

func insertTopic(topic *DataCreateTopic) bool {
	db := DbCon()
	req, _ := db.Prepare("INSERT INTO `topics`(`title`, `subject`, `start_date`, `polity`, `description`, `id_topic_picture`, `id_category`, `id_user`) VALUES (?,?,?,?,?,?,?,?)")
	res, err0 := req.Exec(topic.Title, topic.Subjet, topic.Start_date, topic.Polity, topic.Description, topic.Id_topic_picture, topic.Id_category, topic.Id_user)
	println(res)
	if err0 != nil {
		println("Probleme avec la requete pour creer un topic : ", err0.Error())
		return false
	} else {
		err := db.QueryRow("SELECT LAST_INSERT_ID() AS id").Scan(&topic.Id)
		if err != nil {
			fmt.Println(err.Error())
			return false
		}
	}
	return true
}

func editeDate(date string) string {
	dateArry := strings.Split(date, "-")
	newDate := dateArry[2] + "-" + dateArry[0] + "-" + dateArry[1]
	return newDate
}

func newTopic(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "cookies-ses")
	dt := time.Now()
	ruler, _ := regexp.MatchString("/favicon.ico", r.URL.Path)
	if !ruler {
		queries := r.URL.Query()

		id_pic, _ := strconv.Atoi(queries.Get("topic_picture"))
		id_cate, _ := strconv.Atoi(queries.Get("topic_category"))
		id_user, _ := session.Values["id"].(int)
		date := editeDate(dt.Format("01-02-2006"))

		NewTopic := DataCreateTopic{0,
			queries.Get("topic_titre"),
			queries.Get("topic_sujet"),
			date,
			1,
			queries.Get("topic_description"),
			id_pic,
			id_cate,
			id_user,
		}
		println(NewTopic.Id, NewTopic.Title, NewTopic.Subjet, NewTopic.Start_date, NewTopic.Polity, NewTopic.Description, NewTopic.Id_topic_picture, NewTopic.Id_category, NewTopic.Id_user)
		check := CheckInputCreate(
			[]checkInputs{
				{Value: NewTopic.Title, InputName: "Title-Subjet"},
				{Value: NewTopic.Subjet, InputName: "Title-Subjet"},
				{Value: NewTopic.Description, InputName: "Description"},
				{Value: strconv.Itoa(NewTopic.Id_topic_picture), InputName: "Id-picture"},
				{Value: strconv.Itoa(NewTopic.Id_category), InputName: "Id-category"},
			})
		fmt.Println("valeur de check : ", check)
		if check {
			req := insertTopic(&NewTopic)
			if req {
				println("cest ok pour le topic ", NewTopic.Id)
				path := "/topic?topic=" + strconv.Itoa(NewTopic.Id) + "&page=0"
				http.Redirect(w, r, path, http.StatusMovedPermanently)
			} else {
				erreur.Message = "Serveur injoignable... Veuillez recommencer plus tard !"
				http.Redirect(w, r, "/erreur", http.StatusMovedPermanently)
			}
		} else {
			erreur.Message = "Les données entres dans les champs ne sont pas valides... Veuillez recommencer "
			http.Redirect(w, r, "/erreur", http.StatusMovedPermanently)
		}
	}
}
