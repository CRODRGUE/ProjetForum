package main

type Message struct {
	Id_message      int
	Message         string
	Date_message    string
	Pseudo          string
	Id_user_message int
	Id_Topic        int
}

type LikeAndMessage struct {
	Message Message
	Check   bool
}

func getAllMessage(id_topic int) []Message {
	db := DbCon()
	getMessages, err := db.Query("SELECT m.id_message, m.message, m.date_message, (SELECT u.pseudo FROM users as u WHERE u.id_user = m.id_user) as pseudo, m.id_user FROM `topics` as t join `topics_pictures` as tp on t.id_topic_picture = tp.id_topic_picture join categorys as c on t.id_category = c.id_category LEFT JOIN messages as m on m.id_topic = t.id_topic where t.id_topic = ? ORDER BY m.date_message DESC;", id_topic)
	if err != nil {
		panic(err.Error())
	}

	var allMessage []Message

	for getMessages.Next() {
		var message Message
		err = getMessages.Scan(&message.Id_message, &message.Message, &message.Date_message, &message.Pseudo, &message.Id_user_message)
		if err != nil {
			println("il y a pas de message dans le topic !")
			return allMessage
		}
		allMessage = append(allMessage, message)
	}
	return ReverseArryMess(allMessage)
}

func getIndexMessage(id_topic int, page int) []Message {
	db := DbCon()
	getMessages, err := db.Query("SELECT m.id_message, m.message, m.date_message, (SELECT u.pseudo FROM users as u WHERE u.id_user = m.id_user) as pseudo, m.id_user FROM `topics` as t join `topics_pictures` as tp on t.id_topic_picture = tp.id_topic_picture join categorys as c on t.id_category = c.id_category LEFT JOIN messages as m on m.id_topic = t.id_topic where t.id_topic = ? ORDER BY m.date_message DESC LIMIT 10 OFFSET ?;", id_topic, (10 * page))
	if err != nil {
		panic(err.Error())
	}

	var allMessage []Message

	for getMessages.Next() {
		var message Message
		err = getMessages.Scan(&message.Id_message, &message.Message, &message.Date_message, &message.Pseudo, &message.Id_user_message)
		if err != nil {
			println("il y a pas de message dans le topic !")
			return allMessage
		}
		allMessage = append(allMessage, message)
	}
	return ReverseArryMess(allMessage)
}

func MessageAndLike(id_user, id_topic, page int) []LikeAndMessage {
	var allLikeAndMessage []LikeAndMessage
	Message := getIndexMessage(id_topic, page)
	Like := getLikeUserTopic(id_user, id_topic)
	for _, element := range Message {
		temp := LikeAndMessage{element, false}
		for _, elementLike := range Like {
			if temp.Message.Id_message == elementLike.id_message {
				temp.Check = true
				break
			}
		}
		allLikeAndMessage = append(allLikeAndMessage, temp)
	}
	return allLikeAndMessage
}

func ReverseArryMess(Tab []Message) []Message {
	for i, j := 0, len(Tab)-1; i < j; i, j = i+1, j-1 {
		Tab[i], Tab[j] = Tab[j], Tab[i]
	}
	return Tab
}
