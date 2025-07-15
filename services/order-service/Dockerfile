# Stage 1: Build React app
FROM node:18-alpine as builder
WORKDIR /app
COPY package*.json ./
RUN npm install
COPY . .
RUN npm run build

# Stage 2: Serve via Express
FROM node:18-alpine
WORKDIR /app

# Install only runtime dependencies
COPY package*.json ./
RUN npm install --only=prod

# Copy built frontend
COPY --from=builder /app/build ./build
COPY server.js .

EXPOSE 8000
CMD ["node", "server.js"]
