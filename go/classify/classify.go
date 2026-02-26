package classify

import "itflow/response"

type Classify string

const (
	Login       Classify = "login"
	User        Classify = "user"
	Bug         Classify = "bug"
	Version     Classify = "version"
	Project     Classify = "project"
	Env         Classify = "env"
	StatusGroup Classify = "statusgroup"
	Role        Classify = "role"
	Status      Classify = "status"
	UserGroup   Classify = "usergroup"
	Important   Classify = "important"
	Level       Classify = "level"
	Position    Classify = "position"
)

func (c Classify) String() string {
	return string(c)
}

var CLASSIFY = []response.Option{
	{
		Label: "login",
		Value: 0,
	},
	{
		Label: "user",
		Value: 1,
	},
	{
		Label: "statusgroup",
		Value: 2,
	},
	{
		Label: "role",
		Value: 3,
	},
	{
		Label: "bug",
		Value: 4,
	}, {
		Label: "version",
		Value: 5,
	}, {
		Label: "status",
		Value: 6,
	},
	{
		Label: "usergroup",
		Value: 7,
	},

	{
		Label: "project",
		Value: 8,
	}, {
		Label: "env",
		Value: 9,
	}, {
		Label: "important",
		Value: 10,
	},
	{
		Label: "level",
		Value: 11,
	},
	{
		Label: "position",
		Value: 012,
	},
}
