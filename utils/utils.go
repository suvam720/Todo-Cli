package utils

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Postask struct {
	ID        primitive.ObjectID `json:"_id"`
	CreatedAt time.Time          `json:"created_at"`
	UpdatedAt time.Time          `json:"updated_at"`
	Text      string             `json:"text"`
	Completed bool               `json:"completed"`
}
//getting Primitive object id and storing in slice

func GetId() []primitive.ObjectID {

	var data []Postask

	res, err := http.Get("http://localhost:4000/api/tasks")
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal(body, &data)

	var Id []primitive.ObjectID
	for i := range data {
		str := data[i]
		Id = append(Id, str.ID)
	}

	return Id
}

func ObjectID(index int) primitive.ObjectID{

	var PobjectId primitive.ObjectID
		ObjectID := GetId()
		for i := range ObjectID {
			if index == i+1 {
				PobjectId = ObjectID[i]
			}
		}

	return PobjectId
}