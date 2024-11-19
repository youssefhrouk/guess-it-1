# guess-it-1

## Description 
A program that, given a number as standard input, prints out a range in which the next number provided should be. The application helps in guiding the user to guess a number within a specific range based on previous inputs.

## test
You can test the program by downloading the server from one of these links:

- [link](https://assets.01-edu.org/guess-it/guess-it.zip)
- [docker-app link](https://assets.01-edu.org/guess-it/guess-it-dockerized.zip)

Copy the student folder conatining the student guesser and move it to the server folder.

Make the binaries inside the ai folder executables:
```
chmod +x *
```
## Usage
To use the program, simply run it and enter a number when prompted. The program will then print out a range in which the next number should be
## Running the Application

### Using `Docker-Compose`
```# Start the application
docker-compose up

# Reload the server
docker-compose down -v
docker-compose up --build
```
### Using `Dockerfile`
```# Build the Docker image
docker build -t guesser .

# Run the Docker container
docker run -p 3000:3000 guesser
```
### Using `Node.js`
```# Install dependencies
npm install

# Start the server
node server.js
```

## How It Works

The program takes a number as input and provides a range suggestion for the next number to be guessed. This helps in narrowing down the possible values through an iterative guessing process.
