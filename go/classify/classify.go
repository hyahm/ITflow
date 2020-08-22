package classify

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
	// Login = "restfulproject"
	// Login = "api"
	// Login = "type"
	UserGroup Classify = "usergroup"
	// Login = "header"
	Important Classify = "important"
	Level     Classify = "level"
	Position  Classify = "position"
)

func (c Classify) String() string {
	return string(c)
}

var CLASSIFY = []string{
	"login", "user", "statusgroup", "role",
	"bug", "version", "status", "usergroup",
	"project", "env", "important", "level", "position",
}
