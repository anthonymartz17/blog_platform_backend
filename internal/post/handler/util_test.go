package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	entity "github.com/anthonymartz17/blog_platform_backend.git/internal/post"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestResponseJSON(t *testing.T){

	fixedTime := time.Date(2024, time.January, 1, 12, 0, 0, 0, time.UTC)

t.Run("Success response list of posts", func(t *testing.T) {
	//arrange
	respRec:= httptest.NewRecorder() // simulates an http.NewResponseWriter
	wantBody := []entity.Post{
		{
			ID:        "abc123",
			UserID:    "user_1",
			Content:   "First test post",
			CreatedAt: fixedTime,
			UpdatedAt: fixedTime,
		},
		{
			ID:        "def456",
			UserID:    "user_2",
			Content:   "Second test post",
			CreatedAt: fixedTime,
			UpdatedAt: fixedTime,
		},


	}
	
	//act
	ResponseJSON(respRec,http.StatusOK,wantBody)

	//assert
	assert.Equal(t,"application/json",respRec.Result().Header.Get("Content-Type"),"should set header: Content-Type:application/json")
	assert.Equal(t,http.StatusOK,respRec.Result().StatusCode,"should set status code 200")

	var gotBody []entity.Post

	err:= json.Unmarshal(respRec.Body.Bytes(),&gotBody); 
	require.NoError(t,err,"should not fail to unmarshal test body response")
	

	assert.NotNil(t,gotBody)
	assert.Len(t,gotBody,2)
	assert.Equal(t,wantBody[0].ID,gotBody[0].ID,"should match same post id")

})

t.Run("Success response single post", func(t *testing.T) {
	//arrange
	respRec:= httptest.NewRecorder() // simulates an http.NewResponseWriter
	wantBody:= &entity.Post{
			ID:        "abc123",
			UserID:    "user_1",
			Content:   "First test post",
			CreatedAt: fixedTime,
			UpdatedAt: fixedTime,
		}

	
	//act
	ResponseJSON(respRec,http.StatusOK,wantBody)

	//assert
	assert.Equal(t,"application/json",respRec.Result().Header.Get("Content-Type"),"should set header: Content-Type:application/json")
	assert.Equal(t,http.StatusOK,respRec.Result().StatusCode,"should set status code 200")

	var gotBody entity.Post

	err:= json.Unmarshal(respRec.Body.Bytes(),&gotBody); 
	require.NoError(t,err,"should not fail to unmarshal test body response")
	

	assert.NotNil(t,gotBody)
	assert.Equal(t,wantBody.ID,gotBody.ID,"should match ID")
	assert.Equal(t,wantBody.Content,gotBody.Content,"should match Content")
	assert.Equal(t,wantBody.UserID,gotBody.UserID,"should match UserID")
	assert.Equal(t,wantBody.CreatedAt,gotBody.CreatedAt,"should match CreatedAt")
	assert.Equal(t,wantBody.UpdatedAt,gotBody.UpdatedAt,"should match UpdatedAt")

})


}
func TestResponseError(t *testing.T){
	//arrange
	responseRecorder:= httptest.NewRecorder()
	
	

	//act
	ResponseError(responseRecorder,http.StatusNotFound,"failed to find post")
  var resp map[string]string

	err:= json.Unmarshal(responseRecorder.Body.Bytes(),&resp)
	require.NoError(t,err,"should not fail to unmarshal test response")

	//assert
	assert.Equal(t,"application/json",responseRecorder.Result().Header.Get("Content-Type"))
	assert.Equal(t,http.StatusNotFound,responseRecorder.Result().StatusCode)
	assert.Equal(t,"failed to find post",resp["error"])

}