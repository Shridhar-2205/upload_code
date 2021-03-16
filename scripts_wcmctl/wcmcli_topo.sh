
cd /home/ubuntu/go/src/github.com/cisco-app-networking/wcm-system/system_topo/
./setup_kind_clusters.sh --delete
./setup_kind_clusters.sh
cd ..
cd /home/ubuntu/go/src/github.com/cisco-app-networking/wcm-system/wcmctl 
./wcmctl install central-cluster-components --connectivitydomain-operator-appdns-org containers.cisco.com/appn --connectivitydomain-operator-nsr-org containers.cisco.com/appn --connectivitydomain-operator-org containers.cisco.com/appn --connectivitydomain-operator-wcmd-org containers.cisco.com/appn --connectivitydomain-operator-cdinfo-org containers.cisco.com/appn  --wait --kubeconfig /home/ubuntu/kubeconfigs/central/kind-1.kubeconfig
./wcmctl install member-cluster-components --wait --metallb-prefix 254 --central-kubeconfig /home/ubuntu/kubeconfigs/central/kind-1.kubeconfig --membercore-operator-nsediscovery-org containers.cisco.com/appn --membercore-operator-nsm-org containers.cisco.com/appn --membercore-operator-org containers.cisco.com/appn --nse-operator-nse-org containers.cisco.com/appn --nse-operator-org containers.cisco.com/appn  --kubeconfig /home/ubuntu/kubeconfigs/nsm/kind-2.kubeconfig
./wcmctl install member-cluster-components --wait --metallb-prefix 253 --central-kubeconfig /home/ubuntu/kubeconfigs/central/kind-1.kubeconfig --membercore-operator-nsediscovery-org containers.cisco.com/appn --membercore-operator-nsm-org containers.cisco.com/appn --membercore-operator-org containers.cisco.com/appn --nse-operator-nse-org containers.cisco.com/appn --nse-operator-org containers.cisco.com/appn  --kubeconfig /home/ubuntu/kubeconfigs/nsm/kind-3.kubeconfig