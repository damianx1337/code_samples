package main

import (
	"flag"
	"fmt"
	"os"

	// necessary for k8s
	"context"
	"strings"
	"log"
	"path/filepath"

	// necessary for k8s
	//clientcmd "k8s.io/client-go/1.5/tools/clientcmd"
	clientcmd "k8s.io/client-go/tools/clientcmd"
	kubernetes "k8s.io/client-go/kubernetes"

	batchv1 "k8s.io/api/batch/v1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func main() {
	jobName := flag.String("jobname", "test-job", "The name of the job")
	containerImage := flag.String("image", "ubuntu:latest", "Name of the container image")
	entryCommand := flag.String("command", "ls", "The command to run inside the container")

	flag.Parse()

	fmt.Printf("Args : %s %s %s\n", *jobName, *containerImage, *entryCommand)

	clientset := connectK8sCluster()
	launchK8sCronJob(clientset, jobName, containerImage, entryCommand)
}

func connectK8sCluster() *kubernetes.Clientset {
	home, exists := os.LookupEnv("HOME")
	if !exists {
		home = "/root"
	}

	configPath := filepath.Join(home, ".kube", "config")

	config, err := clientcmd.BuildConfigFromFlags("", configPath)
	if err != nil {
		log.Panicln("failed to create k8s config")
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Panicln("Failed to generate k8s clientset")
	}

	return clientset
}

func launchK8sCronJob(clientset *kubernetes.Clientset, jobName *string, image *string, cmd *string) {
	jobs := clientset.BatchV1().Jobs("default")
	var backOffLimit int32 = 0

	jobSpec := &batchv1.Job{
		ObjectMeta: metav1.ObjectMeta{
			Name: *jobName,
			Namespace: "default",
		},
		Spec: batchv1.JobSpec{
			Template: v1.PodTemplateSpec{
				Spec: v1.PodSpec{
					Containers: []v1.Container{
						{
							Name: *jobName,
							Image: *image,
							Command: strings.Split(*cmd, ""),
						},	
					},
					RestartPolicy: v1.RestartPolicyNever,
				},
			},
			BackoffLimit: &backOffLimit,
		},
	}

	_, err := jobs.Create(context.TODO(), jobSpec, metav1.CreateOptions{})
	if err != nil {
		log.Fatalln("failed to create k8s job")
	}
	
	log.Println("successfully created k8s job")
}
