resource "kubernetes_deployment" "rdkapi" {
  metadata {
    name = "rdkapi"

    labels = {
      app = "rdkapi"
    }
  }

  spec {
    replicas = 1

    selector {
      match_labels = {
        app = "rdkapi"
      }
    }

    template {
      metadata {
        labels = {
          app = "rdkapi"
        }
      }

      spec {
        container {
          name  = "rdkapi"
          image = "randokiak/randokiak-api:0.0.1"

          port {
            container_port = 8082
          }

          env {
            name  = "GEN_HOST"
            value = "rdkgen"
          }

          env {
            name  = "GEN_PORT"
            value = "50051"
          }

          env {
            name  = "NB_FETCH_PROFILE"
            value = "10"
          }

          env {
            name  = "PORT"
            value = "8082"
          }

          env {
            name  = "PULSAR_HOST"
            value = "pulsar"
          }

          env {
            name  = "PULSAR_PORT"
            value = "6650"
          }

          env {
            name  = "TOPIC"
            value = "randokiak-topic"
          }
        }
      }
    }
  }
}

