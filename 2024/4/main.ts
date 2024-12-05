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
    let total = 0;
    const directions = [
        [-1, 0],  [1, 0],   // vertical
        [0, -1],  [0, 1],   // horizontal
        [-1, -1], [-1, 1],  // diagonal up
        [1, -1],  [1, 1]    // diagonal down
    ];

    const isInBounds = (y: number, x: number) => 
        y >= 0 && y < input.length && x >= 0 && x < input[0].length;

    // For each row
    for (let row = 0; row < input.length; row++) {
        // For each column
        for (let col = 0; col < input[row].length; col++) {
            // Check each 'X'
            if (input[row][col] === 'X') {
                // Check each direction
                directions.forEach(([dy, dx]) => {
                    let y = row, x = col;
                    const sequence = ['M', 'A', 'S'];
                    // Check if the sequence 'M', 'A', 'S' is in the direction
                    if (sequence.every(char => {
                        y += dy;
                        x += dx;
                        // Check if the position is in bounds and the character is the expected one
                        return isInBounds(y, x) && input[y][x] === char;
                    })) {
                        // If so, increment the total
                        total++;
                    }
                });
            }
        }
    }
    return total;
}

// Part 2 solution
function solvePart2(input: string[]): number {
    let total = 0;
    const crossesTypes = [
        // Possible Cross 1, same letters horizontally
        // M M
        //  A
        // S S
        [[[-1, -1], [-1, 1]], // Diagonal up
        [[1, -1], [1, 1]]],   // Diagonal down
        // Possible Cross 2, same letters vertically
        // S M
        //  A
        // S M
        [[[-1, -1], [1, -1]], // Diagonal left
        [[-1, 1], [1, 1]]]   // Diagonal right
    ];

    const isInBounds = (y: number, x: number) => 
        y >= 0 && y < input.length && x >= 0 && x < input[0].length;

    // For each row
    for (let row = 0; row < input.length; row++) {
        // For each column
        for (let col = 0; col < input[row].length; col++) {
            // Check each 'A' as center of cross
            if (input[row][col] === 'A') {
                // For the two possible types of crosses
                crossesTypes.forEach(cross => {
                    // Check both possibilities for each cross type
                    const isValidCross = (
                        // Check M's in first position, S's in second position
                        (cross[0].every(([dy, dx]) => {
                            const y = row + dy, x = col + dx;
                            return isInBounds(y, x) && input[y][x] === 'M';
                        }) &&
                        cross[1].every(([dy, dx]) => {
                            const y = row + dy, x = col + dx;
                            return isInBounds(y, x) && input[y][x] === 'S';
                        })) ||
                        // Check S's in first position, M's in second position
                        (cross[0].every(([dy, dx]) => {
                            const y = row + dy, x = col + dx;
                            return isInBounds(y, x) && input[y][x] === 'S';
                        }) &&
                        cross[1].every(([dy, dx]) => {
                            const y = row + dy, x = col + dx;
                            return isInBounds(y, x) && input[y][x] === 'M';
                        }))
                    );

                    if (isValidCross) {
                        total++;
                    }
                });
            }
        }
    }
    return total;
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
