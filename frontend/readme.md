# Frontend

I am trying to create a svelt app that creates static files that can be run in golang

I am following this tutorial https://codeandlife.com/2022/02/12/combine-golang-and-sveltekit-for-gui/

I started but running this

```bash
mkdir app
cd app
npx sv create
bun i -D @sveltejs/adapter-static@next
```

I am choosing to run it with [bun](https://bun.sh/)

Notes after the create

```bash
# Run the app
bun run dev --open
# Stuck? Visit us at https://svelte.dev/chat

# Use drizzle to control db (removed because demo was an issue)
bun run db:start # start the docker container
bun run db:push # update your database schema

npm run storybook

#Lucia auth demo (removed because demo was an issue)
http://localhost:5173/demo/lucia/login

# paraglide i18n
http://localhost:5173/demo/paraglide
# Edit messages in messages/en.json
```

https://svelte.dev/docs/kit/adapter-static
https://khromov.se/the-missing-guide-to-understanding-adapter-static-in-sveltekit/