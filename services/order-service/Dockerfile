FROM node:20-alpine

WORKDIR /app

# Copy package files and install dependencies
COPY package*.json ./
RUN npm install

# Copy everything else including the src folder
COPY . .

# Set working directory to src
WORKDIR /app/src

# Start the application
CMD ["node", "index.js"]