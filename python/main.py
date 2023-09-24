#!python3

import argparse
import sys
import traceback
import ckmeans

from trace import Trace


def main():
    if len(sys.argv) < 2:
        usage()
        return -1

    parser = argparse.ArgumentParser(description='ckmeans_1d_dp example')

    parser.add_argument('--data', type=str, help='data filepath')

    parser.add_argument('--debug',
                        action=argparse.BooleanOptionalAction,
                        default=False,
                        help='enables verbose debugging information')

    args = parser.parse_args()
    datafile = args.data
    debug = args.debug

    try:
        data = read(datafile)
        weights = None
        ckmeans.ckmeans_1d_dp(data, weights)
    except Exception as x:
        print()
        print(f'*** ERROR  {x}')
        print()
        if debug:
            print(traceback.format_exc())

        sys.exit(1)


def read(file):
    data = []
    with open(file, 'r', newline='') as f:
        for line in f:
            tokens = line.split()
            for token in tokens:
                try:
                    print(token, float(token))
                    data.append(float(token))
                except ValueError:
                    pass

    return data


def usage():
    print()
    print('  Usage: python3 main.py --data <file>')
    print()
    print()


if __name__ == '__main__':
    main()