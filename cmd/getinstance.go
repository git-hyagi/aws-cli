package cmd

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"log"
)

func getInstances(inst_name, inst_owner string) (*ec2.EC2, *ec2.DescribeInstancesOutput) {

	sess := session.Must(session.NewSessionWithOptions(session.Options{SharedConfigState: session.SharedConfigEnable}))
	svc := ec2.New(sess)

	// filter instances by tag and owned by inst_owner
	params := &ec2.DescribeInstancesInput{
		Filters: []*ec2.Filter{
			{
				Name:   aws.String("tag:Name"),
				Values: []*string{aws.String(inst_name)},
			},
			{
				Name:   aws.String("owner-id"),
				Values: []*string{aws.String(inst_owner)},
			},
		},
	}

	// gather instances information
	result, err := svc.DescribeInstances(params)
	if err != nil {
		fmt.Println("Error", err)
		log.Fatal(err.Error())
	}

	return svc, result
}
