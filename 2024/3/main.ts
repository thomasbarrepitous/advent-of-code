import * as fs from 'fs';
import * as path from 'path';

// Parse input file
function parseInput(filename: string): string[] {
    return fs.readFileSync(filename, 'utf-8')
        .trim()
        .split('\n');
}

// Part 1 solution
function solvePart1(input: string[]): number {
    // Extract all the mul(x,y) pairs with numbers up to 3 digits
    const mulPairs = input.map(line => line.match(/mul\(\d{1,3},\d{1,3}\)/g));
    
    // Extract both x and y values from the pairs
    const mulValues = mulPairs.map(pair => pair ? pair.map(p => {
        const [x, y] = p.slice(4, -1).split(','); // Remove 'mul(' and ')' and split by comma
        return [parseInt(x), parseInt(y)];
    }) : []);

    // Multiply the pairs together
    const mulResults = mulValues.map(pair => pair ? pair.map(([x, y]) => x * y) : []);

    // Sum up all the results
    return mulResults.flat().reduce((acc, curr) => acc + curr, 0);
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

    console.log('Part 1 Test:', solvePart1(testInput));
    console.log('Part 1:', solvePart1(input));

    console.log('Part 2 Test:', solvePart2(testInput));
    console.log('Part 2:', solvePart2(input));
}

main();
