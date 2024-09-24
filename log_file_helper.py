#!/usr/bin/env python3
import base64
import enum
import itertools
import json
import os
import sys
from typing import Any, Iterable, TextIO

EVENT_START = "start"
EVENT_PLACE_TILE = "place"
EVENT_SCORE = "score"
_JsonObj = dict[str, Any]


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
    transform_name = sys.argv[4]
    transformer = TRANSFORMS[transform_name]()

    with open(output_file, "w", encoding="utf-8") as fp:
        transformer.transform_file(iter_decoded(input_file), fp)


def encode(data: _JsonObj) -> str:
    data = data.copy()
    data["content"] = base64.b64encode(_to_json(data["content"]).encode()).decode()
    return f"{_to_json(data)}\n"


def _to_json(data: _JsonObj) -> str:
    return json.dumps(data, separators=(",", ":"))


def iter_decoded(path: os.PathLike[str] | str) -> Iterable[_JsonObj]:
    with open(path, "r", encoding="utf-8") as fp:
        for line in fp:
            data = json.loads(line)
            data["content"] = json.loads(base64.b64decode(data["content"]))
            yield data


class Transformer:
    def transform_event(self, data: _JsonObj) -> Iterable[_JsonObj]:
        return [data]

    def transform_log(self, decoded_events: Iterable[_JsonObj]) -> Iterable[_JsonObj]:
        for data in decoded_events:
            yield from self.transform_event(data)

    def transform_file(self, decoded_events: Iterable[_JsonObj], fp: TextIO) -> None:
        for transformed in self.transform_log(decoded_events):
            fp.write(encode(transformed))


class AddMeeplesEverywhereTransformer(Transformer):
    def transform_event(self, data: _JsonObj) -> Iterable[_JsonObj]:
        if data["event"] == EVENT_SCORE:
            return []
        if data["event"] == EVENT_PLACE_TILE:
            for idx, feature in enumerate(data["content"]["move"]["Features"], start=1):
                meeple = feature["Meeple"]
                meeple["Type"] = 1
                meeple["PlayerID"] = idx
        return [data]


# Log V1 was based on v0.0.0-20240902151828-926a89e4df8c.
# Log V2 has been introduced by v0.0.0-20240908171157-2f353db5fffb.
# No further changes have been made to the log as of v0.0.0-20240923073901-5669151e0436.
class LogV1ToLogV2Transformer(Transformer):
    def transform_event(self, data: _JsonObj) -> Iterable[_JsonObj]:
        if data["event"] != EVENT_PLACE_TILE:
            return [data]
        for idx, feature in enumerate(data["content"]["move"]["Features"], start=1):
            feature["Meeple"] = {
                "Type": feature.pop("Type"),
                "PlayerID": feature.pop("PlayerID"),
            }
        return [data]


COMMANDS = {
    "decode-content": decode_content,
    "encode-content": encode_content,
    "transform": transform,
    "decode-sides": decode_sides,
}
TRANSFORMS = {
    "add-meeples-everywhere": AddMeeplesEverywhereTransformer,
    "convert-log-v1-to-v2": LogV1ToLogV2Transformer,
}


def main() -> None:
    cmd_name = sys.argv[1]
    COMMANDS[cmd_name]()


if __name__ == "__main__":
    main()
