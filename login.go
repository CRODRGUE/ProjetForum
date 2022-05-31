package main

import (
	"fmt"
	"net/http"
	"regexp"
)

type UserLogin struct {
	email    string
	password string
	id       int
}

type checkInputs struct {
	InputName string
	Value     string
}

func checkInputLogin(input []checkInputs) bool {
	email := regexp.MustCompile(`^[A-Za-z0-9._-]+[@]+[A-Za-z0-9]+[.]+[A-Za-z0-9]+$`)

	for i := 0; i < len(input); i++ {
		println("TEST POUR LA VALEUR : ", input[i].Value)
		switch input[i].InputName {
		case "email":
			if !email.MatchString(input[i].Value) {
				println("email non")
				return false
			}
			println("email ok")
		case "password":
			println("taille du pwd : ", len(input[i].Value))
			if !(8 <= len(input[i].Value) && len(input[i].Value) <= 64) {
				println("pwd non")
				return false
			}
			println("pwd ok")
		}
	}
	return true
}

func getLogin(Uemail string) UserLogin {
	db := DbCon()
	var user UserLogin
	err := db.QueryRow("SELECT email, password, id_user FROM users WHERE email = ?", Uemail).Scan(&user.email, &user.password, &user.id)
	if err != nil {
		fmt.Println("err getLogin : ", err.Error())
	}
	return user
}

func checkString(str_1, str_2 string) bool {
	if len(str_1) == len(str_2) {
		for index := range str_1 {
			if str_1[index] != str_2[index] {
				return false
			}
		}
		return true
	} else {
		return false
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	println("Page /log")
	ruler, _ := regexp.MatchString("/favicon.ico", r.URL.Path)
	if !ruler {
		queries := r.URL.Query()
		UserEmail := queries.Get("user_mail")
		UserPassword := queries.Get("user_password")
		println(UserEmail, UserPassword)

		check := checkInputLogin([]checkInputs{
			{Value: UserEmail, InputName: "email"},
			{Value: UserPassword, InputName: "password"},
		})

		println("check = ", check)
		if check {
			user := getLogin(UserEmail)
			test_Email := checkString(UserEmail, user.email)
			test_Password := CheckPasswordHash(UserPassword, user.password)

			println("test de correspondance (mail puis pws) : ", test_Email, test_Password)
			if test_Email && test_Password {

				session, _ := store.Get(r, "cookies-ses")
				session.Values["auth"] = true
				session.Values["id"] = user.id
				session.Save(r, w)

				// test pour observer les cookies !
				auth, _ := session.Values["auth"].(bool)
				id, _ := session.Values["id"].(int)
				println("valeur de auth : ", auth, " valeur de l'id : ", id)
				println("ok cookies set !")

				http.Redirect(w, r, "/accueil", 301)

			} else {
				erreur.Message = "L'utilisateur n'est pas bon ! Le mot de pass ou bien l'email est incorrect... Veuillez recommencer !"
				http.Redirect(w, r, "/erreur", 301)
			}

		} else {
			erreur.Message = "Il y a des problemes avec les données entrées... Veuillez recommencer !"
			http.Redirect(w, r, "/erreur", 301)
		}
	}
}
