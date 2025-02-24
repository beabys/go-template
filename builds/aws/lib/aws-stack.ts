import * as cdk from 'aws-cdk-lib';
import { Construct } from 'constructs';
import * as lambda from 'aws-cdk-lib/aws-lambda';
// import * as sqs from 'aws-cdk-lib/aws-sqs';

export class AwsStack extends cdk.Stack {
  constructor(scope: Construct, id: string, props?: cdk.StackProps) {
    super(scope, id, props);

    // The code that defines your stack goes here

    // example resource
    // const queue = new sqs.Queue(this, 'AwsQueue', {
    //   visibilityTimeout: cdk.Duration.seconds(300)
    // });

    const api = new lambda.Function(this, 'API', {
      runtime: lambda.Runtime.PROVIDED_AL2023,
      handler: 'main',
      code: lambda.Code.fromAsset('../assets'),
      environment: {
        CONFIG_FILE: "./config.yaml",
        STAGE: 'my-stage',
      }
    });

    const functionUrl = api.addFunctionUrl({
      authType: lambda.FunctionUrlAuthType.NONE,
      cors: {
        allowedOrigins: ['*'],
        allowedMethods: [lambda.HttpMethod.GET],
        allowedHeaders: ['*'],
        // maxAge: cdk.Duration.days(10),
        // allowCredentials: true,
      }
    });

    new cdk.CfnOutput(this, 'APIUrl', {
      value: functionUrl.url,
    }
    )
  }
}
