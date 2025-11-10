/**
 * Student Class
 * Defines a student with a map (JavaScript Object) of subject marks
 * and methods to calculate total marks and check passing status.
 */
class Student {
    /**
     * Constructor for the Student class.
     * @param {Object<string, number>} marks - A map where keys are subjects (string) and values are scores (number).
     */
    constructor(marks) {
        // The 'Marks' field from the Go struct becomes a property of the class
        this.Marks = marks || {}; // Initialize with provided marks or an empty object
        this.MIN_PASSING_MARK = 40; // Define the passing threshold
    }

    /**
     * Calculates and returns the cumulative sum of all marks.
     * @returns {number} The total marks.
     */
    Total() {
        let totalMarks = 0;
        // Object.values() gets an array of all the scores
        // We use the reduce method to sum them up
        totalMarks = Object.values(this.Marks).reduce((sum, mark) => sum + mark, 0);
        return totalMarks;
    }

    /**
     * Returns true if ALL marks are 40 or greater. Returns false otherwise.
     * @returns {boolean} True if the student passed all subjects.
     */
    Passed() {
        // We use the every() method which checks if all elements in an array
        // satisfy the condition (mark >= MIN_PASSING_MARK)
        return Object.values(this.Marks).every(mark => mark >= this.MIN_PASSING_MARK);
    }
}

// --- Example Usage ---

// 1. Example of a student who passed all subjects
const passedStudent = new Student({
    "Math": 95,
    "Science": 88,
    "History": 72,
});

// 2. Example of a student who failed one subject
const failedStudent = new Student({
    "Math": 60,
    "Science": 35, // Below 40
    "History": 80,
});

// console.log("--- Passed Student Data ---");
// console.log("Marks:", passedStudent.Marks);
// console.log("Total Marks:", passedStudent.Total());      // Output: 255
// console.log("Passed All Subjects:", passedStudent.Passed()); // Output: true

// console.log("\n--- Failed Student Data ---");
// console.log("Marks:", failedStudent.Marks);
// console.log("Total Marks:", failedStudent.Total());      // Output: 175
// console.log("Passed All Subjects:", failedStudent.Passed()); // Output: false

// // To run this code, save it as 'student.js' and execute: node student.js