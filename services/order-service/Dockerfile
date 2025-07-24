# Use official Node.js base image
FROM node:20

# Set working directory
WORKDIR /app

# Copy package.json and package-lock.json
COPY package*.json ./

# Install dependencies
RUN npm install

# Copy the entire source code
COPY . .

# Expose the port your app runs on (adjust if needed)
EXPOSE 3000

# Set the correct start command
CMD ["node", "src/index.js"]
