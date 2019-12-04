func main() {
    m := map[string]int{
        "key1": 10,
        "key2": 20,
        "key3": 300,
        "key4": 400,
    }

    fmt.Println(m)

    // mapから値を取得
    fmt.Println(m["key2"])
    fmt.Println(m["key3"])
    fmt.Println(m["key99"])
}
