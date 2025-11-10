// student.js
export class Student {
  constructor(name, marks) {
    this.name = name;
    this.marks = marks;
  }

  total() {
    return this.marks.reduce((sum, mark) => sum + mark, 0);
  }

  passed() {
    return this.marks.every(mark => mark >= 40);
  }
}
