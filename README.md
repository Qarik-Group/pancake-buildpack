# Pancake Buildpack

This tiny buildpack does a simple thing for your application - it flattens the crazy `$VCAP_SERVICES` service instance credentials into many specific environment variables.

In the following example, in addition to `$VCAP_SERVICES` the application will also have variables starting with `MYSQL_` for each credential (such as `MYSQL_HOSTNAME`, `MYSQL_USERNAME`, `MYSQL_URI`, for credentials `hostname`, `username`, and `uri`):

```plain
cf cs p-mysql 10mb db
cf push -p fixtures/phpapp -f fixtures/phpapp/manifest.yml
```

This `phpinfo` sample app will show that your `p-mysql` service instance credentials are available as dedicated environment variables:

![phpinfo](https://cl.ly/fd9532d45a9b/download/phpinfo-cf-pancake.png)

The environment variables are available in many permutations for your ease of usage:

* Prefixed by each tag name. For example tags `[mysql, sql]` would generate environment variables `MYSQL_...` and `SQL_...`.
* Prefixed by the service name. For example if the service is named `p-mysql` then environment variables would be prefixed `P_MYSQL_...`.

To see the generated environment variables on your application:

```plain
$ cf ssh phpapp -c "/tmp/lifecycle/shell /home/vcap/app 'env | sort'"
MYSQL_HOSTNAME=10.144.0.17
MYSQL_JDBCURL=jdbc:mysql://10.144.0.17:3306/cf_ac66eb76_f883_41c5_8840_53fe71e7cb63?user=rhUB9OOJlXqdmwSE&password=W1SRPFb2TeoBP9GM
MYSQL_NAME=cf_ac66eb76_f883_41c5_8840_53fe71e7cb63
MYSQL_PASSWORD=W1SRPFb2TeoBP9GM
MYSQL_URI=mysql://rhUB9OOJlXqdmwSE:W1SRPFb2TeoBP9GM@10.144.0.17:3306/cf_ac66eb76_f883_41c5_8840_53fe71e7cb63?reconnect=true
MYSQL_USERNAME=rhUB9OOJlXqdmwSE
...
P_MYSQL_HOSTNAME=10.144.0.17
P_MYSQL_JDBCURL=jdbc:mysql://10.144.0.17:3306/cf_ac66eb76_f883_41c5_8840_53fe71e7cb63?user=rhUB9OOJlXqdmwSE&password=W1SRPFb2TeoBP9GM
P_MYSQL_NAME=cf_ac66eb76_f883_41c5_8840_53fe71e7cb63
P_MYSQL_PASSWORD=W1SRPFb2TeoBP9GM
P_MYSQL_URI=mysql://rhUB9OOJlXqdmwSE:W1SRPFb2TeoBP9GM@10.144.0.17:3306/cf_ac66eb76_f883_41c5_8840_53fe71e7cb63?reconnect=true
P_MYSQL_USERNAME=rhUB9OOJlXqdmwSE
```

## Usage

Simple add `pancake_buildpack` or `https://github.com/starkandwayne/pancake-buildpack` to the start of your `manifest.yml` buildpacks list.

For example, from `fixtures/phpapp/manifest.yml`:

```yaml
applications:
- name: phpapp
  buildpacks:
  - pancake_buildpack
  - php_buildpack
  services:
  - db
```

If you use `pancake_buildpack` then your Cloud Foundry administrator will need to install the `pancake_buildpack` using `cf create-buildpack`.

If you cannot get a Cloud Foundry administrator to do this, then use the Git URL:

```yaml
applications:
- name: phpapp
  buildpacks:
  - https://github.com/starkandwayne/pancake-buildpack
  - php_buildpack
  services:
  - db
```

## Buildpack Developer Documentation

To build this buildpack, run the following command from the buildpack's directory:

1. Source the .envrc file in the buildpack directory.

    ```bash
    source .envrc
    ```

    To simplify the process in the future, install [direnv](https://direnv.net/) which will automatically source .envrc when you change directories.

1. Install buildpack-packager

    ```bash
    ./scripts/install_tools.sh
    ```

1. Build the buildpack

    ```bash
    buildpack-packager build -stack cflinuxfs3 -cached
    ```

1. Use in Cloud Foundry

    Upload the buildpack to your Cloud Foundry.

    ```bash
    cf create-buildpack pancake_buildpack pancake_buildpack-*.zip 1
    cf cs p-mysql 10mb db
    cf push -p fixtures/phpapp -f fixtures/phpapp/manifest.yml
    ```


### Testing

Buildpacks use the [Cutlass](https://github.com/cloudfoundry/libbuildpack/cutlass) framework for running integration tests.

To test this buildpack, run the following command from the buildpack's directory:

1. Source the .envrc file in the buildpack directory.

    ```bash
    source .envrc
    ```

    To simplify the process in the future, install [direnv](https://direnv.net/) which will automatically source .envrc when you change directories.

1. Run integration tests

    ```bash
    cf cs p-mysql 10mb db
    ./scripts/integration.sh
    ```

    To run integration tests against CFDev:

    ```bash
    cf login -a https://api.dev.cfdev.sh --skip-ssl-validation -u admin -p admin
    cf cs p-mysql 10mb db
    CUTLASS_SCHEMA=https CUTLASS_SKIP_TLS_VERIFY=true ./scripts/integration.sh
    ```

    More information can be found on Github [cutlass](https://github.com/cloudfoundry/libbuildpack/cutlass).

### Reporting Issues

Open an issue on this project.