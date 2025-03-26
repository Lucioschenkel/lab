# Intro

This folder contains my experiments with Jenkins in a local environment. I used Docker to provision a Jenkins container and a Gitea container.

# Setup

If you're using Windows, you should use WSL2 to ensure compatibility with the instructions below.

1. If you're on Linux, copy `/etc/timezone` under this directory under the name `timezone`. macOS users can copy their current timezone file, from `/usr/share/zoneinfo`, e.g.:

```shell
cp /usr/share/zoneinfo/America/New_York timezone
```

2. Create two directories: `config` and `data`. Those are used to persist Gitea data across restarts.
3. Run `docker compose up -d`
