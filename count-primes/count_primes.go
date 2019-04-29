package count_primes

func countPrimes(n int) int {
    var r int
    o := make([]bool, n)
    
    for i:=2; i<n; i++ {
        if o[i] == false {
            r++
            
            for j:=2; i*j<n; j++ {
                o[j*i] = true
            }
        }
    }
    
    return r
}