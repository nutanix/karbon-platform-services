import time
import logging
import random

logging.getLogger().setLevel(logging.DEBUG)

def getDog():
    dog_array = [
        [" (___()'`;", "/,    /`", "\\\\\"--\\\\"],
        ["^..^      /", "/_/\_____/", "   /\   /\\", "  /  \ /  \\"],
        ["  __      _", "o'')}____//", " `_/      )", " (_(_/-(_/"]
    ]

    dog_logo = [
        "     _                 ",
        "  __| | ___   __ _ ___ ",
        " / _` |/ _ \ / _` / __|",
        "| (_| | (_) | (_| \__ \\",
        " \__,_|\___/ \__, |___/",
        "             |___/     "

    ]

    return dog_array, dog_logo

def getCat():
    cat_array = [
        [" /\_/\\", "( o.o )", " > ^ <"],
        ["|\---/|", "| o_o |", " \_^_/"],
        [" /\_/\\", "( o o )", "==_Y_==", "  `-'"]
    ]

    cat_logo = [
        "  ___ __ _| |_ ___ ",
        " / __/ _` | __/ __|",
        "| (_| (_| | |_\__ \\",
        " \___\__,_|\__|___/"
    ]

    return cat_array, cat_logo

animal_array, animal_logo = getDog()
for line in animal_logo:
    logging.info(line)

logging.info("-------------------")
while True:
    animal = random.choice(animal_array)
    for l in animal:
        logging.info(l)
    logging.info("~~~~~~~~")
    time.sleep(3)