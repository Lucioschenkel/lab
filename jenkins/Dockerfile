FROM jenkins/jenkins:lts-jdk17

USER root

RUN apt-get update && apt-get install cowsay jq -y

USER jenkins

ENV PATH="$PATH:/usr/games"
