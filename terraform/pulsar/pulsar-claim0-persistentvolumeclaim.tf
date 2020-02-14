resource "kubernetes_persistent_volume_claim" "pulsar_claim_0" {
  metadata {
    name = "pulsar-claim0"

    labels = {
      app = "pulsar-claim0"
    }
  }

  spec {
    access_modes = ["ReadWriteOnce"]

    resources {
      requests = {
        storage = "100Mi"
      }
    }
  }
}

