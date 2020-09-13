package v1

import (
	"context"
	v1 "feedback-generator/pkg/api/v1/feedbackreqpb"
	"fmt"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	apiVersion = "v1"
)

type feedbackServiceServer struct {
	client *mongo.Client
	db     *mongo.Database
	logger *logrus.Logger
}

//Initializing feedback mapping comments, this information we can have in DB, currently having it in code.
//TODO move this mapping in DB
var feedbackMapping = map[string]string{
	"CodeCompiled":  "written compilable and executable code",
	"PseudoCode":    "able to write pseudo code",
	"AlgoEfficient": "it was efficient and candidate has considered space and Time complexity, while implementing the solution",
	//This will go in summary and skill comments
	"Proxy":                "using proxy and someone else was giving Interview on behalf of him, since it was proxy hence I have done some basic discussion on each skill sets.",
	"Whiteboard_explained": "in white-boarding session, candidate has performed very well, explained the solution with proper diagram and flow.",
	"Whiteboard_partial":   "in white-boarding session, candidate was partially able to explain the solution",
	"Coding_Standards":     "well-versed with coding standards and followed the same while writing the code",
	"s-1":                  "substantial development in %v skil and have to work a lot, candidate was missing fundamentals",
	"e-1":                  "no experience in this skill, unable to demonstrate his experience",
	"s-2":                  "some training to bring competency up to standards, have some basic understanding but missing some other fundamentals",
	"e-2":                  "limited experience, close supervision will be needed for him",
	"s-3":                  "competent and can perform his task, no additional training is required at this time",
	"e-3":                  "competent and can complete assignments with reasonable supervision",
	"s-4":                  "above average, and competent in this skill, no training required",
	"e-4":                  "considerable experience and can perform his tasks with very minimal supervision",
	"s-5":                  "expert in this skill and can teach and mentor others in the team",
	"e-5":                  "extensive experience and can work independently",
	"HaveTheoretical":      "theoretically clear and explained the concepts of %v very well",
	"NoTheoretical":        " not clear with theoretical part of %v(topic-name), unable to explain %v(theory question)",
	"InDepthUnderstanding": "deep understanding of the concepts and explained the concepts with example",
	"AbleToExplain":        "explained the concept very well with example",
	"Partially Explained":  "able to partially explain the concepts, missing in-depth understanding of the concept, this will cause challenge in debugging/troubleshooting the problems",
	"Hnads-On":             "hands-on with the skill",
	"ScenarioQuestioned":   "I have covered scenarion questions",
	"ScenarioExplained":    "explained the scenario question (%v) very well and how to solve the problem in such cases",
	"ScenarioNotExplained": "unable to explain the scenario question (%v), seems to me not much hands-on in this skill",
}

var candidate = "Candidate"
var was = "was"
var has = "has"
var needs = "needs"
var is = "is"
var can = "can"
var topicsCovered = "We have covered following topics and this is how candidate has performed:"

// NewFeedbackServiceServer creates FeedbackServiceServer
func NewFeedbackServiceServer(client *mongo.Client, db *mongo.Database, logger *logrus.Logger) v1.FeedbackServiceServer {
	return &feedbackServiceServer{client: client, db: db, logger: logger}
}

func (fs *feedbackServiceServer) checkAPI(api string) error {
	if len(api) > 0 {
		if apiVersion != api {
			return status.Errorf(codes.Unimplemented,
				"unsupported API version: service implements API version '%s', but asked for '%s'", apiVersion, api)
		}
	}
	return nil
}

func (fs *feedbackServiceServer) Create(ctx context.Context, req *v1.FeedbackRequest) (*v1.FeedbackResponse, error) {

	if err := fs.checkAPI(req.Api); err != nil {
		fmt.Println(err)
		return new(v1.FeedbackResponse), err
	}
	fReq := req.GetFeedbackReq()
	// Create new request Id
	fReq.Id = primitive.NewObjectID().Hex()

	coll := fs.db.Collection("feedback_request")
	//Insert created FeedbackRequest
	result, err := coll.InsertOne(ctx, fReq)

	if result.InsertedID == nil && err != nil {
		fs.logger.WithFields(logrus.Fields{
			"request": fReq,
			"status":  http.StatusInternalServerError,
			"Error":   err,
		}).Error("Unable Insert the Document!")

		fRes := &v1.FeedbackResponse{
			Api:         "v1",
			StatusCode:  http.StatusInternalServerError,
			RequestId:   "",
			Message:     fmt.Sprintln(err),
			FeedbackRes: nil,
		}
		return fRes, nil
	}

	fResult := coll.FindOne(ctx, bson.D{primitive.E{Key: "id", Value: fReq.Id}})
	feedbackRes := v1.Feedback{}

	if err := fResult.Decode(&feedbackRes); err != nil {
		logrus.Errorf("Unable to read document for request id: %v", fReq.Id)
		return nil, nil
	}

	fRes := &v1.FeedbackResponse{
		Api:         "v1",
		StatusCode:  http.StatusCreated,
		RequestId:   fReq.Id,
		Message:     "Document Inserted Successfuly",
		FeedbackRes: &feedbackRes,
	}
	return fRes, nil
}

func (fs *feedbackServiceServer) Read(ctx context.Context, req *v1.ReadFeedbackRequest) (*v1.FeedbackResponse, error) {
	coll := fs.db.Collection("feedback_request")
	fResult := coll.FindOne(ctx, bson.D{primitive.E{Key: "id", Value: req.RequestId}})
	feedbackRes := v1.Feedback{}
	fRes := &v1.FeedbackResponse{}
	fRes.Api = "v1"
	fRes.RequestId = req.RequestId

	if err := fResult.Decode(&feedbackRes); err != nil {
		logrus.Errorf("Unable to read document for request id: %v", req.RequestId)
		fRes.StatusCode = http.StatusNotFound
		fRes.Message = fmt.Sprint(fResult.Decode(bson.M{}))
		return fRes, nil
	}

	fRes.StatusCode = http.StatusCreated
	fRes.Message = "Document Inserted Successfuly"
	fRes.FeedbackRes = &feedbackRes

	return fRes, nil
}
func (fs *feedbackServiceServer) GenerateFeedbackForRequest(ctx context.Context, req *v1.FeedbackRequest) (*v1.GeneratedFeedbackResponse, error) {

	return new(v1.GeneratedFeedbackResponse), nil
}

func (fs *feedbackServiceServer) Delete(ctx context.Context, req *v1.DeleteFeedbackRequest) (*v1.DeleteFeedbackResponse, error) {
	coll := fs.db.Collection("feedback_request")
	result, err := coll.DeleteOne(ctx, bson.D{primitive.E{Key: "id", Value: req.RequestId}})

	if err != nil {
		fs.logger.Errorf("Unable to delete document for request id: %v Deleted Count: %v", req.RequestId, result.DeletedCount)
		return nil, nil
	}
	var fRes = &v1.DeleteFeedbackResponse{}
	fRes.Api = "v1"
	if result.DeletedCount == 0 {
		fRes.StatusCode = http.StatusNotFound
	} else {
		fRes.StatusCode = http.StatusOK
	}

	return fRes, nil
}
