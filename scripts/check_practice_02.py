#!/usr/bin/env python3
"""Проверка демонстраций и каркаса ПР № 2. Требуются Python 3.9+ и Go.

Запуск из корня репозитория: python scripts/check_practice_02.py
Этот скрипт предназначен для сопровождающего курса, не проверяет работы студентов.
"""

from pathlib import Path
import os
import shutil
import subprocess
import sys
import tempfile

ROOT = Path(__file__).resolve().parents[1]
PRACTICE = ROOT / "practices" / "02-control-flow-functions"
EXPECTED = {
    "conditions": "Попытки остались: да\nСтатус: в работе\n",
    "loops": (
        "Обычный for:\n"
        "1: new\n2: done\n3: in_progress\n"
        "for range:\n"
        "1: new\n2: done\n3: in_progress\n"
        "Осталось: 0\n"
    ),
    "functions": "Максимум: 7\nПоложительных: 3\n",
    "collections": (
        "Массив: 3 элемента\n"
        "Срез: 4 элемента\n"
        "new: 2, done: 2\n"
        "blocked: значение=0, ключ существует=false\n"
    ),
}


def run(args, cwd):
    """Run a command and fail with the complete output on errors."""
    result = subprocess.run(
        args,
        cwd=cwd,
        text=True,
        encoding="utf-8",
        errors="replace",
        stdout=subprocess.PIPE,
        stderr=subprocess.PIPE,
        timeout=120,
    )
    if result.returncode:
        raise RuntimeError(
            f"Command failed ({result.returncode}): {' '.join(map(str, args))}\n"
            f"{result.stdout}{result.stderr}"
        )
    return result.stdout


def main():
    if not shutil.which("go") or not shutil.which("gofmt"):
        raise RuntimeError("Go and gofmt must be installed and available in PATH.")

    print(run(["go", "version"], ROOT).strip())
    for folder in ("examples", "starter"):
        path = PRACTICE / folder
        sources = sorted(str(source) for source in path.rglob("*.go"))
        if not sources:
            raise RuntimeError(f"No Go source files in {path}")
        unformatted = run(["gofmt", "-l", *sources], path).strip()
        if unformatted:
            raise RuntimeError(f"Unformatted Go source files:\n{unformatted}")

        run(["go", "build", "./..."], path)
        run(["go", "vet", "./..."], path)
        print(f"PASS: {folder} formatting, build, vet")

    with tempfile.TemporaryDirectory(prefix="elite-practice-02-") as temporary:
        for name, expected in EXPECTED.items():
            executable = Path(temporary) / (name + (".exe" if os.name == "nt" else ""))
            run(
                ["go", "build", "-o", str(executable), f"./{name}"],
                PRACTICE / "examples",
            )
            actual = run([str(executable)], ROOT)
            if actual.replace("\r\n", "\n") != expected:
                raise AssertionError(
                    f"{name}: output differs.\nExpected: {expected!r}\nActual: {actual!r}"
                )
            print(f"PASS: {name} scenario")

    print("PASS: practice 02 examples and starter verified.")


if __name__ == "__main__":
    try:
        main()
    except (OSError, RuntimeError, AssertionError, subprocess.TimeoutExpired) as error:
        print(f"FAIL: {error}", file=sys.stderr)
        sys.exit(1)
