package v1

import (
	"testing"

	"github.com/32leaves/ruruku/pkg/types"
)

func TestConvertParticipant(t *testing.T) {
	orig := types.Participant{Name: "foo"}

	conv := ConvertParticipant(&orig)
	assertEqual(t, orig.Name, conv.Name)

	convb := conv.Convert()
	assertEqual(t, orig.Name, convb.Name)
}

func TestConvertTestcase(t *testing.T) {
	orig := types.Testcase{
		ID:             "foo",
		Name:           "name",
		Group:          "grp",
		Description:    "desc",
		Steps:          "stps",
		MustPass:       true,
		MinTesterCount: 42,
		Annotations: map[string]string{
			"foo": "bar",
			"baz": "5",
		},
	}

	conv := ConvertTestcase(&orig)
	assertEqual(t, orig.ID, conv.Id)
	assertEqual(t, orig.Name, conv.Name)
	assertEqual(t, orig.Group, conv.Group)
	assertEqual(t, orig.Description, conv.Description)
	assertEqual(t, orig.Steps, conv.Steps)
	assertEqual(t, orig.MustPass, conv.MustPass)
	assertEqual(t, orig.MinTesterCount, conv.MinTesterCount)
	assertMapEqual(t, orig.Annotations, conv.Annotations)

	convb := conv.Convert()
	assertEqual(t, orig.ID, convb.ID)
	assertEqual(t, orig.Name, convb.Name)
	assertEqual(t, orig.Group, convb.Group)
	assertEqual(t, orig.Description, convb.Description)
	assertEqual(t, orig.Steps, convb.Steps)
	assertEqual(t, orig.MustPass, convb.MustPass)
	assertEqual(t, orig.MinTesterCount, convb.MinTesterCount)
	assertMapEqual(t, orig.Annotations, conv.Annotations)
}

func TestConvertTestcaseRunResult(t *testing.T) {
	orig := types.TestcaseRunResult{
		Comment:     "foobar",
		Participant: types.Participant{Name: "foo"},
		State:       types.Undecided,
	}

	conv := ConvertTestcaseRunResult(&orig)
	assertEqual(t, orig.Comment, conv.Comment)
	assertEqual(t, orig.Participant.Name, conv.Participant.Name)
	if conv.State != TestRunState_UNDECIDED {
		t.Errorf("State was not converted. Was \"%v\" but should be undecided", conv.State)
	}

	convb := conv.Convert()
	assertEqual(t, orig.Comment, convb.Comment)
	assertEqual(t, orig.Participant.Name, convb.Participant.Name)
	if convb.State != types.Undecided {
		t.Errorf("State was not converted. Was \"%v\" but should be undecided", convb.State)
	}
}

func TestConvertTestcaseStatus(t *testing.T) {
	t.Skip("Not implemented")
}

func TestConvertTestRunStatus(t *testing.T) {
	t.Skip("Not implemented")
}

func TestConvertTestPlan(t *testing.T) {
	orig := types.TestPlan{
		ID:          "foo",
		Name:        "nme",
		Description: "desc",
		Case: []types.Testcase{
			{},
		},
	}

	conv := ConvertTestPlan(&orig)
	assertEqual(t, orig.ID, conv.Id)
	assertEqual(t, orig.Name, conv.Name)
	assertEqual(t, orig.Description, conv.Description)
	assertEqual(t, len(orig.Case), len(conv.Case))

	convb := conv.Convert()
	assertEqual(t, orig.ID, convb.ID)
	assertEqual(t, orig.Name, convb.Name)
	assertEqual(t, orig.Description, convb.Description)
	assertEqual(t, len(orig.Case), len(convb.Case))
}

func TestConvertTestRunState(t *testing.T) {
	if ConvertTestRunState(types.Passed) != TestRunState_PASSED {
		t.Errorf("ConvertTestRunState does not convert \"passed\"")
	}
	if ConvertTestRunState(types.Undecided) != TestRunState_UNDECIDED {
		t.Errorf("ConvertTestRunState does not convert \"undecided\"")
	}
	if ConvertTestRunState(types.Failed) != TestRunState_FAILED {
		t.Errorf("ConvertTestRunState does not convert \"failed\"")
	}

	if TestRunState_PASSED.Convert() != types.Passed {
		t.Errorf("Converting TestRunState_PASSED yields wrong result")
	}
	if TestRunState_UNDECIDED.Convert() != types.Undecided {
		t.Errorf("Converting TestRunState_UNDECIDED yields wrong result")
	}
	if TestRunState_FAILED.Convert() != types.Failed {
		t.Errorf("Converting TestRunState_FAILED yields wrong result")
	}
}

func TestConvertPermission(t *testing.T) {
	for _, p := range types.AllPermissions {
		r := ConvertPermission(p).Convert()
		if r != p {
			t.Errorf("Converting %v maps to wrong permission %v", p, r)
		}
	}
}

func assertEqual(t *testing.T, expected interface{}, actual interface{}) {
	if expected != actual {
		t.Errorf("Values do not match: (expected) %v != %v (actual)", expected, actual)
	}
}

func assertMapEqual(t *testing.T, expected map[string]string, actual map[string]string) {
	for ekey, eval := range expected {
		if aval, ok := actual[ekey]; !ok {
			t.Errorf("Did not find %s in actual map", ekey)
		} else if eval != aval {
			t.Errorf("Values did not match for %s: expected %v, actual %v", ekey, eval, aval)
		}
	}

	for ekey, aval := range actual {
		if eval, ok := expected[ekey]; !ok {
			t.Errorf("Did not find %s in expected map", ekey)
		} else if eval != aval {
			t.Errorf("Values did not match for %s: expected %v, actual %v", ekey, eval, aval)
		}
	}
}
