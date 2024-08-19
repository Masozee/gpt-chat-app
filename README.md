# GPT-like Chat Application

This is a WebSocket-based chat application that integrates with AI models like ChatGPT and Claude.

## Setup

1. Ensure you have Go 1.16 or later installed.
2. Clone this repository.
3. Run `go mod download` to fetch dependencies.
4. Create a `config.json` file in the project root with your desired configuration.

## Running the Application

1. Build the application:
   ```
   ./scripts/build.sh
   ```
2. Run the server:
   ```
   ./gpt-chat-app
   ```

## Testing

Run the tests using:
```
./scripts/test.sh
```

## Docker

To build and run using Docker:

```
docker build -t gpt-chat-app .
docker run -p 8080:8080 gpt-chat-app
```

## Configuration

Edit `config.json` to configure the application:

```json
{
  "port": 8080,
  "max_clients": 300,
  "ai_provider": "chatgpt"
}
```

## License

[MIT License](LICENSE)