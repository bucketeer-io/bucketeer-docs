---
title: API keys
# sidebar_position: 
slug: /using-feature-flags/api-keys
description: Describes what API keys are, what they are used for, and how to create them.
tags: ['integration', 'feature-flag']
---

import CenteredImg from '@site/src/components/centered-img/CenteredImg';

On this page your find information about what are API Keys, how to create them and their role on Bucketeer system.

## What are API Keys

When working with flags in Bucketeer, API keys are essential. They authorize and control access to the feature flag management system, providing authentication capabilities that ensure secure connections between your application and the Bucketeer management service. Each API key is associated with a specific project and environment, allowing you to identify and link your application to the desired user group. By requiring a valid API key for accessing the flag information, the security and integrity of the system is maintained, reducing the risk of unauthorized entry or manipulation with feature flag configurations.

When setting up the SDK on your application, you must include a valid API key. This key is necessary for the server to recognize that the request is valid when the SDK performs a server request. If you provide an invalid API key during configuration, the server will deny the request. For more information on configuring the SDK, refer to the [SDK's documentation](../sdk).

## Creating an API Key

To create an API Key, you need to be an Admin or have an environment account with the Owner role. Other members can see the APY keys but can't manage them.

To generate a new API key in the Bucketeer dashboard, navigate to the **API Keys** tab. This section will present a list of all currently available API keys. If you wish to create a new one, click the **+ Add** button. A side panel will appear on the right side of your screen. Enter a name for the new key, limited to one hundred characters, and finalize the process by clicking **Submit**. The new key will show up at the top of the **API Keys** list.

<CenteredImg
  imgURL="img/using-feature-flags/api-keys/create-api-key.png"
  alt="Account dashboard tab"
  wSize="100%"
/>

The Bucketeer team advises giving each API key a meaningful name based on its associated feature or project to make it easy to manage. Avoid using randomly generated names as they can cause confusion and make it difficult to identify where the API key is being used.

## Manage your API Keys

Once an API key is created, it becomes immediately available for use and is set to the **ON** state. In addition, it grants access to all feature flags within the current environment. The Bucketeer dashboard does not offer a direct option to delete API keys. However, if a key is no longer required, you can toggle its state to **OFF** using the available switch button in the **API Keys** tab. This action will result in denying all SDK requests using that specific key.


