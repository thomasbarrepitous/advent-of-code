import os
import sys
import requests
from pathlib import Path
import datetime

def get_session_cookie():
    # Try to get session cookie from environment variable
    session = os.getenv('AOC_SESSION')
    if session:
        return session
        
    # Look for session cookie in .env file
    env_path = Path(__file__).parent.parent / '.env'
    if env_path.exists():
        with open(env_path, 'r') as f:
            for line in f:
                if line.startswith('AOC_SESSION='):
                    return line.strip().split('=')[1]
    
    raise Exception("No session cookie found. Please set AOC_SESSION environment variable or create .env file")

def create_day_folder(year: int, day: int):
    # Get the project root directory (parent of utils)
    root_dir = Path(__file__).parent.parent

    if root_dir.name != 'advent-of-code':
        raise Exception("Root directory is not advent-of-code")
    
    # Create year directory if it doesn't exist
    year_dir = root_dir / str(year)
    year_dir.mkdir(exist_ok=True)
    
    # Create day directory
    day_dir = year_dir / str(day)
    day_dir.mkdir(exist_ok=True)
    
    return day_dir

def create_template_file(day_dir: Path, day: int, year: int):
    template = f'''import * as fs from 'fs';
import * as path from 'path';

// Parse input file
function parseInput(filename: string): string[] {{
    return fs.readFileSync(filename, 'utf-8')
        .trim()
        .split('\\n');
}}

// Part 1 solution
function solvePart1(input: string[]): number {{
    // TODO: Implement part 1 solution
    return 0;
}}

// Part 2 solution
function solvePart2(input: string[]): number {{
    // TODO: Implement part 2 solution
    return 0;
}}

// Main execution
function main() {{
    const testInput = parseInput(path.join(__dirname, 'test.txt'));
    const input = parseInput(path.join(__dirname, 'input.txt'));

    console.log('Part 1 Test:', solvePart1(testInput));
    console.log('Part 1:', solvePart1(input));

    console.log('Part 2 Test:', solvePart2(testInput));
    console.log('Part 2:', solvePart2(input));
}}

main();
'''
    
    with open(day_dir / 'main.ts', 'w') as f:
        f.write(template)

def download_input(year: int, day: int, output_path: Path):
    session = get_session_cookie()
    url = f'https://adventofcode.com/{year}/day/{day}/input'
    
    headers = {
        'Cookie': f'session={session}',
        'User-Agent': 'github.com/thomasbarrepitous/advent-of-code by barrepitousthomas@gmail.com'
    }
    
    response = requests.get(url, headers=headers)
    
    if response.status_code != 200:
        raise Exception(f"Failed to download input: HTTP {response.status_code}")
    
    with open(output_path / 'input.txt', 'w') as f:
        f.write(response.text)

def create_empty_test_file(day_dir: Path):
    with open(day_dir / 'test.txt', 'w') as f:
        f.write('')

def main():
    if len(sys.argv) != 3:
        print("Usage: python new_day.py <day> <year>")
        sys.exit(1)
    
    try:
        day = int(sys.argv[1])
        year = int(sys.argv[2])
        
        # Validate day and year
        if not (1 <= day <= 25):
            raise ValueError("Day must be between 1 and 25")
        
        current_year = datetime.datetime.now().year
        if not (2015 <= year <= current_year):
            raise ValueError(f"Year must be between 2015 and {current_year}")
        
        # Create folder structure
        day_dir = create_day_folder(year, day)
        
        # Download input
        download_input(year, day, day_dir)
        
        # Create template file
        create_template_file(day_dir, day, year)
        
        # Create empty test.txt file
        create_empty_test_file(day_dir)
        
        print(f"Successfully created day {day} folder for year {year}, downloaded input, and created template files")
        
    except ValueError as e:
        print(f"Error: {e}")
        sys.exit(1)
    except Exception as e:
        print(f"Error: {e}")
        sys.exit(1)

if __name__ == "__main__":
    main()
