package main

import (
	"context"
	"feedback-generator/internal/config"
	f "feedback-generator/pkg/api/v1/feedbackreqpb"
	"fmt"
	"time"

	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/ptypes"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

const (
	// apiVersion is version of API is provided by server
	apiVersion = "v1"
)

var serverAddr = "localhost:9090"

func main() {
	logger, err := config.CreateDefaultLogConfiguration()
	if err != nil {
		logger.WithFields(logrus.Fields{
			"filename": "logger",
			"status":   500,
			"Error":    err,
		}).Fatal("Unable to read the Config file given!")
	}
	//serverAddr := flag.String("server", "localhost:9090", "The server address in the format of host:port")
	//Connect to grpc server
	conn, err := grpc.Dial("localhost:9090", grpc.WithInsecure())
	defer conn.Close()
	if err != nil {
		logger.WithFields(logrus.Fields{
			"server": "localhost:9090",
			"status": 500,
			"Error":  err,
		}).Fatalln("Unable to connect to grpc server")
	}

	client := f.NewFeedbackServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	t := time.Now().In(time.UTC)
	createDate, _ := ptypes.TimestampProto(t)

	fReq := f.FeedbackRequest{
		Api: apiVersion,
		FeedbackReq: &f.Feedback{
			CandidateName:           "Deepak Singh",
			RecruiterName:           "Amanda",
			CreatDate:               createDate,
			UpdateDate:              createDate,
			IsCodingRequired:        true,
			IsAbleToWritePseudoCode: true,
			IsAlgoEfficient:         true,
			IsCodeCompiled:          false,
			IsIdRequired:            false,
			IsProxy:                 false,
			IsWhiteboardingRequired: false,
			IsWhiteboardDone:        false,
			JobType:                 "Full-Stack Java Developer",
			MyComments:              "Good Candidate",
			TechSkills: []*f.TechSkill{
				&f.TechSkill{
					SkillName:        "Java",
					ExperienceRating: 3,
					SkillRating:      3,
					Topics: []*f.Topic{
						&f.Topic{
							TopicName:                 "OOPs concept",
							HaveTheroreticalKnowledge: true,
							InDepthUnderstanding:      true,
							IsAbleToExaplain:          true,
							IsAbleToExplainScenario:   true,
							IsScenarioCovered:         true,
							IsHandsOn:                 true,
							PartiallyExplained:        false,
							WhatSceanrioQuestion:      "How OOPs concept used in his project, explain with example",
						},
					},
				},
			},
		},
	}

	//fRes, err := client.CreateSimpleRequest(ctx, &sReq)
	fRes, err := client.Create(ctx, &fReq)
	fmt.Println(fRes)

	if err != nil {
		marshaler := jsonpb.Marshaler{}
		jsonReq, errMarshaling := marshaler.MarshalToString(&fReq)
		if errMarshaling != nil {
			logger.WithFields(logrus.Fields{
				"reqMessage": fReq.FeedbackReq,
				"status":     500,
				"Error":      err,
			}).Fatalln("Error in Transforming Request message to Json String!")
		}
		logger.WithFields(logrus.Fields{
			"request": jsonReq,
			"status":  500,
			"Error":   err,
		}).Fatalln("Unable to connect to grpc server")
	}

}
