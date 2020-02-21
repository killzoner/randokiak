resource "kubernetes_ingress" "lb_ingress" {
  metadata {
    name      = "lb-ingress"

    annotations = {
      "nginx.ingress.kubernetes.io/rewrite-target" = "/$2"
    }
  }

  spec {
    rule {
      host = "rdk.io"
      http {
        path {
          path = "/rdkapi(/|$)(.*)"
          backend {
            service_name = "rdkapi"
            service_port = "8082"
          }
        }
      }
    }
  }
}

