package main

import (
	"os"

	"github.com/aws/aws-cdk-go/awscdk"
	ec2 "github.com/aws/aws-cdk-go/awscdk/awsec2"
	ecs "github.com/aws/aws-cdk-go/awscdk/awsecs"
	"github.com/aws/constructs-go/constructs/v3"
	"github.com/aws/jsii-runtime-go"
)

type StreamingservicebotInfraStackProps struct {
	awscdk.StackProps
}

func NewStreamingservicebotInfraStack(scope constructs.Construct, id string, props *StreamingservicebotInfraStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	// Create VPC and Cluster
	vpc := ec2.NewVpc(stack, jsii.String("StreamingServiceBotVpc"), &ec2.VpcProps{
		MaxAzs: jsii.Number(1),
	})

	cluster := ecs.NewCluster(stack, jsii.String("StreamingServiceBotECSCluster"), &ecs.ClusterProps{
		ClusterName: jsii.String("streamingservicebot-cluster"),
		Vpc: vpc,
	})

	// Create Task Definition
	taskDef := ecs.NewFargateTaskDefinition(stack, jsii.String("StreamingServiceBotTaskDef"), &ecs.FargateTaskDefinitionProps{
		MemoryLimitMiB: jsii.Number(512),
		Cpu:            jsii.Number(256),
	})
	container := taskDef.AddContainer(jsii.String("StreamingServiceBotContainer"), &ecs.ContainerDefinitionOptions{
		Environment: &map[string]*string{
			"SSB_REDDIT_ID": jsii.String(os.Getenv("SSB_REDDIT_ID")),
			"SSB_REDDIT_PASSWORD": jsii.String(os.Getenv("SSB_REDDIT_PASSWORD")),
			"SSB_REDDIT_SECRET": jsii.String(os.Getenv("SSB_REDDIT_SECRET")),
			"SSB_REDDIT_SUBREDDITS": jsii.String("hqtrackbot+streamingservicebot"),
			"SSB_REDDIT_USERNAME": jsii.String(os.Getenv("SSB_REDDIT_USERNAME")),
			"SSB_SPOTIFY_CLIENT_ID": jsii.String(os.Getenv("SSB_SPOTIFY_CLIENT_ID")),
			"SSB_SPOTIFY_CLIENT_SECRET": jsii.String(os.Getenv("SSB_SPOTIFY_CLIENT_SECRET")),
			"SSB_YOUTUBE_KEY": jsii.String(os.Getenv("SSB_YOUTUBE_KEY")),
		},
		Image: ecs.ContainerImage_FromRegistry(jsii.String("ghcr.io/scottbrenner/streamingservicebot:main"), &ecs.RepositoryImageProps{}),
	})

	container.AddPortMappings(&ecs.PortMapping{
		ContainerPort: jsii.Number(80),
		Protocol:      ecs.Protocol_TCP,
	})

	// Create Fargate Service
	service := ecs.NewFargateService(stack, jsii.String("StreamingServiceBotService"), &ecs.FargateServiceProps{
		Cluster:        cluster,
		ServiceName:    jsii.String("streamingservicebot-service"),
		TaskDefinition: taskDef,
	})

	awscdk.NewCfnOutput(stack, jsii.String("Service"), &awscdk.CfnOutputProps{Value: ecs.FargateService.ServiceName(service)})

	return stack
}

func main() {
	app := awscdk.NewApp(nil)

	NewStreamingservicebotInfraStack(app, "StreamingservicebotInfraStack", &StreamingservicebotInfraStackProps{
		awscdk.StackProps{
			Env: env(),
		},
	})

	app.Synth(nil)
}

// env determines the AWS environment (account+region) in which our stack is to
// be deployed. For more information see: https://docs.aws.amazon.com/cdk/latest/guide/environments.html
func env() *awscdk.Environment {
	// If unspecified, this stack will be "environment-agnostic".
	// Account/Region-dependent features and context lookups will not work, but a
	// single synthesized template can be deployed anywhere.
	//---------------------------------------------------------------------------
	// return nil

	// Uncomment if you know exactly what account and region you want to deploy
	// the stack to. This is the recommendation for production stacks.
	//---------------------------------------------------------------------------
	// return &awscdk.Environment{
	//  Account: jsii.String("123456789012"),
	//  Region:  jsii.String("us-east-1"),
	// }

	// Uncomment to specialize this stack for the AWS Account and Region that are
	// implied by the current CLI configuration. This is recommended for dev
	// stacks.
	//---------------------------------------------------------------------------
	return &awscdk.Environment{
		Account: jsii.String(os.Getenv("AWS_DEFAULT_ACCOUNT")),
		Region:  jsii.String(os.Getenv("AWS_DEFAULT_REGION")),
	}
}
