# GO Notebook

## Local development

Start database:

```ps
docker compose up -d
```

Apply migrations. Nothing fancy jet, just copy and paste the script
in `migrations\0001.up.sql` in a psql shell:

```psql
docker exec -it db psql -U go-notebook -d go-notebook
```

```sql
go-notebook=# CREATE TABLE "Notebooks" (
  ...
```

## Testing

Uses test-contianers. Running requires docker.
