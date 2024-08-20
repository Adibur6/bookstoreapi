
# BookAPI CLI

`BookAPI CLI` is a command-line tool built with Cobra to manage the `BookAPI` server. It allows you to start, stop, and manage the server with various options.

## Features

- Start the server with custom configuration.
- Stop the server gracefully.
- Manage server settings through a `config.json` file.

## Installation

### Prerequisites

- Go 1.16+ installed on your machine.

### Clone the Repository

```bash
git clone https://github.com/yourusername/bookapi-cli.git
cd bookapi-cli
``` 

### Build the CLI

To build the `BookAPI CLI`:

```bash
go build -o bookapi
```

This command will create a binary named `bookapi` in your current directory.

### Configuration

Before using the CLI, ensure you have a `config.json` file in the root directory of your project. The file should look like this:

```json
{
  "serverRunning": false,
  "portNumber": 8080
} 
```
## Usage

Here are the main commands provided by the `BookAPI CLI`:

### Starts the server using the port number specified in `config.json`.
```bash
./bookapi start 
```

### Starts the server on a specified port.

```bash
./bookapi start -p 9090
```
### Starts the server in detached mode (runs in the background).
```bash
./bookapi start -d 
```


### Stops the server if it is running.
```bash
./bookapi stop
```