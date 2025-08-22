# PWA Engine Full Build Script

# 1. Prepare Go modules
go mod tidy

# 2. Build and run Docker Compose services
docker compose up --build

# 3. Build TailwindCSS

## Windows (PowerShell)
From the root of your project in Command Prompt or PowerShell, run:

docker run --rm -v <path-to-your-project-root>:/project -w /project/tailwind d3fk/tailwindcss:v3 build -i tailwind-input.css -o ../client/css/output.css --minify

## macOS / Linux (bash/zsh)
From the root of your project in Terminal, run:

docker run --rm -v /path/to/your/project:/project -w /project/tailwind d3fk/tailwindcss:v3 build -i tailwind-input.css -o ../client/css/output.css --minify
