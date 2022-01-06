<br />
<p align="center">
  <h2 align="center">Enceladus</h3>
</p>

<!-- CONTENT TABLE -->

## Content Table

- [About](#about)
- [Requirements](#requirements)
- [Get Started](#get-started)
- [Run](#run)

<!-- ABOUT -->

## About

This repository contains a small API designed to test clean architecture principles in GO. The main idea was to use the standard library packages as much as I could to construct the API, but to avoid unnecessary work I used some external packages to help with some tasks.

<!-- REQUIREMENTS -->

## Requirements

* [Golang](https://go.dev/dl/)
* [Docker](https://docs.docker.com/get-docker/)
* [Docker Compose](https://docs.docker.com/compose/install/)

<!-- GET STARTED -->

## Get Started

Choose a folder and clone the repository:

```sh
# ssh
git clone git@github.com:rhuancaetano/enceladus.git

# https
git clone https://github.com/rhuancaetano/enceladus.git
```

<!-- SETUP -->

### Setup

Create a `.env` file with the necessary environment variables. You can use the `.env.dev.template` if you want:

```sh
cp .env.dev.template .env
```

This step is only required if you want to run the application locally. Install the project dependencies:

```sh
go mod download
```

<!-- RUN -->

## Run

There are two different ways to run the application, you can run the api locally and the postgres database in the container or run the containerized application.

### API locally:

To start the database and the api:

```sh
make dev-start
```

To stop the database from running in background:

```sh
make dev-stop
```

### Containerized application:

To start the application:

```sh
make start
```

To stop the application from running in background:

```sh
make stop
```

If you add changes to the api files you will need to rebuild the application, please do:

```sh
make build
```

### Testing:

To run the tests use:

```sh
make test
```

Test coverage is not high yet, the main idea was to understand how they work in golang.
