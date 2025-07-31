# Stage 1: Build React App
FROM node:18-alpine AS build

WORKDIR /app

COPY package.json ./
COPY package-lock.json ./
RUN npm install

COPY . ./
RUN npm run build

# Stage 2: Serve with Express
FROM node:18-alpine

WORKDIR /app

COPY --from=build /app/build ./build
COPY server.js .
COPY package.json ./
COPY package-lock.json ./

RUN npm install --only=production

EXPOSE 3000
CMD ["node", "server.js"]