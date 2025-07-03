#!/bin/bash
set -euo pipefail

DEST_DB="xano_backup_migration"
DUMP_DIR="./db_backup"
EXPORT_DIR="./json_exports"

echo "🔍 Verificando existencia de la base de datos '$DEST_DB'..."

if sudo -u postgres psql -lqt | cut -d \| -f 1 | grep -qw "$DEST_DB"; then
  echo "⚠️  La base '$DEST_DB' ya existe. Eliminando todas las tablas..."

  sudo -u postgres psql "$DEST_DB" <<EOF
DO
\$do\$
DECLARE
    r RECORD;
BEGIN
    EXECUTE 'SET session_replication_role = replica';
    FOR r IN (SELECT tablename FROM pg_tables WHERE schemaname = 'public') LOOP
        EXECUTE 'DROP TABLE IF EXISTS public.' || quote_ident(r.tablename) || ' CASCADE';
    END LOOP;
    EXECUTE 'SET session_replication_role = origin';
END
\$do\$;
EOF

  echo "✅ Todas las tablas existentes fueron eliminadas."
else
  echo "🆕 La base '$DEST_DB' no existe. Creándola..."
  sudo -u postgres createdb "$DEST_DB"
fi

echo "💾 Restaurando dumps desde '$DUMP_DIR'..."

for dump_file in "$DUMP_DIR"/mvpw3_*.db; do
  echo "→ Restaurando desde: $dump_file"
  sudo -u postgres pg_restore \
    --dbname="$DEST_DB" \
    --no-owner \
    --verbose \
    "$dump_file"
done

echo "✅ Restauración finalizada."

# Exportación a JSON
echo "📤 Exportando tablas a JSON..."
mkdir -p "$EXPORT_DIR"

tables=$(sudo -u postgres psql -d "$DEST_DB" -Atc "SELECT tablename FROM pg_tables WHERE schemaname = 'public';")

for table in $tables; do
  echo "→ Exportando tabla: $table"
  sudo -u postgres psql -d "$DEST_DB" -t -A -F '' -c \
  "SELECT jsonb_build_object('id', id) || xdo::jsonb FROM $table;" \
  | jq -s '.' > "$EXPORT_DIR/$table.json"
done

echo "✅ Exportación completa en '$EXPORT_DIR'."
