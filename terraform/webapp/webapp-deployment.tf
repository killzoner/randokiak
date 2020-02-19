resource "kubernetes_deployment" "webapp" {
  metadata {
    name = "webapp"

    labels = {
      app = "webapp"
    }
  }

  spec {
    replicas = 1

    selector {
      match_labels = {
        app = "webapp"
      }
    }

    template {
      metadata {
        labels = {
          app = "webapp"
        }
      }

      spec {
        container {
          name  = "webapp"
          image = "randokiak/webapp:0.0.1"

          env {
            name  = "ELASTIC_HOST"
            value = "elasticsearch"
          }

          env {
            name  = "ELASTIC_PORT"
            value = "9200"
          }

          env {
            name  = "NGINX_LISTEN_PORT"
            value = "4200"
          }

          env {
            name  = "RDKAPI_API_HOST"
            value = "rdkapi"
          }

          env {
            name  = "RDKAPI_API_PORT"
            value = "8082"
          }
        }
      }
    }
  }
}

