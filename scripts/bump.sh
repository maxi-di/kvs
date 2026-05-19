#!/bin/bash
# scripts/bump.sh
set -e

VERSION_FILE="VERSION"
VERSION=$(cat "$VERSION_FILE")

# парсим версию: v1.3.2-rc2 или 1.3.2 или 1.3.2-rc2
VERSION_CLEAN="${VERSION#v}"   # убираем leading v если есть
BASE="${VERSION_CLEAN%%-*}"    # 1.3.2
PRE="${VERSION_CLEAN#"$BASE"}" # -rc2 или пусто
PRE="${PRE#-}"                 # rc2 или пусто

IFS='.' read -r MAJOR MINOR PATCH <<<"$BASE"

usage() {
    echo "Usage: $0 <command>"
    echo ""
    echo "Commands:"
    echo "  major            1.2.3 -> 2.0.0"
    echo "  minor            1.2.3 -> 1.3.0"
    echo "  patch            1.2.3 -> 1.2.4"
    echo "  pre <label>      1.2.3 -> 1.2.3-rc1  (или rc2 если rc уже есть)"
    echo "  pre <label> bump 1.2.3-rc1 -> 1.2.3-rc2"
    echo "  release          1.2.3-rc2 -> 1.2.3  (снять prerelease)"
    echo ""
    echo "Examples:"
    echo "  $0 minor"
    echo "  $0 pre rc"
    echo "  $0 pre rc bump"
    echo "  $0 release"
    exit 1
}

bump_pre() {
    local label="$1" # rc, beta, alpha...

    if [ -z "$PRE" ]; then
        # нет prerelease — начинаем с 1
        echo "${BASE}-${label}1"
    else
        # есть prerelease
        local existing_label existing_num
        # разбиваем rc2 -> rc + 2
        existing_label=$(echo "$PRE" | sed 's/[0-9]*$//')
        existing_num=$(echo "$PRE" | sed 's/[^0-9]*//g')

        if [ "$existing_label" = "$label" ]; then
            # тот же label — инкрементируем номер
            echo "${BASE}-${label}$((existing_num + 1))"
        else
            # другой label — начинаем с 1
            echo "${BASE}-${label}1"
        fi
    fi
}

case "$1" in
major)
    NEW="$((MAJOR + 1)).0.0"
    ;;
minor)
    NEW="${MAJOR}.$((MINOR + 1)).0"
    ;;
patch)
    NEW="${MAJOR}.${MINOR}.$((PATCH + 1))"
    ;;
pre)
    [ -z "$2" ] && usage
    LABEL="$2"
    if [ "$3" = "bump" ]; then
        # bump существующего prerelease
        NEW=$(bump_pre "$LABEL")
    else
        # новый prerelease — сначала поднимаем patch если нет prerelease
        if [ -z "$PRE" ]; then
            BASE="${MAJOR}.${MINOR}.$((PATCH + 1))"
        fi
        NEW="${BASE}-${LABEL}1"
    fi
    ;;
release)
    # снять prerelease suffix
    if [ -z "$PRE" ]; then
        echo "Версия $VERSION уже без prerelease"
        exit 1
    fi
    NEW="$BASE"
    ;;
*)
    usage
    ;;
esac

NEW_FULL="v${NEW}"

echo "  $VERSION -> $NEW_FULL"
read -r -p "Продолжить? [y/N] " confirm
[[ "$confirm" =~ ^[Yy]$ ]] || exit 0

echo "$NEW_FULL" >"$VERSION_FILE"
git add "$VERSION_FILE"
git commit -m "release: bump version to $NEW_FULL"
git tag -a "$NEW_FULL" -m "$NEW_FULL"

echo ""
echo "Готово! Не забудь: git push --follow-tags"
