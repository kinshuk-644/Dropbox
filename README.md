# Dropbox
A simple version of dropbox application meant to upload, download and view files on a UI.

# Tech Stack
Backend: Go + Gin (Framework), GORM (Database ORM), Viper (Configuration)

Frontend: React (with Vite), Axios

Database: PostgreSQL (running via Docker Compose)

File Storage: Local file system (/backend/uploads)

# Software
Ensure the following software is installed on your machine:

Git: For cloning the repository.

Go: Version 1.24

Node.js: Version 20.19+ or 22.12+ (Vite requirement).

Docker & Docker Compose: For running the PostgreSQL database.

# Steps to run on local
1. git clone https://github.com/kinshuk-644/Dropbox.git
2. cd Dropbox

3. docker-compose up -d

4. cd frontend

5. npm install

6. cd ../backend
   
7. go run main.go

8. cd frontend (in new terminal)
   
9. npm run dev

Now access http://localhost:5173 on browser

# API Endpoints:

1. POST		/upload      	  	Uploads a new file. Expects multipart/form-data with a key of file.
2. GET		/files         		Gets a JSON list of all file metadata.
3. GET		/files/:id     		Forces a download for the file with the given ID.
4. GET		/view/:id      		Serves the file "inline" to be viewed in the browser.
