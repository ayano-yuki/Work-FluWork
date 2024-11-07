
# My Project

This project demonstrates how to use `docker-compose` to simulate `make` command-like behavior for building and running a simple C program.

## Project Structure

```
my_project/
├── Dockerfile
├── docker-compose.yml
├── Makefile
└── src/
    └── main.c
```

- **Dockerfile**: Defines the Docker image and sets up the environment for running `make` commands.
- **docker-compose.yml**: Configures the `docker-compose` environment to build and run the application within a container.
- **Makefile**: Contains the build instructions (such as compiling the C program) for the project.
- **src/main.c**: A sample C source file to be compiled.

## Prerequisites

- Docker
- Docker Compose

Make sure you have Docker and Docker Compose installed on your machine.

## How to Use

1. Clone this repository to your local machine.

2. Navigate to the project directory:

    ```bash
    cd my_project
    ```

3. Build and run the project using `docker-compose`:

    ```bash
    docker-compose up --build
    ```

    This command will:
    - Build the Docker image using the `Dockerfile`.
    - Mount the current project directory into the container.
    - Execute the `make` command within the container to compile the C program.

4. After the build is complete, you should see the output of the compilation.

## Customization

- Modify the `src/main.c` file to change the C program.
- Update the `Makefile` if you have additional build requirements.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
