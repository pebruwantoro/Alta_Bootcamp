# Getting started with Okteto (Kubernetes)

* Download kubectl from kubernetes's official website
* Create Okteto account (cloud.okteto.com)
* Set `KUBECONFIG` environment:
	- Powershell (windows) user: $Env:KUBECONFIG=("$HOME\Downloads\okteto-kube.config;$Env:KUBECONFIG;$HOME\.kube\config")
	* Check your current kubectl 