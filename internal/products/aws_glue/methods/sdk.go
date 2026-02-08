package aws_glue_methods

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/glue"
)

type SDK struct{}

// product: aws_glue
// method: sdk
// args:
//	0 - path STRING

func parseArgs(args []any) (string, error) {
	var jobName string
	var err error
	if len(args) != 1 {
		return "", errors.New("not enough arguments passed into function")
	}
	jobName, ok := args[0].(string)
	if !ok {
		return "", errors.New("argument passed as JOB-NAME cannot be cast to string")
	}
	return jobName, err
}

func (c SDK) Run(id string, args []any) (string, error) {
	// TODO: implememt calling AWS SDK
	logger := slog.Default().With("product", "aws-glue", "method", "sdk", "check_id", id)
	jobName, err := parseArgs(args)
	if err != nil {
		logger.Error("failed to parse argments", "error", err)
		return "", err
	}

	cfg, err := config.LoadDefaultConfig(context.Background())
	if err != nil {
		logger.Error("failed to load credentials for aws account", "error", err)
		return "", err
	}

	glueClient := glue.NewFromConfig(cfg)
	job, err := glueClient.GetJob(
		context.Background(),
		&glue.GetJobInput{
			JobName: aws.String(jobName),
		},
	)
	if err != nil {
		logger.Error("failed to get job information", "error", err)
		return "", err
	}

	fmt.Println("job", job)

	return *job.Job.GlueVersion, nil
}
