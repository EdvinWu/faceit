# How to run

```bash
run 'docker-compose up -d' in root directory of project
```

## You can run go application without docker.

Run from console

```bash
go run cmd/faceit-test/main.go
```

Or you can run it using your favourite IDE

## API

# create USER

POST to localhost:8880/api/v1/user

```json
{
  "first_name": "name1",
  "last_name": "lastname",
  "nick_name": "nick1",
  "password": "pwd1",
  "email": "email",
  "country": "country"
}
```

# update USER

PUT to localhost:8880/api/v1/user/:id <br>
You can only update fields mentioned below

```json
{
  "first_name": "name1",
  "last_name": "lastname",
  "nick_name": "nick1"
}
```

# delete USER

GET to localhost:8880/api/v1/user/:id

# paginated USERS

GET to localhost:8880/api/v1/user?nick=nick&country=country&email=email&page=1&size=5 <br>
Fields available for search:

```text
field name -> parameter name

country -> country
nickname -> nick
email -> email
```

Paging parameters

```text
description -> parameter name

page number -> page
page size -> size
```

# How to run TESTS

```bash
run 'go text ./...' from root directory
```

To regenerate mocks run 'genmocks.sh' in script folder

# Explanations and Assumptions

DB choice: Plain SQL is enough for simple crud on 1 table <br>
DB library: Do not like ORM, so decided to use a bit more powerful sql.db library "sqlx"<br>
Web server choice: Echo provides testing library, possible to test handlers without integration tests<br>
Rabbit: Just to emulate "Have the ability to notify other interested services of changes to User entities"

Assumption: Far from production ready in logging/error handling since it's a test assignment


# Extensions or Improvements
To add:
production level logger,
error handling solution to handle all necessary db errors,
tracing,
config for according env(dev/stage/etc),
CI/CD scripts if needed,
deployment scripts/configuration if needed,
service versioning,

