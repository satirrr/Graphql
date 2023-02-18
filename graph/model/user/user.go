package user

import "Graphql/graph/model/meetup"

type User struct {
	ID       string           `json:"id"`
	Username string           `json:"username"`
	Email    string           `json:"email"`
	Meetups  []*meetup.Meetup `json:"meetups"`
}
