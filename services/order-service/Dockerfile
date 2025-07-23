# Base Node image
FROM node:18-alpine

# Set working dir
WORKDIR /app

# Copy files
COPY package*.json ./
RUN npm install

COPY . .

# Expose port
EXPOSE 3003

CMD ["npm", "start"]
