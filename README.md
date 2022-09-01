# yobol

## Prerequisites

### Frontend

#### [Node.js](https://nodejs.org/en/)

A JavaScript runtime built on Chrome's V8 JavaScript engine.

##### [Downloads](https://nodejs.org/en/download/)

We use node 16.17.0 with npm 8.15.0 as developing environment.

- Windows: https://nodejs.org/dist/v16.17.0/node-v16.17.0-x64.msi

```shell
$ node -v
v16.17.0

$ npm -v
8.15.0


$ npm install -g cnpm --registry=http://registry.npm.taobao.org

added 359 packages in 27s

11 packages are looking for funding
  run `npm fund` for details
npm notice
npm notice New minor version of npm available! 8.15.0 -> 8.19.0
npm notice Changelog: https://github.com/npm/cli/releases/tag/v8.19.0
npm notice Run npm install -g npm@8.19.0 to update!
npm notice

```

#### [Vue.js](https://vuejs.org/)

An approachable, performant and versatile framework for building web user interfaces.

- Approachable: Builds on top of standard HTML, CSS and JavaScript with intuitive API and world-class documentation.

- Performant: Turly reactive, compiler-optimized rendering system that rarely requires manual optimization.

- Versatile: A rich, incrementally adoptable ecosystem that scales between a library and a full-featured framework.

##### [Quick Start](https://vuejs.org/guide/quick-start.html)

Initialize our project using the offical Vue project scaffolding tool '[create-vue](https://github.com/vuejs/create-vue)':

```shell
$ npm init vue@3
Need to install the following packages:
  create-vue@3.3.2
Ok to proceed? (y) y

Vue.js - The Progressive JavaScript Framework

√ Project name: ... yobol
√ Add TypeScript? ... No / [Yes]
√ Add JSX Support? ... [No] / Yes
√ Add Vue Router for Single Page Application development? ... No / [Yes]
√ Add Pinia for state management? ... No / [Yes]
√ Add Vitest for Unit Testing? ... No / [Yes]
√ Add Cypress for End-to-End testing? ... No / [Yes]
√ Add ESLint for code quality? ... No / [Yes]
√ Add Prettier for code formatting? ... No / [Yes]

Scaffolding project in D:\Project\Go\src\github.com\yobol\yobol\yobol...

Done. Now run:

  cd yobol
  npm install
  npm run lint
  npm run dev

# rename the directory yobol as web
$ cd web
$ npm install
$ npm run dev
```

When you are ready to shp your app to production, run the following:

```shell
$ npm run build
```

This will create a production-ready build of your app in the project's `./dist`  directory.

##### [Tips]

- Recommended IDE setup: [Visual Studio Code](https://code.visualstudio.com/) + [Volar extension](https://marketplace.visualstudio.com/items?itemName=Vue.volar)

#### [Ant Design of Vue](https://www.antdv.com/docs/vue/introduce)

The UI component library for Vue.

### Backend