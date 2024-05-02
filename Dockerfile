FROM ubuntu:latest
LABEL authors="tv"

ENTRYPOINT ["top", "-b"]