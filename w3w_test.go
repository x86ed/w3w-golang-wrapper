package w3w

import (
	"testing"
)



TestForward(t *testing.T){
	t.Log("Testing the forward API wrapper")
	w := W3W{"","fartFunction", nil, false, "en", "", false}
	w.Forward("index.home.raft")
}

TestReverse(t *testing.T){
	t.Log("Testing the reverse API wrapper")
}

TestAutoSuggest(t *testing.T){
	t.Log("Testing the reverse API wrapper")
}

TestStandardBlend(t *testing.T){
	t.Log("Testing the standard blend API wrapper")
}

TestGrid(t *testing.T){
	t.Log("Testing the grid API wrapper")
}

TestGetLanguages(t *testing.T){
	t.Log("Testing the languages API wrapper")
}