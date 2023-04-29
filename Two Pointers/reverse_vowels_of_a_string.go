func reverseWords(s string) string {
    b := []byte(s)
    b = trim(b)
    b = reverse(b)
    b = findWordAndReverse(b)
    return string(b)
}

const SPACE = byte(rune(' '))

func trim(s []byte) []byte {
    for s[0] == SPACE {
        for i := 1; i < len(s); i++ {
            s[i - 1] = s[i]
        }
        s = s[0 : len(s) - 1]
    }
    
    l := 1
    swap := 0
    for i := 1; i < len(s); i++ {        
        if s[i] == SPACE && s[i - 1] == SPACE {
            swap++
        } else {
            if i != l {
                s[l] = s[i]
            }
            l++  
        }
    }
    s = s[0 : len(s) - swap]
    
    l = len(s)-1
    for ; l > 0; l-- {
        if s[l] != SPACE {
            break
        }
    } 

    return s[0 : l + 1]
}

func reverse(s []byte) []byte {
    for i := 0; i < len(s) - i; i++ {
        s[i], s[len(s) - i - 1] = s[len(s) - i - 1], s[i]
    }
    return s
}

func findWordAndReverse(s []byte) []byte {
    prev := 0
    for i := 1; i < len(s) + 1; i++ {
        if  i == len(s) || s[i] == SPACE {
            reverse(s[prev : i])
            prev = i + 1
        }
    }  
    return s
}