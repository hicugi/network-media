# Inspired project for hackathon by [Boot.Dev](https://www.boot.dev/)

<p align="center">
  <a href="https://go.dev/" target="_blank">
    <img src="https://img.shields.io/badge/Go-v1.14.1-007d9c" alt="Go">
  </a>
</p>

## Idea

Transforms your device into server and shares your media folder across local network by web.

## How to run the project

1. Install Go v1.14.1
2. Run the command in the root of the project `go build`
3. Move the `medianetwork` into directory where media files located

The [server](http://127.0.0.1:8000) will be running on port `:8000` in the background.<br>
You can check your local IP by next command `ifconfig | grep RUNNING -A 1` and open it in your device.<br>
Be sure to connect in the same network and don't forget to add the PORT `:8000`.<br>
Example of the URL: `http://192.168.0.100:8000/`. Replace the `192.168.0.100` with your server's local IP.

## Requirements

- User should be able to run executable binary file on his server (for now Ubuntu linux/amd64)
- A folder, where the binary file located will be shared by network
- Devices in the same network can open a Web page in a browser and play videos (for now)

## Approach

- [x] Use `net/http` to run web server
- [x] Scan the folder where is the app is running (skip sub folder for now)
- [x] Look for all .mp4 (for now) video files in root of the binary file
- [x] Display list of the files in browser
- [x] Check if it is accessible on other devices
- [x] Update the README doc for step by step instruction - how to run the project

## Extra features when the main [approach](#Approach) is finished

- [ ] Scan sub filders
- [ ] Add **.mkv** format for video files. Check for other formats
- [ ] Make the binary build for Windows 10/11
- [ ] Scan for images (**.jpg, .jpeg, .png, .gif**) and add gallery UI
- [ ] Scan for audio files and add music player
