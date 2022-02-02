package manager

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"test/models"
)

func GetAllPets() []models.Pet {
	pets := make([]models.Pet, 0)
	url := "http://petstore-demo-endpoint.execute-api.com/petstore/pets"

	req, _ := http.NewRequest("GET", url, nil)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(string(body))
	err := json.Unmarshal(body, &pets)
	if err != nil {
		log.Fatal(err)
	}
	return pets
}
func GetPetById(id string) models.Pet {
	var pet models.Pet
	//idAsString := strconv.FormatUint(uint64(id), 10)
	url := "http://petstore-demo-endpoint.execute-api.com/petstore/pets/" + id

	req, _ := http.NewRequest("GET", url, nil)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(string(body))
	err := json.Unmarshal(body, &pet)
	if err != nil {
		log.Fatal(err)
	}
	return pet
}
func saveInfo(req models.SavePet) {

}
