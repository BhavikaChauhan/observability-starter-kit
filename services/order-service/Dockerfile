# Use official Node.js image
FROM node:20

# Create app directory
WORKDIR /app

# Copy package.json and install dependencies
COPY package*.json ./
RUN npm install

# Copy the rest of the code
COPY . .

# Set the working directory to the actual source
WORKDIR /app/src

# Expose port (update if needed)
EXPOSE 3003

# Start the service
CMD ["node", "index.js"]
