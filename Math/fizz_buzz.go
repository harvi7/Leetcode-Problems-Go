func fizzBuzz(n int) []string {
    result := make([]string, 0);
    
    for i := 1 ; i <= n ; i++{   
        if (multipleOf(i, 3) && multipleOf(i, 5) ) {
            result = append(result, "FizzBuzz")
        } else if (multipleOf(i, 3)) {
            result = append(result, "Fizz")
        } else if (multipleOf(i, 5)) {
            result = append(result, "Buzz")
        } else {
            result = append(result, strconv.Itoa(i) )
        }
    }
    
    return result
}

func multipleOf(i, k int) bool{
    return i % k == 0;
}