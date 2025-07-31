# Stage 1: Build React app
FROM node:18-alpine as build

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
COPY server.js ./
COPY package.json ./
RUN npm install --only=prod

EXPOSE 3002

CMD ["node", "server.js"]