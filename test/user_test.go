package test

import "GinHello/initRouter"
import "github.com/stretchr/testify/assert"
import "net/http"
import "net/http/httptest"
import "testing"

func TestUserSave(t *testing.T)  {
	username:="zzichuan"
	router:=initRouter.SetupRouter()
	w:=httptest.NewRecorder()
	req,_:=http.NewRequest(http.MethodGet,"/user/"+username,nil)
	router.ServeHTTP(w,req)
	assert.Equal(t,http.StatusOK,w.Code)
	assert.Equal(t,"用户"+username+"已保存", w.Body.String())

}