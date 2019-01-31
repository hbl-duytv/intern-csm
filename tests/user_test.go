package tests

import (
	"strconv"
	"testing"

	"github.com/hbl-duytv/intern-csm/helper"
	"github.com/hbl-duytv/intern-csm/services"
)

func TestGetAllEditorUser(t *testing.T) {
	dataCorrect := 5
	resultTest, err := services.GetAllEditorUser()
	if len(resultTest) != dataCorrect || err != nil {
		t.Errorf("Result test : %d , Result correct : %d", len(resultTest), dataCorrect)
	}
}
func TestGetUserByUsername(t *testing.T) {
	dataTest := []string{"anhnt1", "anhnt2", "anhnt3", "vip2", "admin"}
	dataCorrect := []int{41, 42, 0, 43, 14}
	for index, value := range dataTest {
		resultTest, _ := services.GetUserByUsername(value)
		if resultTest.ID != dataCorrect[index] {
			t.Errorf("Result test : %d , Result correct : %d", resultTest.ID, dataCorrect[index])
		}
	}
}
func TestGetUserByUsernamePassword(t *testing.T) {
	dataTest := []map[string]string{
		{
			"username": "anhnt1",
			"password": "123456",
		},
		{
			"username": "anhnt2",
			"password": "123456",
		},
		{
			"username": "anhnt3",
			"password": "123456",
		},
		{
			"username": "admin",
			"password": "12345678",
		},
		{
			"username": "admin",
			"password": "123456",
		},
	}
	dataCorrect := []int{41, 42, 0, 0, 14}
	for i, value := range dataTest {
		resultTest, _ := services.GetUserByUsernamePassword(value["username"], value["password"])
		if resultTest.ID != dataCorrect[i] {
			t.Errorf("Result test : %d , Result correct : %d", resultTest.ID, dataCorrect[i])
		}
	}
}
func TestGetUserByEmail(t *testing.T) {
	dataTest := []string{"anhnt1@hblab.vn", "anhnt2@gmail.com", "anhnt3@gmail.com", "vip2@gmail.com", "admin@hblab.vn"}
	dataCorrect := []int{41, 42, 0, 43, 14}
	for index, value := range dataTest {
		resultTest, _ := services.GetUserByEmail(value)
		if resultTest.ID != dataCorrect[index] {
			t.Errorf("Result test : %d , Result correct : %d", resultTest.ID, dataCorrect[index])
		}
	}
}
func TestGetUserByToken(t *testing.T) {
	dataTest := []string{"anhnt1", "anhnt2", "anhnt3", "vip2", "admin"}
	dataCorrect := []int{41, 42, 0, 43, 0}
	for index, value := range dataTest {
		resultTest, _ := services.GetUserByToken(helper.GetToken(value))
		if resultTest.ID != dataCorrect[index] {
			t.Errorf("Result test : %d , Result correct : %d", resultTest.ID, dataCorrect[index])
		}
	}
}
func TestGetUserByID(t *testing.T) {
	dataTest := []int{41, 42, 20, 43, 14}
	dataCorrect := []string{"anhnt1", "anhnt2", "", "vip2", "admin"}
	for index, value := range dataTest {
		resultTest, _ := services.GetUserByID(strconv.Itoa(value))
		if resultTest.Username != dataCorrect[index] {
			t.Errorf("Result test : %s , Result correct : %s", resultTest.Username, dataCorrect[index])
		}
	}
}
func TestCreateUser(t *testing.T) {
	dataTest := []map[string]string{
		{
			"username":     "anhnt1",
			"password":     "123456",
			"name":         "Anh",
			"gender":       "Nam",
			"birthday":     "05/03/1997",
			"phone_number": "0961706497",
			"email":        "anhnt1@hblab.vn",
			"type":         "0",
			"status":       "0",
			"token":        "08970496B6310D4E2E84C99B4064DCCC9B45195D374724F65886AA5BFF97E2C5",
			"confirm":      "1",
		},
		{
			"username":     "anhnt53",
			"password":     "123456",
			"name":         "Anh",
			"gender":       "Nam",
			"birthday":     "05/03/1997",
			"phone_number": "0961706497",
			"email":        "anhnt53@hblab.vn",
			"type":         "0",
			"status":       "0",
			"token":        "A0BDEDD934AB69AFB46F729AB7151DA8652950E90C102913E9DD64F16C83FC9A",
			"confirm":      "1",
		},
		{
			"username":     "anhnt3",
			"password":     "123456",
			"name":         "Anh",
			"gender":       "Nam",
			"birthday":     "05/03/1997",
			"phone_number": "0961706497",
			"email":        "anhnt3@hblab.vn",
			"type":         "0",
			"status":       "0",
			"token":        "A0BDEDD934AB69AFB46F729AB7151DA8652950E90C102913E9DD64F16C83FC9A",
			"confirm":      "1",
		},
		{
			"username":     "kienmd",
			"password":     "12345678",
			"name":         "Kiên",
			"gender":       "Nam",
			"birthday":     "05/03/1997",
			"phone_number": "0961706497",
			"email":        "kienmd@hblab.vn",
			"type":         "0",
			"status":       "0",
			"token":        "2BF4AC2EEB547EF0B643EB2665E96BAAD91655AE8BA97CE16D1D26989C3A01EA",
			"confirm":      "1",
		},
	}
	for _, value := range dataTest {
		phoneNumber, _ := strconv.Atoi(value["phone_number"])
		typeUser, _ := strconv.Atoi(value["type"])
		statusUser, _ := strconv.Atoi(value["status"])
		confirmUser, _ := strconv.Atoi(value["confirm"])
		err := services.CreateUser(value["username"], value["password"], value["email"], value["name"], value["gender"],
			value["birthday"], phoneNumber, typeUser, statusUser, value["token"], confirmUser)
		if err != nil {
			t.Errorf("Error : %+v", err)
		}
	}
}
