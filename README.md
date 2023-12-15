# News API Client in Go

This Go application demonstrates how to fetch and display top headlines from the News API. It makes use of the HTTP package to interact with the News API and parse the JSON response.

## Prerequisites

Before running the application, make sure you have the following:

- News API Key: Replace `YOUR_API_KEY` in `app.go` with your actual News API key. You can obtain one by signing up at [https://newsapi.org/](https://newsapi.org/).

## Usage

1. Open the `app.go` file and replace `"YOUR_API_KEY"` with your actual News API key.

2. Customize the request parameters (optional):
   - `country`: Specify the country for which you want to fetch headlines (e.g., "us", "gb", "in").
   - `category`: Choose a news category (e.g., "general", "business", "technology").

3. Run the application:
   ```bash
   go run app.go
