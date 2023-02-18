package cons

import (
	"Graphql/graph/model/meetup"
	"Graphql/graph/model/user"
)

var Meetups = []*meetup.Meetup{
	{
		ID:          "1",
		Name:        "First meetup",
		Description: "First description",
		UserID:      "1",
	},
	{
		ID:          "2",
		Name:        "Secod meetup",
		Description: "Second description",
		UserID:      "2",
	},
}

var Users = []*user.User{
	{
		ID:       "1",
		Username: "Alek",
		Email:    "alek@email.com",
	},
	{
		ID:       "2",
		Username: "Daiyen",
		Email:    "daiyen@email.com",
	},
}
