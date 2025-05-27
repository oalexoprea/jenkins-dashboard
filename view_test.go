package main

import (
	"testing"
)

func TestExtractFolderName(t *testing.T) {
	testCases := []struct {
		name     string
		url      string
		expected string
	}{
		{
			name:     "typical jenkins job url",
			url:      "https://jenkins.example.com/job/folder/job/my-pipeline/",
			expected: "my-pipeline",
		},
		{
			name:     "url with trailing slash",
			url:      "https://jenkins.example.com/job/another-job/",
			expected: "another-job",
		},
		{
			name:     "url base path",
			url:      "https://jenkins.example.com/",
			expected: "jenkins.example.com",
		},
		{
			name:     "url base path no trailing slash",
			url:      "https://jenkins.example.com",
			expected: "jenkins.example.com",
		},
		{
			name:     "empty string",
			url:      "",
			expected: "",
		},
		{
			name:     "only slashes",
			url:      "///",
			expected: "", // After TrimRight("///", "/"), input to Split is "", Split("", "/") is [""]
		},
		{
			name:     "single folder",
			url:      "folder1/",
			expected: "folder1",
		},
		{
			name:     "job name with hyphens and numbers",
			url:      "https://jenkins.example.com/job/project-alpha-123/",
			expected: "project-alpha-123",
		},
		{
            name:     "root url",
            url:      "/",
            expected: "", // TrimRight("/", "/") is "", Split("", "/") is [""]
        },
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := extractFolderName(tc.url)
			if actual != tc.expected {
				t.Errorf("extractFolderName(%q): expected %q, got %q", tc.url, tc.expected, actual)
			}
		})
	}
}
