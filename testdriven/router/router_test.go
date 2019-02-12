package testdriven

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"testing"

	"github.com/hbl-duytv/intern-csm/routers"

	"github.com/gin-gonic/gin"
	"github.com/hbl-duytv/intern-csm/constant"
)

/* TEST METHOD POST*/
func TestLogin(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := routers.InitRouter()
	resp := httptest.NewRecorder()
	data := url.Values{}
	data.Set("username", "admin")
	data.Set("password", "123456")
	r, err := http.NewRequest("POST", "/login", bytes.NewBufferString(data.Encode()))
	if err != nil {
		t.Errorf("Couldn't create request: %v \n", err)
	}
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))
	router.ServeHTTP(resp, r)
	if resp.Code != constant.DirectStatus {
		t.Errorf("error status code: %v", resp.Code)
	}
}
func TestSendConfirmRegister(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := routers.InitRouter()
	resp := httptest.NewRecorder()
	data := url.Values{}
	dataCorrect := "Gửi mail xác nhận thành công, vui lòng check mail để xác nhận đăng ký tài khoản!"
	data.Set("username", "trungduc09")
	data.Set("password", "123456")
	data.Set("email", "nguyentrugduc248@gmail.com")
	r, err := http.NewRequest("POST", "/register", bytes.NewBufferString(data.Encode()))
	if err != nil {
		t.Errorf("Couldn't create request: %v \n", err)
	}
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))
	router.ServeHTTP(resp, r)
	if resp.Body.String() != dataCorrect {
		t.Errorf("error status code: %s", resp.Body.String())
	}
}
func TestCheckUserExist(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := routers.InitRouter()
	resp := httptest.NewRecorder()
	data := url.Values{}
	data.Set("username", "admin")

	r, err := http.NewRequest("POST", "/check-user-exist", bytes.NewBufferString(data.Encode()))
	if err != nil {
		t.Errorf("Couldn't create request: %v \n", err)
	}
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))
	router.ServeHTTP(resp, r)
	if resp.Code != http.StatusOK {
		t.Errorf("error status code: %v", resp.Code)
	}

}
func TestCheckEmailExist(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := routers.InitRouter()
	resp := httptest.NewRecorder()
	data := url.Values{}
	data.Set("email", "vip1@gmail.com")

	r, err := http.NewRequest("POST", "/check-email-exist", bytes.NewBufferString(data.Encode()))
	if err != nil {
		t.Errorf("Couldn't create request: %v \n", err)
	}
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))
	router.ServeHTTP(resp, r)
	if resp.Code != http.StatusOK {
		t.Errorf("error status code: %v", resp.Code)
	}

}
func TestActiveEditorUser(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := routers.InitRouter()
	resp := httptest.NewRecorder()
	data := url.Values{}
	data.Set("id", "51")
	r, err := http.NewRequest("POST", "/active-editor", bytes.NewBufferString(data.Encode()))
	if err != nil {
		t.Errorf("Couldn't create request: %v \n", err)
	}
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))
	router.ServeHTTP(resp, r)
	if resp.Code != http.StatusOK {
		t.Errorf("error status code: %v", resp.Code)
	}

}
func TestDeactiveEditorUser(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := routers.InitRouter()
	resp := httptest.NewRecorder()
	data := url.Values{}
	data.Set("id", "51")
	r, err := http.NewRequest("POST", "/deactive-editor", bytes.NewBufferString(data.Encode()))
	if err != nil {
		t.Errorf("Couldn't create request: %v \n", err)
	}
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))
	router.ServeHTTP(resp, r)
	if resp.Code != http.StatusOK {
		t.Errorf("error status code: %v", resp.Code)
	}

}
func TestDeleteUser(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := routers.InitRouter()
	resp := httptest.NewRecorder()
	data := url.Values{}
	data.Set("id", "51")
	r, err := http.NewRequest("POST", "/delete-user", bytes.NewBufferString(data.Encode()))
	if err != nil {
		t.Errorf("Couldn't create request: %v \n", err)
	}
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))
	router.ServeHTTP(resp, r)
	if resp.Code != http.StatusOK {
		t.Errorf("error status code: %v", resp.Code)
	}

}
func TestActiveStatusPost(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := routers.InitRouter()
	resp := httptest.NewRecorder()
	data := url.Values{}
	data.Set("id", "16")
	data.Set("comment", "test driven")
	r, err := http.NewRequest("POST", "/active-status-post", bytes.NewBufferString(data.Encode()))
	if err != nil {
		t.Errorf("Couldn't create request: %v \n", err)
	}
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))
	router.ServeHTTP(resp, r)
	if resp.Code != http.StatusOK {
		t.Errorf("error status code: %v", resp.Code)
	}

}
func TestDeActiveStatusPost(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := routers.InitRouter()
	resp := httptest.NewRecorder()
	data := url.Values{}
	data.Set("id", "16")
	data.Set("comment", "test driven lan 2")
	r, err := http.NewRequest("POST", "/deactive-status-post", bytes.NewBufferString(data.Encode()))
	if err != nil {
		t.Errorf("Couldn't create request: %v \n", err)
	}
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))
	router.ServeHTTP(resp, r)
	if resp.Code != http.StatusOK {
		t.Errorf("error status code: %v", resp.Code)
	}
}
func TestDeletePost(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := routers.InitRouter()
	resp := httptest.NewRecorder()
	data := url.Values{}
	data.Set("id", "100")
	data.Set("comment", "test driven lan 2")
	r, err := http.NewRequest("POST", "/delete-post", bytes.NewBufferString(data.Encode()))
	if err != nil {
		t.Errorf("Couldn't create request: %v \n", err)
	}
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))
	router.ServeHTTP(resp, r)
	if resp.Code != http.StatusOK {
		t.Errorf("error status code: %v", resp.Code)
	}

}
func TestCreatePost(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := routers.InitRouter()
	resp := httptest.NewRecorder()
	data := url.Values{}
	data.Set("creator", "44")
	data.Set("title", "test driven")
	data.Set("topic", "test driven topic")
	data.Set("description", "test driven description")
	data.Set("content", "test driven content")
	data.Set("tag", "t1,t2,t3")
	r, err := http.NewRequest("POST", "/create-post", bytes.NewBufferString(data.Encode()))
	if err != nil {
		t.Errorf("Couldn't create request: %v \n", err)
	}
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))
	router.ServeHTTP(resp, r)
	if resp.Code != http.StatusCreated {
		t.Errorf("error status code: %v", resp.Code)
	}
}
func TestUpdateContentPost(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := routers.InitRouter()
	resp := httptest.NewRecorder()
	data := url.Values{}
	data.Set("id", "183")
	data.Set("title", "test driven update")
	data.Set("topic", "update test driven topic")
	data.Set("description", "test driven description")
	data.Set("content", "test driven content")
	data.Set("tag", "t1,t2,t3")
	r, err := http.NewRequest("POST", "/update-content-post", bytes.NewBufferString(data.Encode()))
	if err != nil {
		t.Errorf("Couldn't create request: %v \n", err)
	}
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))
	router.ServeHTTP(resp, r)
	if resp.Code != http.StatusOK {
		t.Errorf("error status code: %v", resp.Code)
	}

}

/*METHOD GET*/
func TestRouterIndex(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := routers.InitRouter()
	router.LoadHTMLGlob("/home/hblab/work/src/github.com/hbl-duytv/intern-csm/views/html/*")
	// router.GET("/", controllers.Index)
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Errorf("Couldn't create request: %v \n", err)
	}
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	if resp.Code == http.StatusOK {
		t.Errorf("want status code: %d but receive status code: %v", http.StatusOK, resp.Body)
	} else {

	}
}
func TestRegisterSuccess(t *testing.T) {
	token := "18a62ae584fe344998438d8957f6bab8f54f82fb236c68232cbda39c58631b1e"
	gin.SetMode(gin.TestMode)
	router := routers.InitRouter()

	resp := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/confirm-register/"+token, nil)
	if err != nil {
		t.Errorf("Couldn't create request: %v \n", err)
	}
	router.ServeHTTP(resp, req)
	if resp.Code != http.StatusOK {
		// t.Errorf("want status code: %d but receive status code: %d", http.StatusOK, resp.Code)
		t.Errorf("err code: %v", resp.Code)
	}

}
func TestRenderEditorManagement(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := routers.InitRouter()

	router.LoadHTMLGlob("/home/hblab/work/src/github.com/hbl-duytv/intern-csm/views/html/*")
	resp1 := httptest.NewRecorder()
	data := url.Values{}
	data.Set("username", "admin")
	data.Set("password", "123456")
	r1, err1 := http.NewRequest("POST", "/login", bytes.NewBufferString(data.Encode()))
	if err1 != nil {
		t.Errorf("Couldn't create request: %v \n", err1)
	}
	r1.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r1.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))
	router.ServeHTTP(resp1, r1)
	if resp1.Code != constant.DirectStatus {
		t.Errorf("Error! , resp1 : %v \n", resp1)
	}
	resp := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/editor-management2", nil)
	if err != nil {
		t.Errorf("Couldn't create request: %v \n", err)
	}
	router.ServeHTTP(resp, req)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	if resp.Code == http.StatusOK {
		t.Errorf("want status code: %d but receive status code: %v", http.StatusOK, resp.Body)
	}
}