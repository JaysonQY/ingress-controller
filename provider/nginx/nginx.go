package nginx

import (
	fmt"
	"github.com/rancher/ingress-controller/config"
	"io"
	"os"
	"os/exec"
	"text/template"
        "github.com/Sirupsen/logrus"
        utils "github.com/rancher/ingress-controller/utils"
        "github.com/rancher/ingress-controller/provider"
)

func init() {
	var config string
	if config = os.Getenv("NGINX_CONFIG"); len(config) == 0 {
		logrus.Info("NGINX_CONFIG is not provided.")
		return
        } else {
		logrus.Info("NGIXN_CONFIG is provided.")
	}
	nginxCfg := &nginxConfig {
		ReloadCmd: "nginx -s reload",
                Config: config,
		Template: "/etc/nginx/nginx_template.cfg"
        }
	lbp := NginxProvider {
		cfg: nginxCfg
        }
	provider.RegisterProvider(lbp.GetName(), &lbp)
}

type NginxProvider struct {
	cfg *nginxConfig
}

type nginxConfig struct {
	Name      string
	ReloadCmd string
	Config    string
	Template  string
}

func (lbp *NginxProvider) IsHealthy() bool {
	return true
}

func (lbp *NginxProvider) Stop() error {
	return nil
}

func (lbp *NginxProvider) GetName() {
	return "nginx"
}

func (lbp *NginxProvider) GetPublicEndpoints(lbName string) []string {
        arr := []string{} 
	return arr
}

func (lbp *NginxProvider) Run(syncEndpointsQueue *utils.TaskQueue) {
        logrus.Infof("shutting down kubernetes-ingress-controller")
}

func (lbp *NginxProvider) CleanupConfig(name string) error {
        return nil 
}

func (lbp *NginxProvider) ApplyConfig(lbConfig []*config.LoadBalancerConfig) error {
	if err := lbp.cfg.write(lbConfig); err != nil {
		return err
	}
	return lbp.cfg.reload()
}

func (cfg *nginxConfig) reload() error {
	output, err := exec.Command("sh", "-c", cfg.ReloadCmd).CombinedOutput()
	msg := fmt.Sprintf("%v -- %v", cfg.Name, string(output))
	if err != nil {
		return fmt.Errorf("error restarting %v: %v", msg, err)
	}
	return nil
}

func (lbp *NginxProvider) write(lbConfig []*config.LoadBalancerConfig) (err error) {
	return err
}
