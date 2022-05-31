package main

import (
	"net/http"
	"regexp"
	"strconv"
)

func CheckInputUpdateProfil(input []checkInputs) bool {
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
			println("Password : ", input[i].Value)
			if !(9 < len(input[i].Value) && len(input[i].Value) < 65) && !(input[i].Value == "") {
				return false
			}
		case "Description":
			println("Description")
			if !(0 < len(input[i].Value) && len(input[i].Value) <= 1500) {
				return false
			}
		case "ProfilPicture":
			println("ProfilPicture : ", input[i].Value)
			nbr, _ := strconv.Atoi(input[i].Value)
			if !(1 <= nbr && nbr <= 5) && !(input[i].Value == "") {
				return false
			}
		}
	}
	return true
}

func validateDataUser(tab []bool, userUp *User, id int) {
	user := getUser(id)
	println("valeur de date de user :", user.Date)
	for _, element := range tab {
		println(element)
		switch element {
		case userUp.Password == "":
			if true {
				userUp.Password = user.Password
			} else {
				println("le mot de pass va etres encrypé !")
				userUp.Password, _ = HashPassword(userUp.Password)
			}
		case userUp.Email == "":
			if true {
				userUp.Email = user.Email
			}
		case userUp.ProfilePicture == "":
			if true {
				userUp.ProfilePicture = user.ProfilePicture
			}
		case userUp.Date == "":
			println("valeur de date dans le validateData : ", userUp.Date)
			if true {
				println("valeur de user.Date : ", user.Date)
				userUp.Date = user.Date
			}
		}
	}
}

func updateProfileUser(id int, user User) bool {
	db := DbCon()
	req, _ := db.Prepare("UPDATE `users` SET `lastname`= ?,`firstname`= ?,`brith_day`= ?,`email`= ?,`password`= ?,`pseudo`= ?,`description`= ?,`profile_picture`= ? WHERE `id_user`= ?")
	res, err := req.Exec(&user.LastName, &user.FirstName, &user.Date, &user.Email, &user.Password, &user.Pseudo, &user.Description, &user.ProfilePicture, id)
	println(res)
	if err != nil {
		println("erreur update profil : ", err.Error())
		return false
	}
	return true
}

func updateProfile(w http.ResponseWriter, r *http.Request) {
	ruler, _ := regexp.MatchString("/favicon.ico", r.URL.Path)

	if !ruler {
		queries := r.URL.Query()
		updateUser := User{
			queries.Get("user_pseudo"),
			queries.Get("user_firstname"),
			queries.Get("user_lastname"),
			queries.Get("user_description"),
			queries.Get("user_picture"),
			queries.Get("user_email"),
			queries.Get("user_password"),
			queries.Get("user_date"),
			0,
		}

		InputsCheck := []checkInputs{
			{Value: updateUser.Pseudo, InputName: "Pseudo"},
			{Value: updateUser.LastName, InputName: "UserInfo"},
			{Value: updateUser.FirstName, InputName: "UserInfo"},
			{Value: updateUser.Description, InputName: "Description"},
			{Value: updateUser.ProfilePicture, InputName: "ProfilPicture"},
			{Value: updateUser.Password, InputName: "Password"},
			{Value: queries.Get("user_password_c"), InputName: "Password"},
			{Value: updateUser.Email, InputName: "Email"},
			{Value: queries.Get("user_email_c"), InputName: "Email"},
		}

		if checkString(updateUser.Password, queries.Get("user_password_c")) && checkString(updateUser.Email, queries.Get("user_email_c")) && CheckInputUpdateProfil(InputsCheck) {
			id_user, _ := strconv.Atoi(queries.Get("user"))
			println("avant : ", updateUser.Password, updateUser.Date, updateUser.ProfilePicture)
			validateDataUser([]bool{updateUser.Password == "", updateUser.Date == "", updateUser.ProfilePicture == "", updateUser.Email == ""}, &updateUser, id_user)
			println("aprés : ", updateUser.Password, updateUser.Date, updateUser.ProfilePicture)
			req := updateProfileUser(id_user, updateUser)
			if req {
				http.Redirect(w, r, r.Header.Get("Referer"), http.StatusMovedPermanently)
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
