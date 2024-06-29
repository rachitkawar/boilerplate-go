
# Project Title

Golang API Boilerplate


## Features

- Authentication: Includes JWT-based authentication out-of-the-box.
- Database Migration: Integrated database migration for easy schema management.
- RESTful API Structure: Pre-defined RESTful routes and handlers.
- Modular Codebase: Organized code structure for easy scalability and maintenance.
- Live Reloading: Using air for hot reloading.
- Configuration Management: Simple and flexible configuration using environment variables.
- Social Login: Social Login Out of the box.


## Prerequisites

- Go 1.18+
- Make (optional, for running commands)
- Go [air](https://github.com/air-verse/air) 

## Installation

Clone the repository https://github.com/rachitkawar/boilerplate-go.git

```bash
  git clone https://github.com/rachitkawar/boilerplate-go.git my-project
  cd my-project
```


    
## Environment Variables

To run this project, you will need to add the following environment variables to your .env file

```
DSN=
GIN_MODE=
MIGRATION_PATH=src/internal/database/migrations/
GOOGLE_CLIENT_ID=
GOOGLE_CLIENT_SECRET=
GOOGLE_REDIRECT_URL=
GOOGLE_LOGIN_URL=
SOCIAL_LOGIN_SECRET_STATE=
JWT_SECRET=
```


## Migrations


To create migration:
```
make migrations-create {table_name}

```

Migration Up:
```
make migrations-up
```


Migration Down:
```
make migrations-down
```


## Run Locally

To run the app:
```
make dev
```




## Deployment

To build this app

```bash
make build
```

To run this app
```
make run
```


## License

This project is licensed under the MIT License. See the [LICENSE](https://github.com/rachitkawar/boilerplate-go/blob/master/LICENSE) file for details.

