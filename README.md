# oso-authz-poc

Proof-of-concept using Postgres and [Oso][oso] for authorization.

## Why Oso?

Oso is unlike most authorization solutions:

- unlike [Zanzibar][zanzibar] because it doesn’t store all/tons of data
- unlike [AWS Verified Permissions][awsvp] because it doesn’t force your HTTP client to send a ton of data as part of its request.

Turns out, doing either of those things is a massive pain in the ass. From Oso’s [docs][distributed-authz]:

> That "send your authorization-relevant data to the service" item masks a lot of **pain**. Because that authorization-relevant data is also application data, you have to **maintain two copies** of it: one copy in your application, and another in the authorization service. Some of that data can be very large … Some of it **changes frequently**, like the jobs that are currently running against repositories in GitHub. Keeping all this data in sync, detecting and **remediating drift**, and ensuring consistency between your application and your authorization service is a **huge hidden cost** that many people don't consider when implementing authorization as a service.
>
> You don't have to send that application data to Oso Cloud. Instead, you can tell the Oso Cloud client in your application how to resolve authorization questions using local data. This allows you to use high-volume or transient data in authorization decisions without incurring the overhead of keeping that data consistent in two places.

Basically, Oso reads your database. You tell it [how][cfg].

[oso]: https://www.osohq.com/
[cfg]: https://github.com/osohq/rails_list_filtering_sample_app/blob/0703315759f5573bed90b0458cf3957a215981fa/config/oso.yml
[awsvp]: https://docs.aws.amazon.com/verifiedpermissions/latest/userguide/getting-started-differences-verifiedpermissions-cedar.html#differences-json-entities
[zanzibar]: https://www.osohq.com/learn/google-zanzibar
[distributed-authz]: https://www.osohq.com/docs/concepts/distributed-authorization

## Getting Started

### Run database

This will restart your containers and run database migrations.

```
pkgx docker-clean stop && pkgx task run
```

### Seed data

```
PGPASSWORD=postgrespassword \
  pkgx psql \
    --host localhost \
    --username postgres \
    --port 15432 \
    --no-password \
    --file seed.sql
```
