# url-shortner
URL shortner is an application which takes url as an input and generates a tiny URL for it, also get the metrics for top 3 most visited domains.

## Prerequisites

Before you can run this application, make sure you have the following tools installed on your system:

- [Docker](https://docs.docker.com/get-docker/)
- [Docker Compose](https://docs.docker.com/compose/install/)

## Installation

### 1. Clone the Repository

First, clone this Git repository to your local machine using the following command:

```sh
git clone https://github.com/vijayaedke/url-shortner
```

### 2. Install dependency of go required to run the app

Use the following command:

```
go mod download
```

### 3. Run the application

To execute the application run the following command:

```sh
docker-compose up
```
