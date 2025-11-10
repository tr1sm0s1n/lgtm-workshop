package main

import (
    "testing"
    "os"
    "io/ioutil"
)

// ...existing code for TestTotal, TestEligibleForFreeShipping, and TestOrderItem...

func TestMain(t *testing.T) {
    // Capture stdout
    oldStdout := os.Stdout
    r, w, _ := os.Pipe()
    os.Stdout = w

    main()

    // Restore stdout
    w.Close()
    out, _ := ioutil.ReadAll(r)
    os.Stdout = oldStdout

    expected := "Total Order Amount: $100.50\nEligible for Free Shipping: true\n"
    if string(out) != expected {
        t.Errorf("Expected output %q, got %q", expected, string(out))
    }
}