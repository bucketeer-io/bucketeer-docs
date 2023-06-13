---
title: Integrate Bucketeers
# sidebar_position: 
slug: /getting-started/integrate-bucketeers
description: Describes the steps to integrate the user application with Bucketeer.
tags: ['guide','integration','sdk']
---

Integrating Bucketeer with your application is a simple process that is the same across all SDKs. Bucketeer has multiple SDKs available, and you can refer to the [Choose an SDK](./choose-sdk.md) guide to determine which one best suits your needs. To integrate Bucketeer into your application, follow these steps:

1. **Install the Bucketeer SDK**: Use your project's dependency manager to install the Bucketeer SDK in your application, allowing your application to access the Bucketeer SDK and its features. 

2. **Import the Bucketeer Client**: In your application's code, import the Bucketeer client, the primary interface to interact with the Bucketeer SDK, and communicate with the Bucketeer service. This step and the previous one are covered in details on each SDK guide. Access the [SDKs](../sdk/) page to get more information.

3. **Configure the Bucketeer Client**: Provide the credentials for your environment to configure the Bucketeer client, including setting up the API key and specifying the endpoint URL. These credentials uniquely identify your project and environment and authorize your application to connect with Bucketeer. Check the [API keys](../using-feature-flags/api-keys.md) to learn how to access your credentials.

4. **Assign Feature Flag Variations**: Use the feature flag ID to associate specific flag variations with different users. Each feature flag has a ID, and by using it, you can control which variation of the flag a user will see in your application.

By following these steps, you can successfully integrate Bucketeer into your application and take advantage of its powerful feature flag management capabilities. If you run into any problems during the integration process or need further assistance, contact the Bucketeer team.
