package request

var (
	url = "https://jsonplaceholder.typicode.com/"
	rc  = New("", "")
)

/*

func TestFetchPosts(t *testing.T) {
	result, err := rc.Fetch("", Query{})
	if err != nil {
		t.Fatalf(err.Error())
	}
	if result != expected {
		t.Fatalf(fmt.Sprintf("result %s does not match the expected value %s", result, expected))
	}
}

/*

func TestFetchPost(t *testing.T) {
	result, err := rc.Fetch("", Query{})
	if err != nil {
		t.Fatalf(err.Error())
	}
	if result != expected {
		t.Fatalf(fmt.Sprintf("result %s does not match the expected value %s", result, expected))
	}
}
*/
