# JavaScript

Snyk는 [코드 분석을 위한 JavaScript](javascript-for-code-analysis.md) 및 [오픈 소스용 JavaScript](javascript-for-open-source.md)를 지원합니다. SCM 통합을 사용하여 프로젝트를 가져오는 방법에 대한 정보는 [Git 저장소 및 JavaScript](git-repositories-and-javascript.md)를 참조하십시오.

[JavaScript 및 Node.js에 대한 지침](best-practices-for-javascript-and-node.js.md)이 제공됩니다.

## 적용 가능성

Snyk는 [코드 분석을 위한 JavaScript](javascript-for-code-analysis.md)와 [오픈 소스용 JavaScript](javascript-for-open-source.md)를 지원합니다.&#x20;

SCM 통합을 사용하여 프로젝트를 가져오는 방법에 대한 자세한 정보는 [Git 저장소 및 JavaScript](git-repositories-and-javascript.md)를 참조하십시오.

[JavaScript 및 Node.js에 대한 지침](best-practices-for-javascript-and-node.js.md)이 있습니다.

Snyk 제품을 사용하여 가져올 수 있는 언어의 가용성을 확인하십시오.

가능한 기능:

* SCM 가져오기, Snyk 오픈 소스 및 Snyk 코드에서 사용 가능합니다.
* CLI 및 IDE를 통해 앱을 테스트하거나 모니터링, Snyk 오픈 소스 및 Snyk 코드에서 사용 가능합니다.
* `pkg:npm`을 사용하여 앱의 SBOM 테스트
* `pkg:npm`을 사용하여 앱의 패키지 테스트

코드 분석을 위해 Snyk CLI를 사용하는 방법에 대한 정보는 [Snyk CLI를 사용한 Snyk 코드 스캔](../../snyk-cli/scan-and-maintain-projects-using-the-cli/snyk-cli-for-snyk-code/)을 참조하십시오.

## 패키지 관리자 및 지원되는 파일 확장자

Snyk는 JavaScript를 위해 npm, pnpm 및 Yarn을 패키지 관리자로 지원하며 다음 버전을 지원합니다:&#x20;

- npm: `Lockfile 1`, `Lockfile 2`, `Lockfile 3, 7.*`
- pnpm: `pnpm 7`, `pnpm 8`, `pnpm 9`
- Yarn: `Yarn 1`, `Yarn 2`, `Yarn 3`

패키지 레지스트리로 [npmjs.org](https://www.npmjs.org/)를 지원합니다.

Snyk는 JavaScript를 위해 다음 파일 형식을 지원합니다:

- Snyk 오픈 소스:
  - npm 패키지 관리자: `package.json`, `package-lock.json`
  - pnpm 패키지 관리자: `pnpm-lock.yaml`
  - Yarn 패키지 관리자: `yarn.lock`
- Snyk 코드: `.ejs`, `.es`, `.es6`, `.htm`, `.html`, `.js`, `.jsx`, `.ts`, `.cts`, `.mts`, `.tsx`, `.vue`, `.mjs`, `.cjs`

## 프레임워크 및 라이브러리

다음 프레임워크 및 라이브러리가 Snyk의 JavaScript에서 지원됩니다:

- @Google Drive/generative-ai - 포괄적
- @anthropic-ai/sdk - 포괄적
- @huggingface/inference - 포괄적
- @mistralai/mistralai - 포괄적
- axios - 포괄적
- Angular - 부분
- apollo-server - 부분
- bcrypt-nodejs - 포괄적
- cross-spawn - 포괄적
- crypto-js - 포괄적
- date-fns - 포괄적
- dayjs - 포괄적
- dompurify - 포괄적
- electron - 부분
- ejs - 부분
- execa - 포괄적
- express - 포괄적
- express-graphql - 부분
- express-jwt - 부분
- fs - 포괄적
- fs-extra - 포괄적
- fs-plus - 포괄적
- graceful-fs - 포괄적
- graphql-js - 부분
- jQuery - 포괄적
- js-yaml - 포괄적
- jzip - 포괄적
- koa - 포괄적
- koa-graphql - 포괄적
- libxml - 포괄적
- libxmljs - 포괄적
- lodash - 포괄적
- luxon - 포괄적
- minimongo - 포괄적
- minimist - 포괄적
- mongodb - 포괄적
- Mongoose - 포괄적
- mercurius - 부분
- Nestjs - 부분
- Node Crypto - 포괄적
- node-buffer - 부분
- node-cmd - 포괄적
- Node Crypto - 포괄적
- node-dir - 포괄적
- node-forge - 포괄적
- node-pty - 포괄적
- node-serialize - 포괄적
- octokit - 포괄적
- openai - 포괄적
- pg - 포괄적
- pg-promise - 포괄적
- React - 부분
- request-promise - 포괄적
- restler - 부분
- rimraf - 포괄적
- sanitize-html - 포괄적
- shelljs - 포괄적
- Stanford JS Crypto - 포괄적
- superagent - 포괄적
- tar-stream - 포괄적
- unirest - 포괄적
- unzip - 포괄적
- underscore - 포괄적
- url - 포괄적
- vm - 포괄적
- webstomp-client - 부분
- WebCryptoAPI - 포괄적
- xpath - 포괄적
- yargs - 포괄적

## 기능

다음의 기능이 Snyk의 JavaScript에서 지원됩니다:

| Snyk 오픈 소스                                                      | Snyk 코드                                                                  |
| -------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| <ul><li>PR 수정 </li><li>라이센스 스캔 </li><li>보고서</li></ul> | <ul><li>보고서</li><li>사용자 정의 규칙 </li><li>간파 분석</li></ul> |

도움이 필요하시면 [Snyk 지원팀에 문의](https://support.snyk.io)하십시오.