package main

import (
	"fmt"
	"log"

	//"time"
	//"sync"
	//"regexp"

	"github.com/vmware-labs/yaml-jsonpath/pkg/yamlpath"
	"gopkg.in/yaml.v3"
)

func main() {

	y := `---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: sample-deployment
spec:
  template:
    spec:
      containers:
      - name: nginx
        image: nginx
        ports:
        - containerPort: 80
      - name: nginy
        image: nginy
        ports:
        - containerPort: 81
`

	var n yaml.Node

	err := yaml.Unmarshal([]byte(y), &n)
	if err != nil {
		log.Fatalf("cannot unmarshal data: %v", err)
	}

	var pattern string
	fmt.Scanln(&pattern)

	p, err := yamlpath.NewPath(pattern) //e.g., "$..spec.containers[*].image"
	if err != nil {
		log.Fatalf("cannot create path: %v", err)
	}

	q, err := p.Find(&n)
	if err != nil {
		log.Fatalf("unexpected error: %v", err)
	}

	for _, i := range q {
		fmt.Println(i.Value)
	}
}
