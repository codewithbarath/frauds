
package api

import (
	"busbooking/db"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
)

// Signup request details
type signupRequest struct {
	Username string `json:"username"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Database structure
type userstruct struct {
	UserID   int    `gorm:"column:userid;primaryKey;autoIncrement" json:"userId"`
	Username string `gorm:"column:username" json:"username"`
	Phone    string `gorm:"column:phone" json:"phone"`
	Email    string `gorm:"column:email" json:"email"`
	Password string `gorm:"column:password" json:"-"`
}

// Response structure
type signupResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	ErrorMsg string `json:"errorMsg"`
}

// Specify the database table name
func (userstruct) TableName() string {
	return "userss"
}

// Main signup API handler
func Signup(writer http.ResponseWriter, request *http.Request) {

	log.Println("Signup(+)")

	writer.Header().Set("Content-Type", "application/json")
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	// Handle browser preflight request
	if request.Method == http.MethodOptions {
		writer.WriteHeader(http.StatusNoContent)
		log.Println("Signup(-)")
		return
	}

	var lRespRec signupResponse

	// Check the request method
	if request.Method != http.MethodPost {
		writer.WriteHeader(http.StatusMethodNotAllowed)

		lRespRec.Status = "E"
		lRespRec.ErrorMsg = "Invalid Method"

		completeSignupMethod(writer, lRespRec)

		log.Println("Signup(-)")
		return
	}

	// Collect signup details
	lReqRec, lErr := collectSignupMethod(request)

	if lErr != nil {
		writer.WriteHeader(http.StatusBadRequest)

		lRespRec.Status = "E"
		lRespRec.ErrorMsg = "ASU:001 " + lErr.Error()

		completeSignupMethod(writer, lRespRec)

		log.Println("Signup(-)")
		return
	}

	// Communicate with the database
	lErr = communicateSignupMethod(lReqRec)

	if lErr != nil {
		writer.WriteHeader(http.StatusInternalServerError)

		lRespRec.Status = "E"
		lRespRec.ErrorMsg = "ASU:002 " + lErr.Error()

		completeSignupMethod(writer, lRespRec)

		log.Println("Signup(-)")
		return
	}

	// Construct successful response
	lRespRec.Status = "S"
	lRespRec.Message = "User created successfully"

	// Complete response
	completeSignupMethod(writer, lRespRec)

	log.Println("Signup(-)")
}

// Collect signup details from Vue
func collectSignupMethod(
	request *http.Request,
) (lReqRec signupRequest, lErr error) {

	log.Println("collectSignupMethod(+)")

	lByte, lErr := io.ReadAll(request.Body)

	if lErr != nil {
		log.Println("collectSignupMethod(-)")
		return lReqRec, fmt.Errorf("ASCO:001 %s", lErr.Error())
	}

	lErr = json.Unmarshal(lByte, &lReqRec)

	if lErr != nil {
		log.Println("collectSignupMethod(-)")
		return lReqRec, fmt.Errorf("ASCO:002 %s", lErr.Error())
	}

	log.Println("collectSignupMethod(-)")

	return lReqRec, nil
}

// Call the database-working method
func communicateSignupMethod(
	lReqRec signupRequest,
) (lErr error) {

	log.Println("communicateSignupMethod(+)")

	lErr = CreateUser(lReqRec)

	if lErr != nil {
		log.Println("communicateSignupMethod(-)")
		return errors.New("ASCM:001 " + lErr.Error())
	}

	log.Println("communicateSignupMethod(-)")

	return nil
}

// Insert signup details into MariaDB
func CreateUser(
	lReqRec signupRequest,
) (lErr error) {

	log.Println("CreateUser(+)")

	// Connect to the database
	lDB, lErr := db.ConnectDB()

	if lErr != nil {
		log.Println("CreateUser(-)")
		return errors.New("ACU:001 " + lErr.Error())
	}

	// Assign the received details to the database structure
	lUserData := userstruct{
		Username: lReqRec.Username,
		Phone:    lReqRec.Phone,
		Email:    lReqRec.Email,
		Password: lReqRec.Password,
	}

	// Insert the record into the userss table
	lErr = lDB.Create(&lUserData).Error

	if lErr != nil {
		log.Println("CreateUser(-)")
		return errors.New("ACU:002 " + lErr.Error())
	}

	log.Println("CreateUser(-)")

	return nil
}

// Send the response back to Vue
func completeSignupMethod(
	writer http.ResponseWriter,
	lRespRec signupResponse,
) {

	log.Println("completeSignupMethod(+)")

	lRespByte, lErr := json.Marshal(lRespRec)

	if lErr != nil {
		log.Println("Error while marshalling response:", lErr)
		return
	}

	_, lErr = writer.Write(lRespByte)

	if lErr != nil {
		log.Println("Error while writing response:", lErr)
	}

	log.Println("completeSignupMethod(-)")
}

