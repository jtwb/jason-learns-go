package main

func main() {
    x := 120.0
    z := 1.0
    steps := 20 * 1000 * 1000;
    for i := 0; i < steps; i++ {
        z -= (z*z - x) / (2 * x)
    }
    print(z)
}
