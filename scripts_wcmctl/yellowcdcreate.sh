cd /home/ubuntu/go/src/github.com/cisco-app-networking/wcm-system/wcmctl 
./wcmctl create connectivitydomain --wait --kubeconfig /home/ubuntu/kubeconfigs/central/kind-1.kubeconfig --name yellow --prefixPool 172.45.0.0/16 --memberConfig=/home/ubuntu/kubeconfigs/nsm/kind-2.kubeconfig,/home/ubuntu/kubeconfigs/nsm/kind-3.kubeconfig
