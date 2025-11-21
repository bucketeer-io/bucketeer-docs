---
title: Goals
# sidebar_position: 
slug: /feature-flags/testing-with-flags/goals
description: Describes how to create goals, what to consider when creating them, and how to use them.
tags: ['goals']
---

import CenteredImg from '@site/src/components/centered-img/CenteredImg';

Goals in Bucketeer are used to measure user behaviors and track the user journey within experiments. By defining goals, you can measure how your feature flags influence specific user actions, such as clicking a button, completing a purchase, or reaching a page.

Goals are essential for both creating and evaluating experiments.

## Create a goal

To create goals in Bucketeer, navigate to the **Goals** tab on the dashboard, which provides an overview of all existing goals within the current environment.

<CenteredImg
  imgURL="img/experimentation/experiments/goal-list.png"
  alt="List of goals in the dashboard"
  wSize="650px"
  borderWidth="1px"
/>

You can search for goals by name, ID, or description. To create a new goal, click the **+ New Goal** button.

### Required Fields

Provide the following information when creating a goal:

- **Goal ID**: A unique identifier used to report user progress in your application
- **Name**: A descriptive name for the goal
- **Description**: Detailed explanation of what the goal measures and why

<CenteredImg
  imgURL="img/experimentation/experiments/goal-create.png"
  alt="Goal creation form"
  wSize="450px"
  borderWidth="1px"
/>

:::tip Description
The Bucketeer team strongly recommends using descriptive names and thoroughly documenting the goal's description to ensure better comprehension of the data during future analyses of the experiments.
:::


### Goals examples

You can define goals for a wide range of scenarios related to your system. To improve your understanding of possible scenarios, the following subsections present goal descriptions related to A/B and multivariable tests.

#### A/B test example

- **Goal ID**: ab_test_contact_button_position
- **Goal Name**: Compare Contact Button Position (Top vs. Bottom)
- **Goal Description**: The goal is to assess the performance and user engagement of two different positions for the contact button on the webpage: positioned at the top or at the bottom. By tracking user interactions with the contact button under both variants, we can determine which position yields higher click-through rates, conversions, and overall user satisfaction. The results will provide insights to inform the optimal placement of the contact button for improved user experience and desired business outcomes.

#### Multivariable example

- **Goal ID**: multivariable_test_cta_performance
- **Goal Name**: Evaluate CTA Performance (Four Defined CTAs)
- **Goal Description**: The goal aims to evaluate the effectiveness and impact of four different Call-to-Action (CTA) variations within the feature. The experiment will measure user interactions, click-through rates, and conversion rates for each CTA variant. By analyzing the data, we can identify which CTA design, wording, or visual elements perform better regarding user engagement, conversions, and desired actions. The results will provide valuable insights to optimize the CTA strategy and enhance overall user engagement and conversion rates.

## How to use goals

To track goals, use the `track` function provided by the Bucketeer SDK. This registers when a user reaches a goal in your defined journey.

### Basic Goal Tracking

For simple events like button clicks or page views, track the goal without a value:

```js showLineNumbers
client.track("GOAL_ID");
```

### Goal Tracking with Values

For goals that have measurable values (like revenue, time spent, or items purchased), pass a numeric value:

```js showLineNumbers
// Track a purchase with the amount spent
client.track("purchase_completed", 49.99);

// Track time spent on a page (in seconds)
client.track("page_engagement", 120);

// Track number of items added to cart
client.track("items_added", 3);
```

The value will appear in the experiment results as:
- **Value total**: Sum of all values across all events
- **Value/User**: Average value per unique user

This is particularly useful for measuring the impact of variations on revenue, engagement time, or other quantifiable metrics.

:::tip When to use values
Use goal values when you want to measure not just *if* users completed an action, but *how much* of something (revenue, time, quantity) resulted from that action.
:::

For further details regarding goal tracking, check the **Reporting customer events** section in the SDK documentation for your programming language.

