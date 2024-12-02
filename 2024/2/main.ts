import * as fs from 'fs';
import * as path from 'path';

// Parse input file
function parseInput(filename: string): string[] {
    return fs.readFileSync(filename, 'utf-8')
        .trim()
        .split('\n');
}

function parseInputAsMatrix(input: string[]): number[][] {
    return input.map(line => 
        line.split(' ')
            .map(Number)
            .filter(n => !isNaN(n))
    );
}

function isRowIncreasing(row: number[]): boolean {
    return row.every((value, index, array) => index === 0 || (value > array[index - 1] && value <= array[index - 1] + 3));
}

function isRowDecreasing(row: number[]): boolean {
    return row.every((value, index, array) => index === 0 || (value < array[index - 1] && value >= array[index - 1] - 3));
}

function isRowSafe(row: number[]): boolean {
    return isRowIncreasing(row) != isRowDecreasing(row);
}

// Part 1 solution
function solvePart1(input: string[]): number {
    // Convert input into a matrix of numbers
    const matrix = parseInputAsMatrix(input);

    // Count the number of safe rows
    const safeRows = matrix.filter(row => isRowSafe(row));

    return safeRows.length;
}

// Part 2 solution
function solvePart2(input: string[]): number {
    // TODO: Implement part 2 solution
    return 0;
}

// Main execution
function main() {
    const testInput = parseInput(path.join(__dirname, 'test.txt'));
    const input = parseInput(path.join(__dirname, 'input.txt'));

    console.log(testInput)

    console.log('Part 1 Test:', solvePart1(testInput));
    console.log('Part 1:', solvePart1(input));

    // console.log('Part 2 Test:', solvePart2(testInput));
    // console.log('Part 2:', solvePart2(input));
}

main();
