[![Review Assignment Due Date](https://classroom.github.com/assets/deadline-readme-button-24ddc0f5d75046c5622901739e7c5dd533143b0c8e959d652212380cedb1ea36.svg)](https://classroom.github.com/a/h2qMnoBZ)
[![Open in Visual Studio Code](https://classroom.github.com/assets/open-in-vscode-718a45dd9cf7e7f842a935f5ebbe5719a5e09af4491e668f4dbf3b35d5cca122.svg)](https://classroom.github.com/online_ide?assignment_repo_id=11438886&assignment_repo_type=AssignmentRepo)

# GoGo Form

This is a simple API project built with Go language that provides functionalities for creating and managing forms and answers.

## Getting Started

### Prerequisites
* Go 1.10
* MongoDB
OR
* Docker

### Installation

1. Clone the repository:
   ```bash
   git clone git@github.com:ufrn-golang/trabalho-3-desenvolvimento-de-api-restful-team-pablo-api.git gogo-form
   cd gogo-form
   ```
2. Install the dependencies:
    ```bash
    go mod download
    ```
3. Run
    ```bash
    go run main.go
    ```
### Using dokcer
You can easilly run the application with Docker using the command
* `make build` to build the image
* `make run` to run the application from the container
* `make stop` to stop the container