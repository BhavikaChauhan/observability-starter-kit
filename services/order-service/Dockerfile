# Use official Node.js LTS image
FROM node:18

# Set working directory
WORKDIR /app

# Copy package.json and install dependencies
COPY package.json ./
RUN npm install

# Copy rest of the app
COPY . .

# Expose port (ensure it matches your app's listen port)
EXPOSE 3003

# Start the app
CMD ["npm", "start"]