---
title: Experiment results
slug: /feature-flags/testing-with-flags/experiment-results
description: Learn how to analyze experiment results, understand key metrics, and make data-driven decisions using Bayesian inference.
tags: ['experiments', 'feature-flag', 'test', 'bayesian']
---

import CenteredImg from '@site/src/components/centered-img/CenteredImg';

After creating an experiment, it automatically becomes available on the dashboard **Experiments** tab. Experiments can have one of the following states:

- **Waiting**: The start date hasn't been reached yet. No data is being collected
- **Running**: The experiment is actively collecting data about user behavior according to the defined goals
- **Stopped**: The experiment has been manually stopped or reached its end date

You can access experiment results at any time by clicking the **Result** button, regardless of the experiment's current status.

<CenteredImg
  imgURL="img/experimentation/experiments/experiment-list-running.png"
  alt="List of experiments with different states"
  wSize="650px"
  borderWidth="1px"
/>

You can also check experiment results from the flag details page by selecting the **Experiment** tab.

## Analyzing experiment results

The experiment result page provides comprehensive data to help you make informed decisions:

### Page Components

1. **Experiment header**: Shows the current state (Waiting, Running, Stopped) and evaluation period
2. **Goal selector**: Choose which goal to analyze if you have multiple goals
3. **Evaluation data table**: View user and event counts for each variation
4. **Conversion data table**: Compare performance metrics across variations using Bayesian inference
5. **Data visualization chart**: Track trends over time for any metric

### Evaluation Metrics

<CenteredImg
  imgURL="img/experimentation/experiments/experiment-result-evaluation.png"
  alt="Experiment evaluation metrics table"
  borderWidth="1px"
/>

The evaluation table shows fundamental data for each variation:

- **Evaluation user**: Number of unique users who received this variation from the server
- **Goal total**: Total count of goal event occurrences, including multiple triggers by the same user
- **Goal user**: Number of unique users who fired the goal event (counted once per user)
- **Conversion rate**: Percentage of users who completed the goal (Goal user / Evaluation user)
- **Value total**: Sum of all values assigned to goal events. Used when tracking metrics like revenue or time spent
- **Value/User**: Average value per user (Value total / Goal user)

### Conversion Rate Analysis

When you have sufficient data, Bucketeer displays a confidence indicator showing which variation is winning. This banner appears at the top of the results and provides a quick summary of the experiment's outcome.

<CenteredImg
  imgURL="img/experimentation/experiments/experiment-result-converstion-rate.png"
  alt="Experiment results showing confidence banner, conversion rate chart, and Bayesian analysis table"
  wSize="750px"
  borderWidth="1px"
/>

The conversion rate table uses Bayesian inference to help identify the best-performing variation:

- **Conversion Rate** or **Value/User**: The primary metric being analyzed
- **Improvement**: How much better this variation performs compared to the baseline. Calculated by comparing the range of values for the variation against the baseline range
- **Probability to Beat Baseline**: The estimated likelihood that this variation outperforms the baseline. We recommend a minimum of 95% confidence
- **Probability to Be Best**: The probability that this variation is the top performer among all variations. We recommend a minimum of 95% confidence
- **Expected Loss**: The average opportunity cost of choosing this variation if it's not actually the best. A lower expected loss means less risk of missing out on better performance

:::info Understanding Expected Loss

Expected Loss helps you quantify the risk of choosing a variation. For example:
- Variation A: Expected Loss 2.5% means you might miss out on 2.5% better performance by choosing this
- Variation B: Expected Loss 0.1% means minimal risk - this is likely the best option

Lower expected loss indicates higher confidence that you're making the right choice.

:::

:::tip How to select the best variation

The Bucketeer team recommends selecting a variation with:
- **At least 95% Probability to Beat Baseline**
- **At least 95% Probability to Be Best**
- **Low Expected Loss** (ideally below 1%)

If no variation meets these criteria, continue running the experiment to collect more data.

:::

### Data Visualization

The chart at the bottom allows you to visualize any metric over time. You can:

- Select which metric to display (Conversion Rate, Goal Users, Value/User, etc.)
- Toggle variations on/off to focus your analysis
- Observe trends and patterns throughout the experiment duration

## Making Decisions

Use Bayesian inference results to make data-driven decisions:

1. **Clear winner**: If one variation has >95% probability on both metrics and low expected loss, it's ready to roll out
2. **Needs more time**: If no clear winner emerges, extend the experiment or wait for more data
3. **No significant difference**: If variations perform similarly, consider other factors (implementation complexity, maintenance cost)
4. **Multiple goals**: Compare results across different goals to ensure the winning variation doesn't negatively impact other metrics
