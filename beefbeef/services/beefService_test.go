package services

import (
	"beefbeef/ports/mocks"
	"github.com/stretchr/testify/mock"
	"io"
	"net/http"
	"strings"
	"testing"
)

func TestBeefService(t *testing.T) {
	t.Run("Count bacon", func(t *testing.T) {
		mockClient := new(mocks.HttpClient)
		srv := NewBeefService(mockClient)
		sr := strings.NewReader(`"beef beef ribeye beef ribeye"`)
		src := io.NopCloser(sr)
		resp := &http.Response{Status: "200", Body: src}
		mockClient.On("Do", mock.Anything).Return(resp, nil).Once()
		actual := srv.Count()

		if actual["ribeye"] != 2 {
			t.Error("ribeye should be 2")
		}
	})

	t.Run("Count", func(t *testing.T) {
		mockClient := new(mocks.HttpClient)
		srv := NewBeefService(mockClient)
		sr := strings.NewReader(`Fatback t-bone t-bone, pastrami  ..   t-bone.  pork, meatloaf jowl enim.  Bresaola t-bone.`)
		src := io.NopCloser(sr)
		resp := &http.Response{Status: "200", Body: src}
		mockClient.On("Do", mock.Anything).Return(resp, nil).Once()
		actual := srv.Count()

		if actual["t-bone"] != 4 {
			t.Error("t-bone should be 4")
		}
	})
}
