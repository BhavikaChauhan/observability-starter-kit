# Dockerfile (for server.js)
FROM node:20-alpine

WORKDIR /app

COPY package.json .

RUN npm install

COPY . .

CMD ["node", "server.js"]

