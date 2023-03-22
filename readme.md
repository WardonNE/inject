Usage
```
go get -u github.com/wardonne/inject
```

Init
```
container := inject.NewContainer()
```

Example
```golang
type User struct {
    ID int
    Name string
    Tags []string `inject:""`
    Attributes map[string]string `inject:""`
    Group *UserGroup `inject:"usergroup"`
}

func (u *User) Init() {
    u.ID = 1
    u.Name = "user"
    u.Tags = append(u.Tags, "tag-a")
    u.Attributes["gender"] = "female"
}

type UserGroup struct {
    ID int
    Name string
}

func (u *UserGroup) Init() {
    u.ID = 1
    u.Name = "group1"
}

container := inject.NewContainer()
user := new(User)
container.Porvide("user", user)

fmt.Println(user.ID) // 1
fmt.Println(user.Name) // user
fmt.Println(user.Tags[0]) // tag-a
fmt.Println(user.Attributes["gender"]) // female
fmt.Println(user.Group.Name) // group1
```