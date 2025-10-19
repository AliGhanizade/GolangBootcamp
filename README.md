# Golang Bootcamp

This repository was created as part of my internship at **Tosee Saman** company and their amazing development team.  
I would like to express my sincere gratitude to **Tosee Saman** and their team for giving me this great opportunity to learn and grow as a developer.  

As the name suggests, **Golang Bootcamp** is a collection of my training projects, exercises, and experiments completed during the internship period.  
It reflects my learning journey with the **Go programming language**, covering both fundamental and practical concepts.  

All the materials are organized clearly to make it easy to explore and understand my progress throughout the bootcamp.

## Project Overview

The **Golang Bootcamp** project represents my practical journey in learning the Go programming language.  
It includes several small and medium-sized projects developed during my internship at **Tosee Saman**, where I focused on:

- Understanding Go fundamentals (syntax, data types, functions)
- Learning about Go’s concurrency model using goroutines and channels
- Building real-time communication systems with **WebSocket**
- Developing RESTful APIs using the **Gin** framework
- Working with databases using **Gorm**, **PostgreSQL**, and **SQLite**
- Writing modular, clean, and maintainable code
- Applying version control and collaboration best practices with Git

> [!TIP]
> This repository serves as both a learning record and a personal reference for future Go development.

## How to Run the Project

To run the project, follow these simple steps:

1. **Clone the repository**  
   First, download the project to your local system using Git.
   
   ```bash
   git clone https://github.com/AliGhanizade/GolangBootcamp.git
   cd GolangBootcamp
2. **Install dependencies**  
  Use the following command to install all required Go modules.

    ```bash
    go mod tidy
3. **Run the main program**  
  The main entry point of the project is located in the cmd directory.
  You can start the program by running:  
    ```bash
    go run ./cmd/main.go
4. **Once the program starts**, it will prompt you to enter a week number :  
    ```scss
    Enter a Week(Number):
  Each week corresponds to a different Go project or exercise developed during the bootcam

### Notes on Submodules
   
   To properly start the project and **get all submodules**:
   
   1. Clone the repository **including submodules**:

      ```bash
      git clone --recurse-submodules https://github.com/AliGhanizade/GolangBootcamp.git
      cd GolangBootcamp
   2. If you already **cloned without submodules**, initialize and update them:

      ```bash
      git submodule update --init --recursive
> This ensures all Mini Projects (**BookStore**, **Todo** and **Algorithms**) are available and ready to use.




## Project Structure

The Golang Bootcamp repository is organized into three main sections:

|Section        | Contents / Notes |
|-------------------|------------------|
| **Algorithms**   | Exercises and practice problems to strengthen Go fundamentals and problem-solving skills. |
| **Mini Projects** | Contains two subfolders: **BookStore** and **Todo** (both are Git submodules). Important branches are already included. |
| **Weeks**     | Week folders (`week_1` to `week_6`) containing exercises and projects for each week. Some weeks include web server projects using **Gin** or **WebSocket**. |

## Technologies Used

This project uses the following technologies and tools:
- ![Go](https://img.shields.io/badge/Go-1.25-blue?logo=go&logoColor=white) 
- ![Gin](https://img.shields.io/badge/Gin-Framework-00A1DE?logo=gin&logoColor=white) 
- ![Gorm](https://img.shields.io/badge/Gorm-ORM-2C3E50?logo=go&logoColor=white)
- ![PostgreSQL](https://img.shields.io/badge/PostgreSQL-blue?logo=postgresql&logoColor=white) 
- ![SQLite](https://img.shields.io/badge/SQLite-3-07405E?logo=sqlite&logoColor=white) 
- ![WebSocket](https://img.shields.io/badge/WebSocket-FF5722?logo=websocket&logoColor=white) 
- ![Git](https://img.shields.io/badge/Git-F05032?logo=git&logoColor=white)

## Features

Key features of the Golang Bootcamp project:

- CLI program with week-based exercises
- Web servers using Gin and WebSocket
- Database CRUD operations with PostgreSQL & SQLite
- Modular and reusable code structure
- Clear and educational examples for learning Go


## Contributing

Contributions are welcome! To contribute:

1. Fork the repository
2. Create a branch for your changes
3. Make your modifications
4. Submit a pull request for review

Please make sure any new code follows the same structure and formatting as the existing project.

## License

This project is intended for educational purposes only.  
It is open for personal learning, experimentation, and practice.  
You are free to use and study the code, but not for commercial purposes.
