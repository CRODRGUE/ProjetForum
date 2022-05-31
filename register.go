package main

import (
	"net/http"
	"regexp"
	"strconv"
)

type UserRegister struct {
	Pseudo         string `json:"pseudo"`
	FirstName      string `json:"firstname"`
	LastName       string `json:"lastname"`
	Description    string `json:"description"`
	ProfilePicture string `json:"profile_picture"`
	Email          string `json:"email"`
	Password       string `json:"password"`
	Date           string `json:"brith_day"`
}

func CheckInputRegister(input []checkInputs) bool {
	userInfo := regexp.MustCompile(`^([A-Za-z0-9-]{1,255})+$`)
	email := regexp.MustCompile(`^[A-Za-z0-9._-]+[@]+[A-Za-z0-9]+[.]+[A-Za-z0-9]+$`)
	userPseudo := regexp.MustCompile(`^([A-Za-z0-9._@-]{1,64})+$`)

	for i := 0; i < len(input); i++ {
		switch input[i].InputName {
		case "UserInfo":
			println("UserInfo")
			if !userInfo.MatchString(input[i].Value) && !(0 < len(input[i].Value) && len(input[i].Value) < 255) {
				return false
			}
		case "Pseudo":
			println("Pseudo")
			if !userPseudo.MatchString(input[i].Value) && !(0 < len(input[i].Value) && len(input[i].Value) < 65) {
				return false
			}
		case "Email":
			println("Email")
			if !email.MatchString(input[i].Value) {
				return false
			}
		case "Password":
			println("Password")
			if !(9 < len(input[i].Value) && len(input[i].Value) < 65) {
				return false
			}
		case "Description":
			println("Description")
			if !(0 < len(input[i].Value) && len(input[i].Value) <= 1500) {
				return false
			}
		case "ProfilPicture":
			println("ProfilPicture")
			nbr, _ := strconv.Atoi(input[i].Value)
			if !(1 <= nbr && nbr <= 5) {
				return false
			}
		}
	}
	return true
}

func CreateUser(User UserRegister) bool {
	db := DbCon()
	req, _ := db.Prepare("INSERT INTO `users`(`lastname`, `firstname`, `email`, `password`, `pseudo`, `description`, `profile_picture`,`brith_day`) VALUES (?, ?, ?, ?, ?, ?, ?, ?)")
	res, err0 := req.Exec(User.LastName, User.FirstName, User.Email, User.Password, User.Pseudo, User.Description, User.ProfilePicture, User.Date)
	println(res)
	if err0 != nil {
		println("Probleme avec la requete pour creer un compte : ", err0.Error())
		return false
	}
	return true
}

func register(w http.ResponseWriter, r *http.Request) {
	ruler, _ := regexp.MatchString("/favicon.ico", r.URL.Path)
	if !ruler {
		queries := r.URL.Query()

		PassWord := queries.Get("user_password")
		LenPassWord := len(PassWord)
		if 9 < LenPassWord && LenPassWord < 65 {
			PassWord, PassChek := HashPassword(PassWord)
			NewUser := UserRegister{queries.Get("user_pseudo"),
				queries.Get("user_firstname"),
				queries.Get("user_lastname"),
				queries.Get("user_description"),
				queries.Get("badge"),
				queries.Get("user_email"),
				PassWord,
				queries.Get("user_date"),
			}

			check := CheckInputRegister(
				[]checkInputs{
					{Value: NewUser.Pseudo, InputName: "Pseudo"},
					{Value: NewUser.LastName, InputName: "UserInfo"},
					{Value: NewUser.FirstName, InputName: "UserInfo"},
					{Value: NewUser.Description, InputName: "Description"},
					{Value: NewUser.Password, InputName: "Password"},
					{Value: NewUser.ProfilePicture, InputName: "ProfilPicture"},
					{Value: NewUser.Email, InputName: "Email"},
				})
			println(NewUser.Date)
			if check && PassChek == nil {
				req := CreateUser(NewUser)
				if req {
					println("Utilisateur créer !")
					http.Redirect(w, r, "/login", 301)
				} else {
					erreur.Message = "Serveur injoignable... Veuillez recommencer plus tard !"
					http.Redirect(w, r, "/erreur", 301)
				}
			}
		} else {
			erreur.Message = "Les données entres dans les champs ne sont pas valides... Veuillez recommencer !"
			http.Redirect(w, r, "/erreur", 301)
		}
	}

}
