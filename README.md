# Docker Multi-Stage Samples

Minimize your docker images with multi-stage Dockerfiles. This Repo contains several examples created for the Docker Meetup in Berlin.

Each folder contains a Dockerfile for multi-stage build and normal/fat build, each for alpine and debian.

Every project is a simple HTTP server serving: `{ message: "Hallo Docker Berlin! 🐳" }` as JSON. We use popular frameworks on each platform.

Run `./show-me.sh` to build all containers and display the image sizes.
