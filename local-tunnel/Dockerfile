FROM node:lts
RUN npm install -g localtunnel
RUN apt-get update -q && apt-get install -qy dumb-init && rm -rf /var/lib/apt/lists/*
ENTRYPOINT ["/usr/bin/dumb-init", "--", "lt"]
