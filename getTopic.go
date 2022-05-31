package main

type DisplayTopic struct {
	Id          int
	Title       string
	Subjet      string
	Category    string
	Description string
	Start_date  string
	Picture     string
	NbrMessage  int
	Pseudo      string
	Id_user     int
}

func getTopicsIndex(page int) []DisplayTopic {
	db := DbCon()
	getTopic, err := db.Query("SELECT t.`id_topic`, `title`, `subject`, `start_date`, `description`, tp.path, COUNT(m.message), (SELECT pseudo FROM users WHERE id_user = t.id_user)as pseudo FROM `topics` as t join `topics_pictures` as tp on t.id_topic_picture = tp.id_topic_picture join categorys as c on t.id_category = c.id_category LEFT JOIN messages as m on m.id_topic = t.id_topic where t.polity = \"1\" GROUP BY t.id_topic ORDER BY t.start_date DESC LIMIT 10 OFFSET ?;", (10 * page))
	if err != nil {
		panic(err.Error())
	}

	var allTopic []DisplayTopic

	for getTopic.Next() {
		var topic DisplayTopic
		err = getTopic.Scan(&topic.Id, &topic.Title, &topic.Subjet, &topic.Start_date, &topic.Description, &topic.Picture, &topic.NbrMessage, &topic.Pseudo)
		if err != nil {
			panic(err.Error())
		}
		allTopic = append(allTopic, topic)
	}

	return allTopic
}

func getTopicsCategory(category string, page int) []DisplayTopic {
	db := DbCon()
	getTopic, err := db.Query("SELECT t.`id_topic`, `title`, `subject`, `start_date`, `description`, tp.path, COUNT(m.message),(SELECT pseudo FROM users WHERE id_user = t.id_user)as pseudo FROM `topics` as t join `topics_pictures` as tp on t.id_topic_picture = tp.id_topic_picture join categorys as c on t.id_category = c.id_category LEFT JOIN messages as m on m.id_topic = t.id_topic where c.category = ? AND t.polity = \"1\" GROUP BY t.id_topic ORDER BY t.start_date DESC LIMIT 10 OFFSET ?;", category, (10 * page))
	if err != nil {
		panic(err.Error())
	}

	var allTopic []DisplayTopic

	for getTopic.Next() {
		var topic DisplayTopic
		err = getTopic.Scan(&topic.Id, &topic.Title, &topic.Subjet, &topic.Start_date, &topic.Description, &topic.Picture, &topic.NbrMessage, &topic.Pseudo)
		if err != nil {
			panic(err.Error())
		}
		allTopic = append(allTopic, topic)
	}

	return ReverseArryTopics(allTopic)
}

func getTopicsId(id int) DisplayTopic {
	db := DbCon()
	getTopic := db.QueryRow("SELECT t.`id_topic`, `title`, `subject`, `start_date`, `description`, tp.path, c.category, COUNT(m.message),(SELECT pseudo FROM users WHERE id_user = t.id_user)as pseudo FROM `topics` as t join `topics_pictures` as tp on t.id_topic_picture = tp.id_topic_picture join categorys as c on t.id_category = c.id_category LEFT JOIN messages as m on m.id_topic = t.id_topic where t.id_topic = ? AND t.polity = \"1\";", id)

	var topic DisplayTopic

	err := getTopic.Scan(&topic.Id, &topic.Title, &topic.Subjet, &topic.Start_date, &topic.Description, &topic.Picture, &topic.Category, &topic.NbrMessage, &topic.Pseudo)
	if err != nil {
		println("erreur getTopicsId ==> ", err.Error())
	}
	return topic
}

func getTopicUp(id int) DataCreateTopic {
	db := DbCon()
	getTopic := db.QueryRow("SELECT `title`, `subject`, `description`, `id_topic_picture`, `id_category` FROM `topics` WHERE `id_topic` = ?;", id)

	var topic DataCreateTopic

	err := getTopic.Scan(&topic.Title, &topic.Subjet, &topic.Description, &topic.Id_topic_picture, &topic.Id_category)
	if err != nil {
		println("getTopicUp ", err.Error())
	}
	return topic
}

func getIdUserTopic(id int) int {
	db := DbCon()
	getTopic := db.QueryRow("SELECT id_user FROM `topics` WHERE `id_topic` = ?;", id)

	var id_user int

	err := getTopic.Scan(&id_user)
	if err != nil {
		println("getTopicUp ", err.Error())
		return 0
	}
	return id_user
}

func ReverseArryTopics(Tab []DisplayTopic) []DisplayTopic {
	for i, j := 0, len(Tab)-1; i < j; i, j = i+1, j-1 {
		Tab[i], Tab[j] = Tab[j], Tab[i]
	}
	return Tab
}
