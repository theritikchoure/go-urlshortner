# URL Shortener and Redirection Service

![Project Logo](logo.png)
_Shorten Your Links, Expand Your Reach_

## Table of Contents

- [Introduction](#introduction)
- [Features](#features)
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
- [Usage](#usage)
- [Configuration](#configuration)
- [API Documentation](#api-documentation)
- [Contributing](#contributing)
- [License](#license)

## Introduction

Welcome to the URL Shortener and Redirection Service! This project is a simple yet powerful tool for shortening long URLs and redirecting users to their intended destinations. Whether you need concise links for sharing or tracking click statistics, our service has you covered.

### Features

- **URL Shortening:** Convert long, unwieldy URLs into short and user-friendly aliases.
- **Custom Short URLs:** Create custom short URLs for branding and easy recognition.
- **Rate Limiting:** Implement rate limiting to protect your service from abuse.
- **Domain Validation:** Ensure that only valid external URLs are shortened.
- **Usage Statistics:** Track click statistics and user engagement.
- **Easy Integration:** Use our API to integrate URL shortening into your applications.
- **Highly Configurable:** Customize settings via environment variables.

## Getting Started

### Prerequisites

Before you begin, ensure you have met the following requirements:

- [Go (Golang)](https://golang.org/) installed on your system.
- [Docker](https://www.docker.com/) installed on your system.
- [Redis](https://redis.io/) database server running and accessible.
- A [Fiber](https://gofiber.io/) web framework installed (`go get -u github.com/gofiber/fiber/v2`).
- [Go-Redis](https://github.com/go-redis/redis/v8) library (`go get -u github.com/go-redis/redis/v8`).
- Clone this repository to your local machine.

### Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/your-username/url-shortener.git
   cd url-shortener
   ```

2. Set up your environment variables:

   - Create a .env file in the project root.
   - Define the following environment variables in the .env file:
     - `DB_ADDR`: Address of your Redis database server.
     - `DB_PASS`: Password for connecting to Redis.
     - `APP_PORT`: Port on which the application will run.
     - `DOMAIN`: Your domain name for URL generation.
     - `APP_QUOTA`: Rate limiting quota (e.g., "10" requests per minute).

3. Build and run the application:
   ```bash
   go run main.go
   ```

Your URL shortener service should now be up and running.

## Usage
To shorten a URL, make a POST request to /api/v1 with a JSON payload containing the long URL:

```bash
curl -X POST http://localhost:3000/api/v1 -H "Content-Type: application/json" -d '{"url": "https://example.com"}'
```

To access a shortened URL, simply visit `http://your-domain.com/short-alias`.

## Configuration
The application is highly configurable through environment variables defined in the .env file. Adjust these variables to fit your specific needs.

- `DB_ADDR`: Address of the Redis database server.
- `DB_PASS`: Redis database password.
- `APP_PORT`: Port on which the application will listen.
- `DOMAIN`: Your domain name for URL generation.
- `APP_QUOTA`: Rate limiting quota (requests per minute).

## Contributing
We welcome contributions from the community! If you'd like to contribute to this project.