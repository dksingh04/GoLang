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
