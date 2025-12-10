# Advent of Code - Day 6: Cephalopod Math

## 1. The Problem
We are trapped in a garbage smasher (classic movie trope) and need to help a squid with its math homework to escape. The input is a "worksheet" of arithmetic problems, but the layout is bizarre.

**The Input:**
A visual representation of math problems, arranged side-by-side.
```text
123   328
 45    64
  +     *
````

  * **Part 1 (Human Style):** Read numbers horizontally.
      * Left problem: `123 + 45`.
      * Right problem: `328 * 64`.
  * **Part 2 (Cephalopod Style):** Read numbers **vertically** within columns, and read columns **Right-to-Left**.
      * Left problem becomes:
          * Col 3 (Rightmost): `3`, `5` $\rightarrow$ `35`
          * Col 2: `2`, `4` $\rightarrow$ `24`
          * Col 1: `1` $\rightarrow$ `1`
          * Result: `35 + 24 + 1`

## 2\. The Core Concept: Text as a 2D Grid

Standard string manipulation tools (like `split(" ")`) fail here because the problems are aligned vertically. A space on line 1 doesn't necessarily mean a break in the data if line 2 has a digit there.

To solve this, we treat the text block as a **2D Matrix of Characters** (a `[][]rune`).

1.  **Row ($y$):** The line number.
2.  **Column ($x$):** The character index within the line.

### Step 1: Normalization (Padding)

Input lines might be different lengths. To avoid "Index Out of Range" errors, we first pad every line with spaces so they form a perfect rectangle.

### Step 2: Finding "Rivers"

We need to separate the distinct math problems. We do this by scanning columns from left to right.

  * If an entire column is empty (contains only spaces), it is a **Gap**.
  * Any non-empty columns between Gaps form a **Block**.

## 3\. The Logic

### Part 1 Parsing (Horizontal)

Once we isolate a Block (e.g., the left problem), extracting numbers for Part 1 is simple:

1.  Iterate through the rows.
2.  Parse the string into an integer.
3.  The last row contains the Operator (`+`, `*`).

### Part 2 Parsing (Vertical & Reversed)

This is trickier. We need to "read" down the columns.

1.  **Iterate Columns:** Go through the block from Left to Right ($x=0$ to $x=width$).
2.  **Build Digits:** For each column, scan top-to-bottom. If you see a digit, append it to a string buffer.
      * *Example:* Column has `1` on row 1 and `5` on row 2 $\rightarrow$ String "15" $\rightarrow$ Int `15`.
3.  **Reverse:** The problem says to read columns Right-to-Left. Our loop went Left-to-Right.
      * Action: `slices.Reverse(numbers)`.

## 4\. The Solution Code (Go)

The solution uses a `Task` struct to separate **Data** (the numbers/ops) from **Behavior** (the calculation).

### The Struct

```go
type Task struct {
    Numbers []int
    Op      string
}

// The "Apply" method handles the math. 
// This keeps our parser logic clean of calculation logic.
func (t *Task) Apply() int {
    result := 0
    switch t.Op {
    case "+":
        for _, num := range t.Numbers { result += num }
    case "*":
        result = t.Numbers[0]
        for _, num := range t.Numbers[1:] { result *= num }
    // ... handles -, / as well
    }
    return result
}
```

### The "River" Scanner (`splitVerticalBlocks`)

This function converts the raw text into distinct blocks.

```go
func splitVerticalBlocks(lines []string) [][]string {
    // 1. Convert to a 2D Rune Grid and Pad with spaces
    // ... (grid construction logic) ...

    var blocks [][]string
    startCol := -1

    // 2. Scan every column index (0 to maxWidth)
    for col := 0; col < maxWidth; col++ {
        isGap := true
        
        // Check every row in this specific column
        for row := 0; row < len(grid); row++ {
            if grid[row][col] != ' ' {
                isGap = false // Found a character! Not a gap.
                break
            }
        }

        // State Machine:
        // If we are IN a block and hit a GAP -> Slice the block and save it.
        // If we are NOT in a block and hit a CHAR -> Start a new block.
        if !isGap {
            if startCol == -1 { startCol = col }
        } else {
            if startCol != -1 {
                // We just finished a block. Extract it.
                blocks = append(blocks, extractSlice(grid, startCol, col))
                startCol = -1
            }
        }
    }
    // ... (Handle the final block) ...
    return blocks
}
```

### The Vertical Parser (Part 2)

```go
func extractNumbersFromBlock(block []string) []int {
    // ...
    var numbers []int

    // Iterate Columns
    for col := 0; col < maxWidth; col++ {
        var digitBuffer strings.Builder

        // Iterate Rows
        for _, line := range block {
            if col < len(line) {
                char := rune(line[col])
                if unicode.IsDigit(char) {
                    digitBuffer.WriteRune(char)
                }
            }
        }

        // If this column had digits, parse them into a number
        if digitBuffer.Len() > 0 {
            num, _ := strconv.Atoi(digitBuffer.String())
            numbers = append(numbers, num)
        }
    }

    // Crucial for Part 2: The problem asks for Right-to-Left order
    slices.Reverse(numbers)
    return numbers
}
```

## 5\. Engineering Lesson: Separation of Concerns

Notice how we structured this solution:

1.  **`splitVerticalBlocks`**: Only cares about geometry (splitting the image). It doesn't know what numbers are.
2.  **`Task` Parser**: Only cares about extracting numbers. It doesn't know how to add or multiply.
3.  **`Task.Apply()`**: Only cares about math. It doesn't know where the numbers came from (text, rows, or columns).

This makes the code robust. If Part 3 said "Now divide instead of add," we only change `Apply()`. If the input format changed to be comma-separated, we only change the Parser. We don't have to rewrite the whole program.

