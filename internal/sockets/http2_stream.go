package sockets

import (
	"fmt"
	"slices"

	"github.com/google/uuid"
)

type Http2Stream struct {
	activeFlow *Flow
	// activeUuid is used to keep track of the request uuid so that when a response is processed, its flow can have the same uuid as the request
	activeUuid *string
}

func NewHttp2Stream() *Http2Stream {
	return &Http2Stream{}
}

func (stream *Http2Stream) ProcessFrame(frame *Http2Frame) *Flow {
	// Only accept complete header or data frames, the rest are ignored
	acceptedTypes := []uint8{frameTypeData, frameTypeHeaders}
	if !frame.Complete() || !slices.Contains(acceptedTypes, frame.Type()) {
		return nil
	}

	if frame.HasCompleteHeaders() {
		return stream.processHeaderFrame(frame)

	} else if frame.HasCompleteBody() {
		return stream.processDataFrame(frame)
	}

	return nil
}

func (stream *Http2Stream) processHeaderFrame(frame *Http2Frame) *Flow {
	if frame.IsRequest() {
		stream.activeFlow = NewFlow(
			uuid.NewString(),
			"0.0.0.0",
			"127.0.0.1:80",
			"tcp",
			"http2",
			123,
			5,
			[]byte(frame.HeadersText()),
		)
		stream.activeUuid = &stream.activeFlow.UUID
	} else {
		if stream.activeUuid == nil {
			fmt.Println("ERROR: no active request UUID for this response")
			return nil
		}

		stream.activeFlow = NewFlowResponse(
			*stream.activeUuid,
			"0.0.0.0",
			"127.0.0.1:80",
			"tcp",
			"http2",
			123,
			5,
			[]byte(frame.HeadersText()),
		)
	}

	// If there is no body in the request then send the flow back
	if frame.Flags().EndStream {
		flow := *stream.activeFlow
		stream.clearActiveFlow()
		if len(flow.Response) > 0 {
			stream.clearActiveUuid()
		}

		return &flow
	}

	return nil
}

func (stream *Http2Stream) processDataFrame(frame *Http2Frame) *Flow {
	// TODO: This should check the END_STREAM flag to know if the headers are actually complete

	if stream.activeFlow == nil {
		fmt.Println("ERROR: received http2 data frame but no active Flow")
		return nil
	}

	stream.activeFlow.AddData(frame.Payload())

	// Send the flow back
	flow := *stream.activeFlow
	stream.clearActiveFlow()
	if len(flow.Response) > 0 {
		stream.clearActiveUuid()
	}

	return &flow
}

func (stream *Http2Stream) clearActiveFlow() {
	stream.activeFlow = nil
}

func (stream *Http2Stream) clearActiveUuid() {
	stream.activeUuid = nil
}