#!/usr/bin/env python3
import argparse
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
EVENT_FINAL_SCORE = "final_score"
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


def encode(data: _JsonObj) -> str:
    data = data.copy()
    data["content"] = base64.b64encode(_to_json(data["content"]).encode()).decode()
    return f"{_to_json(data)}\n"


def _to_json(data: _JsonObj) -> str:
    return json.dumps(data, separators=(",", ":"))


def iter_decoded(fp: TextIO) -> Iterable[_JsonObj]:
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
# Log V2.1 has been introduced by v0.0.0-20240924143151-8da9e28521f9.
# No further changes have been made to the log as of v0.0.0-20240924143944-1a1d6e31de89.
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


class LogV2ToLogV2_1Transformer(Transformer):
    def __init__(self) -> None:
        self._totals = {}
        self._player_count = 0
        self._last_returned_meeples = {}

    def transform_event(self, data: _JsonObj) -> Iterable[_JsonObj]:
        content = data["content"]
        if data["event"] == EVENT_PLACE_TILE:
            self._last_returned_meeples = {}
        elif data["event"] == EVENT_SCORE:
            scores = content["scores"]
            self._last_returned_meeples = scores["ReturnedMeeples"]
            for player_id, received_points in scores["ReceivedPoints"].items():
                old_total = self._totals.setdefault(str(player_id))
                self._totals[str(player_id)] = old_total + received_points
        elif data["event"] == EVENT_START:
            for player_id in range(1, content["playerCount"] + 1):
                self._totals.setdefault(str(player_id), 0)
        return [data]

    def transform_log(self, decoded_events: Iterable[_JsonObj]) -> Iterable[_JsonObj]:
        yield from super().transform_log(decoded_events)
        yield {
            "event": EVENT_FINAL_SCORE,
            "content": {
                "scores": {
                    "ReceivedPoints": self._totals,
                    "ReturnedMeeples": self._last_returned_meeples,
                }
            }
        }


class App:
    _TRANSFORMS = {
        "add-meeples-everywhere": AddMeeplesEverywhereTransformer,
        "convert-log-v1-to-v2": LogV1ToLogV2Transformer,
        "convert-log-v2-to-v2.1": LogV2ToLogV2_1Transformer,
    }

    def __init__(self) -> None:
        self._args = argparse.Namespace()

    def run(self) -> None:
        self._args = self._parse_args()
        self._args.func()

    def _parse_args(self) -> argparse.Namespace:
        parser = argparse.ArgumentParser(
            description=(
                "A bunch of utilities for working on Carcassonne-Engine log files."
            ),
        )

        subparsers = parser.add_subparsers()
        decode_content = subparsers.add_parser("decode-content")
        self._add_io_arguments(decode_content)
        decode_content.set_defaults(func=self.decode_content)

        encode_content = subparsers.add_parser("encode-content")
        self._add_io_arguments(encode_content)
        encode_content.set_defaults(func=self.encode_content)

        decode_sides = subparsers.add_parser("decode-sides")
        decode_sides.set_defaults(func=self.decode_sides)

        transform = subparsers.add_parser("transform")
        self._add_io_arguments(transform)
        transform.add_argument("transform_name", choices=self._TRANSFORMS.keys())
        transform.set_defaults(func=self.transform)

        return parser.parse_args()

    def _add_io_arguments(self, parser: argparse.ArgumentParser) -> None:
        parser.add_argument(
            "input_file", type=argparse.FileType("r", encoding="utf-8"),
        )
        parser.add_argument(
            "output_file", nargs="?", type=argparse.FileType("w", encoding="utf-8"), default="-"
        )

    def decode_content(self) -> None:
        with self._args.output_file as fp:
            for data in iter_decoded(self._args.input_file):
                fp.write(_to_json(data))
                fp.write("\n")

    def encode_content(self) -> None:
        with (
            self._args.input_file as fp_input,
            self._args.output_file as fp_output,
        ):
            for line in fp_input:
                fp_output.write(encode(json.loads(line)))


    def decode_sides(self) -> None:
        for line in sys.stdin:
            print(Side(int(line.strip())))


    def transform(self) -> None:
        transformer = self._TRANSFORMS[self._args.transform_name]()

        with self._args.output_file as fp:
            transformer.transform_file(iter_decoded(self._args.input_file), fp)


def main() -> None:
    app = App()
    app.run()


if __name__ == "__main__":
    main()
