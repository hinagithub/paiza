#!/bin/bash

# 対象のベースディレクトリ
ROOT_DIR="golang"

# 再帰的に .go ファイルを検索（.test.go 除外したい場合は name '*.go' ! -name '*_test.go'）
find "$ROOT_DIR" -type f -name "*.go" | while read -r FILEPATH; do
  DIRNAME=$(dirname "$FILEPATH")
  BASENAME=$(basename "$FILEPATH" .go)
  TARGET_DIR="$DIRNAME/$BASENAME"

  # すでに main.go に変換済みならスキップ
  if [[ "$BASENAME" == "main" ]]; then
    continue
  fi

  mkdir -p "$TARGET_DIR"
  mv "$FILEPATH" "$TARGET_DIR/main.go"
  echo "Moved: $FILEPATH → $TARGET_DIR/main.go"
done

