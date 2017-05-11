package main

import ()

var roles map[string]uint

func init() {
  roles = map[string]uint {
    "admin":      1,
    "manager":    2,
    "moderator":  4,
  }
}

type User struct {
  roles uint
}

func NewUser() *User {
  return &User{ 0 }
}

func (user *User) IsAdmin() bool {
  return user.roles & roles["admin"] > 0
}

func (user *User) IsManager() bool {
  return user.roles & roles["manager"] > 0
}

func (user *User) IsModerator() bool {
  return user.roles & roles["moderator"] > 0
}

func (user *User) Update(params []string) {
  user.roles = 0

  for _, role := range params {
    user.roles += roles[role]
  }
}

func (user *User) Roles() (results []string) {
  results = []string{}

  for name, value := range roles {
    if user.roles & value > 0 {
      results = append(results, name)
    }
  }

  return
}
