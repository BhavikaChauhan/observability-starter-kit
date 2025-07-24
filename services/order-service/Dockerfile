# Use official Node.js LTS image
FROM node:20-alpine

# Set working directory inside the container
WORKDIR /app

# Copy package files and install dependencies
COPY package*.json ./
RUN npm install

# Copy the entire project (including index.js now at root)
COPY . .

# Expose port (optional, based on what port your service uses)
EXPOSE 3000

# Start the app (index.js is now in root)
CMD ["node", "index.js"]
