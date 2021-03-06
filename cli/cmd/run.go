package cmd

import (
	"fmt"

	"github.com/dodopizza/kubectl-shovel/internal/kubernetes"
)

func run(
	options *commonOptions,
	tool string,
) error {
	k8s, err := kubernetes.NewClient(options.kubeFlags)
	if err != nil {
		return nil
	}
	pod, err := k8s.GetPodInfo(options.podName)
	if err != nil {
		return err
	}

	jobName := newJobName()
	containerInfo := kubernetes.GetContainerInfo(pod)
	jobVolume := kubernetes.NewJobVolume(containerInfo)
	fmt.Println("Spawning diagnostics job")
	err = k8s.RunJob(
		jobName,
		options.image,
		pod.Spec.NodeName,
		jobVolume,
		[]string{
			tool,
			"--container-id",
			containerInfo.ID,
			"--container-runtime",
			containerInfo.Runtime,
		},
	)
	if err != nil {
		return err
	}

	fmt.Println("Waiting for a job to start")
	jobPodName, err := k8s.WaitPod(map[string]string{"job-name": jobName})
	if err != nil {
		return err
	}
	readCloser, err := k8s.ReadPodLogs(jobPodName)
	if err != nil {
		return err
	}

	handleLogs(readCloser, options.output)
	fmt.Printf("Result successfully written to %s\n", options.output)

	err = k8s.DeleteJob(jobName)
	if err != nil {
		return err
	}

	return nil
}
