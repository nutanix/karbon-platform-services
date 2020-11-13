from __future__ import absolute_import

import argparse
import logging
import uuid
import sys
import time
logging.basicConfig( stream=sys.stderr, format='%(funcName)s:%(levelname)s:%(message)s', level=logging.DEBUG )


def main():
    root_parser = argparse.ArgumentParser()
    subparsers = root_parser.add_subparsers(help='supported commands')

    root_parser.add_argument("-v", "--version", help="CLI version")

    sd_parser = subparsers.add_parser("servicedomain")
    sd_parser.add_argument("-n", "--name", help="Name of the servicedomain")

    node_parser = subparsers.add_parser("node")

    argument = root_parser.parse_args()

if __name__ == "__main__":
  main()