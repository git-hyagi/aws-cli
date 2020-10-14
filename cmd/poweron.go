package cmd

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/spf13/cobra"
	"sync"
)

func powerOn() *cobra.Command {
	return &cobra.Command{
		Use:   "poweron",
		Short: "Power on the instances",
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

						fmt.Println("Starting ", *instance.InstanceId, name)
						startInstance(svc, *instance.InstanceId)

						wg.Done()
					}(instance, svc)
				}
			}
			wg.Wait()

			return nil
		},
	}
}

// function to start an instance
func startInstance(svc *ec2.EC2, instance string) {
	instance_id := &ec2.StartInstancesInput{InstanceIds: []*string{aws.String(instance)}}

	_, err := svc.StartInstances(instance_id)

	// **TO-DO** abort program execution in case of error
	// or, at least, handle the error in a better way
	if err != nil {
		fmt.Println(err)
	}
}
