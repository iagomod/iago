
# Iago
Iago's Another Go ORM.

## Quick Start

Get the module.

```shell
go get github.com/iagomod/iago@latest
go install github.com/iagomod/iago@latest
```

Create a model.
```yaml
user:
  username:
    - type: text
      args:
        max-length: 50
  password:
    - type: text
      args:
        max-length: 100
```

Generate go file.
```shell
iago -o models/user.go -p models user.yaml
```

And work with the model.

```go
package main

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/iagomod/iago/pkg/driver/sqlite3"
	"github.com/iagomod/iago/pkg/iago"
	"github.com/iagomod/iago/pkg/model"
	"github.com/iagomod/iago/pkg/selection"

	"example/models"
)

func main() {
	ctx := context.Background()

	// Prepare Iago's Data Manager
	c, err := sqlite3.NewDriver(ctx, ".test.db", sql.LevelDefault)
	if err != nil {
		panic(err)
	}

	g := iago.NewManager(c)

	// Get User Descriptor
	userD := models.GetUserDescriptor()

	// Get and Apply Migrations
	migrations, err := g.GetMigrationsCtx(ctx, userD)
	if err != nil {
		panic(err)
	}

	if migrations != "" {
		_, err = g.RawExecCtx(ctx, migrations)
		if err != nil {
			panic(err)
		}
	}

	// Create a user
	user := &models.User{
		Username: "iago",
		Password: "iago-pwd",
	}

	// Insert the user into DB
	g.Buffer(user)
	err = g.FlushBufferCtx(ctx)
	if err != nil {
		panic(err)
	}

	// Select the user from DB
	sn := selection.NewSelection(userD).Filter(
		userD.SF_Username().Equal(model.String("iago")).
			And(userD.SF_Password().Equal(model.String("iago"))),
	)

	uu, err := g.SelectCtx(ctx, sn)
	if err != nil {
		panic(err)
	}

	if len(uu) == 0 {
		panic("user not found")
	}

	u := uu[0]

	// Print the user
	fmt.Println(u)
}

```