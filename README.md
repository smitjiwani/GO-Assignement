# Go Programming - List Manipulation with API

#### Objective:
Write a Go program that maintains a list of integers and exposes an API to manipulate the list based on the sign of an input number. The program should append the number to the list if its sign matches the sign of the integers in the list. If the sign is opposite, the program should keep reducing the quantity from the list starting from the first element in the list (FIFO basis) till the input number is exhausted and update the list. (Example given below)

#### Requirements:
1. **List Initialization**: Start with an empty list of integers.
2. **API Endpoint**: Expose an API endpoint that accepts a number as input.
3. **Sign Matching**:
   - If the sign of the input number matches the sign of the integers in the list, append the number to the list.
   - If the sign is opposite, remove the value of the input number starting from the oldest element in the list and update the list.
4. **Example**: Provide an example of how the API works with a sequence of inputs.

#### Example:
Consider the following sequence of inputs:
- Input: `5`  
  Updated List: `[5]`  

- Input: `10`  
  Updated List: `[5, 10]`  

- Input: `-6`  
  Updated List: `[9]`  
  (Sign is opposite, remove `6` quantity from the list starting from the oldest element `5` → `5 - 5 = 0` and `10` → `10 - 1 = 9`)


#### Deliverables:
1. **Go Code**: Implement the logic in Go, including the API endpoint. Use `gin` router.
2. **Example Output**: Provide the output of the program for the example sequence of inputs.
3. **Explanation**: Briefly explain how the code works and how the list is manipulated based on the input.
4. **Edge Case**: Think and solve for any edge cases. Mention the edge cases in the explaination README file.
#### Submission:
- Submit your Go code file(s) along with a README file that includes instructions on how to run the code and the example output.

#### Bonus:
- Add unit tests to verify the functionality of your code.
- Add logs.

---
