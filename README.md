# ColorFun

## Description
This project is designed to provide a fun demo of WebSockets with React Server Components. It includes a rotator component that can be customized by changing the WebSocket URL in `stateProvider.js`. 

## Demo
Check out the live demo at [https://colorfun.fly.dev](https://colorfun.fly.dev).

## Features
- Dynamic color rotation
- WebSocket integration for real-time updates
- Dockerized setup for easy deployment

## Prerequisites
- Docker
- Node.js
- Go
- Any other dependencies

## Installation

1. Clone the repository:
    ```sh
    git clone https://github.com/yourusername/colorfun.git
    cd colorfun
    ```

2. Build the Docker image:
    ```sh
    docker build -t colorfun .
    ```

3. Run the Docker container:
    ```sh
    docker run -p 3000:3000 colorfun
    ```

## Configuration

To change the WebSocket URL in `stateProvider.js`, follow these steps:

1. Open `stateProvider.js` in your preferred text editor.
2. Locate the line containing the WebSocket URL:
    ```js
    const socket = new WebSocket('wss://colorfun.fly.dev:7001/ws');
    ```
3. Replace `'wss://colorfun.fly.dev:7001/ws'` with your desired WebSocket URL.
4. Save the file.

## Deployment

1. Ensure all changes are committed:
    ```sh
    git add .
    git commit -m "Your commit message"
    ```

2. Push the changes to your repository:
    ```sh
    git push origin main
    ```

3. Deploy the Docker container as described in the Installation section.

## License
This project is licensed under the MIT License

