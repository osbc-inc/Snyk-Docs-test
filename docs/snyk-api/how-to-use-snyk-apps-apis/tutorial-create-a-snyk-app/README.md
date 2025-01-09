# 튜토리얼: Snyk App 만들기

이 튜토리얼에서는 TypeScript를 사용하여 간단한 Snyk App을 만들어 사용자에게 Snyk 내의 프로젝트 목록을 보여줍니다.

## 전제 조건

* NodeJS
* Snyk 계정

## 기본 구성

먼저, 기기의 어딘가에 프로젝트를 담을 디렉토리를 생성합니다. 새롭게 생성된 디렉토리 안에서 응용프로그램의 종속성을 추적하고 프로젝트를 이식할 수 있도록 `package.json` 매니페스트를 초기화합니다:

`npm init`을 실행하고 프롬프트에 따릅니다. 여기에 원하는 만큼 많은 정보를 추가할 수 있습니다. 지금은 기본 옵션이 적당합니다.

이제 종속성 정보를 저장할 곳이 있으므로 TypeScript를 개발 종속성으로 설치하기 위해 `npm`을 사용합니다:

```bash
npm install typescript --save-dev
```

이 시점에서 TypeScript가 설치되었지만, `tsc` 이진에 컴파일 옵션을 제공하려면 구성 파일이 필요합니다. 프로젝트의 루트에 `tsconfig.json`이라는 TypeScript 구성 파일을 만듭니다. 다음 템플릿을 사용합니다:

```json
{
  "compilerOptions": {
    "target": "es6",
    "module": "commonjs",
    "declaration": true,
    "sourceMap": true,
    "moduleResolution": "node",
    "outDir": "./dist",
    "rootDir": "./src",
    "strict": true,
    "skipLibCheck": true,
    "esModuleInterop": true
  },
  "exclude": [
    "./tests",
    "./dist"
  ]
}
```

제공한 옵션은 TypeScript가 ES6 자바스크립트를 출력하도록하고 생성할 모듈 코드 유형 및 컴파일된 파일에 대한 소스 맵을 제공할지 여부를 지정하며 몇 가지 편리한 옵션을 지정합니다. 가능한 옵션에 대한 전체 개요는 TypeScript 설명서의 [TSConfig 참조 소개](https://aka.ms/tsconfig.json)를 참조하십시오.

이 튜토리얼의 목적을 위해 주목할만한 옵션은 `rootDir`와 `outDir`입니다. 이러한 옵션은 각각 프로젝트 내에서 소스 `.ts` 파일과 컴파일된 `.js` 파일이 있는 위치를 설명합니다. 설정 값에 의해 참조되는 디렉토리를 만듭니다:

```bash
mkdir ./dist
mkdir ./src
```

## 시험해 보기

기본 매개변수가 설정되었으므로 `./src/index.ts` 파일을 만들어 간단한 Hello World를 만들겠습니다.

```
const world = 'world';

export function hello(world: string = world): string {
  return `Hello ${world}! `;
}
```

이제 모든 것이 올바르게 연결되어 있는지 확인할 수 있습니다. 프로젝트를 컴파일하려면 `npx tsc`를 실행합니다.

모든 것이 성공하면 프로젝트 트리는 다음과 같습니다:

```
my-snyk-app/
 - dist/
   - index.d.ts
   - index.js
   - index.js.map
 - src/
   - index.ts
 - node_modules/
   - <lots of things here>
 - package-lock.json
 - tsconfig.json
```

`tsc` 프로그램이 소스 TypeScript 파일을 ES6 자바스크립트로 컴파일하고 디버깅을 위한 소스 맵을 제공했습니다.

{% hint style="info" %}
컴파일러는 `.ts` 파일의 변경 사항을 지속적으로 폴링하기 위해 감시 모드로 설정할 수도 있습니다: `npx tsc -w`
{% endhint %}
