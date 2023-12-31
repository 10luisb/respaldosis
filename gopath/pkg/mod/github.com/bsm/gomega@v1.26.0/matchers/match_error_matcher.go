package matchers

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/bsm/gomega/format"
)

type MatchErrorMatcher struct {
	Expected interface{}
}

func (matcher *MatchErrorMatcher) Match(actual interface{}) (success bool, err error) {
	if isNil(actual) {
		return false, fmt.Errorf("Expected an error, got nil")
	}

	if !isError(actual) {
		return false, fmt.Errorf("Expected an error.  Got:\n%s", format.Object(actual, 1))
	}

	actualErr := actual.(error)
	expected := matcher.Expected

	if isError(expected) {
		// first try the built-in errors.Is
		if errors.Is(actualErr, expected.(error)) {
			return true, nil
		}
		// if not, try DeepEqual along the error chain
		for unwrapped := actualErr; unwrapped != nil; unwrapped = errors.Unwrap(unwrapped) {
			if reflect.DeepEqual(unwrapped, expected) {
				return true, nil
			}
		}
		return false, nil
	}

	if isString(expected) {
		return actualErr.Error() == expected, nil
	}

	var subMatcher omegaMatcher
	var hasSubMatcher bool
	if expected != nil {
		subMatcher, hasSubMatcher = (expected).(omegaMatcher)
		if hasSubMatcher {
			return subMatcher.Match(actualErr.Error())
		}
	}

	return false, fmt.Errorf(
		"MatchError must be passed an error, a string, or a Matcher that can match on strings. Got:\n%s",
		format.Object(expected, 1))
}

func (matcher *MatchErrorMatcher) FailureMessage(actual interface{}) (message string) {
	return format.Message(actual, "to match error", matcher.Expected)
}

func (matcher *MatchErrorMatcher) NegatedFailureMessage(actual interface{}) (message string) {
	return format.Message(actual, "not to match error", matcher.Expected)
}
