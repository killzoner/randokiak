resource "kubernetes_persistent_volume_claim" "manager_claim_0" {
  metadata {
    name = "manager-claim0"

    labels = {
      app = "manager-claim0"
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

