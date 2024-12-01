const { readFileSync } = require('fs');

// Types
type Input = string[];

// Parse input file
function parseInput(filename: string): Input {
    return readFileSync(filename, 'utf-8')
        .trim()
        .split('\n');
}

// Part 1 solution
function solvePart1(input: Input): number {
    // TODO: Implement part 1 solution
    return 0;
}

// Part 2 solution
function solvePart2(input: Input): number {
    // TODO: Implement part 2 solution
    return 0;
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
