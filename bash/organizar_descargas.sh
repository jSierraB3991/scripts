#!/usr/bin/env bash
set -euo pipefail

# ====== COLORES ======
C_RESET="\033[0m"
C_BOLD="\033[1m"
C_RED="\033[31m"
C_GREEN="\033[32m"
C_YELLOW="\033[33m"
C_BLUE="\033[34m"
C_CYAN="\033[36m"

log()  { echo -e "${C_BLUE}➜${C_RESET} $1"; }
info() { echo -e "${C_CYAN}ℹ${C_RESET} $1"; }
ok()   { echo -e "${C_GREEN}✔${C_RESET} $1"; }
warn() { echo -e "${C_YELLOW}⚠${C_RESET} $1"; }
err()  { echo -e "${C_RED}✖${C_RESET} $1" >&2; }

have() { command -v "$1" >/dev/null 2>&1; }

get_dir() {
  local key="$1"
  if have xdg-user-dir; then
    case "$key" in
      docs) xdg-user-dir DOCUMENTS ;;
      images) xdg-user-dir PICTURES ;;
      videos) xdg-user-dir VIDEOS ;;
      downloads) xdg-user-dir DOWNLOAD ;;
    esac
  else
    case "$key" in
      docs) echo "$HOME/Documentos" ;;
      images) echo "$HOME/Imágenes" ;;
      videos) echo "$HOME/Vídeos" ;;
      downloads) echo "$HOME/Descargas" ;;
    esac
  fi
}

DOWNLOADS="$(get_dir downloads)"
PRE="$DOWNLOADS/pre descargas"

mkdir -p "$PRE"

echo -e "${C_BOLD}=== ORGANIZADOR DE DESCARGAS ===${C_RESET}"
info "Carpeta origen: $DOWNLOADS"

shopt -s nullglob

for file in "$DOWNLOADS"/*; do
  [[ -d "$file" ]] && continue

  name=$(basename "$file")
  mime=$(file -b --mime-type "$file")

  log "Procesando: $name"
  info "Tipo: $mime"

  if [[ "$mime" == application/pdf ||  "$mime" == text/* || \
        "$mime" == application/msword || "$mime" == application/vnd.openxmlformats* ]]; then
    dest="$(get_dir docs)"
  elif [[ "$mime" == image/* ]]; then
    dest="$(get_dir images)"
  elif [[ "$mime" == video/* ]]; then
    dest="$(get_dir videos)"
  elif [[ "$mime" == application/*compressed* || \
        "$mime" == application/zip || \
        "$mime" == application/x-* ]]; then
    dest="$PRE"
    warn "Comprimido → pre descargas"
  else
    dest="$PRE"
    warn "Desconocido → pre descargas"
  fi

  mkdir -p "$dest"

  mv "$file" "$dest/$name"
  ok "Movido → $dest"
done

echo
ok "Organización completada"
