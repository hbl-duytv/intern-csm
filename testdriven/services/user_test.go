package testdriven

import (
	"testing"

	"github.com/hbl-duytv/intern-csm/services"

	"github.com/hbl-duytv/intern-csm/libfn"
	"github.com/hbl-duytv/intern-csm/models"
)

func TestGetAllEditorUser(t *testing.T) {
	want := []models.User{
		{
			ID:          40,
			Username:    "vip1",
			Password:    "e10adc3949ba59abbe56e057f20f883e",
			Email:       "vip1@gmail.com",
			Type:        0,
			Name:        "vip1",
			Gender:      "Nam",
			PhoneNumber: 235235123,
			Status:      1,
			CreatedAt:   libfn.ParseTime("2018-05-25 10:36:13"),
			UpdatedAt:   libfn.ParseTime("2019-01-29 16:14:02"),
			Token:       "d397d6dd4c1df859c6f5530006dd561357d95cd4360715110e8df263416fc467",
			Confirm:     1,
			Birthday:    "25/09/1994",
			TimeConfirm: 8,
		},
		{
			ID:          41,
			Username:    "anhnt1",
			Password:    "e10adc3949ba59abbe56e057f20f883e",
			Email:       "anhnt1@hblab.vn",
			Type:        0,
			Name:        "Tuấn Anh",
			Gender:      "Nam",
			PhoneNumber: 12345678,
			Status:      1,
			CreatedAt:   libfn.ParseTime("2019-01-25 10:44:47"),
			UpdatedAt:   libfn.ParseTime("2019-01-29 16:14:02"),
			Token:       "08970496b6310d4e2e84c99b4064dccc9b45195d374724f65886aa5bff97e2c5",
			Confirm:     1,
			Birthday:    "06/06/1993",
			TimeConfirm: 8,
		},
		{
			ID:          44,
			Username:    "vip3",
			Password:    "e10adc3949ba59abbe56e057f20f883e",
			Email:       "vip3@gmail.com",
			Type:        0,
			Name:        "Vip3",
			Gender:      "Nam",
			PhoneNumber: 12345678,
			Status:      1,
			CreatedAt:   libfn.ParseTime("2019-01-25 11:12:08"),
			UpdatedAt:   libfn.ParseTime("2019-01-29 16:14:02"),
			Token:       "d879966df1d5a4d57208916aca869157148ac5d025d15ec678f4d8bbbda340e3",
			Confirm:     1,
			Birthday:    "03/03/1993",
			TimeConfirm: 8,
		},
	}
	if input, err := services.GetAllEditorUser(); err == nil {
		if !libfn.CompareTwoArrUser(want, input) {
			t.Error("get all editor user failed")
		}
	} else {
		t.Errorf("error not found: %s", err)
	}

}
func TestGetUserByToken(t *testing.T) {
	want := models.User{
		ID:          40,
		Username:    "vip1",
		Password:    "e10adc3949ba59abbe56e057f20f883e",
		Email:       "vip1@gmail.com",
		Type:        0,
		Name:        "vip1",
		Gender:      "Nam",
		PhoneNumber: 235235123,
		Status:      1,
		CreatedAt:   libfn.ParseTime("2018-05-25 10:36:13"),
		UpdatedAt:   libfn.ParseTime("2019-01-29 16:14:02"),
		Token:       "d397d6dd4c1df859c6f5530006dd561357d95cd4360715110e8df263416fc467",
		Confirm:     1,
		Birthday:    "25/09/1994",
		TimeConfirm: 8,
	}
	if input, err := services.GetUserByToken("d397d6dd4c1df859c6f5530006dd561357d95cd4360715110e8df263416fc467"); err == nil {
		if !libfn.CompareTwoUser(want, input) {
			t.Error("get user by token failed")
		}
	} else {
		t.Errorf("error not found: %s", err)
	}

}
func TestGetUserByUsername(t *testing.T) {
	want := models.User{
		ID:          40,
		Username:    "vip1",
		Password:    "e10adc3949ba59abbe56e057f20f883e",
		Email:       "vip1@gmail.com",
		Type:        0,
		Name:        "vip1",
		Gender:      "Nam",
		PhoneNumber: 235235123,
		Status:      1,
		CreatedAt:   libfn.ParseTime("2018-05-25 10:36:13"),
		UpdatedAt:   libfn.ParseTime("2019-01-29 16:14:02"),
		Token:       "d397d6dd4c1df859c6f5530006dd561357d95cd4360715110e8df263416fc467",
		Confirm:     1,
		Birthday:    "25/09/1994",
		TimeConfirm: 8,
	}
	if input, err := services.GetUserByUsername("vip1"); err == nil {
		if !libfn.CompareTwoUser(want, input) {
			t.Error("get user by username failed")
		}
	} else {
		t.Errorf("error not found: %s", err)
	}

}
func TestGetUserByEmail(t *testing.T) {
	want := models.User{
		ID:          40,
		Username:    "vip1",
		Password:    "e10adc3949ba59abbe56e057f20f883e",
		Email:       "vip1@gmail.com",
		Type:        0,
		Name:        "vip1",
		Gender:      "Nam",
		PhoneNumber: 235235123,
		Status:      1,
		CreatedAt:   libfn.ParseTime("2018-05-25 10:36:13"),
		UpdatedAt:   libfn.ParseTime("2019-01-29 16:14:02"),
		Token:       "d397d6dd4c1df859c6f5530006dd561357d95cd4360715110e8df263416fc467",
		Confirm:     1,
		Birthday:    "25/09/1994",
		TimeConfirm: 8,
	}
	if input, err := services.GetUserByEmail("vip1@gmail.com"); err == nil {
		if !libfn.CompareTwoUser(input, want) {
			t.Error("get user by email failed")
		}
	} else {
		t.Errorf("error not found: %s", err)
	}

}
func TestUpdateStatusUser(t *testing.T) {
	input := services.UpdateStatusUser(44, 1)
	if input != nil {
		t.Errorf("update status user failed: %s", input)
	}

}
func TestGetUserByID(t *testing.T) {
	want := models.User{
		ID:          44,
		Username:    "vip3",
		Password:    "e10adc3949ba59abbe56e057f20f883e",
		Email:       "vip3@gmail.com",
		Type:        0,
		Name:        "Vip3",
		Gender:      "Nam",
		PhoneNumber: 12345678,
		Status:      1,
		CreatedAt:   libfn.ParseTime("2019-01-25 11:12:08"),
		UpdatedAt:   libfn.ParseTime("2019-01-31 14:27:54"),
		Token:       "d879966df1d5a4d57208916aca869157148ac5d025d15ec678f4d8bbbda340e3",
		Confirm:     1,
		Birthday:    "03/03/1993",
		TimeConfirm: 8,
	}
	if input, err := services.GetUserByID(44); err == nil {
		if !libfn.CompareTwoUser(input, want) {
			t.Error("get user by id failed")
		}
	} else {
		t.Errorf("error not found: %s", err)
	}
}

func TestConfirmRegisterUser(t *testing.T) {
	err := services.ConfirmRegisterUser(44)
	if err != nil {
		t.Errorf("confirm register user failed: %v", err)
	}

}
func TestCreateUser(t *testing.T) {
	input := services.CreateUser("trungduc081", "123456", "nguyentrugduc248@gmail.com", 1, 0)
	if input != nil {
		t.Errorf("create user failed: %s", input)
	}
}
func TestDeleteUser(t *testing.T) {
	input := services.DeleteUser(49)
	if input != nil {
		t.Errorf("delete user failed: %s", input)
	}

}
func TestRequireLogin(t *testing.T) {
	want := models.User{
		ID:          40,
		Username:    "vip1",
		Password:    "e10adc3949ba59abbe56e057f20f883e",
		Email:       "vip1@gmail.com",
		Type:        0,
		Name:        "vip1",
		Gender:      "Nam",
		PhoneNumber: 235235123,
		Status:      1,
		CreatedAt:   libfn.ParseTime("2018-05-25 10:36:13"),
		UpdatedAt:   libfn.ParseTime("2019-01-29 16:14:02"),
		Token:       "d397d6dd4c1df859c6f5530006dd561357d95cd4360715110e8df263416fc467",
		Confirm:     1,
		Birthday:    "25/09/1994",
		TimeConfirm: 8,
	}
	if input, err := services.RequireLogin("vip1", "123456"); err == nil {
		if !libfn.CompareTwoUser(want, input) {
			t.Error("require login failed")
		}
	} else {
		t.Errorf("error not found: %s", err)
	}

}
func TestGetTimeCreateUSer(t *testing.T) {
	want := []int{1, 2019}
	if month, year, err := services.GetTimeCreateUSer(14); err == nil {
		if month != want[0] || year != want[1] {
			t.Error("get time create user failed")
		}
	} else {
		t.Errorf("error not found: %s", err)
	}
}
func TestHasLimitTimeConfirm(t *testing.T) {
	check, err := services.HasLimitTimeConfirm(48)
	if err == nil {
		if !check {
			t.Error("limit time to confirm")
		}
	}
	// if !services.HasLimitTimeConfirm(48) {
	// 	t.Error("limit time to confirm")
	// }

}
