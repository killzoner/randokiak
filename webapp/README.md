# Webapp

A webapp in Angular, because why not!

Steps for install :

- `npm install -g yarn`
- `yarn`
- (OPTIONAL) To generate typescript definitions from protos, run `make` (does not seem to work in windows, even in bash compatible shell, too bad. Use Ubuntu for windows for that!)

## Development server

Run `npm start` for a dev server. Navigate to `http://localhost:4200/`. The app will automatically reload if you change any of the source files.

## Build

Run `npm run build`

Note: with bazel, file are located under `bazel-out/dist/bin/src/prodapp`
Build is not passing in travis, but is ok locally. Should move TravisCI to bazel based build...
