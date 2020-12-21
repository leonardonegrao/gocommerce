# GoCommerce
> A commerce inspired project using Go and Postgres to exercise CRUD with Golang.

## About

GoCommerce is a products management dashboard, where it's possible to perform basic CRUD operations. The project purpose is to be a introduction to Go and how it works when creating a basic web server.

The project, besides Go, uses HTML and Bootstrap to render pages, and PostgreSQL to store data.

## 🎨 Features

You're able to:

- Create new products;
- Update existent products;
- Remove products

## 🚀 How to run

To run this project in your machine, first it's necessary to have Golang and PostgreSQL. It's recommended to use Docker to create a PostgreSQL database to test the application. The application expects a database named "gocommerce" and "admin" as password, but that can be changed in the `database/db.go` file.

Once everything is set-up, you can run, by executing at the root of the project:

```shell
go run main.go
```

The project then can be found on `http://localhost:8000`.

## 🧰 Technologies used

- [Golang](https://golang.org/)
- [PostgreSQL](https://www.postgresql.org/)
- [Postgres drivers](https://github.com/lib/pq)
