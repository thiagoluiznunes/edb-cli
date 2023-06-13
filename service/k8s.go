package service

import (
	"context"
	"edb-assessment/consts"
	"fmt"
	"math"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/rodaine/table"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

type Service struct {
	clientSet *kubernetes.Clientset
}

func NewK8SService() *Service {
	return &Service{
		clientSet: nil,
	}
}

func (s *Service) InitClusterConfig(config *rest.Config) (err error) {

	s.clientSet, err = kubernetes.NewForConfig(config)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) GetPods(namespace string) {

	if strings.TrimSpace(namespace) == "" {
		namespace = "default"
	}

	pods, err := s.clientSet.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		fmt.Println(err)
		return
	}
	if len(pods.Items) < 1 {
		fmt.Println(consts.NO_RESOURCES)
		return
	}

	headerFmt := color.New(color.FgGreen).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()

	tbl := table.New("NAME", "READY", "STATUS", "RESTARTS", "AGE")
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

	for _, pod := range pods.Items {
		currentTime := time.Now()
		podAge := currentTime.Sub(pod.Status.StartTime.Time).Minutes()
		podAgeLabel := fmt.Sprintf("%vm", math.Trunc(podAge))
		readyContainers := GetReadyContainers(pod)
		restarContainers := GetRestartContainers(pod)

		tbl.AddRow(pod.Name, readyContainers, pod.Status.Phase, restarContainers, podAgeLabel)
	}
	tbl.Print()
}

func GetRestartContainers(pod v1.Pod) (restartContainers int32) {

	restartContainers = 0
	for _, containerStatus := range pod.Status.ContainerStatuses {
		restartContainers += containerStatus.RestartCount
	}
	return restartContainers
}

func GetReadyContainers(pod v1.Pod) string {

	amountOfContainers := len(pod.Spec.Containers)
	readyContainers := 0
	for _, containerStatus := range pod.Status.ContainerStatuses {
		if containerStatus.Ready {
			readyContainers += 1
		}
	}
	return fmt.Sprintf("%v/%v", readyContainers, amountOfContainers)
}
