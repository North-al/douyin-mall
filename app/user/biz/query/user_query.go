package query

import (
	"context"
	"math/rand"
	"strconv"
	"strings"

	"github.com/North-al/douyin-mall/app/user/model"
	"gorm.io/gorm"
)

type UserQuery struct {
	ctx context.Context
	db  *gorm.DB
}

func NewUserQuery(ctx context.Context, db *gorm.DB) *UserQuery {
	return &UserQuery{ctx: ctx, db: db}
}

func (q *UserQuery) GetUserByEmail(email string) (user model.User, err error) {
	err = q.db.WithContext(q.ctx).Model(&model.User{}).Where("email = ?", email).First(&user).Error
	return
}

func (q *UserQuery) GetUserById(id int32) (user model.User, err error) {
	err = q.db.WithContext(q.ctx).Model(&model.User{}).Where("id = ?", id).First(&user).Error
	return
}

func (q *UserQuery) CreateUser(email, password string) (user model.User, err error) {
	user = model.User{
		Email:    email,
		Password: password,
		Username: "user_" + strings.Split(email, "@")[0] + "_" + strconv.Itoa(rand.Intn(1000000)),
	}

	err = q.db.Create(&user).Error
	return
}
