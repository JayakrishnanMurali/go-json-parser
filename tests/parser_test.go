package tests

import (
	"jsonparser/pkg/lexer"
	"jsonparser/pkg/parser"
	"testing"
)

func TestParseJson(t *testing.T) {
	input := `{
		"name": "John",
		"age": 30,
		"isStudent": false,
		"grades": [90, 80, 70],
		"address": {
			"city": "New York",
			"zip": 10001,
			"test": null,
			"test2": {
				"test3": null,
				"details": {
					"building": "A",
					"floor": 5,
					"contact": {
						"email": "john@example.com",
						"phones": ["123-456-7890", "098-765-4321"]
					}
				}
			}
		},
		"projects": [
			{
				"name": "Project X",
				"completed": true,
				"tasks": [
					{"task": "Design", "status": "completed"},
					{"task": "Development", "status": "in-progress"}
				]
			},
			{
				"name": "Project Y",
				"completed": false,
				"tasks": [
					{"task": "Research", "status": "not started"},
					{"task": "Implementation", "status": "not started"}
				]
			}
		]
	}`

	l := lexer.NewLexer(input)
	p := parser.NewParser(l)

	obj := p.ParseJson()
	if obj == nil {
		t.Fatalf("ParseJson returned nil: %v", p.Errors())
	}

	parsedObj, ok := obj.(*parser.Object)
	if !ok {
		t.Fatalf("Expected Object, got %T", obj)
	}

	if parsedObj.Pairs["name"].(*parser.StringLiteral).Value != "John" {
		t.Errorf("Expected 'John', got %s", parsedObj.Pairs["name"].(*parser.StringLiteral).Value)
	}

	if parsedObj.Pairs["age"].(*parser.IntegerLiteral).Value != 30 {
		t.Errorf("Expected 30, got %d", parsedObj.Pairs["age"].(*parser.IntegerLiteral).Value)
	}

	if parsedObj.Pairs["isStudent"].(*parser.BooleanLiteral).Value != false {
		t.Errorf("Expected false, got %t", parsedObj.Pairs["isStudent"].(*parser.BooleanLiteral).Value)
	}

	grades, ok := parsedObj.Pairs["grades"].(*parser.Array)
	if !ok {
		t.Fatalf("Expected Array for 'grades', got %T", parsedObj.Pairs["grades"])
	}

	if len(grades.Elements) != 3 {
		t.Errorf("Expected 3 elements in grades, got %d", len(grades.Elements))
	}

	expectedGrades := []int64{90, 80, 70}
	for i, grade := range grades.Elements {
		if intGrade, ok := grade.(*parser.IntegerLiteral); ok {
			if intGrade.Value != expectedGrades[i] {
				t.Errorf("Expected grade %d, got %d", expectedGrades[i], intGrade.Value)
			}
		} else {
			t.Fatalf("Expected IntegerLiteral, got %T", grade)
		}
	}

	address, ok := parsedObj.Pairs["address"].(*parser.Object)
	if !ok {
		t.Fatalf("Expected Object for 'address', got %T", parsedObj.Pairs["address"])
	}

	if address.Pairs["city"].(*parser.StringLiteral).Value != "New York" {
		t.Errorf("Expected 'New York', got %s", address.Pairs["city"].(*parser.StringLiteral).Value)
	}

	if address.Pairs["zip"].(*parser.IntegerLiteral).Value != 10001 {
		t.Errorf("Expected 10001, got %d", address.Pairs["zip"].(*parser.IntegerLiteral).Value)
	}

	if _, ok := address.Pairs["test"].(*parser.NullLiteral); !ok {
		t.Errorf("Expected null, got %T", address.Pairs["test"])
	}

	test2, ok := address.Pairs["test2"].(*parser.Object)
	if !ok {
		t.Fatalf("Expected Object for 'test2', got %T", address.Pairs["test2"])
	}

	if _, ok := test2.Pairs["test3"].(*parser.NullLiteral); !ok {
		t.Errorf("Expected null, got %T", test2.Pairs["test3"])
	}

	details, ok := test2.Pairs["details"].(*parser.Object)
	if !ok {
		t.Fatalf("Expected Object for 'details', got %T", test2.Pairs["details"])
	}

	if details.Pairs["building"].(*parser.StringLiteral).Value != "A" {
		t.Errorf("Expected 'A', got %s", details.Pairs["building"].(*parser.StringLiteral).Value)
	}

	if details.Pairs["floor"].(*parser.IntegerLiteral).Value != 5 {
		t.Errorf("Expected 5, got %d", details.Pairs["floor"].(*parser.IntegerLiteral).Value)
	}

	contact, ok := details.Pairs["contact"].(*parser.Object)
	if !ok {
		t.Fatalf("Expected Object for 'contact', got %T", details.Pairs["contact"])
	}

	if contact.Pairs["email"].(*parser.StringLiteral).Value != "john@example.com" {
		t.Errorf("Expected 'john@example.com', got %s", contact.Pairs["email"].(*parser.StringLiteral).Value)
	}

	phones, ok := contact.Pairs["phones"].(*parser.Array)
	if !ok {
		t.Fatalf("Expected Array for 'phones', got %T", contact.Pairs["phones"])
	}

	expectedPhones := []string{"123-456-7890", "098-765-4321"}
	for i, phone := range phones.Elements {
		if phoneStr, ok := phone.(*parser.StringLiteral); ok {
			if phoneStr.Value != expectedPhones[i] {
				t.Errorf("Expected phone %s, got %s", expectedPhones[i], phoneStr.Value)
			}
		} else {
			t.Fatalf("Expected StringLiteral, got %T", phone)
		}
	}

	projects, ok := parsedObj.Pairs["projects"].(*parser.Array)
	if !ok {
		t.Fatalf("Expected Array for 'projects', got %T", parsedObj.Pairs["projects"])
	}

	if len(projects.Elements) != 2 {
		t.Errorf("Expected 2 elements in projects, got %d", len(projects.Elements))
	}

	projectX, ok := projects.Elements[0].(*parser.Object)
	if !ok {
		t.Fatalf("Expected Object for projectX, got %T", projects.Elements[0])
	}

	if projectX.Pairs["name"].(*parser.StringLiteral).Value != "Project X" {
		t.Errorf("Expected 'Project X', got %s", projectX.Pairs["name"].(*parser.StringLiteral).Value)
	}

	if projectX.Pairs["completed"].(*parser.BooleanLiteral).Value != true {
		t.Errorf("Expected true, got %t", projectX.Pairs["completed"].(*parser.BooleanLiteral).Value)
	}

	tasks, ok := projectX.Pairs["tasks"].(*parser.Array)
	if !ok {
		t.Fatalf("Expected Array for 'tasks', got %T", projectX.Pairs["tasks"])
	}

	expectedTasks := []struct {
		task   string
		status string
	}{
		{"Design", "completed"},
		{"Development", "in-progress"},
	}

	for i, taskObj := range tasks.Elements {
		task, ok := taskObj.(*parser.Object)
		if !ok {
			t.Fatalf("Expected Object for task, got %T", taskObj)
		}

		if task.Pairs["task"].(*parser.StringLiteral).Value != expectedTasks[i].task {
			t.Errorf("Expected task %s, got %s", expectedTasks[i].task, task.Pairs["task"].(*parser.StringLiteral).Value)
		}

		if task.Pairs["status"].(*parser.StringLiteral).Value != expectedTasks[i].status {
			t.Errorf("Expected status %s, got %s", expectedTasks[i].status, task.Pairs["status"].(*parser.StringLiteral).Value)
		}
	}

	projectY, ok := projects.Elements[1].(*parser.Object)
	if !ok {
		t.Fatalf("Expected Object for projectY, got %T", projects.Elements[1])
	}

	if projectY.Pairs["name"].(*parser.StringLiteral).Value != "Project Y" {
		t.Errorf("Expected 'Project Y', got %s", projectY.Pairs["name"].(*parser.StringLiteral).Value)
	}

	if projectY.Pairs["completed"].(*parser.BooleanLiteral).Value != false {
		t.Errorf("Expected false, got %t", projectY.Pairs["completed"].(*parser.BooleanLiteral).Value)
	}

	tasksY, ok := projectY.Pairs["tasks"].(*parser.Array)
	if !ok {
		t.Fatalf("Expected Array for 'tasks', got %T", projectY.Pairs["tasks"])
	}

	expectedTasksY := []struct {
		task   string
		status string
	}{
		{"Research", "not started"},
		{"Implementation", "not started"},
	}

	for i, taskObj := range tasksY.Elements {
		task, ok := taskObj.(*parser.Object)
		if !ok {
			t.Fatalf("Expected Object for task, got %T", taskObj)
		}

		if task.Pairs["task"].(*parser.StringLiteral).Value != expectedTasksY[i].task {
			t.Errorf("Expected task %s, got %s", expectedTasksY[i].task, task.Pairs["task"].(*parser.StringLiteral).Value)
		}

		if task.Pairs["status"].(*parser.StringLiteral).Value != expectedTasksY[i].status {
			t.Errorf("Expected status %s, got %s", expectedTasksY[i].status, task.Pairs["status"].(*parser.StringLiteral).Value)
		}
	}
}
