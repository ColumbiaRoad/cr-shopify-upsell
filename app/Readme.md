# Developing the Svelte app

The built files from this app are served from `/backend/templates`

First time you need to install dependencies with `npm i`

After you make changes:

1) run `npm run build` , a build folder will be created on `/public/build`
2) copy the contents of `/public/build` to `backend/templates/build`

Note that the path from global.css is different in /app and /backend, take that into account when modifying it

## TODO

* Make an automated build step to copy files
