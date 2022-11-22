package sms

import (
	"errors"
	"reflect"
	"testing"
)

import (
	gomock "github.com/golang/mock/gomock"
)

func TestSendMessageHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	m := NewMockSMSSender(ctrl)

	var testcase = []struct {
		expectedError error
		to            string
		msg           string
		mockCall      *gomock.Call
	}{
		{expectedError: nil, to: "8439043294", msg: "This is a new message", mockCall: m.EXPECT().Send(gomock.Any(), gomock.Any()).Return(nil)},
		{expectedError: errors.New("invalid phone"), to: "2494234ufytsjvhjk", msg: "This is a new message smnf mdsbfsdf sdfjdsfhhsd sdffsdf fdsfsd fdsf sd fsdff", mockCall: nil},
		{expectedError: errors.New("invalid sms message"), to: "8439043294", msg: "This is a new messageyftukgilhjrqghou;i5324 huoijrtwhih4iouti452io5io4oi5h445ioh4oi5", mockCall: nil},
	}

	for _, tt := range testcase {
		receivedError := New(m).SendMessage(tt.to, tt.msg)
		if !reflect.DeepEqual(receivedError, tt.expectedError) {
			t.Errorf("Failed")
		}
	}

}
