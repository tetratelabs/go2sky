package go2sky

import (
	"github.com/tetratelabs/go2sky/propagation"
	"sync/atomic"

	"github.com/tetratelabs/go2sky/pkg"
	"github.com/tetratelabs/go2sky/reporter/grpc/common"
	v2 "github.com/tetratelabs/go2sky/reporter/grpc/language-agent-v2"
)

func newSegmentSpan(defaultSpan *defaultSpan, parentSpan segmentSpan) (s segmentSpan) {
	ssi := &segmentSpanImpl{
		defaultSpan:    *defaultSpan,
	}
	ssi.createSegmentContext(parentSpan)
	if parentSpan == nil || !parentSpan.segmentRegister() {
		rs := newSegmentRoot(ssi)
		rs.createRootSegmentContext(parentSpan)
		s = rs
	} else {
		s = ssi
	}
	return
}

type SegmentContext struct {
	TraceID         []int64
	SpanID          int32
	SegmentID       []int64
	ParentSpanID    int32
	ParentSegmentID []int64
	collect         chan<- ReportedSpan
	refNum          *int32
	spanIDGenerator *int32
}

// Span is accessed by Reporter to load reported data
type ReportedSpan interface {
	Context() *SegmentContext
	Refs() []*propagation.SpanContext
	StartTime() int64
	EndTime() int64
	OperationName() string
	Peer() string
	SpanType() common.SpanType
	SpanLayer() common.SpanLayer
	IsError() bool
	Tags() []*common.KeyStringValuePair
	Logs() []*v2.Log
}

type segmentSpan interface {
	Span
	context() SegmentContext
	segmentRegister() bool
}

type segmentSpanImpl struct {
	defaultSpan
	SegmentContext
}

// For Span

func (s *segmentSpanImpl) End() {
	s.defaultSpan.End()
	go func() {
		s.collect <- s
	}()
}

// For Span

func (s *segmentSpanImpl) Context() *SegmentContext {
	return &s.SegmentContext
}

func (s *segmentSpanImpl) Refs() []*propagation.SpanContext {
	return s.defaultSpan.Refs
}

func (s *segmentSpanImpl) StartTime() int64 {
	return pkg.Millisecond(s.startTime)
}

func (s *segmentSpanImpl) EndTime() int64 {
	return pkg.Millisecond(s.endTime)
}

func (s *segmentSpanImpl) OperationName() string {
	return s.operationName
}

func (s *segmentSpanImpl) Peer() string {
	return s.peer
}

func (s *segmentSpanImpl) SpanType() common.SpanType {
	return common.SpanType(s.spanType)
}

func (s *segmentSpanImpl) SpanLayer() common.SpanLayer {
	return s.layer
}

func (s *segmentSpanImpl) IsError() bool {
	return s.isError
}

func (s *segmentSpanImpl) Tags() []*common.KeyStringValuePair {
	return s.tags
}

func (s *segmentSpanImpl) Logs() []*v2.Log {
	return s.logs
}

func (s *segmentSpanImpl) context() SegmentContext {
	return s.SegmentContext
}

func (s *segmentSpanImpl) segmentRegister() bool {
	for {
		o := atomic.LoadInt32(s.refNum)
		if o < 0 {
			return false
		}
		if atomic.CompareAndSwapInt32(s.refNum, o, o+1) {
			return true
		}
	}
}

func (s *segmentSpanImpl) createSegmentContext(parent segmentSpan) {
	if parent == nil {
		s.SegmentContext = SegmentContext{}
		if len(s.defaultSpan.Refs) > 0 {
			s.TraceID = s.defaultSpan.Refs[0].TraceID
		} else {
			s.TraceID = pkg.GenerateGlobalID()
		}
	} else {
		s.SegmentContext = parent.context()
		s.ParentSegmentID = s.SegmentID
		s.ParentSpanID = s.SpanID
		s.SpanID = atomic.AddInt32(s.spanIDGenerator, 1)
	}
}

type rootSegmentSpan struct {
	*segmentSpanImpl
	notify  <-chan ReportedSpan
	segment []ReportedSpan
	doneCh  chan int32
}

func (rs *rootSegmentSpan) End() {
	rs.defaultSpan.End()
	go func() {
		rs.doneCh <- atomic.SwapInt32(rs.refNum, -1)
	}()
}

func (rs *rootSegmentSpan) createRootSegmentContext(parent segmentSpan) {
	rs.SegmentID = pkg.GenerateScopedGlobalID(int64(rs.tracer.instanceID))
	i := int32(0)
	rs.spanIDGenerator = &i
	rs.SpanID = i
	rs.ParentSpanID = -1
}

func newSegmentRoot(segmentSpan *segmentSpanImpl) *rootSegmentSpan {
	s := &rootSegmentSpan{
		segmentSpanImpl: segmentSpan,
	}
	var init int32
	s.refNum = &init
	ch := make(chan ReportedSpan)
	s.collect = ch
	s.notify = ch
	s.segment = make([]ReportedSpan, 0, 10)
	s.doneCh = make(chan int32)
	go func() {
		total := -1
		defer close(ch)
		defer close(s.doneCh)
		for {
			select {
			case span := <-s.notify:
				s.segment = append(s.segment, span)
			case n := <-s.doneCh:
				total = int(n)
			}
			if total == len(s.segment) {
				break
			}
		}
		s.tracer.reporter.Send(append(s.segment, s))
	}()
	return s
}
