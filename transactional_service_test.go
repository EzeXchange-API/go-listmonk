package listmonk

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"
)

type transactionalServiceTestSuite struct {
	suite.Suite
}

func TestTransactionalService(t *testing.T) {
	suite.Run(t, new(transactionalServiceTestSuite))
}

func (s *transactionalServiceTestSuite) TestPostTransactionalService() {
	mockResponse := []byte(`{
		"data": true,
	}`)

	mockClient := newMockClient(mockResponse)
	service := &PostTransactionalService{
		c: mockClient,
		// set other fields if needed
	}

	service.TemplateId(1)
	service.SubscriberIds([]uint{1, 2})
	service.SubscriberEmails([]string{"test2", "test2"})
	service.Data(map[string]string{"key1": "value1", "key2": "value2"})
	service.Headers(map[string]string{"key1": "value1", "key2": "value2"})

	err := service.Do(context.Background())

	s.Nil(err)

	s.Equal(mockClient.Calls[0].Arguments.Get(1).(*request).method, "POST")
	s.Equal(mockClient.Calls[0].Arguments.Get(1).(*request).json["template_id"], uint(1))
	s.Equal(mockClient.Calls[0].Arguments.Get(1).(*request).json["subscriber_ids"], "[1, 2]")
	s.Equal(mockClient.Calls[0].Arguments.Get(1).(*request).json["subscriber_emails"], "[\"test2\", \"test2\"]")
	s.Equal(mockClient.Calls[0].Arguments.Get(1).(*request).json["data"], "{\"key1\":\"value1\",\"key2\":\"value2\"}")
	s.Equal(mockClient.Calls[0].Arguments.Get(1).(*request).json["headers"], "[{\"key1\": \"value1\"}, {\"key2\": \"value2\"}]")
}
