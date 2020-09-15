package pubsubwrapper

import (
	"context"
	"time"

	"github.com/stretchr/testify/mock"
)

// Client mock
type ClientMock struct {
	mock.Mock
}

func NewMockedClient() Client {
	return &ClientMock{}
}

func (c *ClientMock) CreateTopic(ctx context.Context, topicID string) (Topic, error) {
	args := c.Called(ctx, topicID)
	return args.Get(0).(Topic), args.Error(1)
}

func (c *ClientMock) Topic(id string) Topic {
	args := c.Called(id)
	return args.Get(0).(Topic)
}

func (c *ClientMock) Topics(ctx context.Context) ([]Topic, error) {
	args := c.Called(ctx)
	return args.Get(0).([]Topic), args.Error(1)
}

func (c *ClientMock) CreateSubscription(ctx context.Context, id string, cfg SubscriptionConfig) (Subscription, error) {
	args := c.Called(ctx, id, cfg)
	return args.Get(0).(Subscription), args.Error(1)
}

func (c *ClientMock) Subscription(id string) Subscription {
	args := c.Called(id)
	return args.Get(0).(Subscription)
}

func (c *ClientMock) Subscriptions(ctx context.Context) ([]Subscription, error) {
	args := c.Called(ctx)
	return args.Get(0).([]Subscription), args.Error(1)
}

func (c *ClientMock) embedToIncludeNewMethods() {}

// Topic mock
type TopicMock struct {
	mock.Mock
}

func NewMockedTopic() Topic {
	return &TopicMock{}
}

func (t *TopicMock) String() string {
	args := t.Called()
	return args.String(0)
}

func (t *TopicMock) Publish(ctx context.Context, msg Message) PublishResult {
	args := t.Called(ctx, msg)
	return args.Get(0).(PublishResult)
}

func (t *TopicMock) Exists(ctx context.Context) (bool, error) {
	args := t.Called(ctx)
	return args.Bool(0), args.Error(1)
}

func (t *TopicMock) Delete(ctx context.Context) error {
	args := t.Called(ctx)
	return args.Error(0)
}

func (t *TopicMock) embedToIncludeNewMethods() {
}

// subscription
type SubscriptionMock struct {
	mock.Mock
	MessageMock Message
}

func NewMockedSubscription() Subscription {
	return &SubscriptionMock{}
}

func (s *SubscriptionMock) String() string {
	return s.Called().String(0)
}

func (s *SubscriptionMock) Exists(ctx context.Context) (bool, error) {
	args := s.Called(ctx)
	return args.Bool(0), args.Error(1)
}

func (s *SubscriptionMock) Receive(ctx context.Context, f func(context.Context, Message)) error {
	args := s.Called(ctx, f)
	err := args.Error(0)
	if err == nil {
		f(ctx, s.MessageMock) // apply function
	}
	return err
}

func (s *SubscriptionMock) Delete(ctx context.Context) error {
	args := s.Called(ctx)
	return args.Error(0)
}

func (s *SubscriptionMock) embedToIncludeNewMethods() {}

// message
type MessageMock struct {
	mock.Mock
}

func NewMockedMessage() Message {
	return &MessageMock{}
}

func (m *MessageMock) ID() string {
	return m.Called().String(0)
}

func (m *MessageMock) Data() []byte {
	return m.Called().Get(0).([]byte)
}

func (m *MessageMock) Attributes() map[string]string {
	return m.Called().Get(0).(map[string]string)
}

func (m *MessageMock) PublishTime() time.Time {
	return m.Called().Get(0).(time.Time)
}

func (m *MessageMock) Ack() {
	m.Called()
}

func (m *MessageMock) Nack() {
	m.Called()
}

func (m *MessageMock) embedToIncludeNewMethods() {}

// publish result
type PublishResultMock struct {
	mock.Mock
}

func NewMockedPublishResult() PublishResult {
	return &PublishResultMock{}
}

func (p *PublishResultMock) Get(ctx context.Context) (serverID string, err error) {
	args := p.Called()
	return args.String(0), args.Error(1)
}

func (p *PublishResultMock) embedToIncludeNewMethods() {}
