# Express.js 설정

## Express를 설정하여 응용프로그램 제공

지금까지 프로젝트 디렉터리를 구성하고 TypeScript 컴파일러 옵션을 설정하고, `Hello World`를 출력하는 매우 간단한 TypeScript 응용프로그램을 만들었습니다.

이 섹션에서는 응용프로그램을 확장하여 `ExpressJS`를 추가하여 응용프로그램을 제공하고 다양한 종류의 요청을 처리하기 위해 미들웨어를 구성합니다.

## 새로운 종속성 설치

이 섹션에서 필요한 새 패키지를 설치하려면 다음을 실행합니다. 각 패키지에 대해 자세히 설명하겠습니다.

```
npm install --save \
  express \
  express-rate-limit \
  express-session \
  http
```

패키지의 TypeScript 타입 및 인터페이스 정의를 개발 디펜던시로 설치합니다.

```
npm install --save-dev \
  @types/express \
  @types/express-rate-limit \
  @types/express-session \
  @types/node
```

## 응용프로그램 캡슐화

> 이 자습서에서는 객체지향 프로그래밍(OOP) 개념을 사용합니다. OOP에 익숙하지 않은 경우 [이 MDN 페이지](https://developer.mozilla.org/en-US/docs/Learn/JavaScript/Objects/Object-oriented_JS)를 참조하여 간략한 소개를 받아보세요.

이 튜토리얼은 이전 모듈에서 간단한 'Hello world' 앱을 구성한 경우에 다룰 애플리케이션 코드를 담을 새 파일을 만드는 것으로 시작합니다. `index.ts` 파일에 방금 'Hello world' 앱을 구성한 경우에는 파일을 그대로 두시면 됩니다.

```bash
touch ./src/app.ts
```

새 파일(`./src/app.ts`) 맨 위에 Express 및 그 TypeScript 타입/인터페이스 정의를 가져오는 import 문을 추가한 후, 응용프로그램을 실행하는 데 필요한 모든 관련 함수와 구성을 담을 `App` 클래스를 만듭니다.

클래스 생성자는 해당 클래스가 인스턴스화될 때 포트 3000에서 Express를 초기화합니다. 생성자가 호출할 `listen()`이라는 비공개 함수를 생성하여 Express 서버의 실제 시작을 처리합니다.

마지막에 클래스를 내보내도록 잊지 마세요. 이후에 필요하게 됩니다.

```typescript
import express from 'express';
import type { Application } from 'express';
import type { Server } from 'http';

class App {
  public app: Application;
  private server: Server;

  constructor() {
    this.app = express();
    this.server = this.listen(3000);
  }

  private listen(port: number) {
    this.server = this.app.listen(port, () => {
      console.log(`App listening on port: ${port}`);
    });

    return this.server;
  }

}

export default App;
```

`npx tsc`를 사용하여 새 파일을 테스트하고 `./dist` 디렉터리에 `app.js` 파일이 성공적으로 생성되는지 확인합니다. 오류가 없다면, 잘 하셨습니다! 이제 라우트를 추가할 준비가 되었습니다.

## 기본 라우트

Express는 지정된 포트에 대한 모든 요청을 수신할 수 있지만 요청을 받으면 _무엇을 해야 하는지_ 알지 못합니다. Express에게 받은 요처를 _어떻게_ 처리해야할지 알려야 합니다. 이를 구성하는 방법은 응용프로그램의 아키텍처 및 응용프로그램이 요청을 받을 때 어느 경로/페이지에 응답해야 하는지에 크게 의존합니다.

이 자습서에서는 간단하게 유지합니다. 최종 목표는 사용자에게 Snyk내에서 프로젝트 목록을 제공하는 것이므로 최소한 두 개의 경로가 필요합니다: 초기 연결을 처리하는 인덱스 경로(`/`) 및 Snyk에서 데이터를 제공할 프로젝트 경로(`/projects`).

프로젝트 구성을 유지하기 위해 경로 모음을 보유할 `routes` 디렉터리를 만들고, 각 경로에 해당하는 경로의 주 컨트롤러 파일과 필요한 추가 항목을 포함시킬 서브디렉터리를 만듭니다.

이 모듈에서는 인덱스 경로에 초점을 맞춥니다.

```bash
mkdir -p ./src/routes/index
touch ./src/routes/index/indexController.ts
```

인덱스 컨트롤러를 작업하기 전에 모든 컨트롤러 데이터에 대한 공통 형태를 설명하는 [TypeScript 인터페이스](https://www.typescriptlang.org/docs/handbook/interfaces.html) 정의를 만들어야합니다. 이렇게 하면 컨트롤러를 일관되게 유지하고 TypeScript가 컨트롤러에 중요한 내용이 누락되었을 때 미리 경고할 수 있습니다.

우리가 경로에 대해 결정한 패턴과 유사한 분리 패턴을 사용하겠습니다. TypeScript에서 인터페이스 정의는 일반적으로 독립적이고 기술적이므로 각 정의에 대한 디렉터리를 만들지 않고 생성한 인터페이스 파일들을 동등한 위치에 저장합니다.

```bash
mkdir -p ./src/interfaces
touch ./src/interfaces/Controller.ts
```

우리가 만드는 모든 경로 컨트롤러는 경로 값과 Express에서 가져온 라우터를 최소한으로 가져야합니다. `./src/interfaces/Controller.ts`를 수정하고 아래 내용으로 채웁니다.

```typescript
import type { Router } from 'express';

export interface Controller {
  path: string;
  router: Router;
}
```

이제 이 정의를 완료하고, 에디터에서 경로 컨트롤러 파일 `./src/routes/index/indexController`를 열고 다음을 추가합니다:

```typescript
import type { Controller } from '../../interfaces/Controller';
import type { Request, Response, NextFunction } from 'express';
import { Router } from 'express';

class IndexController implements Controller {
  public path = '/';
  public router = Router();

  constructor() {
    this.initRoutes();
  }

  private initRoutes() {
    this.router.get(`${this.path}`, this.indexPage);
  }

  private indexPage(req: Request, res: Response, next: NextFunction) {
    return res.send('index route!');
  }
}

export default IndexController;
```

여기에서 몇 가지 TypeScript 타입/인터페이스 정의와 Express 패키지의 Router 개체가 가져옵니다.

`IndexController`를 사용하여 `/` 요청을 `indexPage` 함수로 응답하도록 설정했습니다. 그러나 우리의 App 클래스에는 인스턴스화될 때 경로를 등록하는 방법이 아직 없습니다. `./src/app.ts` 파일로 돌아가서 인스턴스화될 때 일련의 컨트롤러 및 포트 값을 인수로 하는 `initRoute()` 함수를 추가하세요.

또한 App의 생성자 함수를 컨트롤러 배열 및 포트 값으로 업데이트합니다.

`./src/app.ts` 파일을 편집하고 다음 내용으로 업데이트하세요:

```typescript
import express from 'express';
import type { Application } from 'express';
import type { Server } from 'http';
import type { Controller } from './interfaces/Controller';

class App {
  public app: Application;
  private server: Server;

  constructor(controllers: Controller[], port: number) {
    this.app = express();
    this.initRoutes(controllers);
    this.server = this.listen(port);
  }

  private initRoutes(controllers: Controller[]) {
    controllers.forEach((controller: Controller) => {
      this.app.use('/', controller.router);
    });
  }

  private listen(port: number) {
    this.server = this.app.listen(port, () => {
      console.log(`App listening on port: ${port}`);
    });

    return this.server;
  }

}

export default App;
```

## Express 서버 실행

지금까지 인스턴스화되면 `/`에서 요청에 응답하는 Express 서버를 실행하는 App 클래스를 만들었으며, 모든 컨트롤러가 동일한 데이터 패턴을 따르도록 보장하는 인터페이스 정의를 추가했습니다.

이제 프로젝트를 실행하고 확인해봅시다!

이 단계에서 컴파일된 `./dist/app.js`를 실행하면 흥미로운 일이 없습니다. 이는 내보낸 클래스를 인스턴스화하지 않았기 때문입니다. 이전 모듈에서 `package.json`을 초기화했습니다. 해당 명령어를 실행하면 앱에 관한 질문이 제공됩니다. 질문 중 하나는 프로젝트의 진입점(entrypoint)이 무엇인지에 관한 것입니다. 따라서 따라오시고 기본 값을 선택한 경우 `package.json`에서 `index.js`로 설정되어 있을 것입니다. 해당 진입점은 App 클래스를 인스턴스화할 위치입니다.

`./src/index.ts` 파일(여기에 'Hello world' 코드가 있던 곳)을 열고 내용이 남아 있는 경우 지우고, App 및 필요한 다른 객체에 대한 import 문을 추가하세요.

진입점은 한 가지만 수행하며, `new` 키워드를 사용하여 App 클래스의 인스턴스화를 하는 것 뿐입니다:

```typescript
import IndexController from './routes/indexController';
import App from './app';

new App(
  [
    new IndexController()
  ],
  3000
);
```

그게 다입니다!

`npx tsc`를 사용하여 프로젝트를 빌드하고 컴파일된 진입점을 `node`로 실행하세요:

```bash
npx tsc && node ./dist/index.js
```

모든 작업이 성공적이었다면 콘솔에서 `App listening on port: 3000`가 표시됩니다.

이제 브라우저에서 http://localhost:3000을 방문하거나 curl을 사용하여 응답을 확인하세요. `index route!`를 보면 축하합니다! Express가 응용프로그램을 제공하고 경로에 응답합니다!

다음 자습서 모듈에서는 라우트 및 인증에 좀 더 깊게 들어가고 Snyk API와 작업을 시작할 것입니다.
