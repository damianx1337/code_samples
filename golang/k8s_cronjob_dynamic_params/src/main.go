package main

import (
	"flag"
	"fmt"
	"context"
	"strings"
	"log"

	kubernetes "k8s.io/client-go/kubernetes"
	cruntimeconfig "sigs.k8s.io/controller-runtime/pkg/client/config"

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
	return kubernetes.NewForConfigOrDie(cruntimeconfig.GetConfigOrDie())
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
