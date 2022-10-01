# gcloud app

## Setup

1. Install the [`gcloud` CLI](https://cloud.google.com/sdk/docs/install)
2. [Create gcloud Project](https://console.cloud.google.com/projectcreate)
3. Note the Project ID

```bash
export PROJECT_ID="FIXME"
```

3. Authenticate and configure `gcloud`:

```
gcloud auth login
gcloud config set project ${PROJECT_ID}
gcloud config set run/region us-west1
```

4. Select an application name e.g. `gcloud-app`

```bash
export APP_NAME="FIXME"
```

5. Deploy this application:
    - Enter `y` to enable things for the first time
    - Allow unauthenticated access `Allow unauthenticated invocations to [gcloud-app] (y/N)?  y` (this is for IAM, not oauth2)

```bash
gcloud run deploy
```

6. Note the URL it is deployed at:

```bash
export EXTERNAL_URI="FIXME"
```

7. Create oauth credentials: https://console.cloud.google.com/apis/credentials
    - Make the redirect URL `echo ${EXTERNAL_URI}/auth/redirect`
    - Note Client ID and Client Secret.

```bash
export OAUTH_CLIENT_ID="FIXME"

export OAUTH_CLIENT_SECRET="FIXME"
```

8. Set these env vars on the service

```bash
gcloud run services update $APP_NAME --update-env-vars "EXTERNAL_URI=${EXTERNAL_URI},OAUTH_CLIENT_ID=${OAUTH_CLIENT_ID},OAUTH_CLIENT_SECRET=${OAUTH_CLIENT_SECRET}"
```

9. Log into your app!

```bash
open ${EXTERNAL_URI}
```
