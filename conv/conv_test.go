package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestConvertTemp(t *testing.T) {
	testCases := []struct {
		name         string
		temp         float64
		inUnit       string
		outUnit      string
		expected     float64
		errorMessage string
	}{
		{
