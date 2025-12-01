# **Advent of Code \- Day 1 Solution Documentation**

**Date:** December 1, 2025

**Language:** Go

**Topic:** Modular Arithmetic, Linear Projection

## **Problem Summary**

The puzzle involves a safe dial with numbers 0 through 99\.

* **Start:** The dial points to 50\.  
* **Inputs:** A list of instructions (e.g., L50, R10). L moves toward lower numbers, R moves toward higher numbers.  
* **Part 1 Goal:** Count how many times the dial lands *exactly* on 0 at the end of a rotation instruction.  
* **Part 2 Goal:** Count how many times the dial passes or lands on 0 *during* the rotation (ticks).

## **Core Architectural Decision: Linear Projection**

Instead of keeping the dial variable constrained between 0 and 99 (using modulo % 100 at every step), this solution projects the circular movement onto an infinite linear number line.

* **The Circle:** 0, 1, ... 99, 0, 1 ...  
* **The Line:** ... \-100, ... 0, ... 50, ... 100, ... 200 ...

In this system, any number where x % 100 \== 0 represents the dial pointing at 0\.

* Moving R (Right) increases the absolute value.  
* Moving L (Left) decreases the absolute value.

This simplifies the math significantly because we don't have to handle "wrapping around" logic manually; we just track the total distance traveled from the start.

## **Code & Logic Breakdown**

### **1\. Data Parsing**

We read the file and iterate through lines. We separate the direction char (L or R) from the magnitude.

* If L, we negate the magnitude (num \= \-num).  
* We calculate the sum, which is the projected end position after the move.

### **2\. Part 1 Logic (Variable a)**

**Goal:** Does the rotation end on 0?

sum := dial \+ num  
if sum % 100 \== 0 {  
    a \+= 1  
}

Because we are using an infinite number line, 0, 100, 200, \-100 all represent the dial position 0\. A simple modulo check determines if we landed on the target.

### **3\. Part 2 Logic (Variable b)**

**Goal:** How many times did we click past 0?

This relies on calculating how many "Centuries" (multiples of 100\) we crossed between the dial (start) and sum (end).

#### **The General Formula**

rounds := abs(floorDiv(sum, 100\) \- floorDiv(dial, 100))  
b \+= rounds

* floorDiv(x, 100\) tells us which "100-block" we are in.  
* If we move from 50 (Block 0\) to 150 (Block 1), the difference is 1\. We crossed 100 once.  
* abs ensures we count crossings regardless of direction.

#### **The Edge Case: Moving Left (Negative)**

The simple difference formula has an off-by-one error when moving towards negative numbers due to how Floor handles integer boundaries vs. how the physical dial behaves.

The code implements a correction block when sum \< dial:

if (sum \< dial) {  
    dialC := dial % 100 \== 0 // Did we start exactly on 0?  
    sumC := sum % 100 \== 0   // Did we end exactly on 0?  
      
    // ... setup integers for boolean logic ...

    x := sumI \- dialI   
    b \+= x  
}

**Why is this needed?**

1. **Ending on 0 (sumI):** If we move from 110 to 100, floor(1.1) is 1 and floor(1.0) is 1\. The difference is 0, so the main formula misses that we landed on 0\. We add 1 to correct this.  
2. **Starting on 0 (dialI):** If we move from 100 to 90 (Left 10), floor(1.0) is 1 and floor(0.9) is 0\. The difference is 1\. However, purely starting on 0 and moving away shouldn't count as a new click on 0\. We subtract 1 to correct this over-counting.
