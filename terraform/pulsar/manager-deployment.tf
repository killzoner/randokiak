resource "kubernetes_deployment" "pulsar_manager" {
  metadata {
    name = "pulsar-manager"

    labels = {
      app = "manager"
    }
  }

  spec {
    replicas = 1

    selector {
      match_labels = {
        app = "manager"
      }
    }

    template {
      metadata {
        labels = {
          app = "manager"
        }
      }

      spec {
        volume {
          name = "manager-claim0"

          persistent_volume_claim {
            claim_name = "manager-claim0"
          }
        }

        container {
          name  = "manager"
          image = "apachepulsar/pulsar-manager:v0.1.0"
          args  = ["/bin/sh"]

          port {
            container_port = 9527
          }

          env {
            name  = "DRIVER_CLASS_NAME"
            value = "org.postgresql.Driver"
          }

          env {
            name  = "LOG_LEVEL"
            value = "DEBUG"
          }

          env {
            name  = "PASSWORD"
            value = "pulsar"
          }

          env {
            name  = "REDIRECT_HOST"
            value = "http://172.17.0.1"
          }

          env {
            name  = "REDIRECT_PORT"
            value = "9527"
          }

          env {
            name  = "URL"
            value = "jdbc:postgresql://127.0.0.1:5432/pulsar_manager"
          }

          env {
            name  = "USERNAME"
            value = "pulsar"
          }

          volume_mount {
            name       = "manager-claim0"
            mount_path = "/data"
          }
        }
      }
    }
  }
}

