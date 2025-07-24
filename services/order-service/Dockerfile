FROM node:20-alpine

WORKDIR /app

# Install dependencies
COPY package*.json ./
RUN npm install

# Copy the entire service folder into the image
COPY . .

# Tell Docker to use the src folder as working directory
WORKDIR /app/src

# Set the correct entry point
CMD ["node", "index.js"]
