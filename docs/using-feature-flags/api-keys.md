---
title: API keys
# sidebar_position: 
slug: /using-feature-flags/api-keys
description: Describes what API keys are, what they are used for, and how to create them.
tags: ['integration', 'feature-flag']
---

On this page your find information about what are API Keys, how to create them and their role on Bucketeer system.

## What are API Keys

When working with flags in Bucketeer, API keys are essential. They authorize and control access to the feature flag management system, providing authentication capabilities that ensure secure connections between your application and the Bucketeer management service. Each API key is associated with a specific project and environment, allowing you to identify and link your application to the desired user group. By requiring a valid API key for accessing the flag information, the security and integrity of the system is maintained, reducing the risk of unauthorized entry or manipulation with feature flag configurations.

When setting up the SDK on your application, you must include a valid API key. This key is necessary for the server to recognize that the request is valid when the SDK performs a server request. If you provide an invalid API key during configuration, the server will deny the request. For more information on configuring the SDK, refer to the [SDK's documentation](../sdk).

## Creating an API Key

To create an API Key you need to be an Admin or have an enviroment account with the Owner role. Other members can see the APY keys, but can't manage them.


