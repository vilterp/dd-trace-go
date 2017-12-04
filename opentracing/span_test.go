package opentracing

import (
	"testing"
	"time"

	opentracing "github.com/opentracing/opentracing-go"
	"github.com/stretchr/testify/assert"
)

func TestSpanBaggage(t *testing.T) {
	assert := assert.New(t)

	span := NewSpan("web.request")
	span.SetBaggageItem("key", "value")
	assert.Equal("value", span.BaggageItem("key"))
}

func TestSpanContext(t *testing.T) {
	assert := assert.New(t)

	span := NewSpan("web.request")
	assert.NotNil(span.Context())
}

func TestSpanSetTag(t *testing.T) {
	assert := assert.New(t)

	span := NewSpan("web.request")
	span.Span.Meta = make(map[string]string)
	span.SetTag("component", "tracer")
	assert.Equal("tracer", span.Meta["component"])
}

func TestSpanOperationName(t *testing.T) {
	assert := assert.New(t)

	span := NewSpan("web.request")
	span.SetOperationName("http.request")
	assert.Equal("http.request", span.Span.Name)
}

func TestSpanFinish(t *testing.T) {
	assert := assert.New(t)

	span := NewSpan("web.request")
	span.Finish()

	assert.True(span.Span.Duration > 0)
}

func TestSpanFinishWithTime(t *testing.T) {
	assert := assert.New(t)

	finishTime := time.Now().Add(10 * time.Second)
	span := NewSpan("web.request")
	span.FinishWithOptions(opentracing.FinishOptions{FinishTime: finishTime})

	duration := finishTime.UnixNano() - span.Span.Start
	assert.Equal(duration, span.Span.Duration)
}
