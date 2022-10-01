# gcloud app
Instructions for deploying a Go app protected by G Suite OAuth via Google Cloud Run so only your coworkers can access it.

## Setup

1. Install the [`gcloud` CLI](https://cloud.google.com/sdk/docs/install)
2. [Create gcloud Project](https://console.cloud.google.com/projectcreate)
3. Ensure you have selected your new project in the top left, then note the **"Project ID"** under the "Project info" panel.

```bash
export PROJECT_ID="FIXME"
```

3. Authenticate and configure `gcloud`:

```
gcloud auth login
gcloud config set project ${PROJECT_ID}
gcloud config set run/region us-west1
```

4. Choose an application name e.g. `gcloud-app`

```bash
export APP_NAME="FIXME"
```

5. Deploy this application, selecting `y` for all prompts.
    - Note: `Allow unauthenticated invocations to [gcloud-app] (y/N)?` is for IAM, not oauth2, so still respond with `y`.

```bash
gcloud run deploy ${APP_NAME}
```

6. Note the URL it is deployed at e.g. `https://gcloud-app-random-xy.a.run.app`:

```bash
export EXTERNAL_URI="FIXME"

# or if you have jq
export EXTERNAL_URI=$(gcloud run services describe bravo-app --format=json | jq -r .status.url)
echo $EXTERNAL_URI
```

7. [Create oauth credentials](https://console.cloud.google.com/apis/credentials) by clicking "+ New Credentials" then "Oauth client ID".
    - Follow the prompt to configure the oauth consent screen first; make it internal, and only fill in required fields.
    - You may need select "Credentials" on the left and "+ New Credentials" again
    - Select "Web application" for application type
    - Run `echo ${EXTERNAL_URI}/auth/redirect` and use the output as an "Authorized redirect URIs"
    - and note the **Client ID** and **Client Secret**

```bash
export OAUTH_CLIENT_ID="FIXME"

export OAUTH_CLIENT_SECRET="FIXME"
```

8. Set these env vars on the service

```bash
gcloud run services update ${APP_NAME} --update-env-vars "EXTERNAL_URI=${EXTERNAL_URI},OAUTH_CLIENT_ID=${OAUTH_CLIENT_ID},OAUTH_CLIENT_SECRET=${OAUTH_CLIENT_SECRET}"
```

9. Now you can login to your app!

```bash
open ${EXTERNAL_URI}
```

You can logout by going to `${EXTERNAL_URI}/logout`.

10. Finally verify that only your coworkers have access to the application by trying to login using a Google account that is not part of the Google workspace. You should get a screen from Google that says "Access blocked: <app name> can only be used within its organization."
