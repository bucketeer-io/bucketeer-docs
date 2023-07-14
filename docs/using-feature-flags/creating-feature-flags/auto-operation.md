---
title: Auto operation 
# sidebar_position: 
slug: /using-feature-flags/auto-operation
description: The Auto Operation feature handles the manipulation of feature flag configurations on your behalf. It executes specific operations automatically based on the conditions set by the user.
tags: ['guide','automation', 'operation','feature-flag']
---

import CenteredImg from '@site/src/components/centered-img/CenteredImg';

The auto operation feature handles the manipulation of feature flag configurations on your behalf. It executes specific operations automatically based on the conditions set by the user.

To access the auto operation page on the Bucketeer dashboard, select the **Feature Flags** tab, choose the desired flag, and click on its name. Select the **Auto Ops Rules** tab at the top of the new page. You can check the active and completed operations on the Auto Ops Rules page.

## How auto operation works

The auto operation feature consists of two key elements:

1. **Operation**: An operation is an action that is triggered by the condition. The available operations are:
   - *Enable feature*: Activates the flag.
   - *Kill feature*: Desables the flag.

2. **Condition**: A condition specifies the circumstances under which the operation is triggered. Multiple conditions can be associated with a single killing operation. On the other hand, only one condition can be related to enabling a feature. It's important to note that the corresponding operation is triggered if any conditions are satisfied.

After a condition is triggered, the related operation changes from active to complete. As a result, the operation is moved from the active panel to the completed panel.

## Conditions

The following conditions are available for defining auto operations:

1. **Schedule**: This condition evaluates whether the specified date and time have been reached. You must always use a future date and time to schedule the operation. It is useful, for example, to release a feature at a specific date and time or to turn off a beta feature after a certain period.

<CenteredImg
  imgURL="img/using-feature-flags/auto-ops-rules/schedule-condition.png"
  alt="Event rate panel"
  wSize="500px"
/>

:::info Limit of schedule conditions

You can only have two active operations based on schedule conditions. One will define when to enable and the other when to kill the flag.

:::

2. **Event Rate**: This condition triggers the operation based on the event rate reaching a specified threshold. The event rate is only available for killing operations. It allows you to automatically turn off a flag if it is causing a high error rate or turn it off once enough data has been collected for analysis.

<CenteredImg
  imgURL="img/using-feature-flags/auto-ops-rules/event-rate-condition.png"
  alt="Event rate panel"
  wSize="500px"
/>

Below, you will find a description of the configuration options for setting up the event rate condition:

- **Variation**: The specific variation observed by the target user.
- **Percentage**: The threshold for the event rate.
- **Operator**: The comparison operator used to assess the event rate. The available options are greater than or equal to (>=) or smaller than or equal to (<=).
- **Minimum Count**: Minimum number of occurrences required for the condition to be enabled.

## How the event rate is calculated

The event rate is calculated using the following formula:

**Event Rate = (Unique Counter of Goal Events) / (Unique Counter of Evaluation Events)**

Refer to the experiment and evaluation count section for a better understanding of the user count of goals and evaluations.

## Setting up auto operation

Configuring auto operations is a straightforward process. Follow these steps:

1. Go to the Auto Ops Rules page.
2. Click the **Add operation** button.
3. The operation setting panel will appear. Choose the desired operation. For this example, the **Enable feature** option was selected.
4. Set up the desired condition. For this example **Schedule** was selected.
5. Click the **Submit** button to create the auto operation rule. The image below summarizes the configuration described in the previous steps.

<CenteredImg
  imgURL="img/using-feature-flags/auto-ops-rules/create-condition.png"
  alt="Event rate panel"
  wSize="500px"
/>

When the scheduled time arrives, the feature flag will be automatically enabled.