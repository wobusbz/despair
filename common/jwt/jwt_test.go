package jwt

import (
	"testing"
	"time"
)

type User struct {
	Id       int
	UserName string
}

func TestNewAuthToken(t *testing.T) {
	users := &User{Id: 1, UserName: "wobusbzzzzzzzzzzzz"}
	authToken := NewAuthToken()
	token, err := authToken.EncodeToken(users)

	if err != nil {
		t.Errorf("create token failed error %s\n", err)
	}

	var (
		curTime = time.Unix(time.Now().Unix(), 0)
	)

	if autoToken, err := authToken.DecodeToken(token); err == nil {
		if user, ok := autoToken.userInter.(*User); !ok {
			t.Errorf("NotBefore: %v  want = %v\n", user, users)
		}
		if curTime.Hour() != time.Unix(autoToken.IssuedAt, 0).Hour() {
			t.Errorf("NotBefore: %s  want = %d\n", time.Unix(autoToken.NotBefore, 0).Format("2006-01-02 15:04:05"), curTime.Day())
		}
		if curTime.Hour() != time.Unix(autoToken.IssuedAt, 0).Hour() {
			t.Errorf("IssuedAt: %s  want = %d\n", time.Unix(autoToken.IssuedAt, 0).Format("2006-01-02 15:04:05"), curTime.Day())
		}
		if time.Unix(autoToken.ExpiresAt, 0).Day() != curTime.Day()+7 {
			t.Errorf("ExpiresAt: %s want = %s\n", time.Unix(autoToken.ExpiresAt, 0).Format("2006-01-02 15:04:05"), curTime.Format("2006-01-02 15:04:05"))
		}
	} else {
		t.Errorf("decode token failed error %s", err)
	}
}
