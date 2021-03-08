package utils

import (
	"fmt"
	"os"
	"os/exec"
)

//docker镜像制作
func DockerBuild(imagesName string, dockerFile string) {
	fmt.Printf("[i] Building docker image %v\n", imagesName)

	cmdArgs := []string{"build", "-f", dockerFile, "-t", imagesName}
	cmdArgs = append(cmdArgs, ".")
	cmd := newCommand("docker", cmdArgs...)
	err := cmd.Run()
	if err != nil {
		fmt.Println("[i] Error.")
	} else {
		fmt.Println("[i] Done.")
	}
}

//docker tag创建
func DockerTag(o, d string) error {
	c := newCommand("docker", "tag", o, d)
	return c.Run()
}

//docker push
func DockerPush(imagesName string) error {
	c := newCommand("docker", "push", imagesName)
	return c.Run()
}

func newCommand(name string, arg ...string) *exec.Cmd {
	c := exec.Command(name, arg...)
	c.Stdin = os.Stdin
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	return c
}
