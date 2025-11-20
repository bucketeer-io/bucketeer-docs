---
title: Create a feature flag
# sidebar_position: 
slug: /feature-flags/creating-feature-flags/create-feature-flag
description: Presents how to create a feature flag. The page will show the feature flag tab and cover the fields required to create a feature flag.
tags: ['create','guide','feature-flag']
---

import CenteredImg from '@site/src/components/centered-img/CenteredImg';

The first step to creating feature flags on Bucketeer is to access the dashboard. If your team runs your Bucketeer systems, ask your Admin to provide the dashboard URL. After accessing the dashboard, navigate to the **Feature Flags** tab, where you will find all existing flags in the current environment. Creating a new flag adds it to the list.

If you're looking for a specific flag, you can search it using its name, description, or ID. On the other hand, if you want to create a new flag, click the **+ Create Flag** button.

## Create a feature flag

To create the flag, you need to fulfill the fields on the creation panel, which appear after clicking the **+ Create Flag**. The first fields, **Flag ID**, **Name**, and **Description**, are used for identification purposes. All have a length limit of 100 characters. While **Name** and **Description** are used only to identify and help find the flag after its creation, the **Flag ID** is also used to determine the flag when integrating it into your application. In the [SDKs](/sdk) section, the **Flag ID** is referenced as `FLAG_ID`.

<CenteredImg
  imgURL="img/getting-started/quickstart/create-feature-flag.png"
  alt="create feature flag panel"
  borderWidth="1px"
/>

Configure the following fields to create your flag:

### Tags

Define at least one tag for your flag. Tags serve two purposes:
1. Search and filter flags on the dashboard
2. Filter SDK responses to reduce data transfer

Providing more specific tags reduces the information transferred from server to application, improving response time and reducing bandwidth usage.

:::note
Tags are referenced as `featureTag` in the SDKs. Learn more: [Optimize with tags](/best-practices/optimize-with-tags).
:::

### Flag Type

Select from five types: boolean, string, number, JSON, or YAML. See [Feature flags types](/feature-flags/creating-feature-flags/create-feature-flag#feature-flags-types) for guidance on which to choose.

### Variations

- **Boolean flags**: Automatically include `true` and `false` variations
- **Other types**: Use **Add Variation** to create custom variations

For each variation, provide:
- **Value** (required)
- **Name** (required)
- **Description** (optional, recommended for documentation)

### Default States

Set which variation to return for:
- **ON variation**: When the flag is enabled
- **OFF variation**: When the flag is disabled

Click **Create Flag** to finish. You'll be redirected to the Targeting page to configure [targeting rules](/feature-flags/creating-feature-flags/targeting).

## Manage feature flags

Once you have created a flag, it will automatically appear on the **Feature Flags** page with its initial state set to **OFF**. Consequently, if an SDK request is made for this flag, the Bucketeer system will return the variation associated with the **OFF** state. To change the state, use the switch button. In addition to the state, you will also find the **Flag ID** and related tags displayed on the flag card. Taking the flag card example below, its **Flag ID** is **feature-javascript-e2e-boolean**, associated with the **javascript** and **web** tags.

<CenteredImg
  imgURL="img/getting-started/quickstart/created-feature-flag.png"
  alt="created feature flag"
  borderWidth="1px"
/>

You can clone and archive flags as well. To perform these actions, click the three-dot button on the desired flag.

:::info Feature flag status

Feature flags may have one of the following three statuses:

- **Never Used**: A recently created flag that hasn't yet received any request.
- **Receiving requests**: Active flags that have received requests within the last 7 days.
- **No Recent Traffic**: Flags that haven't received requests for over 7 days. It's important to note that both ON and OFF flags can become inactive.

The flag status helps identify the lifecycle stage of your flag, indicating if it's time to archive it. Access the [feature flag lifecycle page](/best-practices/feature-flag-lifecycle) to learn more about the best practices for the flag's usage.

:::

:::tip

When a flag is archived, any SDK requests related to that flag will return the default value, typically the value associated with the **OFF** state. To ensure proper functionality, we recommend removing flag evaluations from your code when archiving the flag in the Bucketeer system.

:::

## Feature flags types

The Bucketeer system provides five types of flags: boolean, string, number, JSON, and YAML. Each type serves a specific purpose to accommodate different scenarios when implementing flags. Here's an overview of when and why you should use each type:

- **Boolean flags** are ideal for simple on/off scenarios. They allow you to enable or disable a feature based on a binary state. Use a boolean flag to control the visibility of a beta feature in your application, such as a button, toggling it on or off for specific users or groups.
- **String flags** are suitable when defining multiple variations or states for a feature. They provide flexibility in categorizing and managing feature variations. You could use a string flag to define your application's visual themes or layouts. Thus you can find the best performing style, for example.
- **Number flags** are helpful when assigning numeric values to feature variations, enabling more granular control over feature behavior. They can be used to adjust the speed or intensity of an animation or to define different levels of access or permissions within a feature.
- **JSON flags** offer the most flexibility and complexity, allowing you to define custom configurations and structures for your feature variations using JSON objects. You might employ a JSON flag to dynamically change the layout and content of a specific section in your application based on complex business logic, defining more than one modification related to the same flag variation.
- **YAML flags** provide a more readable alternative to JSON for complex configurations. YAML values are compiled into a JSON structure before user evaluation and are returned to the client SDK through the object variation interface, just like JSON flags. This allows you to write more human-friendly configuration while maintaining full compatibility with JSON flag functionality.

:::info JSON and YAML Compatibility
Both JSON and YAML flags are returned to the client SDK through the same object variation interface. YAML is compiled into JSON format during evaluation, making these types fully interchangeable from the SDK's perspective. Choose YAML when readability is important during flag configuration, or JSON when you prefer the more compact format.
:::
