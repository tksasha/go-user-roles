package main_test

import (
  "testing"

  "github.com/stretchr/testify/assert"

  . "github.com/tksasha/user"
)

func TestIsAdmin(t *testing.T) {
  user := NewUser()

  assert.False(t, user.IsAdmin())

  user.Update([]string{ "admin" })

  assert.True(t, user.IsAdmin())
}

func TestIsManager(t *testing.T) {
  user := NewUser()

  assert.False(t, user.IsManager())

  user.Update([]string{ "manager" })

  assert.True(t, user.IsManager())
}

func TestIsModerator(t *testing.T) {
  user := NewUser()

  assert.False(t, user.IsModerator())

  user.Update([]string{ "moderator" })

  assert.True(t, user.IsModerator())
}

func TestUpdate(t *testing.T) {
  user := NewUser()

  assert.False(t, user.IsAdmin())

  assert.False(t, user.IsManager())

  assert.False(t, user.IsModerator())

  user.Update([]string{ "admin", "manager", "moderator" })

  assert.True(t, user.IsAdmin())

  assert.True(t, user.IsManager())

  assert.True(t, user.IsModerator())
}

func TestRoles(t *testing.T) {
  user := NewUser()

  assert.Equal(t, []string{}, user.Roles())

  user.Update([]string{ "admin" })

  assert.Equal(t, []string{ "admin" }, user.Roles())

  user.Update([]string{ "admin", "manager" })

  assert.Equal(t, []string{ "admin", "manager" }, user.Roles())
}
