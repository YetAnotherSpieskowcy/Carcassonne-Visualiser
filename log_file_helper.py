#!/usr/bin/env python3
import base64
import enum
import json
import os
import sys
from typing import Any, Iterator

EVENT_START = "start"
EVENT_PLACE_TILE = "place"
EVENT_SCORE = "score"


class Side(enum.IntFlag):
    TOP = 0b1100_0000
    RIGHT = 0b0011_0000
    BOTTOM = 0b0000_1100
    LEFT = 0b0000_0011

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
            fp.write(_to_json(data))
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
    data["content"] = base64.b64encode(_to_json(data["content"]).encode()).decode()
    return f"{_to_json(data)}\n"


def _to_json(data: dict[str, Any]) -> str:
    return json.dumps(data, separators=(",", ":"))


def iter_decoded(path: os.PathLike[str] | str) -> Iterator[dict[str, Any]]:
    with open(path, "r", encoding="utf-8") as fp:
        for line in fp:
            data = json.loads(line)
            data["content"] = json.loads(base64.b64decode(data["content"]))
            yield data


def add_meeples_everywhere_transform(data: dict[str, Any]) -> bool:
    if data["event"] == EVENT_SCORE:
        return False
    if data["event"] == EVENT_PLACE_TILE:
        for idx, feature in enumerate(data["content"]["move"]["Features"], start=1):
            meeple = feature["Meeple"]
            meeple["Type"] = 1
            meeple["PlayerID"] = idx
    return True


# Log V1 was based on v0.0.0-20240902151828-926a89e4df8c.
# Log V2 has been introduced by v0.0.0-20240908171157-2f353db5fffb.
# No further changes have been made to the log as of v0.0.0-20240923073901-5669151e0436.
def convert_log_v1_to_v2(data: dict[str, Any]) -> bool:
    if data["event"] != EVENT_PLACE_TILE:
        return True
    for idx, feature in enumerate(data["content"]["move"]["Features"], start=1):
        feature["Meeple"] = {
            "Type": feature.pop("Type"),
            "PlayerID": feature.pop("PlayerID"),
        }
    return True


COMMANDS = {
    "decode-content": decode_content,
    "encode-content": encode_content,
    "transform": transform,
    "decode-sides": decode_sides,
}
TRANSFORMS = {
    "add-meeples-everywhere": add_meeples_everywhere_transform,
    "convert-log-v1-to-v2": convert_log_v1_to_v2,
}


def main() -> None:
    cmd_name = sys.argv[1]
    COMMANDS[cmd_name]()


if __name__ == "__main__":
    main()
