package post_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/anthonymartz17/blog_platform_backend.git/internal/post"
	"github.com/anthonymartz17/blog_platform_backend.git/internal/post/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)



func TestGetPosts(t *testing.T){

	tt:= []struct{
		name string
		msg string
		data []post.Post
		wantErr error
		expectsSrvCall bool
	}{
		{
			name:"valid two item list",
			msg:"should succeed on valid request",
			data: []post.Post{
				{
					ID:        "1",
					UserID:    "123",
					Content:   "first post",
					CreatedAt: time.Now(),
				},
				{
					ID:        "2",
					UserID:    "124",
					Content:   "second post",
					CreatedAt: time.Now(),
				},
			},
			wantErr:nil,
			expectsSrvCall:true,
		},
		{
			name:"empty list",
			msg:"should succeed even if list is empty",
			data: []post.Post{},
			wantErr:nil,
			expectsSrvCall:true,
		},
		{
			name:"PostService error",
			msg:"should fail on PostService error",
			data: []post.Post{},
			wantErr:errors.New("failed to retrieve data from repository"),
			expectsSrvCall:false,
		},
	}


	for _,tc:=  range tt{
		t.Run(tc.name,func(t *testing.T) {
			//arrange
			 rr:= httptest.NewRecorder()
			 req:= httptest.NewRequest(http.MethodGet,"/post",nil)
		
			 ctrl := gomock.NewController(t)
			 defer ctrl.Finish()
		
			 PostService:= mocks.NewMockPostService(ctrl)
   
       expectedData:= tc.data
       var expectedErr error
			 
			 if tc.wantErr != nil{
				 expectedData = nil
				 expectedErr = tc.wantErr
			 }
				 PostService.EXPECT().GetPosts(gomock.Any()).Return(expectedData,expectedErr)



			 handler:= post.NewHandler(PostService)

			 //act
		   handler.GetPosts(rr,req)
			

			 
				 //assert
         if tc.wantErr != nil{
					 assert.Equal(t,http.StatusInternalServerError,rr.Result().StatusCode,tc.msg)
					 return
					 
					}
					
					var data []post.Post
					 err:= json.Unmarshal(rr.Body.Bytes(),&data)
					 require.NoError(t,err,"should not fail unmarshalling test response data")

					 assert.NotNil(t,data,tc.msg)
					 assert.Len(t,data,len(tc.data),tc.msg)
				 
		})

	}
	
}
// func TestDecodeReqBody(t *testing.T){

// }
// func TestCreate(t *testing.T){

// }
