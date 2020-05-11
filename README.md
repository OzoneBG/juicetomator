# Juice Server

## Description
A simple http wrapper around a python script that uses selenium and youtube-dl to download songs

## Python external packages
* selenium
* webdriver-manager
* youtube-dl

## Setup
1. Build for the correct OS
2. Copy the build directory to the host machine
3. Run the binary

## Usage

### Getting multiple songs
`curl http://localhost:8080/get-juices -X POST -d '{"username": "wolf", "songs": ["otilia diamante y3mr$", "lucifer in the air tonight"]}' --output songs.zip`

### Getting single song
Not Implemented Yet