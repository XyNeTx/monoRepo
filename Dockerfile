# 1. Use official Node.js version 22 as the base image
FROM node:22.11.0-alpine

# 2. Create /app and set it as working directory
WORKDIR /app

# 3. Copy ONLY package.json & package-lock.json to Workdir
# (this allows Docker to cache npm install)
COPY package*.json ./

# 4. Install dependencies (cached on rebuild)
RUN npm install

# 5. Copy all project files into /app
COPY . .

# 6. Document that the app listens on port 3000
# (EXPOSE does NOT publish the port)
EXPOSE 3000

# 7. Start your Node.js app when the container starts
CMD ["npm", "run", "dev"]
