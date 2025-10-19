package main

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	// "github.com/aws/aws-cdk-go/awscdk/v2/awssqs"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdynamodb"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type LambdaAtScaleStackProps struct {
	awscdk.StackProps
}

func NewLambdaAtScaleStack(scope constructs.Construct, id string, props *LambdaAtScaleStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	// create DB table here
	table := awsdynamodb.NewTable(stack, jsii.String("myUserDynamoTable"), &awsdynamodb.TableProps{
		PartitionKey: &awsdynamodb.Attribute{
			Name: jsii.String("username"),
			Type: awsdynamodb.AttributeType_STRING,
		},
		// this can be binded inside a .env
		TableName: jsii.String("userTable"),
	})

	myLambdaFunction := awslambda.NewFunction(stack, jsii.String("myLambdaFunction"), &awslambda.FunctionProps{
		Runtime: awslambda.Runtime_PROVIDED_AL2023(),
		Code:    awslambda.AssetCode_FromAsset(jsii.String("lambda/function.zip"), nil),
		Handler: jsii.String("main"),
	})

	table.GrantReadWriteData(myLambdaFunction)

	return stack
}

func main() {
	defer jsii.Close()

	app := awscdk.NewApp(nil)

	NewLambdaAtScaleStack(app, "LambdaAtScaleStack", &LambdaAtScaleStackProps{
		awscdk.StackProps{
			Env: env(),
		},
	})

	app.Synth(nil)
}

func env() *awscdk.Environment {
	return nil
}
