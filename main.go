package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/sessions"
)

type err struct {
	Message string
}

var erreur err

var (
	key   = []byte("KBhymFF5daRASq9LVxES")
	store = sessions.NewCookieStore(key)
)

func main() {
	// Mise en place des templates !
	temp, errTemp := template.ParseGlob("./temp/*.html")
	if errTemp != nil {
		println(">>>>>>>>>>", errTemp.Error(), "<<<<<<<<<<<<")
	}

	// Mise en place des routes pour le forum !
	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		println("Page /login")

		session, _ := store.Get(r, "cookies-ses")
		auth, _ := session.Values["auth"].(bool)
		id, _ := session.Values["id"].(int)

		println("valeur bool (auth puis test valeur) : ", auth)
		println("valeur id_user (id puis test valeur) : ", id)

		if auth || id != 0 {
			http.Redirect(w, r, "/accueil", http.StatusMovedPermanently)
		}

		temp.ExecuteTemplate(w, "login", nil)
	})
	http.HandleFunc("/log", login) //Permet la connexion de l'utilisateur...

	//Mise en place de la route register pour permettre à l'utilisateur de s'inscrire sur la plateforme
	http.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		temp.ExecuteTemplate(w, "register", nil)
	})
	http.HandleFunc("/res", register) //Permet l'enregitrement des données et valider l'inscription de l'utilisateur !

	http.HandleFunc("/logout", func(w http.ResponseWriter, r *http.Request) {
		println("Page /logout")

		session, _ := store.Get(r, "cookies-ses")
		session.Values["auth"] = false
		session.Values["id"] = 0
		session.Save(r, w)
		println("cookies update (false && 0)")

		http.Redirect(w, r, "/login", http.StatusMovedPermanently)

		// test pour regarder l'etat des cookies !

	})

	// Page d'erreur qui affiche un message !
	http.HandleFunc("/erreur", func(w http.ResponseWriter, r *http.Request) {
		temp.ExecuteTemplate(w, "erreur", erreur)
	})
	// Route principale qui gere la page d'erreur 404 (page introuvable)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			erreur.Message = "Page introuvable"
			temp.ExecuteTemplate(w, "erreur", erreur)
		}
		http.Redirect(w, r, "/accueil", http.StatusMovedPermanently)
	})

	type accueil struct {
		Id_ses int
		Cat    []category
	}
	// Route vers la page principale du forum (apres connexion...)
	http.HandleFunc("/accueil", func(w http.ResponseWriter, r *http.Request) {
		println("Page /accueil")

		session, _ := store.Get(r, "cookies-ses")
		auth, checkAuth := session.Values["auth"].(bool)
		id, _ := session.Values["id"].(int)

		println("valeur bool (auth puis test valeur) : ", auth)
		println("valeur id_user (id puis test valeur) : ", id)

		if !checkAuth || !auth || (id == 0) {
			http.Redirect(w, r, "/login", http.StatusMovedPermanently)
		}

		Category := accueil{id, getCategory()}
		temp.ExecuteTemplate(w, "accueil", Category)
	})

	//Route qui permet la création de topic (vue)
	http.HandleFunc("/createTopic", func(w http.ResponseWriter, r *http.Request) {

		session, _ := store.Get(r, "cookies-ses")
		auth, checkAuth := session.Values["auth"].(bool)
		id, _ := session.Values["id"].(int)

		println("valeur bool (auth puis test valeur) : ", auth)
		println("valeur id_user (id puis test valeur) : ", id)

		if !checkAuth || !auth || (id == 0) {
			http.Redirect(w, r, "/accueil", http.StatusMovedPermanently)
		}

		DataPage := pageNewTopic{getCategory(), getTopicPicture(), id}
		temp.ExecuteTemplate(w, "new-topic", DataPage)
	})
	http.HandleFunc("/create", newTopic) //Route qui permet de verifié et valider les données ! pour créer un topic

	type DisplayCategory struct {
		Category  string
		DataTopic []DisplayTopic
		Path      string
		PrevPage  int
		NextPage  int
		Page      int
		Id_ses    int
	}

	http.HandleFunc("/categorie/", func(w http.ResponseWriter, r *http.Request) {
		cat := strings.Split(r.URL.Path[11:], "?")
		println(cat[0])
		querise := r.URL.Query()

		session, _ := store.Get(r, "cookies-ses")
		auth, checkAuth := session.Values["auth"].(bool)
		id, _ := session.Values["id"].(int)

		println("valeur bool (auth puis test valeur) : ", auth)
		println("valeur id_user (id puis test valeur) : ", id)

		if !checkAuth || !auth || (id == 0) {
			http.Redirect(w, r, "/login", http.StatusMovedPermanently)
		}

		page, err := strconv.Atoi(querise.Get("page"))
		if err != nil {
			erreur.Message = "Page introuvable... Veuillez recommencer plus tard !"
			http.Redirect(w, r, "/erreur", http.StatusMovedPermanently)
		}
		println(r.URL.Path)
		Data := DisplayCategory{strings.Title(cat[0]), getTopicsCategory(cat[0], page), cat[0], PrevPage(page - 1), (page + 1), page, id}
		temp.ExecuteTemplate(w, "category", Data)
	})

	http.HandleFunc("/lastTopic", func(w http.ResponseWriter, r *http.Request) {
		querise := r.URL.Query()
		page, _ := strconv.Atoi(querise.Get("page"))

		session, _ := store.Get(r, "cookies-ses")
		auth, checkAuth := session.Values["auth"].(bool)
		id, _ := session.Values["id"].(int)

		println("valeur bool (auth puis test valeur) : ", auth)
		println("valeur id_user (id puis test valeur) : ", id)

		if !checkAuth || !auth || (id == 0) {
			http.Redirect(w, r, "/login", http.StatusMovedPermanently)
		}

		Data := DisplayCategory{"Les derniers topics", getTopicsIndex(page), "", PrevPage(page - 1), (page + 1), page, id}
		temp.ExecuteTemplate(w, "allTopic", Data)
	})

	type DisplayTopicId struct {
		DataTopic    DisplayTopic
		LikeMessages []LikeAndMessage
		Id_user      int
		Page         int
		PrevPage     int
		NextPage     int
		Id_ses       int
	}

	http.HandleFunc("/topic", func(w http.ResponseWriter, r *http.Request) {
		querise := r.URL.Query()

		page, errP := strconv.Atoi(querise.Get("page"))
		topic, errT := strconv.Atoi(querise.Get("topic"))
		if errP != nil || errT != nil {
			erreur.Message = "Page introuvable... Veuillez recommencer plus tard !"
			http.Redirect(w, r, "/erreur", http.StatusMovedPermanently)
		}

		session, _ := store.Get(r, "cookies-ses")
		auth, checkAuth := session.Values["auth"].(bool)
		id, _ := session.Values["id"].(int)

		println("valeur bool (auth puis test valeur) : ", auth)
		println("valeur id_user (id puis test valeur) : ", id)

		if !checkAuth || !auth || (id == 0) {
			http.Redirect(w, r, "/login", http.StatusMovedPermanently)
		}

		Data := DisplayTopicId{getTopicsId(topic), MessageAndLike(id, topic, page), id, page, PrevPage(page - 1), (page + 1), id}
		temp.ExecuteTemplate(w, "topic", Data)

	})
	http.HandleFunc("/send/", sendMessage)
	http.HandleFunc("/like", LikeMessages)

	type UpTopic struct {
		CreateTopic pageNewTopic
		TopicInfo   DisplayTopic
		Id          int
	}

	http.HandleFunc("/topic/edite/", func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		id_topic, _ := strconv.Atoi(r.URL.Path[13:])
		println("valeur de l'id_user : ", id_topic, path)
		Data := UpTopic{pageNewTopic{getCategory(), getTopicPicture(), 0}, getTopicsId(id_topic), getIdUserTopic(id_topic)}

		session, _ := store.Get(r, "cookies-ses")
		auth, _ := session.Values["auth"].(bool)
		id, _ := session.Values["id"].(int)
		if !auth || id == 0 || id != Data.Id {
			http.Redirect(w, r, "/login", http.StatusMovedPermanently)
		}

		temp.ExecuteTemplate(w, "editeTopic", Data)
	})
	http.HandleFunc("/upTopic", updateTopic)
	http.HandleFunc("/delTopic", DeleteTopic)

	type Profil struct {
		UserInfo User
		Topics   []DisplayTopic
		Messages []GetUserMessage
		Likes    []GetUserMessage
		Id       int
	}

	http.HandleFunc("/profil/", func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		id_user, _ := strconv.Atoi(r.URL.Path[8:])
		println("valeur de l'id_user : ", id_user, path)

		session, _ := store.Get(r, "cookies-ses")
		auth, _ := session.Values["auth"].(bool)
		id, _ := session.Values["id"].(int)
		if !auth || id == 0 || id != id_user {
			http.Redirect(w, r, "/accueil", http.StatusMovedPermanently)
		}

		Data := Profil{getUser(id_user), getTopicUser(id_user), getMessageUser(id_user), getMessageLikedUser(id_user), id}
		println("valeur de getUser : ", Data.UserInfo.Pseudo)
		temp.ExecuteTemplate(w, "profil", Data)
	})
	http.HandleFunc("/delMess", DeleteMessage)

	http.HandleFunc("/profil/edite/", func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		id_user, _ := strconv.Atoi(r.URL.Path[14:])
		println("valeur de l'id_user : ", id_user, path)

		session, _ := store.Get(r, "cookies-ses")
		auth, checkAuth := session.Values["auth"].(bool)
		id, _ := session.Values["id"].(int)

		println("valeur bool (auth puis test valeur) : ", auth, checkAuth)
		println("valeur id_user (id puis test valeur) : ", id)

		if !auth || id == 0 || id != id_user {
			http.Redirect(w, r, "/accueil", http.StatusMovedPermanently)
		}

		Data := getUser(id_user)
		temp.ExecuteTemplate(w, "eprofil", Data)
	})
	http.HandleFunc("/upProfil", updateProfile)

	// Page de test pour les problemes de cookies (go + chome + cookies = c'est compliqué)
	http.HandleFunc("/ts", func(w http.ResponseWriter, r *http.Request) {
		session, _ := store.Get(r, "cookies-ses")
		auth, ok := session.Values["auth"].(bool)
		id, ok2 := session.Values["id"].(int)

		if !ok || !auth || !ok2 || id == 0 {
			fmt.Fprintln(w, "tu es deco l'ami !")
		} else {
			fmt.Fprintln(w, "tu es connecter l'ami !")
		}
		session.Values["auth"] = false
		session.Values["id"] = 0
		session.Save(r, w)
		auth, ok = session.Values["auth"].(bool)
		id, ok2 = session.Values["id"].(int)
		println("valeur bool (auth puis test valeur) : ", auth, ok)
		println("valeur id_user (id puis test valeur) : ", id, ok2)

	})

	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		session, _ := store.Get(r, "cookies-ses")
		auth, checkAuth := session.Values["auth"].(bool)
		id, _ := session.Values["id"].(int)

		println("valeur bool (auth puis test valeur) : ", auth, checkAuth)
		println("valeur id_user (id puis test valeur) : ", id)
	})

	http.HandleFunc("/co", func(w http.ResponseWriter, r *http.Request) {
		session, _ := store.Get(r, "cookies-ses")

		session.Values["auth"] = true
		session.Values["id"] = 10
		session.Save(r, w)
		auth, ok := session.Values["auth"].(bool)
		id, ok2 := session.Values["id"].(int)
		println("valeur bool (auth puis test valeur) : ", auth, ok)
		println("valeur id_user (id puis test valeur) : ", id, ok2)

	})

	http.HandleFunc("/dco", func(w http.ResponseWriter, r *http.Request) {
		session, _ := store.Get(r, "cookies-ses")

		session.Values["auth"] = false
		session.Values["id"] = 0
		session.Save(r, w)

		auth, ok := session.Values["auth"].(bool)
		id, ok2 := session.Values["id"].(int)
		println("valeur bool (auth puis test valeur) : ", auth, ok)
		println("valeur id_user (id puis test valeur) : ", id, ok2)

	})

	// Mise en place du serveur de fichiers !
	rootDoc, _ := os.Getwd()
	fileserver := http.FileServer(http.Dir(rootDoc + "/asset"))
	http.Handle("/static/", http.StripPrefix("/static/", fileserver))

	println("serveur on !")
	http.ListenAndServe("localhost:8080", nil)

}
