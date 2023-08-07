func replaceSpace(s string) string {
    var res string
    for _, c := range s {
        if c == ' ' {
            res += "%20"
        }else{
            res += string(c)
        }
    }
    return res
}