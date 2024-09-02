import json
import sys
import os

def convert_gocov_to_lcov(gocov_json):
    with open(gocov_json, 'r') as f:
        try:
            data = json.load(f)
        except json.JSONDecodeError:
            print("Error: Invalid JSON format.")
            sys.exit(1)

    lines = []

    if 'Packages' in data:
        for package in data['Packages']:
            for function in package.get('Functions', []):
                filename = function.get('File', '')
                # Make sure to normalize paths if necessary
                if not os.path.isabs(filename):
                    filename = os.path.abspath(filename)
                
                for statement in function.get('Statements', []):
                    start = statement.get('Start', '')
                    end = statement.get('End', '')
                    reached = statement.get('Reached', '')
                    if filename and start and end and reached is not None:
                        lines.append(f'{filename}:{start}:{reached}')

    if not lines:
        print("Error: No coverage data found.")
        sys.exit(1)

    lcov_content = '\n'.join(lines)
    return lcov_content

if __name__ == "__main__":
    if len(sys.argv) != 3:
        print("Usage: python gocov2lcov.py <input_json> <output_lcov>")
        sys.exit(1)

    input_json = sys.argv[1]
    output_lcov = sys.argv[2]

    lcov_content = convert_gocov_to_lcov(input_json)
    with open(output_lcov, 'w') as f:
        f.write(lcov_content)

