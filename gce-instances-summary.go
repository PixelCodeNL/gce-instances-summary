package main

import (
	"fmt"
	"github.com/mgutz/ansi"
	"golang.org/x/net/context"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/cloudresourcemanager/v1"
	"google.golang.org/api/compute/v1"
	"os"
)

// DefaultComputeZone Default zone for all instances
const DefaultComputeZone = "europe-west1-d"

func main() {
	colorTitle := ansi.ColorCode("white+h:grey")
	colorGrey := ansi.ColorCode("grey")
	colorGreen := ansi.ColorCode("green")
	colorRed := ansi.ColorCode("red")
	colorReset := ansi.ColorCode("reset")

	ctx := context.Background()
	client, err := google.DefaultClient(ctx, compute.ComputeScope)
	if err != nil {
		fmt.Printf("\nError: %s\n", err.Error())
		os.Exit(1)
	}

	resourceManager, _ := cloudresourcemanager.New(client)
	projectsResponse, _ := resourceManager.Projects.List().Do()
	compute, _ := compute.New(client)
	for _, project := range projectsResponse.Projects {
		zonesResponse, err := compute.Zones.List(project.ProjectId).Do()
		if err == nil {
			for _, zone := range zonesResponse.Items {
				instancesResponse, err := compute.Instances.List(project.ProjectId, zone.Name).Do()
				if err == nil {
					if len(instancesResponse.Items) > 0 {
						fmt.Printf("%s\n\nProject %s (%s)%s\n\n", colorTitle, project.ProjectId, zone.Name, colorReset)
						for _, instance := range instancesResponse.Items {
							statusColorStart := colorGrey
							switch instance.Status {
							case "RUNNING":
								statusColorStart = colorGreen
							case "STOPPED", "STOPPING", "SUSPENDED", "SUSPENDING", "TERMINATED":
								statusColorStart = colorRed
							}
							fmt.Printf("%30s - %s%s%s\n", instance.Name, statusColorStart, instance.Status, colorReset)
						}
					}
				}
			}
		}
	}
}
