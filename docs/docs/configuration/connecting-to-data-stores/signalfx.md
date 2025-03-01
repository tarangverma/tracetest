# SignalFx

Tracetest fetches traces from [SignalFx's realm and token](https://docs.splunk.com/Observability/references/organizations.html).

:::tip
Examples of configuring Tracetest can be found in the [`examples` folder of the Tracetest GitHub repo](https://github.com/kubeshop/tracetest/tree/main/examples). 
:::

## Configure Tracetest to use SignalFx as a Trace Data Store

Configure Tracetest to be aware that it has to fetch trace data from SignalFx. 

:::tip
Need help configuring the OpenTelemetry Collector so send trace data from your application to SignalFx? Read more in [the reference page here](../opentelemetry-collector-configuration-file-reference)). 
:::

## Connect Tracetest to SignalFx with the Web UI

In the Web UI, open settings, and select SignalFx.

You need your SignalFx:

- **Realm**
- **Token**

Follow this [guide](https://docs.splunk.com/Observability/references/organizations.html).

![](https://res.cloudinary.com/djwdcmwdz/image/upload/v1674644337/Blogposts/Docs/screely-1674644332529_cks0lw.png)


## Connect Tracetest to SignalFx with the CLI

Or, if you prefer using the CLI, you can use this file config.

```yaml
type: DataStore
spec:
  name: SignalFX
  type: signalFx
  isDefault: true
  signalFx:
    realm: us1
    token: mytoken
```

Proceed to run this command in the terminal, and specify the file above.

```bash
tracetest datastore apply -f my/data-store/file/location.yaml
```
