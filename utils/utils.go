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

func GetId() []string {
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
	var Id []string
	for i := range data {
		str := data[i]
		Id = append(Id, str.ID.String())
	}
	return Id

}

// func Get(ctx context.Context, url string) ([]byte, error) {

// 	req, err := http.NewRequest(http.MethodGet, url, nil)
// 	if err != nil {
// 		return []byte{}, err
// 	}
// 	req = req.WithContext(ctx)
// 	resp, err := http.DefaultClient.Do(req)
// 	if err != nil {
// 		return []byte{}, err
// 	}
// 	defer resp.Body.Close()
// 	return checkStatus(resp)
// }

// func Post(ctx context.Context, url string) ([]byte, error) {

// 	req, err := http.NewRequest(http.MethodPost, url, nil)
// 	if err != nil {
// 		return []byte{}, err
// 	}
// 	req = req.WithContext(ctx)
// 	resp, err := http.DefaultClient.Do(req)
// 	if err != nil {
// 		return []byte{}, err
// 	}
// 	defer resp.Body.Close()
// 	return checkStatus(resp)
// }

// func Put(ctx context.Context, url string) ([]byte, error) {

// 	req, err := http.NewRequest(http.MethodPut, url, nil)
// 	if err != nil {
// 		return []byte{}, err
// 	}
// 	req = req.WithContext(ctx)
// 	resp, err := http.DefaultClient.Do(req)
// 	if err != nil {
// 		return []byte{}, err
// 	}
// 	defer resp.Body.Close()
// 	return checkStatus(resp)
// }

// func Delete(ctx context.Context, url string) ([]byte, error) {

// 	req, err := http.NewRequest(http.MethodGet, url, nil)
// 	if err != nil {
// 		return []byte{}, err
// 	}
// 	req = req.WithContext(ctx)
// 	resp, err := http.DefaultClient.Do(req)
// 	if err != nil {
// 		return []byte{}, err
// 	}
// 	defer resp.Body.Close()
// 	return checkStatus(resp)
// }

// func checkStatus(h *http.Response) ([]byte, error) {
// 	if h.StatusCode == http.StatusOK {
// 		dataByte, err := ioutil.ReadAll(h.Body)
// 		if err != nil {
// 			return []byte{}, err
// 		}
// 		return dataByte, nil
// 	} else {
// 		return []byte{}, errors.New(strconv.Itoa(http.StatusNotFound))
// 	}
// }
