#i/bin.sh

# exit status return other 0
set -e

# first step
echo "run db migration"
/app/migrate -path /app/migration -database "$DB_SOURCE" -verbose up

echo "start the app"
exec "$@"
