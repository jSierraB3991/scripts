#!/usr/bin/env bash
set -euo pipefail

# ====== COLORES ======
C_RESET="\033[0m"
C_BOLD="\033[1m"
C_DIM="\033[2m"
C_RED="\033[31m"
C_GREEN="\033[32m"
C_YELLOW="\033[33m"
C_BLUE="\033[34m"
C_MAGENTA="\033[35m"
C_CYAN="\033[36m"

log()   { echo -e "${C_BLUE}➜${C_RESET} $1"; }
info()  { echo -e "${C_CYAN}ℹ${C_RESET} $1"; }
ok()    { echo -e "${C_GREEN}✔${C_RESET} $1"; }
warn()  { echo -e "${C_YELLOW}⚠${C_RESET} $1"; }
err()   { echo -e "${C_RED}✖${C_RESET} $1" >&2; }

usage() {
  echo "Uso: $(basename "$0") -u URL [-d CARPETA]"
  exit 1
}

have() { command -v "$1" >/dev/null 2>&1; }

get_default_dir() {
  local kind="$1"
  if have xdg-user-dir; then
    case "$kind" in
      docs) xdg-user-dir DOCUMENTS ;;
      images) xdg-user-dir PICTURES ;;
      videos) xdg-user-dir VIDEOS ;;
      downloads) xdg-user-dir DOWNLOAD ;;
    esac
  else
    case "$kind" in
      docs) printf "%s/Documentos" "$HOME" ;;
      images) printf "%s/Imágenes" "$HOME" ;;
      videos) printf "%s/Vídeos" "$HOME" ;;
      downloads) printf "%s/Descargas" "$HOME" ;;
    esac
  fi
}

strip_archive_ext() {
  local f="$1"
  f="${f%.tar.gz}"; f="${f%.tgz}"
  f="${f%.tar.bz2}"; f="${f%.tbz2}"
  f="${f%.tar.xz}"; f="${f%.txz}"
  f="${f%.zip}"; f="${f%.tar}"
  f="${f%.7z}"; f="${f%.gz}"
  f="${f%.bz2}"; f="${f%.xz}"
  printf '%s' "$f"
}

is_archive_by_name() {
  case "$1" in
    *.zip|*.tar|*.tar.gz|*.tgz|*.tar.bz2|*.tbz2|*.tar.xz|*.txz|*.7z|*.gz|*.bz2|*.xz) return 0 ;;
    *) return 1 ;;
  esac
}

url=""
dest=""

while getopts ":u:d:h" opt; do
  case "$opt" in
    u) url="$OPTARG" ;;
    d) dest="$OPTARG" ;;
    h) usage ;;
    *) usage ;;
  esac
done

[[ -z "${url}" ]] && usage

if ! have curl || ! have file; then
  err "Faltan dependencias: curl y file"
  exit 1
fi

echo -e "${C_BOLD}${C_MAGENTA}=== DOWNX ===${C_RESET}"

log "Validando URL"
if ! curl --head --silent --fail "$url" >/dev/null; then
  err "URL inválida o no accesible"
  exit 1
fi
ok "URL válida"

tmpdir="$(mktemp -d)"
trap 'rm -rf "$tmpdir"' EXIT

fname="${url%%\?*}"
fname="${fname##*/}"
[[ -z "$fname" || "$fname" == "/" ]] && fname="download-$(date +%s)"

tmpfile="$tmpdir/$fname"

log "Descargando"
curl -L --progress-bar --retry 2 --connect-timeout 10 \
  -o "$tmpfile" "$url"
ok "Descarga completada"

size=$(du -h "$tmpfile" | cut -f1)
info "Tamaño: $size"

mime="$(file -b --mime-type "$tmpfile")"
info "Tipo MIME: $mime"

if [[ -n "${dest}" ]]; then
  final_dir="$dest"
  info "Destino (manual): $final_dir"
else
  case "$mime" in
    application/pdf) final_dir="$(get_default_dir docs)" ;;
    image/*)         final_dir="$(get_default_dir images)" ;;
    video/*)         final_dir="$(get_default_dir videos)" ;;
    *)               final_dir="$(get_default_dir downloads)" ;;
  esac
  info "Destino (auto): $final_dir"
fi

mkdir -p "$final_dir"

if is_archive_by_name "$fname"; then
  warn "Archivo comprimido detectado"
  extract_root="$final_dir/$(strip_archive_ext "$fname")"
  mkdir -p "$extract_root"

  log "Extrayendo → $extract_root"

  case "$fname" in
    *.zip) unzip -q "$tmpfile" -d "$extract_root" ;;
    *.tar) tar -xf "$tmpfile" -C "$extract_root" ;;
    *.tar.gz|*.tgz) tar -xzf "$tmpfile" -C "$extract_root" ;;
    *.tar.bz2|*.tbz2) tar -xjf "$tmpfile" -C "$extract_root" ;;
    *.tar.xz|*.txz) tar -xJf "$tmpfile" -C "$extract_root" ;;
    *.7z) 7z x -y -bd "-o$extract_root" "$tmpfile" >/dev/null ;;
    *.gz) gunzip -c "$tmpfile" > "$extract_root/${fname%.gz}" ;;
    *.bz2) bunzip2 -c "$tmpfile" > "$extract_root/${fname%.bz2}" ;;
    *.xz) xz -dc "$tmpfile" > "$extract_root/${fname%.xz}" ;;
  esac

  ok "Extracción completada"
  echo -e "${C_GREEN}${C_BOLD}→ Resultado:${C_RESET} $extract_root"
else
  log "Moviendo archivo"
  mv "$tmpfile" "$final_dir/$fname"
  ok "Archivo guardado"
  echo -e "${C_GREEN}${C_BOLD}→ Resultado:${C_RESET} $final_dir/$fname"
fi
