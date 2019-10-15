import argparse
import sys

parser = argparse.ArgumentParser()
parser.add_argument('--input-file', help='path to input file', required=True)
args = parser.parse_args()

current_frequency = 0
frequency_adjustments = []
frequencies_seen = set()

with open(args.input_file) as f:
    for line in f:
        frequency_adjustments.append(int(line))

while True:
    for current_frequency_adjustment in frequency_adjustments:
        frequencies_seen.add(current_frequency)
        current_frequency += current_frequency_adjustment

        if current_frequency in frequencies_seen:
            print(current_frequency)
            sys.exit(0)
