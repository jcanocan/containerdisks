package testutil

import (
	"context"
	"hash"
	"os"

	"kubevirt.io/containerdisks/pkg/http"
)

type mockGetter struct {
	mockFile string
}

func (m *mockGetter) GetAll(_ string) ([]byte, error) {
	return os.ReadFile(m.mockFile)
}

func (m *mockGetter) GetAllWithContext(_ context.Context, _ string) ([]byte, error) {
	return os.ReadFile(m.mockFile)
}

func (m *mockGetter) GetWithChecksum(_ string, _ func() hash.Hash) (http.ReadCloserWithChecksum, error) {
	panic("implement me")
}

func (m *mockGetter) GetWithChecksumAndContext(_ context.Context, _ string, _ func() hash.Hash) (http.ReadCloserWithChecksum, error) {
	panic("implement me")
}

func NewMockGetter(mockFile string) *mockGetter {
	return &mockGetter{mockFile: mockFile}
}
