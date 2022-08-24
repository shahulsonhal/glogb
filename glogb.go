package glogb

import "k8s.io/klog/v2"

func LogB() {
	klog.Info("Starting transaction...BBB")
	klog.ContextualLogger(true)
}
