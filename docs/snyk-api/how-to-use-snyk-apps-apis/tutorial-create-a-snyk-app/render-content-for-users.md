# 사용자를 위한 콘텐츠 렌더링

이전 모듈에서는 Snyk 앱을 등록하고 권한 부여 흐름을 설정하며 앱 내에서 사용자 권한을 처리하는 방법을 다루었습니다. 이러한 주제들은 모두 모든 Snyk 앱의 기능에 대하여 중요하지만, 그것들은 모두 '보이지 않는' 주제라고 할 수 있습니다.

이번 모듈에서는 사용자에게 콘텐츠를 표시하는 데 초점을 맞춥니다. 특히, 권한이 부여된 사용자에게는 Snyk에서 가져온 프로젝트 목록을 표시하고, 권한이 부여되지 않은 사용자에게는 권한 부여를 할 수 있는 큰 버튼을 보여주고자 합니다.

## Snyk 앱에 템플릿 엔진 추가

Express는 화면에 내용을 출력하고 HTML 서버 측에서 렌더링할 수 있지만, 템플릿 엔진을 사용하면 더욱 쉽게 작업할 수 있습니다. 이 튜토리얼에서는 [EJS](https://ejs.co)를 사용합니다.

먼저, 이 튜토리얼의 이 부분에서 필요한 노드 패키지를 설치하세요:

```bash
npm install --save ejs
```

다음으로, 이전 모듈에서 생성한`initGlobalMiddlewares()` 함수를 수정하여 express에게 EJS와 같은 _뷰 엔진_을 사용하고 뷰 템플릿을 저장할 위치를 알려줍니다. 우리는 EJS 템플릿을 `./src/views`에 저장하고 이미지 및 CSS와 같은 일반적인 파일은 `./src/public`에 저장할 것입니다.

먼저 새 디렉토리를 생성하세요.

```bash
mkdir -p ./src/views/partials
mkdir -p ./src/public
```

이제 `./src/app.ts`를 업데이트할 수 있습니다:

```typescript
// ./src/app.ts
...

class App {

  ...

  private initGlobalMiddlewares() {

    ...

    this.app.set("views", path.join(__dirname, "/views"));
    this.app.set("view engine", "ejs");
    this.app.use('/public', express.static(path.join(__dirname, '/public')));

    ...

  }

  ...

}
```

각 루트에 대해 템플릿을 제공할 때 해당 컨트롤러를 수정하고 `res.render("<템플릿 이름>")`을 사용하여 `res.send()`와 같이 더 간단한 것 대신에 사용하도록 해야 합니다.

예시:

```typescript
...

private initRoutes() {
  this.router.get(`${this.path}`, this.indexPage);
}
private indexPage(req: Request, res: Response, next: NextFunction) {
  // 이 부분이 Express가 EJS 템플릿을 렌더링하도록 하는 부분입니다.
  return res.render("index");
}

...
```

이게 전부입니다.

EJS 템플릿은 부분 포함 개념을 지원합니다. 엄격하게 필요한 것은 아니지만, 루트 템플릿과는 구분을 위해 헤더 및 푸터와 같은 부분 템플릿을 저장할 `./src/views`에 하위 디렉토리를 추가하는 것이 좋습니다. 이 튜토리얼에서는 `./src/views/partials`를 사용하여 이러한 템플릿을 저장할 것입니다.

## 기본 EJS 템플릿

첫 번째로 생성할 템플릿은 부분 템플릿이며, 다른 템플릿에 포함될 것입니다. `<head>`에 속하는 스타일시트 및 기타 정보들을 링크하는 곳인 `header.ejs`입니다.

```ejs
// ./views/partials/header.ejs

<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8" />
    <meta http-equiv="X-UA-Compatible" content="IE=edge" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <link rel="preconnect" href="https://fonts.googleapis.com" />
    <link rel="preconnect" href="https://fonts.gstatic.com" crossorigin />
    <link href="https://fonts.googleapis.com/css2?family=Roboto:wght@400;500&display=swap" rel="stylesheet" />
    <link href="https://raw.githubusercontent.com/snyk/snyk-apps-demo/main/src/public/css/styles.css" />
    <link rel="shortcut icon" href="https://raw.githubusercontent.com/snyk/snyk-apps-demo/main/src/public/images/snyk_dog.svg" type="image/x-icon" />

    <title>Snyk Apps Tutorial</title>
  </head>
</html>
```

이 `index.ejs` 템플릿은 기본 `/` 루트를 다룰 것입니다.

```ejs
// ./views/index.ejs

<%- include('./partials/header.ejs') %>

<body>
  <div class="index-page">
    <img class="index-page__snyk-logo" src="https://github.com/snyk/snyk-apps-demo/raw/main/src/public/images/snykLogoWithDog.svg" alt="snyk-logo" />
    <div class="index-page__card">
      <h1 class="index-page__title">Add Demo App</h1>
      <p class="index-page__description">Authorize this App to connect to your Snyk account.</p>
      <button class="button" onclick="location.href='/auth';">Install App</button>
    </div>
  </div>
</body>
```

`callback.ejs`는 성공적인 사용자 권한 부여에 대해 렌더링됩니다.

```ejs
// ./views/callback.ejs

<%- include('./partials/header.ejs') %>
<body>
  <div>
    <h2 class="main__heading">Snyk Apps Tutorial</h2>
  </div>
  <div class="card__30">
    <div class="callback-page__success-box">
      <img class="snyk-con-img" src="https://github.com/snyk/snyk-apps-demo/raw/main/src/public/images/success_check.svg" alt="success" />
        <div>
          <h2 class="callback-page__success-text">Successfully connected to Snyk!</h2>
        </div>
        <button class="button" onclick="location.href='/projects';">List Projects</button>
      </div>
    </div>
  </div>
</body>
```

위의 템플릿들을 사용하여 새로운 루트를 만드는 경우 스스로 템플릿을 추가하기 시작하는 데 충분합니다. EJS를 계속 사용한다면 제공되는 기능에 대한 정보는 문서를 참조하세요.

Snyk 앱에 대한 콘텐츠를 렌더링하는 것은 당신이 원하는 만큼 간단하거나 복잡할 수 있습니다. JavaScript를 다루고 있기 때문에 선택사항이 매우 유연합니다!

## 사용자에게 프로젝트 목록 표시

기본 템플릿이 구성되었으므로, 사용자의 Snyk 데이터를 사용하여 Snyk 앱에 기능을 추가하는 방법을 살펴보겠습니다. 이 튜토리얼에서는 사용자가 Snyk 앱 내에서 Snyk의 모든 프로젝트를 볼 수 있도록 앱을 설정했습니다.

이것은 기본적이고 쉽게 확장 가능한 기능입니다.

우리는 다음을 만들어야 합니다:

* 새로운 루트 컨트롤러
* 프로젝트 데이터를 가져오는 기능(또는 기능들)
* 프로젝트를 표시하는 EJS 템플릿

API 작업을 시작합니다. 이전 모듈에서 생성한 `callSnykApi()` 함수를 사용하여 진행합니다. 이것은 특정 루트와 직접 관련이 있기 때문에 해당 컨트롤러와 함께 이 파일을 저장할 것입니다. 이 튜토리얼 모듈 전체에서 사용한 패턴을 따라, 이 파일들을 모두 `./src/routes/projects/`에 생성합니다.

```typescript
// ./src/routes/projects/projectsHandler.ts

import { readFromDb } from "../../util/DB";
import { callSnykApi } from "../../util/apiHelpers";
import { EncryptDecrypt } from "../../util/encrypt-decrypt";
import { AuthData } from "../../interfaces/DB";
import { APIVersion } from "../../interfaces/API";
import { ENCRYPTION_SECRET } from "../../app";

/**
 * 사용자 프로젝트를 모두 가져오는 프로젝트 핸들러
 * 이는 사용자가 Snyk API에서 사용자 접근 토큰을 사용하여
 * 지정된 조직의 모든 프로젝트를 가져오는 예제 목적의 함수입니다.
 * 실제 제품에서는 토큰 스코프에 따라 액세스 가능한 내용이 달라집니다.
 * @returns 사용자 프로젝트 목록 또는 빈 배열
 */
export async function getProjectsFromApi(): Promise<unknown[]> {
  // DB에서 데이터 읽기
  const db = await readFromDb();
  const data = mostRecent(db.installs);
  // 데이터가 없으면 빈 배열 반환
  if (!data) return [];

  // 데이터(액세스 토큰) 복호화
  const eD = new EncryptDecrypt(ENCRYPTION_SECRET as string);
  const access_token = eD.decryptString(data?.access_token);
  const token_type = data?.token_type;
  // Snyk API v1에 대해 구성된 axios 인스턴스 호출
  const result = await callSnykApi(
    token_type,
    access_token,
    APIVersion.V1
  ).post(`/org/${data?.orgId}/projects`);

  return result.data.projects || [];
}

/**
 *
 * @param {AuthData[]} installs 설치 목록에서 최근 설치 가져오기
 * @returns 최근 설치 또는 없음
 */
export function mostRecent(installs: AuthData[]): AuthData | void {
  if (installs) {
    return installs[installs.length - 1];
  }
  return;
}
```

다음으로 루트 컨트롤러를 작성합니다. 패턴을 따라, `./src/routes/projects/projectsController.ts`를 만듭니다.

```typescript
// ./src/routes/projects/projectsController.ts

import type { Controller } from "../../interfaces/Controller";
import type { NextFunction, Request, Response } from "express";
import { Router } from "express";
import { getProjectsFromApi } from "./projectsHandler";

export class ProjectsController implements Controller {
  // 이 컨트롤러의 기본 URL 경로
  public path = "/projects";
  // 이 컨트롤러의 Express 라우터
  public router = Router();

  constructor() {
    this.initRoutes();
  }

  private initRoutes() {
    // 모든 사용자 프로젝트 목록을 렌더링하는 루트
    this.router.get(`${this.path}`, this.getProjects);
  }

  private async getProjects(req: Request, res: Response, next: NextFunction) {
    try {
      const projects = await getProjectsFromApi();
      return res.render("projects", {
        projects,
      });
    } catch (error) {
      return next(error);
    }
  }
}
```

새로운 루트 컨트롤러를 추가할 때마다 `./index.ts`를 업데이트해야 합니다.

```typescript
// ./src/index.ts

import IndexController from "./routes/index/indexController";
import AuthController from "./routes/auth/authController";
import CallbackController from "./routes/callback/callbackController";
import ProjectsController from "./routes/projects/projectsController";
import App from "./app";

new App([
   new IndexController(),
   new AuthController(),
   new CallbackController(),
   new ProjectsController()],
  3000
);
```

## 요약

이 모듈에서 생성한 프로젝트 API 핸들러 및 컨트롤러를 사용하여 자체 사용자 지정 코드를 작성하고 Snyk App을 원하는 대로 작동하도록 만드는 데 필요한 모든 것을 가졌어야 합니다.

우리는 v1 API를 사용했지만 Snyk의 REST API를 주시하십시오. 추가 기능이 추가될 때, Snyk 앱에서 사용할 새로운 또는 효율적인 엔드포인트를 찾을 수 있습니다.