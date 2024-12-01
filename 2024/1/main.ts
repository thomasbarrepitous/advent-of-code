const { readFileSync } = require('fs');

// Types
type Input = string[];

// Parse input file
function parseInput(filename: string): Input {
    return readFileSync(filename, 'utf-8')
        .trim()
        .split('\n');
}

function parseInputAsNumbers(input: Input): [number[], number[]] {
    const column1: number[] = [];
    const column2: number[] = [];

    input.forEach(line => {
        const [num1, num2] = line.trim().split(/\s+/).map(Number);
        column1.push(num1);
        column2.push(num2);
    });

    return [column1, column2];
}

function findDistance(column1: number[], column2: number[]): number {
    let distance = 0;
    for (let i = 0; i < column1.length; i++) {
        distance += Math.abs(column1[i] - column2[i]);
    }
    return distance;
}

function findSimilarity(column1: number[], column2: number[]): number {
    // The similarity is the number of occurences of the left column number in the right column
    let similarity = 0;
    for (let i = 0; i < column1.length; i++) {
        similarity += column1[i] * column2.filter(num => num === column1[i]).length;
        console.log(column1[i], similarity);
    }
    return similarity;
}

// Part 1 solution
function solvePart1(input: Input): number {
    const [column1, column2] = parseInputAsNumbers(input);
    // Sort both columns
    column1.sort((a, b) => a - b);
    column2.sort((a, b) => a - b);

    // Find the distance between the two columns
    return findDistance(column1, column2);
}

// Part 2 solution
function solvePart2(input: Input): number {
    const [column1, column2] = parseInputAsNumbers(input);

    // Find the similarity between the two columns
    return findSimilarity(column1, column2);
}

// Main execution
function main() {
    const testInput = parseInput('./test.txt');
    const input = parseInput('./input.txt');

    console.log('Part 1 Test:', solvePart1(testInput));
    console.log('Part 1:', solvePart1(input));

    console.log('Part 2 Test:', solvePart2(testInput));
    console.log('Part 2:', solvePart2(input));
}

main();
