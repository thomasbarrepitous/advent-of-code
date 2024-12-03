import * as fs from 'fs';
import * as path from 'path';

// Parse input file
function parseInput(filename: string): string[] {
    return fs.readFileSync(filename, 'utf-8')
        .trim()
        .split('\n');
}

function parseInstruction(instruction: string): [number, number] {
    const [x, y] = instruction.slice(4, -1).split(',');
    return [parseInt(x), parseInt(y)];
}

function parseInstructionArray(instructions: string[]): [number, number][] {
    return instructions.map(instruction => parseInstruction(instruction));
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
    // Extract all the mul(x,y) pairs with numbers up to 3 digits and do() and don't()
    const mulPairs = input.map(line => line.match(/(?:mul\(\d{1,3},\d{1,3}\)|do\(\)|don't\(\))/g));

    // Remove all instructions between don't() and do()
    let enabled = true;
    const filteredInstructions: string[] = [];
    mulPairs.forEach((line, index) => {
        if (line) {
            line.forEach(instruction => {
                if (instruction === "don't()") {
                    enabled = false;
                } 
                if (instruction === "do()") {
                    enabled = true;
                }
                // Ignore don't() and do()
                if (enabled && instruction !== "don't()" && instruction !== "do()") {
                    filteredInstructions.push(instruction);
                }
            });
        }
    });

    const parsedInstructions = parseInstructionArray(filteredInstructions);

    // Multiply the pairs together
    const mulResults = parsedInstructions.map(([x, y]) => x * y);

    // Sum up all the results
    return mulResults.reduce((acc, curr) => acc + curr, 0);
}

// Main execution
function main() {
    const testInput = parseInput(path.join(__dirname, 'test.txt'));
    const testInput2 = parseInput(path.join(__dirname, 'test2.txt'));
    const input = parseInput(path.join(__dirname, 'input.txt'));

    console.log('Part 1 Test:', solvePart1(testInput));
    console.log('Part 1:', solvePart1(input));

    console.log('Part 2 Test:', solvePart2(testInput2));
    console.log('Part 2:', solvePart2(input));
}

main();
