package action

import (
	"fmt"

	"github.com/urfave/cli"
)

// User demo
type User struct {
}

// Create demo
func (c *User) Create(cli *cli.Context) {
	uid := cli.Int("uid")
	username := cli.String("username")
	fmt.Println(uid, username)
}

// Delete demo
func (c *User) Delete(cli *cli.Context) {
	uid := cli.Int("uid1")
	username := cli.String("username1")
	uid = 1
	username = "delete username"
	fmt.Println(uid, username)
}
