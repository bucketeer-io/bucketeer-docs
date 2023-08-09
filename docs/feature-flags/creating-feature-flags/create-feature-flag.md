---
title: Create a feature flag
# sidebar_position: 
slug: /feature-flags/creating-feature-flags/create-feature-flag
description: Presents how to create a feature flag. The page will show the feature flag tab and cover the fields required to create a feature flag.
tags: ['create','guide','feature-flag']
---

import CenteredImg from '@site/src/components/centered-img/CenteredImg';

The first step to creating feature flags on Bucketeer is to access the dashboard. If your team runs your Bucketeer systems, ask your Admin to provide the dashboard URL. After accessing the dashboard, navigate to the **Feature Flags** tab, where you will find all existing flags in the current environment. Creating a new flag adds it to the list.

If you're looking for a specific flag, you can search it using its name, description, or ID. On the other hand, if you want to create a new flag, click the **+ Add**.

## Create a feature flag

To create the flag, you need to fulfill the fields on the creation panel, which appear after clicking the **+ Add**. The image below presents an example of the creation panel. The first three fields, **ID**, **Name**, and **Description**, are used for identification purposes. All have a length limit of 100 characters. While **Name** and **Description** are used only to identify and help find the flag after its creation, the **ID** is also used to determine the flag when integrating it into your application. In the [SDKs](/sdk) section, the **ID** is also referenced as `FLAG_ID`.

<CenteredImg
  imgURL="img/getting-started/quickstart/create-feature-flag.png"
  alt="create feature flag panel"
  wSize="350px"
  borderWidth="1px"
/>

You also need to define the **Tags** related to the new flag. You can specify any number of tags. However, you need to provide at least one value. Tags have two functions on Buketeer. First, then help while searching flags on the dashboard. Second, they help filter the information the server returns after a request from an SDK. Therefore, providing more specific tags helps reduce the amount of information transferred from the server to your application, reducing response time and the amount of data transferred. For additional information regarding the optimization of Bucketeer based on the use of tags, access the [Optimize Bucketeer with tags](/best-practices/optimize-with-tags) guide. It's important to notice that during the integration and in the [SDKs](/sdk) section, tags are also referenced as `featureTag`.

Once you have defined the **Tags** for your flag, the next step is to select the **Flag type**. Refer to the [Feature flags types](/feature-flags/creating-feature-flags/create-feature-flag#feature-flags-types) to understand the differences among the four available types. If you choose the boolean type, only two variations will be available, `true`and `false`, with their respective predefined values. On the other hand, if you select a different type, you can add more variations by clicking **Add variation** button. Additionally, you need to provide the corresponding value for each variation. While the name and description for variations are optional, we highly recommend using them as they facilitate understanding and future review of your results.

Select the default variations for the **ON** and **OFF** states to finalize the flag creation process. This is especially important to ensure which variation will be returned by the flag if no targeting is defined. Once you have defined the **ON** and **OFF** variations, click **Submit** to create the flag.

After creating the new flag, you will be redirected to the Targeting page. For further information regarding flag targeting, refer to the [Targeting with feature flags](/feature-flags/creating-feature-flags/targeting) page.

## Manage feature flags

Once you have created a flag, it will automatically appear on the **Feature Flags** page with its initial state set to **OFF**. Consequently, if an SDK request is made for this flag, the Bucketeer system will return the variation associated with the **OFF** state. To change the state, use the switch button. In addition to the state, you will also find the flag **ID** and related tags displayed on the flag card. Taking the flag card example below, its **ID** is **feature-javascript-e2e-boolean**, associated with the **javascript** and **web** tags.

<CenteredImg
  imgURL="img/getting-started/quickstart/created-feature-flag.png"
  alt="created feature flag"
  borderWidth="1px"
/>

You can duplicate and archive flags as well. To perform these actions, click the three-dot button on the desired flag.

:::tip

When a flag is archived, any SDK requests related to that flag will return the default value, typically the value associated with the **OFF** state. To ensure proper functionality, we recommend removing flag evaluations from your code when archiving the flag in the Bucketeer system.

:::

## Feature flags types

The Bucketeer system provides four types of flags: boolean, string, number, and JSON. Each type serves a specific purpose to accommodate different scenarios when implementing flags. Here's an overview of when and why you should use each type:

- **Boolean flags** are ideal for simple on/off scenarios. They allow you to enable or disable a feature based on a binary state. Use a boolean flag to control the visibility of a beta feature in your application, such as a button, toggling it on or off for specific users or groups.
- **String flags** are suitable when defining multiple variations or states for a feature. They provide flexibility in categorizing and managing feature variations. You could use a string flag to define your application's visual themes or layouts. Thus you can find the best performing style, for example.
- **Number flags** are helpful when assigning numeric values to feature variations, enabling more granular control over feature behavior. They can be used to adjust the speed or intensity of an animation or to define different levels of access or permissions within a feature.
- **JSON flags** offer the most flexibility and complexity, allowing you to define custom configurations and structures for your feature variations using JSON objects. You might employ a JSON flag to dynamically change the layout and content of a specific section in your application based on complex business, defining more than one modification related to the same flag variation.
