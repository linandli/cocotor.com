package news



type NewData struct {
	Title   string        `bson:"title"`
	Publish_time     string        `bson:"publish_time"`
	Website string        `bson:"website"`
	Group    string     `bson:"group"`
	Url    string     `bson:"url"`
}