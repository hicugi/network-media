# Inspired project for hackathon by [Boot.Dev](https://www.boot.dev/)

<p align="center">
  <a href="https://go.dev/" target="_blank">
    <img src="https://img.shields.io/badge/Go-v1.14.1-007d9c" alt="Go">
  </a>
</p>

## Idea

Watch media by network across multiple devices.

## Requirements

- User should be able to run executable binary file on his server (for now Ubuntu linux/amd64)
- A folder, where the binary file located will be shared by network
- Devices in the same network can open a Web page in a browser and play videos (for now)

## Approach

- Use `net/http` to run web server
- Scan the folder where is the app is running (skip sub folder for now)
- Look for all .mp4 (for now) video files in root of the binary file
- Display list of the files in browser
- Check if it is accessible on other devices
- Update the README doc for step by step instruction - how to run the project

## Extra features when the main [approach](#Approach) is finished

- Scan sub filders
- Add **.mkv** format for video files. Check for other formats
- Make the binary build for Windows 10/11
- Scan for images (**.jpg, .jpeg, .png, .gif**) and add gallery UI
- Scan for audio files and add music player
