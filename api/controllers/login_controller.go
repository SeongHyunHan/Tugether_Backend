package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/SeongHyunHan/Tugether/api/auth"
	"github.com/SeongHyunHan/Tugether/api/models"
	"github.com/SeongHyunHan/Tugether/api/responses"
	"github.com/SeongHyunHan/Tugether/api/utils/formaterror"
	"golang.org/x/crypto/bcrypt"
)

func (server *Server) Login(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	account := models.Account{}
	err = json.Unmarshal(body, &account)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	account.Prepare()
	err = account.Validate("login")

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	token, err := server.SignIn(account.UserName, account.Password)

	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusUnprocessableEntity, formattedError)
		return
	}
	responses.JSON(w, http.StatusOK, token)
}

func (server *Server) SignIn(username, password string) (string, error) {
	var err error

	account := models.Account{}

	err = server.DB.Debug().Model(models.Account{}).Where("username = ?", username).Take(&account).Error
	if err != nil {
		return "", err
	}
	err = models.VerifyPassword(account.Password, password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	return auth.CreateToken(account.ID)
}
