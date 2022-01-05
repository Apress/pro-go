package main

func Find(slice []string, vals... string) (matches bool) {
    for _, s1 := range slice {
        for _, s2 := range vals {
            if s1 == s2 {
                matches = true
                return
            }
        }
    }
    return
}
