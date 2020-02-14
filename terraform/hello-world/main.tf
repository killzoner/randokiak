resource "kubernetes_pod" "nginx" {
  metadata {
    name = "nginx-example"
    labels = {
      App = "nginx"
    }
  }

  spec {
    container {
      image = "nginx:1.7.8"
      name  = "example"

      port {
        container_port = 80
      }
    }
  }
}

resource "kubernetes_service" "nginx" {
  metadata {
    name = "nginx-example"
  }
  spec {
    selector = {
      App = kubernetes_pod.nginx.metadata[0].labels.App
    }
    port {
      port        = 80
      target_port = 80
    }

    //type = "LoadBalancer" // LoadBalancer is not available on minikube or kind
    type = "NodePort"
  }
}

//LoadBalancer is not available on minikube or kind
/*
output "lb_ip" {
  value = kubernetes_service.nginx.load_balancer_ingress[0].ip
}
*/