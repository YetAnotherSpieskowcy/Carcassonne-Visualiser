#!/usr/bin/env python3
import base64
import enum
import json
import os
import sys
from typing import Any, Iterator


class Side(enum.IntFlag):
    TOP_LEFT_EDGE = 0b1000_0000
    TOP_RIGHT_EDGE = 0b0100_0000
    RIGHT_TOP_EDGE = 0b0010_0000
    RIGHT_BOTTOM_EDGE = 0b0001_0000
    BOTTOM_RIGHT_EDGE = 0b0000_1000
    BOTTOM_LEFT_EDGE = 0b0000_0100
    LEFT_BOTTOM_EDGE = 0b0000_0010
    LEFT_TOP_EDGE = 0b0000_0001


def decode_content() -> None:
    input_file = sys.argv[2]
    output_file = sys.argv[3]

    with open(output_file, "w", encoding="utf-8") as fp:
        for data in iter_decoded(input_file):
            fp.write(json.dumps(data))
            fp.write("\n")


def encode_content() -> None:
    input_file = sys.argv[2]
    output_file = sys.argv[3]

    with (
        open(input_file, "r", encoding="utf-8") as fp_input,
        open(output_file, "w", encoding="utf-8") as fp_output,
    ):
        for line in fp_input:
            fp_output.write(encode(json.loads(line)))


def decode_sides() -> None:
    for line in sys.stdin:
        print(Side(int(line.strip())))


def transform() -> None:
    input_file = sys.argv[2]
    output_file = sys.argv[3]
    transform_func_name = sys.argv[4]

    with open(output_file, "w", encoding="utf-8") as fp:
        for data in iter_decoded(input_file):
            transform_func = TRANSFORMS[transform_func_name]

            if transform_func(data):
                fp.write(encode(data))


def encode(data: dict[str, Any]) -> str:
    data = data.copy()
    data["content"] = base64.b64encode(json.dumps(data["content"]).encode()).decode()
    return f"{json.dumps(data)}\n"


def iter_decoded(path: os.PathLike[str] | str) -> Iterator[dict[str, Any]]:
    with open(path, "r", encoding="utf-8") as fp:
        for line in fp:
            data = json.loads(line)
            data["content"] = json.loads(base64.b64decode(data["content"]))
            yield data


def add_meeples_everywhere_transform(data: dict[str, Any]) -> bool:
    if data["event"] == "score":
        return False
    if data["event"] == "place":
        for idx, feature in enumerate(data["content"]["move"]["Features"], start=1):
            feature["Type"] = 1
            feature["PlayerID"] = idx
    return True


COMMANDS = {
    "decode-content": decode_content,
    "encode-content": encode_content,
    "transform": transform,
    "decode-sides": decode_sides,
}
TRANSFORMS = {
    "add-meeples-everywhere": add_meeples_everywhere_transform,
}


def main() -> None:
    cmd_name = sys.argv[1]
    COMMANDS[cmd_name]()


if __name__ == "__main__":
    main()
