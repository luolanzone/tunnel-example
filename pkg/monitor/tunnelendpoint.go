package monitor

import (
	"context"
	"os"
	"time"

	v1alpha1 "github.com/luolanzone/tunnel-example/apis/gateway/v1alpha1"
	clientset "github.com/luolanzone/tunnel-example/pkg/client/clientset/versioned"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/kubernetes"
	"k8s.io/klog"
)

var (
	controllerName = "tunnelEndpointController"
)

type tunnelEndpointMonitor struct {
	client            clientset.Interface
	tunnelEndpointCRD *v1alpha1.TunnelEndpoint
	k8sClient         kubernetes.Interface
	ctx               context.Context
	//nodeInformer   coreinformers.NodeInformer
	//nodeListerSynced cache.InformerSynced
}

func NewTunnelEndpointController(client clientset.Interface, k8sClient kubernetes.Interface, ctx context.Context) *tunnelEndpointMonitor {
	return &tunnelEndpointMonitor{
		client:            client,
		k8sClient:         k8sClient,
		tunnelEndpointCRD: nil,
		ctx:               ctx,
	}
}

func (m *tunnelEndpointMonitor) Run(stopCh <-chan struct{}) {
	klog.Infof("Starting %s", controllerName)
	defer klog.Infof("Shutting down %s", controllerName)
	wait.Until(m.syncTunnelEndpointCRD, time.Minute, stopCh)
}

func (m *tunnelEndpointMonitor) syncTunnelEndpointCRD() {
	var err error
	if m.tunnelEndpointCRD != nil {
		if m.tunnelEndpointCRD, err = m.updateTunnelEndpointCRD(); err == nil {
			return
		}
		klog.Errorf("Failed to update tunnelendpoint CRD: %v", err)
		m.tunnelEndpointCRD = nil
	} else {
		m.tunnelEndpointCRD, err = m.createTunnelEndpointCRD()
		if err != nil {
			klog.Errorf("Failed to create tunnelendpoint CRD: %v", err)
			m.tunnelEndpointCRD = nil
			return
		}
	}
}

func (m *tunnelEndpointMonitor) updateTunnelEndpointCRD() (*v1alpha1.TunnelEndpoint, error) {
	role := m.ctx.Value("role")
	if role != nil {
		m.tunnelEndpointCRD.Spec.Role = role.(string)
	}
	return m.client.GatewayV1alpha1().TunnelEndpoints().Update(context.TODO(), m.tunnelEndpointCRD, metav1.UpdateOptions{})
}

func (m *tunnelEndpointMonitor) getTunnelEndpointCRD(crdname string) (*v1alpha1.TunnelEndpoint, error) {
	return m.client.GatewayV1alpha1().TunnelEndpoints().Get(context.TODO(), crdname, metav1.GetOptions{})
}

func (m *tunnelEndpointMonitor) createTunnelEndpointCRD() (*v1alpha1.TunnelEndpoint, error) {
	tunnelEndpoint := new(v1alpha1.TunnelEndpoint)
	hostname, err := os.Hostname()
	tunnelEndpoint.Name = hostname
	if err != nil {
		klog.Fatal("fail to get node hostname")
	}
	publicIP, err := GetPublicIP()
	if err != nil {
		klog.Error("fail to get node public ip")
	}
	tunnelEndpoint.Spec = v1alpha1.TunnelEndpointSpec{
		Role:      "Unknown",
		ClusterID: "clusterid",
		Hostname:  hostname,
		Subnets:   nil,
		PrivateIP: GetOutboundIP().String(),
		PublicIP:  publicIP,
	}
	return m.client.GatewayV1alpha1().TunnelEndpoints().Create(context.TODO(), tunnelEndpoint, metav1.CreateOptions{})
}
