# Docker Templates for Geographical Information Systems APIs

This repository provides a set of Docker templates for developing Geographical Information Systems APIs using different programming languages and frameworks.

## Getting Started

### Prerequisites

To use the templates in this repository, you need to have the following software installed on your machine:

- Docker
- Git

### Installation

To use these templates, follow the steps below:

1. Clone this repository to your local machine using the command:

git clone <https://github.com/MxDui/gis-api-docker-template.git>

2. Navigate to the directory of the template you want to use:

cd {template-directory}

3. Build the Docker image:

docker build -t {image-name} .

4. Run the Docker container:

docker run -p {host-port}:{container-port} {image-name}

## Available Templates

Currently, the following templates are available:

- \*_Python Django_ - a simple web application framework for Python
- \*_Go Gin_ - a web framework written in Go (Golang)
- \*_Node Express_ - a web application framework for Node.js
