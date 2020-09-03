package main

import (
	"feedback-generator/pkg/api/v1/feedbackreqpb"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
)

func main() {
	fmt.Println("Feedback generator!")
	fr := createFeedbackRequest()
	saveFeedbackRequest("feedback.bin", fr)

	frRead := &feedbackreqpb.FeedbackRequest{}

	readSavedFeedbackRequest("feedback.bin", frRead)

	fmt.Println(frRead)

}
func createFeedbackRequest() *feedbackreqpb.FeedbackRequest {
	fr := feedbackreqpb.FeedbackRequest{
		CandidateName:           "Deepak Singh",
		InterviewDate:           "some Date",
		RecruiterName:           "Karthik Harish",
		TypeOfJob:               "Full-Stack Developer",
		IsProxy:                 false,
		IsIdRequired:            false,
		IsCodingRequired:        true,
		IsAbleToWritePseudoCode: true,
		IsAlgoEfficient:         true,
		IsWhiteboardingRequired: false,
		IsWhiteboardDone:        false,
		CreatDate:               ptypes.TimestampNow(),
		UpdateDate:              ptypes.TimestampNow(),
		NotesComment:            "Notes",
	}

	fmt.Println(fr)
	return &fr
}

func saveFeedbackRequest(fname string, reqpb proto.Message) error {
	out, err := proto.Marshal(reqpb)
	if err != nil {
		log.Fatalln("Error in Serializing Byte!", err)
		return err
	}

	if err := ioutil.WriteFile(fname, out, 0644); err != nil {
		log.Fatalln("Cannot write to file!", err)
		return err
	}
	fmt.Println("Data written to File Successfully!")
	return nil
}

func readSavedFeedbackRequest(fname string, reqpb proto.Message) error {
	in, err := ioutil.ReadFile(fname)
	if err != nil {
		log.Fatalln("Error in Reading Serializing Byte!", err)
		return err
	}
	err = proto.Unmarshal(in, reqpb)
	if err != nil {
		log.Fatalln("Could not put serialized byte into pb message", err)
		return err
	}

	return nil
}
