package main

type GetUserMessage struct {
	Id_message      int
	Message         string
	Date_message    string
	Pseudo          string
	Id_user_message int
	Id_Topic        int
	Tiltle          string
}

type User struct {
	Pseudo         string `json:"pseudo"`
	FirstName      string `json:"firstname"`
	LastName       string `json:"lastname"`
	Description    string `json:"description"`
	ProfilePicture string `json:"profile_picture"`
	Email          string `json:"email"`
	Password       string `json:"password"`
	Date           string `json:"brith_day"`
	Id             int
}

func getUser(id int) User {
	println("voici la valeur de l'id : ", id)
	db := DbCon()
	var user User
	err := db.QueryRow("SELECT `lastname`, `firstname`, `email`, `password`, `pseudo`, `description`, `profile_picture`, `id_user`, `brith_day` FROM `users` WHERE id_user = ?", id).Scan(&user.LastName, &user.FirstName, &user.Email, &user.Password, &user.Pseudo, &user.Description, &user.ProfilePicture, &user.Id, &user.Date)
	if err != nil {
		println("getUser", err.Error())
	}
	return user
}

func getTopicUser(id int) []DisplayTopic {
	db := DbCon()
	getTopic, err := db.Query("SELECT t.`id_topic`, `title`, `subject`, `start_date`, `description`, tp.path, t.`id_user` FROM `topics` as t join `topics_pictures` as tp on t.id_topic_picture = tp.id_topic_picture join categorys as c on t.id_category = c.id_category LEFT JOIN messages as m on m.id_topic = t.id_topic where t.id_user = ? AND t.polity = \"1\"", id)
	if err != nil {
		println("getTopicUser", err.Error())
	}

	var allTopic []DisplayTopic

	for getTopic.Next() {
		var topic DisplayTopic
		err = getTopic.Scan(&topic.Id, &topic.Title, &topic.Subjet, &topic.Start_date, &topic.Description, &topic.Picture, &topic.Id_user)
		topic.NbrMessage = 0
		if err != nil {
			println("getTopicUser next", err.Error())
		}
		allTopic = append(allTopic, topic)
	}

	return ReverseArryTopics(allTopic)
}

func getMessageLikedUser(id int) []GetUserMessage {
	db := DbCon()
	getIdMessages, err := db.Query("SELECT m.`id_message`, m.`message`, m.`date_message`,(SELECT u.pseudo FROM users as u WHERE u.id_user = m.id_user) as pseudo, m.`id_user`, t.`title`, t.id_topic  FROM `messages` as m JOIN likes as l ON m.id_message = l.id_message JOIN topics as t ON m.`id_topic` = t.`id_topic` WHERE l.`id_user` = ?;", id)
	if err != nil {
		println("getMessageLikedUser", err.Error())
	}

	var allMessage []GetUserMessage

	for getIdMessages.Next() {
		var message GetUserMessage
		err = getIdMessages.Scan(&message.Id_message, &message.Message, &message.Date_message, &message.Pseudo, &message.Id_user_message, &message.Tiltle, &message.Id_Topic)
		if err != nil {
			println("getMessageLikedUser next", err.Error())
		}
		allMessage = append(allMessage, message)
	}
	return ReverseArryUserMess(allMessage)
}

func getMessageUser(id int) []GetUserMessage {
	db := DbCon()
	getMessages, err := db.Query("SELECT m.id_message, m.message, m.date_message, (SELECT u.pseudo FROM users as u WHERE u.id_user = m.id_user) as pseudo, m.id_user, t.`id_topic`, t.`title` FROM `topics` as t LEFT JOIN messages as m on m.id_topic = t.id_topic where m.id_user = ? ORDER BY m.date_message DESC;", id)
	if err != nil {
		println("getMessageUser", err.Error())
	}

	var allMessage []GetUserMessage

	for getMessages.Next() {
		var message GetUserMessage
		err = getMessages.Scan(&message.Id_message, &message.Message, &message.Date_message, &message.Pseudo, &message.Id_user_message, &message.Id_Topic, &message.Tiltle)
		if err != nil {
			println("getMessageUser next", err.Error())
			return allMessage
		}
		allMessage = append(allMessage, message)
	}
	return ReverseArryUserMess(allMessage)
}

func ReverseArryUserMess(Tab []GetUserMessage) []GetUserMessage {
	for i, j := 0, len(Tab)-1; i < j; i, j = i+1, j-1 {
		Tab[i], Tab[j] = Tab[j], Tab[i]
	}
	return Tab
}
