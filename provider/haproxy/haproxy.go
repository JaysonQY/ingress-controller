package haproxy

import (
	"fmt"
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
	if config = os.Getenv("HAPROXY_CONFIG"); len(config) == 0 {
                logrus.Info("HAPROXY_CONFIG is not provided.")
		return
	} else {
                logrus.Info("HAPROXY_CONFIG is  provided.")
        }
	haproxyCfg := &haproxyConfig{
		ReloadCmd: "/etc/haproxy/haproxy_reload",
		Config:    config,
		Template:  "/etc/haproxy/haproxy_template.cfg",
	}
	lbp := HAProxyProvider{
		cfg: haproxyCfg,
	}
	provider.RegisterProvider(lbp.GetName(), &lbp)
}

type HAProxyProvider struct {
	cfg *haproxyConfig
}

type haproxyConfig struct {
	Name      string
	ReloadCmd string
	Config    string
	Template  string
}

func (cfg *haproxyConfig) write(lbConfig []*config.LoadBalancerConfig) (err error) {
        files := [3]string{"frontend.cfg", "use.cfg", "backend.cfg"}
        for _, value := range files { 
		var w io.Writer
		w, err = os.Create("/etc/haproxy/" + value)
		if err != nil {
			return err
		}
		var t *template.Template
		t, err = template.ParseFiles("/etc/haproxy/haproxy_" + value)
		if err != nil {
			return err
		}
		conf := make(map[string]interface{})
		conf["lbconf"] = lbConfig
		logrus.Info("Get Conf %v", conf) 
		err = t.Execute(w, conf)
        }
        var cmd string
        cmd = "cat /etc/haproxy/frontend.cfg  /etc/haproxy/use.cfg /etc/haproxy/backend.cfg > /etc/haproxy/haproxy_tmp.cfg"
        output, err := exec.Command("sh", "-c", cmd).CombinedOutput()
        exec.Command("sh", "-c", "sed '/^$/d' /etc/haproxy/haproxy_tmp.cfg > /etc/haproxy/haproxy.cfg").CombinedOutput()
        fmt.Sprintf("%v ", string(output))
        return err
}

func (lbc *HAProxyProvider) ApplyConfig(lbConfig []*config.LoadBalancerConfig) error {
	if err := lbc.cfg.write(lbConfig); err != nil {
		return err
	}
	return lbc.cfg.reload()
}

func (lbc *HAProxyProvider) GetName() string {
	return "haproxy"
}

func (lbc *HAProxyProvider) GetPublicEndpoints(lbName string) []string {
        arr := []string{} 
	return arr
}

func (lbc *HAProxyProvider) IsHealthy() bool {
        return true
}

func (lbc *HAProxyProvider) Stop() error{
        return nil
}

func (lbc *HAProxyProvider) Run(syncEndpointsQueue *utils.TaskQueue) {
        logrus.Infof("shutting down kubernetes-ingress-controller")
}

func (lbc *HAProxyProvider) CleanupConfig(name string) error {
        return nil 
}
 
func (cfg *haproxyConfig) reload() error {
	output, err := exec.Command("sh", "-c", cfg.ReloadCmd).CombinedOutput()
	msg := fmt.Sprintf("%v -- %v", cfg.Name, string(output))
	if err != nil {
		return fmt.Errorf("error restarting %v: %v", msg, err)
	}
	return nil
}
