---
title: Goals
# sidebar_position: 
slug: /feature-flags/testing-with-flags/goals
description: Describes how to create goals, what to consider when creating them, and how to use them.
tags: ['goals']
---

import CenteredImg from '@site/src/components/centered-img/CenteredImg';

The goal feature in the Bucketeer system is designed to facilitate the organization and analysis of experiments. Goals serve as a means to group end-users' data based on specific behaviors you intend to monitor. However, the main objective of using goals is to track the user journey. By defining goals, you can effectively measure user behaviors that are influenced by your feature flags within experiments. Goals play a pivotal role in both the creation and evaluation of experiments.

## Create a goal

To create goals in Bucketeer, begin by accessing the dashboard. Once in the dashboard, navigate to the **Goals** tab, which provides an overview of all existing goals within the current environment. 

Suppose you need to locate a particular goal. In that case, Bucketeer offers a search functionality that allows you to find goals based on their name, ID, or description. On the other hand, if you wish to create a new goal, you can click the **+ Add** button.

To create a goal, you must fulfill the fields in the creation panel, which becomes visible after selecting **+ Add**. It's essential to provide the goal's ID, name, and description. You will use the goal's ID to report the end-users progress on your application. The image below presents an example of the creation panel.

<CenteredImg
  imgURL="img/feature-flags/goals/create-goal-v2.png"
  alt="create goal example"
  wSize="400px"
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

To make goals data relevant, you need to report the end-user journey. You will use the reporting solutions provided by the Bucketeer SDK to achieve it. To report goal achievements, you will use the SDK `track` function, which enables you to register when a client (end-user) reaches a goal from the journey you defined. The code block below presents an example of registering the goal achievement by the client using the `track` function in Javascript.

```js showLineNumbers
client.track("GOAL_ID");
```

For further details regarding goals tracking, check the **Reporting customer events** section related to the programing language you use on the SDK documentation.

