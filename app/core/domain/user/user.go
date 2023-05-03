package users

import "kitten-backend-go-app/app/core/domain/common"

// User is the entity about users in application
type User struct {
	ID          string
	Name        *common.Name
	Email       *Email
	Password    *Password
	Age         *Age
	MonthlyWage *Wage
	UserPicture *UserPicture
}
