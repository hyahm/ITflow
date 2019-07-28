package autodb

type Apilist struct {
	Id          int64
	Name        string `gorm:"type:varchar(100)"`
	Pid         int64  `gorm:"type:bigint"`
	Url         string `gorm:"type:varchar(200)"`
	Information string `gorm:"type:varchar(255)"`
	Opts        string `gorm:"type:varchar(100)"`
	Methods     string `gorm:"type:varchar(100)"`
	Resp        string `gorm:"type:text"`
	Result      string `gorm:"type:text"`
	Uid         int64  `gorm:"type:bigint"`
	Hid         int64  `gorm:"type:bigint"`
	Calltype    string `gorm:"type:varchar(20)"`
}

type User struct {
	Id         int64
	Nickname   string `gorm:"type:varchar(30);not null;unique"`
	Password   string `gorm:"type:varchar(40);not null;"`
	Email      string `gorm:"type:varchar(50);not null;unique"`
	Headimg    string `gorm:"type:varchar(100);default:''"`
	Createtime int64  `gorm:"type:bigint;default:0"`
	Createuid  int64  `gorm:"type:bigint;default:0"`
	Realname   string `gorm:"type:varchar(30);not null;unique"`
	Showstatus string `gorm:"type:varchar(200);default:''"`
	Disable    bool   `gorm:"type:boolean;default:false"`
	Bugsid     int64  `gorm:"type:bigint;default:0"`
	Level      int8   `gorm:"type:int8;default:2"`
	Rid        int64  `gorm:"type:bigint;default:0"`
	Jid        int64  `gorm:"type:bigint;default:0"`
}

type Status struct {
	Id   int64
	Name string `gorm:"type:varchar(30);not null;default:'';unique"`
}

type Roles struct {
	Id   int64
	Role string `gorm:"type:varchar(30);not null;default:'';unique"`
}

type Projectname struct {
	Id   int64
	Name string `gorm:"type:varchar(30);not null;default:'';unique"`
}

type Jobs struct {
	Id    int64
	Name  string `gorm:"type:varchar(30);not null;default:'';unique"`
	Level int64  `gorm:"type:bigint;not null;default:2"`
	Hypo  int64  `gorm:"type:varchar(30);not null;default:0"`
}

type Environment struct {
	Id      int64
	Envname string `gorm:"type:varchar(30);not null;default:'';unique"`
}

type Header struct {
	Id     int64
	Name   string `gorm:"type:varchar(30);not null;default:'';unique"`
	Hhids  string `gorm:"type:varchar(100);default:''"`
	Remark string `gorm:"type:varchar(30);default:''"`
}

type Usergroup struct {
	Id   int64
	Name string `gorm:"type:varchar(30);not null;default:'';unique"`
	Ids  string `gorm:"type:varchar(200);default:''"`
	Cuid int64  `gorm:"type:bigint;default:0"`
}

type Statusgroup struct {
	Id   int64
	Name string `gorm:"type:varchar(30);not null;default:'';unique"`
	Sids string `gorm:"type:varchar(200);default:''"`
}

type Defaultvalue struct {
	Status    int64 `gorm:"type:bigint;not null;default:0"`
	Important int64 `gorm:"type:bigint;not null;default:0"`
	Level     int64 `gorm:"type:bigint;not null;default:0"`
}

type Types struct {
	Id      int64
	Name    string `gorm:"type:varchar(30);not null;default:'';unique"`
	Types   int8   `gorm:"column:type;type:int8;not null;default:0"`
	Opts    string `gorm:"type:varchar(200);default:''"`
	Tid     int64  `gorm:"type:bigint;default:0"`
	Default string `gorm:"type:varchar(50);default:''"`
}

type Version struct {
	Id         int64
	Name       string `gorm:"type:varchar(30);not null;unique"`
	Urlone     string `gorm:"type:varchar(30);default:''"`
	Urltwo     string `gorm:"type:varchar(30);default:''"`
	Createtime string `gorm:"type:varchar(30);default:0"`
	Createuid  int64  `gorm:"type:bigint;not null"`
}

type Headerlist struct {
	Id int64
	K  string `gorm:"type:varchar(200);not null;unique"`
	V  string `gorm:"type:text;not null"`
}

type Sharefile struct {
	Id         int64
	Filepath   string `gorm:"type:varchar(200);not null"`
	Readuser   bool   `gorm:"type:boolean;default:false"`
	Rid        int64  `gorm:"type:bigint;default:0"`
	Isfile     bool   `gorm:"type:boolean;default:false"`
	Ownerid    int64  `gorm:"type:bigint;default:0"`
	Wid        int64  `gorm:"type:bigint;default:0"`
	Writeuser  bool   `gorm:"type:boolean;default:false"`
	Size       int64  `gorm:"type:bigint;default:0"`
	Updatetime int64  `gorm:"type:bigint;default:0"`
	Name       string `gorm:"type:varchar(100);not null"`
}

type Importants struct {
	Id   int64
	Name string `gorm:"type:varchar(40);not null;default:'';unique"`
}

type Options struct {
	Id   int64
	Name string `gorm:"type:varchar(50);not null"`
	Info string `gorm:"type:varchar(100);default:''"`
	Tid  int64  `gorm:"type:bigint;default:0"`
	Df   string `gorm:"type:varchar(10);default:''"`
	Need string `gorm:"type:varchar(10);default:''"`
}

type Apiproject struct {
	Id       int64
	Name     string `gorm:"type:varchar(50);not null;default:'';unique"`
	Ownerid  int64  `gorm:"type:varchar(100);default:0"`
	Auth     bool   `gorm:"type:boolean;default:false"`
	Readuser bool   `gorm:"type:boolean;default:false"`
	Edituser bool   `gorm:"type:boolean;default:false"`
	Rid      bool   `gorm:"type:boolean;default:false"`
	Eid      bool   `gorm:"type:boolean;default:false"`
}

type Bugs struct {
	Id         int64
	Uid        int64  `gorm:"type:bigint;not null"`
	Title      string `gorm:"type:varchar(50);not null'"`
	Sid        int64  `gorm:"type:bigint;default:0"`
	Content    string `gorm:"type:text"`
	Ownerid    int64  `gorm:"type:bigint;default:0"`
	Iid        int64  `gorm:"type:bigint;default:0"`
	Createtime int64  `gorm:"type:bigint;default:0"`
	Vid        int64  `gorm:"type:bigint;default:0"`
	Spusers    int64  `gorm:"type:bigint;default:0"`
	Lid        int64  `gorm:"type:bigint;default:0"`
	Eid        int64  `gorm:"type:bigint;default:0"`
	Pid        int64  `gorm:"type:bigint;default:0"`
	Updatetime int64  `gorm:"type:bigint;default:0"`
	Dustbin    bool   `gorm:"type:boolean;default:false"`
}

type Log struct {
	Id       int64
	Exectime int64  `gorm:"type:bigint;default:0"`
	Classify string `gorm:"type:varchar(30);not null;default:''"`
	Content  string `gorm:"type:text"`
	Ip       string `gorm:"type:varchar(40);default:''"`
}

type Level struct {
	Id   int64
	Name string `gorm:"type:varchar(30);not null;default:'';unique"`
}

type Restfulname struct {
	Id   int64
	Name string `gorm:"type:varchar(30);not null;default:'';unique"`
}

type Informations struct {
	Id   int64
	Uid  int64  `gorm:"type:bigint;not null;default:0"`
	Bid  int64  `gorm:"type:bigint;not null;default:0"`
	Info string `gorm:"type:varchar(200);not null;default:''"`
	Time int64  `gorm:"type:bigint;default:0"`
}

type Email struct {
	Id         int64
	Email      string `gorm:"type:varchar(50);not null;unique"`
	Password   string `gorm:"type:varchar(50);default:''"`
	Port       int    `gorm:"type:int;default:25"`
	Createuser bool   `gorm:"type:boolean;default:false"`
	Createbug  bool   `gorm:"type:boolean;default:false"`
	Passbug    bool   `gorm:"type:boolean;default:false"`
}

type Rolegroup struct {
	Id       int64
	Name     string `gorm:"type:varchar(30);not null;default:'';unique"`
	Rolelist string `gorm:"type:varchar(200);default:''"`
}
