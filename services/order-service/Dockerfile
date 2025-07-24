# Use official Node.js image
FROM node:20

# Set working directory
WORKDIR /app

# Copy package.json and install dependencies
COPY package*.json ./
RUN npm install

# Copy all app files
COPY . .

# Set working directory to src so index.js is found
WORKDIR /app/src

# Start the app
CMD [ "node", "index.js" ]
