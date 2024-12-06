import * as fs from 'fs';
import * as path from 'path';

// Parse input file
function parseInput(filename: string): string[] {
    return fs.readFileSync(filename, 'utf-8')
        .trim()
        .split('\n');
}

function parseRules(input: string[]): { [key: string]: string[] } {
    const rules: { [key: string]: string[] } = {};
    input.slice(0, input.indexOf('')).forEach(line => {
        const [left, right] = line.split('|');
        if (!rules[left]) {
            rules[left] = [];
        }
        rules[left].push(right);
    });
    return rules;
}

function parseUpdates(input: string[]): string[][] {
    return input.slice(input.indexOf('') + 1).map(line => {
        return line.split(',');
    });
}

function isViolated(rules: { [key: string]: string[] }, update: string[]): boolean {
    // Check if the previous pages violated the rules
    for (const page of update) {
        // Skip if there's no rule for this page
        if (!rules[page]) continue;
        
        // Check if the previous pages violated the rules
        const previousPages = update.slice(0, update.indexOf(page));
        if (previousPages.some(prevPage => rules[page].includes(prevPage))) {
            return true;
        }
    }
    return false;
}

function orderUpdate(rules: { [key: string]: string[] }, update: string[]): string[] {
    while (isViolated(rules, update)) {
        // Start from the end and try to move each page left
        for (let i = update.length - 1; i >= 0; i--) {
            const currentPage = update[i];
            const previousPage = update[i - 1];

        // If the current page violates the rules, swap it with the previous page
        if (rules[currentPage] && rules[currentPage].includes(previousPage)) {
            update[i] = previousPage;
            update[i - 1] = currentPage;
            }
        }
    }
    return update;
}


// Part 1 solution
function solvePart1(input: string[]): number {
    // Split the input into two parts:
    // 1. The rules
    // 2. The updates
    const rules = parseRules(input);
    const updates = parseUpdates(input);

    let middlePageSum = 0;

    // For each update
    updates.forEach(update => {
        // If the update violates the rules, skip it
        if (isViolated(rules, update)) {
            return;
        }
        // Add middle page to sum for valid updates
        middlePageSum += Number(update[Math.floor(update.length / 2)]);
    });
    return middlePageSum;
}

// Part 2 solution
function solvePart2(input: string[]): number {
    // Split the input into rules and updates
    const rules = parseRules(input);
    const updates = parseUpdates(input);

    let invalidInput: string[][] = [];

    // For each update
    updates.forEach(update => {
        // Fetch all inccorect ordereded updates
        if (isViolated(rules, update)) {
            invalidInput.push(update);
        }
    });

    // Order each invalid update
    invalidInput.forEach(update => {
        update = orderUpdate(rules, update);
    });

    // Find the middle page of each ordered update
    let middlePageSum = 0;
    invalidInput.forEach(update => {
        middlePageSum += Number(update[Math.floor(update.length / 2)]);
    });

    return middlePageSum;
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
