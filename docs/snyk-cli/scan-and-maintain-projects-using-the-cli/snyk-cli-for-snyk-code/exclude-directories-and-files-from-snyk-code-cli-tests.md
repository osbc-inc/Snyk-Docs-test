# {{Snyk Code}} CLI 테스트에서 디렉터리와 파일 제외

{Snyk Code}} 리포지토리를 CLI를 사용하여 테스트할 때, `snyk ignore --file-path` 명령을 사용하여 CLI 테스트에서 특정 디렉터리와 파일을 제외할 수 있습니다. 이 명령을 실행하면 자신이 제외하려는 디렉터리나 파일의 이름이 포함된 `.snyk` 파일이 자동으로 리포지토리에 생성됩니다.

{% hint style="info" %}
* `.snyk` 파일을 수동으로 생성하여 CLI 테스트에서 디렉터리와 파일을 제외할 수도 있습니다. `.snyk` 파일의 수동 생성에 대한 자세한 내용은 [수입 프로세스에서 디렉터리 및 파일 제외](https://docs.snyk.io/products/snyk-code/getting-started-with-snyk-code/activating-snyk-code-using-the-web-ui/step-3-importing-repositories-to-snyk-for-the-snyk-code-testing/excluding-directories-and-files-from-the-import-process)를 참조하세요.
* `snyk ignore --file-path` 명령은 특정 취약점 문제를 무시하지 않습니다. 이 명령은 CLI 테스트에서 디렉터리와 파일만 제외합니다.
* 제작 단계로 컴파일되거나 배포되지 않는 디렉터리와 파일을 제외하는 것을 고려하십시오. 제외된 파일이나 디렉터리에 기존 취약점이 있는 경우, 추적이 해당되면 Snyk는 잠재적 문제를 감지하지 못할 수 있습니다.
{% endhint %}

## **CLI 테스트에서 디렉터리와 파일 제외하기**

다음 단계를 따라 {Snyk Code}} 디렉터리와 파일을 CLI 테스트에서 제외하십시오:

1\. 터미널에서 테스트하려는 폴더로 디렉터리를 변경하십시오.

`snyk ignore --file-path` 명령은 해당 폴더와 해당 폴더의 하위 폴더 및 파일에만 적용됩니다.

2\. 터미널에서 다음 내용을 입력하십시오:

```
snyk ignore --file-path=<디렉터리_또는_파일>
```

여기서 `디렉터리_또는_파일`은 테스트에서 제외하려는 디렉터리나 파일의 이름입니다. 예를 들어 `db.js`입니다.

특정해제된 디렉터리나 파일이 포함된 `.snyk` 파일이 루트 폴더에 생성됩니다.

`.snyk` 파일은 숨김 파일로 생성됩니다. 루트 폴더에 보이지 않는 경우 기계에서 **숨겨진 파일 보기** 옵션을 사용하십시오.

3\. 선택적으로, 제외할 여러 디렉터리나 파일을 지정하려면 다음을 입력하십시오:

```
snyk ignore --file-path=<디렉터리1_또는_파일1> && snyk ignore --file-path=<디렉터리2_또는_파일2> && snyk ignore --file-path=<디렉터리3_또는_파일3>
```

이제 선택한 폴더에서 `snyk code test` 명령을 실행할 때 지정한 디렉터리나 파일이 테스트에서 제외됩니다.

## CLI 테스트에서 제외된 파일 다시 포함하기&#x20;

테스트에서 제외된 디렉터리나 파일을 다시 포함하려면 `.snyk` 파일을 수동으로 편집하거나 삭제하십시오.

1\. `snyk-goof-master` 폴더에서 12개의 문제가 발견되었습니다. 세 개의 다른 파일(`app.js`, `db.js`, `routes/index.js`)에서 발견된 문제입니다:

<figure><img src="../../../.gitbook/assets/snyk Code - CLI - snyk code test - Exclusion - before -2.png" alt="CLI 테스트에서 발견된 문제"><figcaption><p>CLI 테스트에서 발견된 문제</p></figcaption></figure>

2\. `app.js`와 `db.js` 파일을 제외하고 `routes/index.js` 파일에서 발견된 문제만 표시하려면 다음을 입력하십시오:

```
snyk ignore --file-path=app.js && snyk ignore --file-path=db.js
```

<figure><img src="../../../.gitbook/assets/snyk Code - CLI - snyk code test - Exclusion - Example command.png" alt="터미널에서 snyk ignore 명령"><figcaption><p><code>snyk ignore</code> 명령</p></figcaption></figure>

3\. `snyk ignore` 명령을 입력하면 `.snyk` 파일이 `snyk-goof-master` 폴더에 자동으로 생성됩니다:

<figure><img src="../../../.gitbook/assets/snyk Code - CLI - snyk code test - Exclusion - Example - .snyk file.png" alt="폴더에 나열된 .snyk 파일"><figcaption><p><code>.snyk</code> 파일이 폴더에 나열됨</p></figcaption></figure>

이 `.snyk` 파일에는 제외되어야 하는 파일이 포함됩니다:

<figure><img src="../../../.gitbook/assets/snyk Code - CLI - snyk code test - Exclusion - Example - .snyk file - content.png" alt=".snyk 파일의 내용"><figcaption><p>.snyk 파일의 내용</p></figcaption></figure>

4\. 테스트가 다시 실행되면 `app.js`와 `db.js` 파일이 테스트에서 제외되고, 결과에는 `routes/index.js` 파일에서 발견된 문제만 표시됩니다:

<figure><img src="../../../.gitbook/assets/snyk Code - CLI - snyk code test - Exclusion - after - 2.png" alt="ignore 명령 사용 후 발견된 문제"><figcaption><p><code>ignore</code> 명령 사용 후 발견된 문제</p></figcaption></figure>