# known-socially

App for aggregating social media accounts

Commands

```bash
# Start the database
docker compose up db

# Build frontend
cd frontend/app
bun run build

# Build everything
#docker compose up go-builder frontend-builder


# Start the Go service
go build . && PORT=80 ./${PWD##*/}

# Development commands
cd frontend/app && bun run dev    # Run frontend in dev mode
```
