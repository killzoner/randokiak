resource "kubernetes_service" "manager" {
  metadata {
    name = "manager"
  }

  spec {
    selector = {
      app = "manager"
    }
    port {
      port        = 9527
      target_port = 9527
    }

    type = "NodePort"
  }
}

