package main

import ()

// func TestPostNoteRoute(t *testing.T) {
// 	ts := httptest.NewServer(setupServer())
// 	defer ts.Close()

// 	values := map[string]string{"title": "Some title", "content": "Some content"}

// 	jsonValue, _ := json.Marshal(values)
// 	resp, err := http.Post(fmt.Sprintf("%s/note", ts.URL), "application/json; charset=utf-8", bytes.NewBuffer(jsonValue))

// 	if err != nil {
// 		t.Fatalf("Expected no error, got %v", err)
// 	}

// 	if resp.StatusCode != 200 {
// 		t.Fatalf("Expected status code 200, got %v", resp.StatusCode)
// 	}
// }
