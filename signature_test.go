package weixin

import "testing"

func TestValidateURL(t *testing.T) {
	token := "0t37dWsIYg6NsVLgEY1fNuB1rSLyyeQE"
	timestamp := "1449648662"
	nonce := "1862651475"
	signature := "717efa7b4910821c7bd59c1b84bbfc363f7551ef"

	ok := ValidateURL(token, timestamp, nonce, signature)

	if !ok {
		t.Fail()
	}
}

func TestSignature(t *testing.T) {
	token := "0t37dWsIYg6NsVLgEY1fNuB1rSLyyeQE"
	timestamp := "1449648662"
	nonce := "1862651475"
	encrypt := "lvMjItfR0rOPRpWGTG3K/b6zEKg4HDKeMU+/HtH6xqZJPpO0fQS8aSVmtornTIowI394/0xSjfxUNT7fdEJvGYpbgU0c2S8P8fQ/+oinc73tEl1hCJSsButo8tPYhjzKzuVITf9OSw4AcS7oo8W8SQBW5ndhj/Cy//kkRm4B82luwpTGHJ8RVcwXriGHVcnW56tYNnmgbGDie2Y0o3vkXX2Gvl7x0iDQpl8QgenMDm4OhvmAL5irMUtPiCFqvB1YM9LCN/f5dbwxMXFdjcI1XJIc6pY6e3t5SC9v96bH+UxgGls5IQuA/ZjNQOFREUp6G3S9A2cvRiNd/jjI72kLbl10KcJRotw1ozkDL8q+azT1OqisNQecsrC/sJ915FlNXbRzSI14RA9HWEOi8XCvphmJlLcoSIYQ/YyC70724tg="
	msgSign := "1f06de4929383d3244fe6981b75dc12ca4928aa2"
	actual := Signature(token, timestamp, nonce, encrypt)

	if actual != msgSign {
		t.Fail()
	}
}
