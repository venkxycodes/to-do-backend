package service

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"to-do/repo"
)

func Test_userService_GetUserIdByUserName(t *testing.T) {
	type fields struct {
		userRepo            repo.UserRepository
		usernameToUserIdMap *domain.UsernameToUserIdMap
	}
	type args struct {
		username string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr assert.ErrorAssertionFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &userService{
				userRepo:            tt.fields.userRepo,
				usernameToUserIdMap: tt.fields.usernameToUserIdMap,
			}
			got, err := u.GetUserIdByUserName(tt.args.username)
			if !tt.wantErr(t, err, fmt.Sprintf("GetUserIdByUserName(%v)", tt.args.username)) {
				return
			}
			assert.Equalf(t, tt.want, got, "GetUserIdByUserName(%v)", tt.args.username)
		})
	}
}
