package main

import (
	"math"
	"testing"
	"time"
)

func floatEquals(a, b float64) bool {
	return math.Abs(a-b) < 0.001
}

func readWithTimeout(t *testing.T, ch <-chan PricingResult, timeout time.Duration) PricingResult {
	t.Helper()
	select {
	case res := <-ch:
		return res
	case <-time.After(timeout):
		t.Fatalf("Timed out after %v waiting for PricingResult from channel. Did you send to the channel?", timeout)
		return PricingResult{}
	}
}

func TestCalculatePricing_Math(t *testing.T) {
	ch := make(chan PricingResult, 2)

	// Case 1: Under 2000 (No Surcharge)
	CalculatePricing(JobTicket{TicketID: "T1", BasePrice: 1000.0}, ch)
	res1 := readWithTimeout(t, ch, 200*time.Millisecond)
	if !floatEquals(res1.SetupFee, 150.0) {
		t.Errorf("Expected setup fee 150.0, got %v", res1.SetupFee)
	}
	if !floatEquals(res1.Surcharge, 0.0) {
		t.Errorf("Expected surcharge 0.0 for price <= 2000, got %v", res1.Surcharge)
	}
	if !floatEquals(res1.FinalPrice, 1150.0) {
		t.Errorf("Expected final price 1150.0, got %v", res1.FinalPrice)
	}

	// Case 2: Above 2000 (10% Surcharge applies)
	CalculatePricing(JobTicket{TicketID: "T2", BasePrice: 3000.0}, ch)
	res2 := readWithTimeout(t, ch, 200*time.Millisecond)
	if !floatEquals(res2.SetupFee, 450.0) {
		t.Errorf("Expected setup fee 450.0, got %v", res2.SetupFee)
	}
	if !floatEquals(res2.Surcharge, 300.0) {
		t.Errorf("Expected surcharge 300.0 for price > 2000, got %v", res2.Surcharge)
	}
	if !floatEquals(res2.FinalPrice, 3750.0) {
		t.Errorf("Expected final price 3750.0, got %v", res2.FinalPrice)
	}
}

func TestBatchCalculatePricing_AllReceived(t *testing.T) {
	tickets := []JobTicket{
		{TicketID: "RO26-0001", BasePrice: 1000},
		{TicketID: "RO26-0002", BasePrice: 2500},
		{TicketID: "FL26-0003", BasePrice: 1800},
		{TicketID: "ST26-0004", BasePrice: 4000},
	}

	results := BatchCalculatePricing(tickets)
	if len(results) != len(tickets) {
		t.Fatalf("Expected %d results from channel, got %d. Did BatchCalculatePricing collect all results?", len(tickets), len(results))
	}

	found := make(map[string]PricingResult)
	for _, r := range results {
		found[r.TicketID] = r
	}

	for _, req := range tickets {
		r, ok := found[req.TicketID]
		if !ok {
			t.Errorf("Ticket %s missing from collected channel results", req.TicketID)
			continue
		}
		if !floatEquals(r.BasePrice, req.BasePrice) {
			t.Errorf("BasePrice mismatch for %s: expected %v, got %v", req.TicketID, req.BasePrice, r.BasePrice)
		}
	}
}
