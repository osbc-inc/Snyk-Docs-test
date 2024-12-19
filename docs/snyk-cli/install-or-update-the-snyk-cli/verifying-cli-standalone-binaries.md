# CLI 독립 실행 파일 확인

다운로드된 바이너리 파일의 shasum 및 GPG 서명을 모두 확인할 수 있습니다.

`downloads.snyk.io`의 다운로드 위치에 `sha256sums.txt.asc`라는 파일이 있습니다. 직접 다운로드할 수 있습니다: `https://downloads.snyk.io/cli/stable/sha256sums.txt.asc` 또는 특정 버전을 위해 `https://downloads.snyk.io/cli/v1.666.0/sha256sums.txt.asc`와 같은 링크를 사용할 수 있습니다.

다운로드한 파일이 체크섬과 일치하는지 확인하려면 `sha256sum` 명령어를 사용하면 됩니다. 예를 들어:

```bash
grep snyk-macos sha256sums.txt.asc | sha256sum -c -
```

Snyk CLI 독립 실행 파일이 [Snyk CLI GPG 키](https://github.com/snyk/cli/blob/master/help/\_about-this-project/snyk-code-signing-public.pgp)에 대해 확인하려면 먼저 GPG 키를 가져와야 합니다:

```bash
# A22665FB96CAB0E0973604C83676C4B8289C296E는 code-signing@snyk.io에 속한 키입니다
# 이 공개 키의 사본은 이 저장소 /help/_about-this-project/snyk-code-signing-public.pgp에도 있습니다
gpg --keyserver hkps://keys.openpgp.org --recv-keys A22665FB96CAB0E0973604C83676C4B8289C296E
```

그런 다음 파일이 서명되었는지 확인하기 위해 다음을 실행하십시오:

```bash
gpg --verify sha256sums.txt.asc
```

명령어 출력은 다음과 같아야 합니다:

```bash
gpg: Signature made So  8 Jan 14:11:44 2023 CET
gpg:                using EDDSA key A22665FB96CAB0E0973604C83676C4B8289C296E
gpg: Good signature from "Snyk Limited <code-signing@snyk.io>" [unknown]
gpg: WARNING: This key is not certified with a trusted signature!
gpg:          There is no indication that the signature belongs to the owner.
Primary key fingerprint: A226 65FB 96CA B0E0 9736  04C8 3676 C4B8 289C 296E
```