## About
Very simple program to start, stop or print the state of some instances running on aws.
The flags "name" and "owner" are used as filters when making the ec2 actions (poweron/poweroff/state).  
   
For this program to run properly it is necessary to have the aws credentials stored in your home directory ($HOME/.aws/credentials).  
More information about aws credentials: https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-files.html  

Here is an example of ~/.aws/credentials file:
~~~
[default]
aws_access_key_id = 123456...
aws_secret_access_key = sfdlkjs.....
region = us-east-2
~~~

## Installing

* Just clone this repo and compile it:
~~~
git clone https://github.com/git-hyagi/aws-cli.git
cd aws-cli
go install aws-test.go
~~~

## Running

* power off instance:
~~~
aws-test poweroff -n <instance name> -o <owner>
~~~

* For example, to turn off all ec2 instances with name `yagi-*` and aws owner '1234':
~~~
aws-test poweroff -n `yagi-*` -o 1234
~~~

* power on instance:
~~~
aws-test poweron -n <instance name> -o <owner>
~~~

* get instance state:
~~~
aws-test state -n <instance name> -o <owner>
~~~


## Pre-reqs

You should have go installed to compile this program. Here is the steps to download and install go on Linux:
~~~
wget https://golang.org/dl/go1.15.2.linux-amd64.tar.gz
tar -C /usr/local -xzf go1.15.2.linux-amd64.tar.gz
rm -f go1.15.2.linux-amd64.tar.gz
~~~

Configure the env vars:
~~~
cat<<EOF>> ~/.bash_profile

# Go specific environment vars
export PATH=$PATH:/usr/local/go/bin
export GOPATH=$HOME/go
export GOBIN=$(go env GOPATH)/bin
export PATH=$PATH:$(go env GOPATH)/bin
EOF

source ~/.bash_profile
~~~
