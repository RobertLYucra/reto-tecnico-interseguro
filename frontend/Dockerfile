FROM node:22-alpine

WORKDIR /app

COPY package*.json ./
RUN npm install

COPY . .

# Vite default port
EXPOSE 5173

# Host 0.0.0.0 is needed for Docker networking
CMD ["npm", "run", "dev", "--", "--host"]
