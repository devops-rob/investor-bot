# The name of the project
# https://www.waypointproject.io/docs/waypoint-hcl
project = "investor-bot"

# Input Variables
# https://www.waypointproject.io/docs/waypoint-hcl/variables
variable "namespace" {
  type        = string
  default     = "investobot"
  description = "Kubernetes Namespace for the application"
}

variable "image" {
  type        = string
  description = "Image registry location"
}

# Dynamic configuration values
# https://www.waypointproject.io/docs/app-config/dynamic
config {
  env = {
    COINBASE_PRO_SECRET = configdynamic("kubernetes", {
      name = "coinbase" # ConfigMap name
      key  = "secret"   # Key in the ConfigMap
    }),
    COINBASE_PRO_BASEURL = configdynamic("kubernetes", {
      name = "coinbase"
      key  = "baseurl"
    }),
    COINBASE_PRO_SANDBOX = configdynamic("kubernetes", {
      name = "coinbase"
      key  = "sandbox"
    }),
    COINBASE_PRO_PASSPHRASE = configdynamic("kubernetes", {
      name = "coinbase"
      key  = "passphrase"
    }),
    COINBASE_PRO_KEY = configdynamic("kubernetes", {
      name = "coinbase"
      key  = "key"
    })
  }
}

# app Stanza
# https://www.waypointproject.io/docs/waypoint-hcl/app
app "investor-bot" {
  labels = {
    "service" = "investor-bot",
    "env"     = "dev"
  }

  # Describes how to build this application during `waypoint up` or `waypoint build`
  # https://www.waypointproject.io/docs/waypoint-hcl/build
  build {
    use "pack" {}
    registry {
      use "docker" {
        image = var.image
        tag   = gitrefpretty()
      }
    }
  }

  deploy {
    use "kubernetes" {
      namespace    = var.namespace
      probe_path   = "/"
      service_port = 3000
    }
  }

  release {
    use "kubernetes" {
      namespace     = var.namespace
      load_balancer = true
    }
  }
}
