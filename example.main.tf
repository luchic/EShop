terraform {
  required_providers {
    docker = {
      source  = "kreuzwerker/docker"
      version = "~> 4.2.0"
    }
  }
}

provider "docker" {}

# ── Network ──────────────────────────────────────────────────────────

resource "docker_network" "shop_net" {
  name   = "shop-net"
  driver = "bridge"
}

# ── Volumes ──────────────────────────────────────────────────────────

resource "docker_volume" "postgres_data" {
  name = "postgres-data"
}

resource "docker_volume" "redis_data" {
  name = "redis-data"
}

resource "docker_volume" "web_data" {
  name = "web-data"
  driver = "local"
  driver_opts = {
    type   = "none"
    o      = "bind"
    device = "${path.module}/nginx/html"
  }
}

# ── Images ───────────────────────────────────────────────────────────

resource "docker_image" "redis" {
  name         = "redis:latest"
  keep_locally = false
}

resource "docker_image" "postgres" {
  name         = "postgres:17-alpine"
  keep_locally = false
}

resource "docker_image" "adminer" {
  name         = "adminer:4"
  keep_locally = false
}

resource "docker_image" "server" {
  name         = "server:latest"
  keep_locally = false

  build {
    context    = "./application"
    dockerfile = "Dockerfile"
    tag        = ["server:latest"]
  }
}

resource "docker_image" "migration" {
  name         = "migration:latest"
  keep_locally = false

  build {
    context    = "./application"
    dockerfile = "Dockerfile.Migration"
    tag        = ["migration:latest"]
  }
}

resource "docker_image" "nginx" {
  name         = "nginx:latest"
  keep_locally = false

  build {
    context    = "./nginx"
    dockerfile = "Dockerfile"
    tag        = ["nginx:latest"]
  }
}

# ── Containers ───────────────────────────────────────────────────────

resource "docker_container" "shop_redis" {
  name  = "shop-redis"
  image = docker_image.redis.image_id

  restart = "unless-stopped"

  volumes {
    volume_name    = docker_volume.redis_data.name
    container_path = "/data"
  }

  healthcheck {
    test         = ["CMD", "redis-cli", "ping"]
    interval     = "2s"
    timeout      = "5s"
    retries      = 5
    start_period = "10s"
  }

  networks_advanced {
    name = docker_network.shop_net.name
  }
}

resource "docker_container" "shop_db" {
  name  = "shop-postgres"
  image = docker_image.postgres.image_id

  restart = "unless-stopped"

  env = [
    "POSTGRES_DB=shop",
    "POSTGRES_USER=shop",
    "POSTGRES_PASSWORD=shop",
  ]

  volumes {
    volume_name    = docker_volume.postgres_data.name
    container_path = "/var/lib/postgresql/data"
  }

  healthcheck {
    test     = ["CMD-SHELL", "pg_isready -U shop -d shop"]
    interval = "2s"
    timeout  = "5s"
    retries  = 10
  }

  networks_advanced {
    name = docker_network.shop_net.name
  }
}

resource "docker_container" "adminer" {
  name  = "shop-adminer"
  image = docker_image.adminer.image_id

  restart = "unless-stopped"

  env = [
    "ADMINER_DEFAULT_SERVER=postgres",
  ]

  ports {
    internal = 8080
    external = 8081
  }

  networks_advanced {
    name = docker_network.shop_net.name
  }

  depends_on = [docker_container.shop_db]
}

resource "docker_container" "migration" {
  name  = "shop-migration"
  image = docker_image.migration.image_id

  networks_advanced {
    name = docker_network.shop_net.name
  }

  depends_on = [docker_container.shop_db]
}

resource "docker_container" "server" {
  name  = "shop-server"
  image = docker_image.server.image_id

  restart = "unless-stopped"

  ports {
    internal = 8080
    external = 8080
  }

  networks_advanced {
    name = docker_network.shop_net.name
  }

  depends_on = [
    docker_container.shop_redis,
    docker_container.shop_db,
  ]
}

resource "docker_container" "nginx" {
  name  = "nginx"
  image = docker_image.nginx.image_id

  restart = "on-failure"
  max_retry_count = 5

  ports {
    internal = 80
    external = 8082
  }

  volumes {
    host_path      = "${path.cwd}/nginx/nginx.conf"
    container_path = "/etc/nginx/nginx.conf"
  }

  volumes {
    volume_name    = docker_volume.web_data.name
    container_path = "/var/www/html"
  }

  networks_advanced {
    name = docker_network.shop_net.name
  }

  depends_on = [docker_container.server]
}
