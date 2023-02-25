package main

import{
	"os"
	"github.com/aws/aws-skd-go/aws/session"
}

func main(){
	region := os.Getenv("AWS_REGION")
	awsSession, err := session.NewSession(&aws.Config{}})
}

const tableName = "LambdaInGoUser"

func handler(req events.APIGatewayProxyRequest) (*events.APIGatewayProxyRequest, error){
		switch req 
}

