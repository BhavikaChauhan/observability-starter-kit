# Stage 1: Build the React app
FROM node:18-alpine AS build

WORKDIR /app

COPY package.json package-lock.json ./

# Install all dependencies
RUN npm install

# Copy source code
COPY ./public ./public
COPY ./src ./src

# Build React app only (no Express involved here)
RUN npm run build

# Stage 2: Setup Express server to serve build
FROM node:18-alpine

WORKDIR /app

# Only copy server and necessary files
COPY server.js .
COPY package.json package-lock.json ./

# Install only production dependencies (includes express)
RUN npm install --only=production

# Copy build output from previous stage
COPY --from=build /app/build ./build

EXPOSE 3000
CMD ["node", "server.js"]