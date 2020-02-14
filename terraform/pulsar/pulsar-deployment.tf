resource "kubernetes_deployment" "pulsar" {
  metadata {
    name = "pulsar"

    labels = {
      app = "pulsar"
    }
  }

  spec {
    replicas = 1

    selector {
      match_labels = {
        app = "pulsar"
      }
    }

    template {
      metadata {
        labels = {
          app = "pulsar"
        }
      }

      spec {
        volume {
          name = "pulsar-claim0"

          persistent_volume_claim {
            claim_name = "pulsar-claim0"
          }
        }

        container {
          name  = "pulsar"
          image = "apachepulsar/pulsar:2.4.2"
          args  = ["/bin/bash", "-c", "bin/apply-config-from-env.py conf/standalone.conf && bin/pulsar standalone"]

          port {
            container_port = 8080
          }

          port {
            container_port = 6650
          }

          env {
            name  = "PULSAR_MEM"
            value = "\" -Xms512m -Xmx512m -XX:MaxDirectMemorySize=1g\""
          }

          volume_mount {
            name       = "pulsar-claim0"
            mount_path = "/pulsar/data"
          }
        }
      }
    }
  }
}

