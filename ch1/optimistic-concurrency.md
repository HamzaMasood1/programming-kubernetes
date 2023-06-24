
# Pseudo code to demonstrate optimistic concurrency
```
func main() {
	var err error
	for retries := 0; retries < 10; retries++ {
		foo, err = client.Get("foo", metav1.GetOptions{})
		if err != nil {
			break
		}

		// <update-the-world-and-foo>

		_, err = client.Update(foo)
		if err != nil && errors.IsConflict(err) {
			continue
		} else if err != nil {
			break
		}
	}
}
```