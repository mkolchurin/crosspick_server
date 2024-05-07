# docker buildx build --push --platform linux/amd64 --tag mkolchurin/crosspicktest:latest .
rm -rf www
cd decider-front
npx vite build --outDir ../www

# docker run --name pg_latest -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=tr2Jng6H# -d postgres:latest          