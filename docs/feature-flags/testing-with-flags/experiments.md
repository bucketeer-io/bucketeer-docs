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

To create experiments in Bucketeer, begin by accessing the dashboard. Once in the dashboard, navigate to the **Experiments** tab, which provides an overview of all existing experiments within the current environment.

Suppose you need to locate a particular experiment. In that case, Bucketeer offers a search functionality that allows you to find experiments based on their name or description. On the other hand, if you wish to create a new goal, you can click the **+ Add** button.

To create an experiment, you must fulfill the fields in the creation panel, which becomes visible after selecting **+ Add**. The necessary data when creating an experiment are:

- **Name**: Identify the experiment.
- **Description**: Provide in-depth information regarding the experiment objective.
- **Feature flag**: Defines which flag will be evaluated by the experiment. Each experiment can use only one feature flag. However, you can use a flag in various experiments simultaneously.
- **Baseline variation**: Establishes which variation is used as a reference when running the Bayesian inference algorithm to determine the performance of each variation.
- **Goals**: Sets the goals used in the experiment. You can associate multiple goals with the same experiment.
- **Start at**: Specifies the experiment starting date.
- **Stop at**: Specifies the experiment finishing date. The maximum interval to run an experiment is 30 days.

:::tip Recommended test interval
The Bucketeer team recommends running experiments for at least two weeks to group sufficient data and increase the experiment's reliability.
:::

The image below presents an example of experiment creation.

<CenteredImg
  imgURL="img/feature-flags/experiments/create-experiment.png"
  alt="create experiment example"
  wSize="400px"
  borderWidth="1px"
/>

## When use multiple goals

Multiple goals are used to evaluate the user journey, identifying behavior patterns and specific areas where users may encounter difficulties or drop off. Comparing the number of users reaching each goal helps detect problematic transitions related to page or critical process points. Multiple goals allow a comprehensive user flow analysis, highlighting potential bottlenecks or obstacles, and tracking progression helps identify where users lose interest, encounter issues, or abandon the process. You will have access to these data when evaluating the experiment results.
