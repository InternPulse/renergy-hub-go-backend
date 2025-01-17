# Renergy Hub Go Backend API

This is the notification service for Renergyhub built in Go

## Project Overview

Live link: is at http://renergyhub.osinachi.me

Doc link: https://renergyhub.osinachi.me/api/docs/index.html

## Installation Instructions

### Prerequisites

Before setting up the project locally, ensure you have the following prerequisites installed:

- [Go](https://go.dev/dl/) (>=1.20).
- A Database System (e.g., PostgreSQL, MySQL, SQLite)

### How to run API Locally

1. Clone the repository:

```bash
git clone https://github.com/InternPulse/renergy-hub-go-backend.git
```

2. Change into the parent directory:

```bash
cd renergy-hub-go-backend
```

3. Create the .env file from the .env.example file

```sh
cp .env.example .env
```

4. Set appropriate values for the following Compulsory Environment Variables in the .env file:

```.env
PORT=7500
DB_HOST="localhost"
DB_PORT="5432"
DB_NAME="renergybackend"
DB_USER="renergyuser"
DB_PASSWORD="go-is-cool"
DB_SSL_MODE="require"
JWT_SECRET=""
```

5. Install the App dependencies:

```bash
go get
```

5. Start the App:

```bash
go run main.go
```

The API should now be running locally at [http://localhost:7500/](http://localhost:7500/).

## Features

1. Send order created notification
2. Send email verification notification
3. Send order shipped notification

# Commit Standards

## Branches

- **dev** -> pr this branch for everything `backend` related
- **main** -> **dont touch** this branch, this is what is running in production!

## Contributions

property-hive-backend-two is open to contributions, but I recommend creating an issue or replying in a comment to let us know what you are working on first that way we don't overwrite each other.

## Contribution Guidelines

1. Clone the repo `git clone https://github.com/InternPulse/property-hive-backend-two.git`.
2. Open your terminal & set the origin branch: `git remote add origin https://github.com/InternPulse/property-hive-backend-two.git`
3. Pull origin `git pull origin dev`
4. Create a new branch for the task you were assigned to, eg `TicketNumber/(Feat/Bug/Fix/Chore)/Ticket-title` : `git checkout -b BA-001/Feat/Sign-Up-from`
5. After making changes, do `git add .`
6. Commit your changes with a descriptive commit message : `git commit -m "your commit message"`.
7. To make sure there are no conflicts, run `git pull origin dev`.
8. Push changes to your new branch, run `git push -u origin feat-csv-parser`.
9. Create a pull request to the `dev` branch not `main`.
10. Ensure to describe your pull request.
11. > If you've added code that should be tested, add some test examples.

# Merging

Under any circumstances should you merge a pull request on a specific branch to the `dev` or `main` branch

### _Commit CheatSheet_

| Type     |                          | Description                                                                                                 |
| -------- | ------------------------ | ----------------------------------------------------------------------------------------------------------- |
| feat     | Features                 | A new feature                                                                                               |
| fix      | Bug Fixes                | A bug fix                                                                                                   |
| docs     | Documentation            | Documentation only changes                                                                                  |
| style    | Styles                   | Changes that do not affect the meaning of the code (white-space, formatting, missing semi-colons, etc.)     |
| refactor | Code Refactoring         | A code change that neither fixes a bug nor adds a feature                                                   |
| perf     | Performance Improvements | A code change that improves performance                                                                     |
| test     | Tests                    | Adding missing tests or correcting existing tests                                                           |
| build    | Builds                   | Changes that affect the build system or external dependencies (example scopes: gulp, broccoli, npm)         |
| ci       | Continuous Integrations  | Changes to our CI configuration files and scripts (example scopes: Travis, Circle, BrowserStack, SauceLabs) |
| chore    | Chores                   | Other changes that don't modify, backend or test files                                                      |
| revert   | Reverts                  | Reverts a previous commit                                                                                   |

> _Sample Commit Messages_

- `chore: Updated README file`:= `chore` is used because the commit didn't make any changes to the backend or test folders in any way.
- `feat: Added plugin info endpoints`:= `feat` is used here because the feature was non-existent before the commit.
