# GO FIBER STARTER

## Folder Structure

Here's the folder structure used in this starter

```
.
├── backend/
│   ├── apps/
│   │   ├── app.go
│   │   └── // other file needed
│   ├── entities/
│   │   ├── dto/
│   │   │   └── // used for simplifiy complex argument passed to function
│   │   ├── request/
│   │   │   └── // for anything about requests
│   │   ├── response/
│   │   │   └── // for anything about responses
│   │   └── table/
│   │       └── // for anything about table
│   ├── handlers/
│   │   └── // all about accepting request or resolve response goes here
│   ├── repositories/
│   │   └── //all about db goes here
│   ├── routes/
│   │   └── // set the routes here
│   ├── services/
│   │   └── // all about business logic goes here
│   └── utils/
│       └── // fill with anything would used globally
├── cmd/
│   └── main.go
└── config/  // you might differentiate this based on SDLC environments
    ├── db/
    │   └── // db config goes here
    ├── redis
    └── // other config
```

## Features:

- [Fiber v2.x](https://docs.gofiber.io/)
- [Sirupsen/Logrus](https://github.com/sirupsen/logrus)
- [GORM](https://gorm.io/)
- [go-playground/validator](https://github.com/go-playground/validator)
- ...

## Rules:

- Naming Connvention:

  > - Use **camelCase** for naming the variable
  > - Use **PascalCase** for naming the struct and the function
  > - Use **lowercase** for naming the package
  > - Use `"err" + function_name ` for variable which store the error value of a function, like:
  >
  >   > GetList -> errGetList
  >
  > - **NEVER** use **snake_case** for naming anything here

- Application Layer:

  1. Handlers

  > - Everything related to showing response and take any requests goes here
  > - Make sure to validate the request payload if needed
  > - On handler function, there's a parameter for **\*fiber.Ctx**, NEVER bring it into another layer (you could use **&fiber.Ctx{}** instead)

  2. Services

  > - Everything related to business logic/process goes here
  > - On services layer, every request payloads from handler and responses from the deeper layer (repositories layer) were "cooked" and ready to serve either to handlers or repositories

  3. Repositories

  > - Everything related to database and querying gose here
  > - Except for Get action, use `$action + To + $table_name` for the function name, like:
  >   > table `change` -> `SubmitToChange` / `UpdateToChange` / `DeleteToChange`
  > - For Get action, use `Get + information` for the function name, like:
  >   > detail change -> `GetDetailChange`
  >   >
  >   > list change -> `GetListChange`

- More:

  ```
  > 🚧🚧 UNDER CONSTRUCTION 🚧🚧
  ```

- Note:

  > - Make sure adjust the Dockerfile & compose.yaml based on your needs before deployment to docker
