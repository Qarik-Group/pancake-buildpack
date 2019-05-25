# Pancake Buildpack

Flatten VCAP_SERVICES into many environment variables. This buildpack bundles and runs `cf-pancake export` to convert any `$VCAP_SERVICES` bindings into one environment variable per credential.

```plain
cf cs p-mysql 10mb db
cf push -p fixtures/phpapp
```

This `phpinfo` sample app will show that your `p-mysql` service instance credentials are available as dedicated environment variables:

![phpinfo](https://cl.ly/46be09e5d940/phpinfo-cf-pancake.png)

## Buildpack User Documentation

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
    cf push -p fixtures/phpapp
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
    ./scripts/integration.sh
    ```

    More information can be found on Github [cutlass](https://github.com/cloudfoundry/libbuildpack/cutlass).

### Reporting Issues

Open an issue on this project