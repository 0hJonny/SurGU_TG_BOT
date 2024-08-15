# SurGU Telegram Bot

## Prerequisites

- **Docker**: Ensure Docker is installed on your system. You can download it from [here](https://docs.docker.com/get-docker/).

### Build and Run

Follow these steps to build and run the SurGU Telegram Bot:

1. **Build the Docker Image**:

   Clone the repository and navigate to the project directory. Then, build the Docker image using the following command:

   ```bash
   docker build -t surgu_tg_bot .
   ```

   This will create a Docker image named `surgu_tg_bot`.

2. **Run the Docker Container**:

   Start the bot using the following command:

   ```bash
   docker run -e BOT_TOKEN=YourToken -p 8080:8080 surgu_tg_bot
   ```

   Replace `YourToken` with your actual Telegram bot token.

   The bot will now be running and accessible on port `8080`.

### Environment Variables

- `BOT_TOKEN`: **(Required)** The token for your Telegram bot.

### License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
