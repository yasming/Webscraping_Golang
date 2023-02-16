# WebsScrapping with golang

This is a project that will get the teams in cblol from https://www.flashscore.com.br/esports/league-of-legends/cblol/classificacao/#/27z054yL/table using a crawler and provide all the teams that are participating in the championship as a json response from the endpoint ```/api/get-teams``` and addressing some ```DDD techniques``` in the process of develop our api. 

## Prerequisites

```
Docker
```

```
Make
```

## Getting started

1. .ENV

```
cp .env.example .env
```

2. Start the server

```
make start
```

3. Create new migration

```
make create-new-migration table_name=teams
```

4. Run migrations

```
make run-migrations
```

5. Rollback migrations

```
make rollnack-migrations
```

6. Install new package

```
make package_name=github.com/chromedp/chromedp install-new-package
```

7. Run all tests with test coverage

```
make run-all-tests
```

7. Stop the server

```
make stop
```