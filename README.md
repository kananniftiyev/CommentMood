<h1 align="center">CommentMood</h1>
<h4 align="center">CommentMood" is your all-in-one solution for gaining quick insights into the sentiment of Social Media apps comments. With just a URL, this user-friendly tool retrieves and analyzes comments, providing a clear understanding of the emotional tone within seconds</h4>

## Table of Contents
- [Getting Started](#getting-started)
    - [Prerequisites](#prerequisites)
- [Usage](#usage)
    - [Running the API](#running-the-api)
    - [API Endpoints](#api-endpoints)
- [Docker](#docker)
- [Configuration](#configuration)
- [Docker](#docker)

## Getting Started
### Prerequisites
Before you begin, ensure you have met the following requirements:

- Go 1.21 installed on your local machine.
- Docker (optional, for running the API in a Docker container).

The following Go modules will be automatically installed as dependencies:

- github.com/jonreiter/govader v0.0.0-20230129030235-c72a790a959e
- google.golang.org/api v0.144.0
- gorm.io/driver/postgres v1.5.2
- gorm.io/gorm v1.25.4
- (and other indirect dependencies, see `go.mod` for the complete list)

## Usage

- You can use this API to have better insights about your content in these social media apps: Youtube, Instagram, Twitch
### Running the API

##### For running project in local you can do these steps:
1. Open the terminal in your OS.
2. Get to directory where you did download the project
3. Run this Script
```bash
go run main.go
```
## Docker
To run both the application and the database within a Docker environment, follow these steps:
1. Open a terminal and navigate to the root directory of your project.

2. Run the following command to build and start the containers:

   ```bash
   docker-compose up --build
   ```
    The --build flag ensures that Docker Compose rebuilds the images if there have been changes to your project's code.
   
   Your API and database will now be up and running in Docker containers, making it easy to manage your development environment.

***Note***: You may need to adjust the docker-compose.yml file in your project to configure the specific services and settings required for your application.
### API Endpoints

### Analyze YouTube Video

- **Endpoint**: `/api/v1/analyze/youtube`
- **HTTP Method**: POST
- **Description**: Analyzes sentiment for a YouTube video using its video ID.
- **Request Body**: JSON object with the following structure:
  ```json
  {
    "videoId": "your_youtube_video_id"
  }
- **Response**:
   ```json
  {
      "videoID": "video_id",
      "positive": 9,
      "neutral": 7,
      "negative": 7,
      "not": 5
  }
  ```

## Configuration

Add all needed information to .env file please.

