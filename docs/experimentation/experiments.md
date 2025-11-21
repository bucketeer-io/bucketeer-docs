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
- **Baseline variation**: The reference variation used for comparison in the Bayesian inference algorithm
- **Goals**: One or more goals to measure in the experiment. Multiple goals allow you to track different aspects of user behavior
- **Start at**: When the experiment begins collecting data
- **Stop at**: When the experiment stops. Maximum experiment duration is 30 days

<CenteredImg
  imgURL="img/experimentation/experiments/experiment-create.png"
  alt="Experiment creation form"
  wSize="450px"
  borderWidth="1px"
/>

:::tip Recommended test interval
The Bucketeer team recommends running experiments for at least two weeks to collect sufficient data and increase the experiment's reliability.
:::

## When to use multiple goals

Multiple goals help you understand the complete user journey by tracking different stages of user behavior. This allows you to:

- **Identify drop-off points**: See where users lose interest or encounter issues
- **Track conversion funnels**: Measure progression through critical steps
- **Detect side effects**: Ensure winning variations don't negatively impact other metrics

For example, if testing a checkout flow, you might track: "Added to Cart", "Viewed Checkout", and "Completed Purchase". This reveals not just which variation has the best completion rate, but also where users drop off in the process.
