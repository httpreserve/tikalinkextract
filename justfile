# CLI helpers.

# Help
help:
    @just -l

# Run all pre-commit checks
all-checks:
   pre-commit run --all-files
