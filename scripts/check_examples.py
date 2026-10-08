#!/usr/bin/env python3
"""Проверка примеров и сборки заготовок. Требуются Python 3.9+ и Go.

Запуск из корня курса: python scripts/check_examples.py
Скрипт не проверяет решения студентов и не требует внешних Python-пакетов.
"""
from pathlib import Path
import os
import shutil
import subprocess
import sys
import tempfile

ROOT = Path(__file__).resolve().parents[1]
PRACTICE = ROOT / "practices" / "01-introduction"


def command(args, cwd, *, stdin=None):
    """Запустить команду с ограничением времени и обработать ошибку."""
    result = subprocess.run(
        args, cwd=cwd, input=stdin, text=True, encoding="utf-8",
        errors="replace", stdout=subprocess.PIPE, stderr=subprocess.PIPE,
        timeout=90,
    )
    if result.returncode:
        raise RuntimeError(
            f"Command failed ({result.returncode}): {' '.join(map(str, args))}\n"
            f"{result.stdout}{result.stderr}"
        )
    return result.stdout


def main():
    if not shutil.which("go") or not shutil.which("gofmt"):
        raise RuntimeError("Go and gofmt must be available in PATH.")

    print(command(["go", "version"], ROOT).strip())
    for module in ("examples", "starter"):
        path = PRACTICE / module
        files = sorted(str(p) for p in path.rglob("*.go"))
        unformatted = command(["gofmt", "-l", *files], path).strip()
        if unformatted:
            raise RuntimeError(f"Run gofmt for:\n{unformatted}")
        command(["go", "build", "./..."], path)
        command(["go", "vet", "./..."], path)
        print(f"PASS: {module} formatting, build, vet")

    with tempfile.TemporaryDirectory(prefix="elite-go-check-") as temporary:
        binaries = {}
        for package in ("hello", "variables", "input"):
            suffix = ".exe" if os.name == "nt" else ""
            binary = Path(temporary) / (package + suffix)
            command(["go", "build", "-o", str(binary), f"./{package}"], PRACTICE / "examples")
            binaries[package] = binary

        scenarios = [
            ("hello", "", "Привет! Это моя первая программа на Go.\n"),
            ("variables", "", "Дробное деление: 3.40\n"),
            ("input", "12\n", "Будет обработано страниц: 12\n"),
            ("input", "0\n", "Будет обработано страниц: 0\n"),
            ("input", "-1\n", "Ошибка: число не должно быть отрицательным.\n"),
            ("input", "abc\n", "Ошибка: нужно ввести целое число.\n"),
            ("input", "", "Ошибка: нужно ввести целое число.\n"),
        ]
        for number, (package, stdin, expected) in enumerate(scenarios, 1):
            output = command([str(binaries[package])], ROOT, stdin=stdin)
            if not output.endswith(expected):
                raise AssertionError(
                    f"Unexpected output for {package}, input={stdin!r}: {output!r}"
                )
            if package == "input" and stdin in ("-1\n", "abc\n", ""):
                if "Будет обработано страниц" in output:
                    raise AssertionError("Invalid input was used for further processing.")
            print(f"PASS: scenario {number} ({package})")

    print("All example checks passed. Student solutions are not included in this check.")


if __name__ == "__main__":
    try:
        main()
    except (OSError, RuntimeError, AssertionError, subprocess.TimeoutExpired) as exc:
        print(f"FAIL: {exc}", file=sys.stderr)
        sys.exit(1)
