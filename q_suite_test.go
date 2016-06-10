package q_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/sditools/q"
)

const REGION = "us-east-1"
const REDIS_ADDR = "localhost:6379"

var svc *sqs.SQS
var queueUrl string
var queue *Queue

var _ = BeforeSuite(func() {

	sess := session.New()
	svc = sqs.New(sess, aws.NewConfig().WithRegion(REGION))

	timestamp := time.Now().Local().Format("20060102150405")
	queueName := "test-queue-" + timestamp
	resp, err := svc.CreateQueue(&sqs.CreateQueueInput{
		QueueName: &queueName,
	})

	if err != nil {
		fmt.Println(err)
	}

	queueUrl = *resp.QueueUrl

	queue = New(queueUrl, REGION, QueueParams{})
})

var _ = AfterSuite(func() {
	_, err := svc.DeleteQueue(&sqs.DeleteQueueInput{
		QueueUrl: aws.String(queueURL),
	})

	if err != nil {
		fmt.Println(err)
	}
})

func sendTestMessage(message string, queueURL string) {
	svc.SendMessage(&sqs.SendMessageInput{
		MessageBody: aws.String(message),
		QueueUrl:    aws.String(queueURL),
	})
}

func TestQ(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Q Suite")
}
