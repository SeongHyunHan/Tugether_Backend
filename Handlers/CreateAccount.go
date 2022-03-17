package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/SeongHyunHan/Tugether/models"
)

func (h handler) CreateAccount(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	//err := ioutil.ReadAll(r.Body)

	// if err != nil {
	// 	log.Fatalln(err)
	// }

	var accounts models.Accounts
	//json.Unmarshal(body, &accounts)

	accounts = models.Accounts{AutoId: 1, UserName: "Test", Password: "Test"}

	if result := h.DB.Create(&accounts); result.Error != nil {
		fmt.Println(result.Error)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("Account Created")
}
