# 4room

A simple web application for creating and browsing posts, with user authentication and comments functionality.

## Prerequisites

- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)

## Installation

1. Clone this repository:

```bash
git clone https://github.com/meirgenuine/4room.git
cd 4room
```

2. Build the Docker images:

```bash
make build
```

3. Run the Docker containers:

```bash
make run
```

4. Visit the application at http://localhost:8080

## Usage

1.  Register a new user or log in with an existing account.
2.  Create new posts and categories.
3.  Browse and filter posts by category.
4.  View individual posts and add comments.

## Development

1. Stop the Docker containers:

```bash
make stop
```

2. Clean up the containers, images, and volumes:

```bash
make clean
```

3. Make the necessary changes to the source code.

4. Rebuild the Docker images and run the containers again.

## License

This project is licensed under the MIT License. See the LICENSE file for details.