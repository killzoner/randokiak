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
          # image = "apachepulsar/pulsar:2.5.0"
          image = "randokiak/pulsar-elastic-sink:0.0.1"
          args  = ["/bin/bash", "-c", "bin/apply-config-from-env.py conf/standalone.conf && bin/pulsar standalone"]
          # #auto deploy connector function
          ## TODO: see for auto deploy
          lifecycle {
            post_start {
              exec {
                command = ["/bin/bash", "-c", "/pulsar/create-connector-pulsar.sh"]
              }
            }
          }

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

