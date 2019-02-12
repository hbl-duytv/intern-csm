package testdriven

import (
	"testing"

	"github.com/hbl-duytv/intern-csm/constant"
	"github.com/hbl-duytv/intern-csm/services"

	"github.com/hbl-duytv/intern-csm/libfn"
	"github.com/hbl-duytv/intern-csm/models"
)

func TestGetAllEditorUser(t *testing.T) {
	var usersTest []models.User
	if err := services.DBTest.Where("type=? AND confirm=?", constant.DeactiveNumber, constant.ActiveNumber).Find(&usersTest).Error; err != nil {
		t.Errorf("error not found in db test: %s", err)
	}
	if input, err2 := services.GetAllEditorUser(); err2 == nil {
		if !libfn.CompareTwoArrUser(usersTest, input) {
			t.Error("get all editor user failed")
		}
	} else {
		t.Errorf("error not found: %s", err2)
	}

}
func TestGetUserByToken(t *testing.T) {
	var userTest models.User
	if err := services.DBTest.Where("username=?", "anhnt1").Find(&userTest).Error; err != nil {
		t.Errorf("error not found in db test: %s", err)
	}
	if input, err2 := services.GetUserByToken(userTest.Token); err2 == nil {
		if !libfn.CompareTwoUser(userTest, input) {
			t.Error("get user by token failed")
		}
	} else {
		t.Errorf("error not found: %s", err2)
	}

}
func TestGetUserByUsername(t *testing.T) {
	var userTest models.User
	if err := services.DBTest.Where("username=?", "anhnt1").Find(&userTest).Error; err != nil {
		t.Errorf("error not found in db test: %s", err)
	}
	if input, err2 := services.GetUserByUsername(userTest.Username); err2 == nil {
		if !libfn.CompareTwoUser(userTest, input) {
			t.Error("get user by username failed")
		}
	} else {
		t.Errorf("error not found: %s", err2)
	}

}
func TestGetUserByEmail(t *testing.T) {
	var userTest models.User
	if err := services.DBTest.Where("username=?", "anhnt1").Find(&userTest).Error; err != nil {
		t.Errorf("error not found in db test: %s", err)
	}
	if input, err2 := services.GetUserByEmail(userTest.Email); err2 == nil {
		if !libfn.CompareTwoUser(input, userTest) {
			t.Error("get user by email failed")
		}
	} else {
		t.Errorf("error not found: %s", err2)
	}

}
func TestUpdateStatusUser(t *testing.T) {
	input := services.UpdateStatusUser(44, 1)
	if input != nil {
		t.Errorf("update status user failed: %s", input)
	}

}
func TestGetUserByID(t *testing.T) {
	var userTest models.User
	if err := services.DBTest.Where("username=?", "anhnt1").Find(&userTest).Error; err != nil {
		t.Errorf("error not found in db test: %s", err)
	}
	if input, err := services.GetUserByID(userTest.ID); err == nil {
		if !libfn.CompareTwoUser(input, userTest) {
			t.Error("get user by id failed")
		}
	} else {
		t.Errorf("error not found: %s", err)
	}
}

func TestConfirmRegisterUser(t *testing.T) {
	var userTest models.User
	if err := services.DBTest.Where("username=?", "anhnt1").Find(&userTest).Error; err != nil {
		t.Errorf("error not found in db test: %s", err)
	}
	err2 := services.ConfirmRegisterUser(userTest.ID)
	if err2 != nil {
		t.Errorf("confirm register user failed: %v", err2)
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
	var userTest models.User
	if err := services.DBTest.Where("username=?", "anhnt1").Find(&userTest).Error; err != nil {
		t.Errorf("error not found in db test: %s", err)
	}
	if input, err2 := services.RequireLogin(userTest.Username, "123456"); err2 == nil {
		if !libfn.CompareTwoUser(userTest, input) {
			t.Error("require login failed")
		}
	} else {
		t.Errorf("error not found: %s", err2)
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
}
