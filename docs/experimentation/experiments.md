---
title: Experiments
# sidebar_position: 
slug: /feature-flags/testing-with-flags/experiments
description: Describes how to create experiments, what to consider when creating them, and how to use them.
tags: ['test', 'experiments', 'A/B']
---

import CenteredImg from '@site/src/components/centered-img/CenteredImg';

Experiments in the context of feature flags are a way to test new features or changes to existing features with a subset of users before rolling them out to everyone. Experiments allow you to collect data and feedback on the feature before making it available to everyone, which can help you to improve the feature and make sure that it's successful and effective.

There are a few different ways to run experiments with feature flags. Common approaches are A/B and multivariate testing. Bucketeer enables you to create all types of experiments.

## Create an experiment

To create experiments in Bucketeer, navigate to the **Experiments** tab on the dashboard, which provides an overview of all existing experiments within the current environment.

<CenteredImg
  imgURL="img/experimentation/experiments/experiment-list-waiting.png"
  alt="Experiment list showing different experiment states"
  wSize="650px"
  borderWidth="1px"
/>

You can search for experiments based on their name or description. To create a new experiment, click the **+ New Experiment** button.

### Required Fields

When creating an experiment, you need to provide:

- **Name**: A unique identifier for the experiment
- **Description**: Detailed information about the experiment's objective
- **Feature flag**: The flag that will be evaluated by the experiment. Each experiment uses one feature flag, but a flag can be used in multiple experiments
- **Baseline variation**: The reference variation (control group) used for comparison in the Bayesian inference algorithm
- **Goals**: One or more goals to measure in the experiment. The first goal you add is the **primary goal**, and the rest are **secondary goals** (see below). Multiple goals allow you to track different aspects of user behavior
- **Start at**: When the experiment begins collecting data
- **Stop at**: When the experiment stops. Maximum experiment duration is 30 days

<CenteredImg
  imgURL="img/experimentation/experiments/experiment-create.png"
  alt="Experiment creation form"
  wSize="450px"
  borderWidth="1px"
/>

After you select your goals, they appear in a list. The first goal is marked with a **Primary** badge. The primary goal is the metric that decides the experiment — it drives the "Ready to roll out" recommendation on the results page — while the others are tracked for insight only. You can change which goal is primary by adjusting the order in which goals are selected.

:::tip Recommended test interval
The Bucketeer team recommends running experiments for at least two weeks to collect sufficient data and increase the experiment's reliability.
:::

## Primary and secondary goals

When an experiment has more than one goal, exactly one of them drives the decision:

- **Primary goal**: The single metric the experiment is designed to improve (sometimes called the Overall Evaluation Criterion). On the results page it carries a **Primary** badge, and only this goal shows the "Ready to roll out" recommendation and the rollout action. The first goal you attach is the primary goal.
- **Secondary goals**: All other goals. Their full results are still computed and displayed so you can learn from them, but they are informational only and do not drive the ship decision.

This mirrors how leading experimentation platforms work, and it protects you from the **multiple-comparisons problem**: the more metrics you check, the more likely one looks like a winner purely by chance. Deciding on a single primary metric up front keeps your conclusions trustworthy.

:::tip Choose your primary goal before you start
Decide which metric matters most before running the experiment, and add it as the first goal. Looking across many metrics after the fact and picking the one that "won" is a common way to fool yourself.
:::

## When to use multiple goals

Multiple goals help you understand the complete user journey by tracking different stages of user behavior. This allows you to:

- **Identify drop-off points**: See where users lose interest or encounter issues
- **Track conversion funnels**: Measure progression through critical steps
- **Detect side effects**: Ensure winning variations don't negatively impact other metrics (use secondary goals as guardrails)

For example, if testing a checkout flow, you might set "Completed Purchase" as the primary goal and track "Added to Cart" and "Viewed Checkout" as secondary goals. This reveals not just which variation has the best completion rate, but also where users drop off in the process.
