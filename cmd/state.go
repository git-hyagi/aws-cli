package cmd

import (
	"fmt"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/spf13/cobra"
	"sync"
)

func state() *cobra.Command {
	return &cobra.Command{
		Use:   "state",
		Short: "Retrieve the state of the instances",
		RunE: func(cmd *cobra.Command, args []string) error {

			inst_name, _ := rootCmd.Flags().GetString("name")
			inst_owner, _ := rootCmd.Flags().GetString("owner")
			svc, result := getInstances(inst_name, inst_owner)

			// create a waitGroup to control the go routines execution
			wg := sync.WaitGroup{}

			for _, v := range result.Reservations {
				for _, instance := range v.Instances {
					wg.Add(1)

					// go routine to start, stop or print instance status
					go func(instance *ec2.Instance, svc *ec2.EC2) {

						// Gather the instance name by its tags
						var name string
						for _, instName := range instance.Tags {
							if *instName.Key == "Name" {
								name = *instName.Value
								break
							}
						}
						fmt.Println(*instance.InstanceId, *instance.State.Name, name)
						wg.Done()
					}(instance, svc)
				}
			}
			wg.Wait()

			return nil
		},
	}
}
