package models_test

import (
	"demo/models"
	"fmt"
	"slices"
	"strings"
	"testing"
)

func TestUserValidateSuccess(t *testing.T) {
	user := &models.User{Name: "Jiten", Email: "jitenp@outlook.com", ContactNo: "9618558500"}
	err := user.Validate()
	if err != nil {
		t.Fail()
	}
}

func TestUserValidateNameFailure(t *testing.T) {
	user := &models.User{Email: "jitenp@outlook.com", ContactNo: "9618558500"}
	err := user.Validate()
	if err == nil {
		t.Fail()
	}
	fmt.Println(err.Error())
	if !strings.Contains(err.Error(), "invalid name") {
		t.Fail()
	}
}

func TestUserValidateEmailFailure(t *testing.T) {
	user := &models.User{Name: "Jiten", ContactNo: "9618558500"}
	err := user.Validate()
	if err == nil {
		t.Fail()
	}
	fmt.Println(err.Error())
	if !strings.Contains(err.Error(), "invalid email") {
		t.Fail()
	}
}

func TestUserToBytes(t *testing.T) {
	expected := []byte{123, 34, 105, 100, 34, 58, 48, 44, 34, 110, 97, 109, 101, 34, 58, 34, 74, 105, 116, 101, 110, 34, 44, 34, 101, 109, 97, 105, 108, 34, 58, 34, 74, 105, 116, 101, 110, 80, 64, 79, 117, 116, 108, 111, 111, 107, 46, 99, 111, 109, 34, 44, 34, 99, 111, 110, 116, 97, 99, 116, 95, 110, 111, 34, 58, 34, 57, 54, 49, 56, 53, 53, 56, 53, 48, 48, 34, 44, 34, 115, 116, 97, 116, 117, 115, 34, 58, 34, 34, 44, 34, 108, 97, 115, 116, 95, 117, 112, 100, 97, 116, 101, 100, 34, 58, 48, 125}
	user := &models.User{Name: "Jiten", Email: "JitenP@Outlook.com", ContactNo: "9618558500"}
	actualbytes, err := user.ToBytes()
	if err != nil {
		t.Fail()
	}

	areEqual := slices.Equal(expected, actualbytes)
	if !areEqual {
		t.Fail()
	}
}

func TestUserToString(t *testing.T) {
	expected := []byte{123, 34, 105, 100, 34, 58, 48, 44, 34, 110, 97, 109, 101, 34, 58, 34, 74, 105, 116, 101, 110, 34, 44, 34, 101, 109, 97, 105, 108, 34, 58, 34, 74, 105, 116, 101, 110, 80, 64, 79, 117, 116, 108, 111, 111, 107, 46, 99, 111, 109, 34, 44, 34, 99, 111, 110, 116, 97, 99, 116, 95, 110, 111, 34, 58, 34, 57, 54, 49, 56, 53, 53, 56, 53, 48, 48, 34, 44, 34, 115, 116, 97, 116, 117, 115, 34, 58, 34, 34, 44, 34, 108, 97, 115, 116, 95, 117, 112, 100, 97, 116, 101, 100, 34, 58, 48, 125}
	user := &models.User{Name: "Jiten", Email: "JitenP@Outlook.com", ContactNo: "9618558500"}
	actualbytes := user.ToString()
	if string(expected) != actualbytes {
		t.Fail()
	}
}
