terraform {
    required_providers {
        docker = {
            source = "kreuzwerker/docker"
            version = "~> 4.2.0"
        }
    }
}

provider "docker" {}


resource "docker_image" "nginx" {
    name = "nginx:latest"
    keep_locally = false
}

resource "docker_image" "shop-redis" {
    name = "redis:latest"
    keep_locally = false
}

resource "docker_image" "shop-db" {
    name = "postgres:17-alpine"
    keep_locally = false
}

resource "docker_image" "adminer" {
    name = "adminer:4"
    keep_locally = false
}

resource "docker_image" "migration" {
    name = "migration:latest"
    keep_locally = false

    build {
        context = "./application"
        dockerfile = "Dockerfile.Migration"
        tag = ["migration:latest"]
    }
}

resource "docker_image" "server" {
    name
}