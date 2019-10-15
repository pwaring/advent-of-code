import argparse

parser = argparse.ArgumentParser()
parser.add_argument('--input-file', help='path to input file', required=True)
args = parser.parse_args()

frequency_sum = 0

with open(args.input_file) as f:
    for line in f:
        frequency_sum += int(line)

print(frequency_sum)
