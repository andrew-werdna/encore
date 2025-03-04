---
title: Migrating an existing backend to Encore
subtitle: Approaches for adopting Encore
seotitle: How to migrate your existing backend to Encore
seodesc: Learn how to migrate your Go backend application to Encore, and unlock Encore's powerful set of backend development tools for your team.
lang: platform
---

Encore features like [automatic infrastructure provisioning](/docs/platform/infrastructure/infra), [distributed tracing](/docs/ts/observability/tracing), [architecture diagrams](/docs/ts/observability/encore-flow), and [Service Catalog](/docs/ts/observability/service-catalog), rely on the [Encore application model](/docs/ts/concepts/application-model).

Building your backend using Encore's declarative backend frameworks [Encore.ts](/docs/ts) and [Encore.go](/docs/go) is what enables Encore to create the application model. This doesn't mean a complete rewrite is necessary to adopt Encore, and in this guide we look at strategies for both incremental adoption and fully migrating your existing backend to Encore.

## Get help with adopting Encore

If you'd like to ask questions or get advice about how to approach adopting Encore, we're happy to talk through your project. You can [join Discord](https://encore.dev/discord) to ask questions and meet other Encore developers, or you can also [book a 1:1](/book) with a member of our core team.

## Incremental adoption: Build or refactor a single service

We recommend using Encore to build a single service to validate if it would work well for your organization. This could be a new, relatively independent project, or a current service or system that needs refactoring. With this approach you can use all Encore features from the start, and then incrementally migrate more services over time.

Your Encore application will talk to your existing backend through APIs, and can be provisioned in your existing cloud account as pictured below.

<img src="/assets/docs/incremental.png" title="Incremental migration" className="noshadow"/>

### 1. Create an Encore app and integrate with GitHub

The first step in any project is to create an Encore app. If you've not tried Encore before, we recommend starting by following the [Quick Start Guide](/docs/ts/quick-start).

Once you've created you app, [integrate it with your GitHub repository](/docs/platform/integrations/github) and you'll get automatic [Preview Environments](/docs/platform/deploy/preview-environments) for every Pull Request.

### 2. Prototype your new backend system

Once you've created your app, it's time to start building. If you're new to Encore, we recommend trying out some of the tutorials.

If you need help or have questions, join the friendly developer community on [Discord](https://encore.dev/discord).

#### Design your APIs

Since Encore is designed to build distributed systems, it should be straightforward to build a new system that integrates with your existing backend through APIs. See the [defining APIs documentation](/docs/ts/primitives/defining-apis) for more details.

Should you want to accept webhooks, that's simple to do using Encore's [Raw endpoints](/docs/ts/primitives/raw-endpoints).

You can also generate API clients in several languages, which makes it simple to integrate with frontends or other systems. See the [Client Generation documentation](/docs/ts/cli/client-generation) for more details.

#### Storing Secrets

When you need to store secrets, you can use Encore's built-in [secrets manager](/docs/ts/primitives/secrets).
It lets you store and manage secrets for all environments, and will automatically provision a secret manager in your cloud account once you deploy to production.

#### Connect to an existing database

When you create an Encore service and add a [database](/docs/ts/primitives/databases) to it, Encore will automatically provision the necessary infrastructure for you. When migrating, it's common to also want to connect to an existing database. [See this guide](/docs/go/primitives/connect-existing-db) for instructions on how to do that with Encore.

### 3. Deploy to your cloud account

Once you're ready to deploy, you can [connect your cloud account](/docs/platform/deploy/own-cloud) (GCP or AWS) and have Encore deploy and provision your new system directly in your existing account.

See the [infrastructure documentation](/docs/platform/infrastructure/infra#production-infrastructure) for details on how Encore provisions cloud infrastructure for each cloud provider.

### Rinse and repeat

Once you're confident that Encore is a good fit, you can use this incremental migration strategy to move more services over to Encore. This will make Encore benefits like automatic provisioning, preview environments, architecture diagrams, and distributed tracing available for more parts of your system.

## Forklift migration using a catch-all handler

If your existing backend system is built with Go, you can use a forklift migration strategy and move the entire application over to Encore in one shot by wrapping your existing HTTP router in a catch-all handler.

This can be relatively straightforward if your existing system is a monolith, or smaller distributed system, and does not rely on many unsupported cloud primitives.

The benefits of this approach is that you'll get everything in one place from the start, and you'll be able to quickly use features like Encore's CI/CD system and secrets manager, for your entire backend application.

However, you will not immediately be able to use some of the powerful Encore features, like [distributed tracing](/docs/ts/observability/tracing) and [architecture diagrams](/docs/ts/observability/encore-flow), which rely on the [Encore application model](/docs/ts/concepts/application-model).

Once your Encore app is up and running, you'll have something that looks like the image below. Notice how Encore doesn't have complete visibility into the inner workings of your application.

<img src="/assets/docs/forklift.png" title="Forklift migration" className="noshadow"/>

### 1. Create an app and structure your code

To start, create an Encore application and copy over the code from your existing repository. In order to run your application with Encore, it needs to follow the expected [application structure](/docs/go/primitives/app-structure), which involves placing the `encore.app` and `go.mod` files in the repository root. This should be straightforward to do with minor modifications.

As an example, a single service application might look like this on disk:

```
/my-app
├── encore.app                       // ... and other top-level project files
│
└── hello                            // hello service (a Go package)
    ├── migrations                   // hello service db migrations (directory)
    │   └── 1_create_table.up.sql    // hello service db migration
    ├── hello.go                     // hello service code
    └── hello_test.go                // tests for hello service
```
You can also have services nested inside a `backend` folder if you prefer.

### 2. Create a catch-all handler for your HTTP router

Now let's mount your existing HTTP router under a [Raw endpoint](/docs/go/primitives/raw-endpoints), which is an Encore API endpoint type that gives you access to the underlying HTTP request.

Here's a basic code example:

```
//encore:api public raw path=/api/*path
func MigrationHandler(w http.ResponseWriter, req *http.Request) {
// pass request to existing router
}
```

By mounting your existing HTTP router in this way, it will work as a catch-all handler for all HTTP requests and responses. This should make your application deployable through Encore with little refactoring. 

### 3. Iteratively fix remaining compilation errors

Exactly what remains to make your application deployable with Encore will depend on your specific app.
As you run your app locally, using `encore run`, Encore will parse and compile it, and give you compilation errors to inform what needs to be adjusted.

By iteratively making adjustments, you should relatively quickly be able to get your application up and running with Encore.

If you need help or have questions, join the developer community on [Discord](https://encore.dev/discord) or send an email to [help@encore.dev](mailto:help@encore.dev).

### Incrementally start using Encore's Backend Framework

Once your application is deployed, gradually break out specific endpoints using the Encore's [API declarations](/docs/ts/primitives/defining-apis) and introduce infrastructure declarations using the Encore backend frameworks. This will let Encore understand your application and unlock all Encore features.