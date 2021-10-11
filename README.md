# sample-go-echo-api
This is sample todo api.


## Commands
The following list of commands defined in the Makefile can be used.

|  Commands  | description                                                          |
|:----------:|:---------------------------------------------------------------------|
|  `setup`   | setup local environments and install dev dependencies.               |
|    `up`    | build and start containers.                                          |
| `rebuild`  | remove all containers and volumes then rebuild and start containers. |
|   `down`   | remove all containers and volumes.                                   |
|   `logs`   | output container logs.                                               |
| `exec-api` | enter the api container.                                             |
| `swagger`  | generate api documentation for swagger.                              |
| `migrate`  | migrate database using ent.                                          |


## Setup
Clone this repository and run `setup` command to create local environments.
```bash
$ make setup
```


Run `up` command to create a local containers.
```bash [Markdown Styles for all IDEA products · GitHub](https://gist.github.com/MikeMitterer/ff00ad4bb86ccaa4617e963eb0c08cf3#file-custom-darcula-css "Markdown Styles for all IDEA products · GitHub")
$ make up
```

## Migrations
Since this project uses [ent](https://entgo.io/ja/) as the ORM, we will follow the ent usage for things like schema generation and migration.

See the official documentation for details: [https://entgo.io](https://entgo.io/ja/docs/getting-started/).

### Schema generation and migrate
Use init command to generate schema file and add fields like this.
```bash
go run entgo.io/ent/cmd/ent init Todo
```

```go
package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Todo holds the schema definition for the Todo entity.
type Todo struct {
	ent.Schema
}

// Fields of the Todo.
func (Todo) Fields() []ent.Field {
	return []ent.Field{
		field.String("title"),
		field.String("description").Optional(),
		field.Enum("status").Values("TODO", "PROGRESS", "DONE").Default("TODO"),
	}
}

// Edges of the Todo.
func (Todo) Edges() []ent.Edge {
	return nil
}

// Indexes of the Todo.
func (Todo) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("status"),
	}
}
```

Enter the api container and run `migrate` command.
```bash
$ make exec-api
$ make migrate
```
