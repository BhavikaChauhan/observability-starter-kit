# Step 1: Build React app
FROM node:18-alpine as build
WORKDIR /app
COPY client/package*.json ./client/
RUN cd client && npm install
COPY client ./client
RUN cd client && npm run build

# Step 2: Set up Express server
FROM node:18-alpine
WORKDIR /app

# Copy backend files
COPY server/package*.json ./server/
RUN cd server && npm install

# Copy server code
COPY server ./server

# Copy built React files into Express public dir
COPY --from=build /app/client/build ./server/build

WORKDIR /app/server
EXPOSE 4002
CMD ["node", "server.js"]