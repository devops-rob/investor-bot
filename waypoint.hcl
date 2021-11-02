project = "investor-bot"

config {
  env = {
    COINBASE_PRO_SECRET = configdynamic("kubernetes", {
      name = "coinbase" # ConfigMap name
      key  = "secret"
    }),
    COINBASE_PRO_BASEURL = configdynamic("kubernetes", {
      name = "coinbase" # ConfigMap name
      key  = "baseurl"
    }),
    COINBASE_PRO_SANDBOX = configdynamic("kubernetes", {
      name = "coinbase" # ConfigMap name
      key  = "sandbox"
    }),
    COINBASE_PRO_PASSPHRASE = configdynamic("kubernetes", {
      name = "coinbase" # ConfigMap name
      key  = "passphrase"
    }),
    COINBASE_PRO_KEY = configdynamic("kubernetes", {
      name = "coinbase" # ConfigMap name
      key  = "key"
    })
  }
}

app "investor-bot" {
  labels = {
    "service" = "investor-bot",
    "env"     = "dev"
  }

  build {
    use "pack" {}
    registry {
      use "docker" {
        image = "onlydole/investor-bot"
        tag   = gitrefpretty()
      }
    }
  }

  deploy {
    use "kubernetes" {
      probe_path   = "/"
      service_port = 3000
    }
  }

  release {
    use "kubernetes" {
      load_balancer = true
    }
  }

}
