---
title: Bucketeer vocabulary
# sidebar_position: 
slug: /getting-started/bucketeer-vocabulary
description: This page guides the user through the main terms used on Bucketeer's documentation.
tags: ['jargon', 'vocabulary', 'definitions', 'glossary']
---

On this page, you will find a list of essential technical terms commonly used in the Bucketeer solution and documentation. This vocabulary list goal is to help you better understand Bucketeer. If you cannot find a particular vocabulary, feel free to [contact us](https://app.slack.com/client/T08PSQ7BQ/C043026BME1), and we will add it to the list for future reference.

**Environment**: An environment represents an organizational unit within a project. It corresponds to the deployment environments of your code, such as development, staging, and production. Each project can have multiple environments containing the same flags with potentially different states, targets, and rules.

**Evaluation**: Evaluation refers to the process where the Bucketeer SDK receives information about a flag and its associated context from your application's code. The SDK then returns the appropriate flag variation for that context, allowing the flag to be evaluated for a specific customer or context.

**Experiment**: Experiments in Bucketeer are crucial for measuring the impact of a chosen feature on selected users. An experiment combines a feature flag, defined goals, and user segments to analyze how the feature contributes to achieving those goals.

**Feature**: A feature refers to a specification or element of your application, such as a button, picture, or more complex components like a purchase flow or recommendation algorithm.

**Flag**: A flag is the fundamental unit of feature management in Bucketeer. It represents different feature variations and includes rules that determine which entities can access each variation. Entities can be a percentage of your application's traffic, individuals, or software entities with common characteristics.

**Goal**: Goals are important metrics for tracking and making decisions based on certain indexes. They represent conversions or business metrics, such as clicks, page views, or user engagement on a website.

**Iteration**: An iteration is a defined period during which an experiment runs. It can have any length, and you can run multiple iterations of the same experiment to gather data and analyze the impact of a feature.

**Member**: A member, also referred to as an account member, is a person who uses Bucketeer within your organization. They can be employees, contractors, or individuals with access rights to your Bucketeer environment.

**Percentage Rollout**: Percentage rollout is a targeting rule that gradually exposes a specific flag variation to a specified percentage of contexts. This rule allows you to incrementally increase the share of customers targeted by a flag until 100% receive a particular variation.

**Project**: A project represents an organizational unit for flags in your Bucketeer account. It allows you to organize flags based on your desired structure, such as creating a project for each product your company develops. Projects can contain multiple environments.

**Role**: A role defines the level of access a member has within Bucketeer. Built-in roles include Viewer, Editor, Owner, and Admin, each providing different levels of access to information within Bucketeer.

**SDK**: The Bucketeer SDK (Software Development Kit) is a set of tools and libraries to integrate Bucketeer with your application's code. It allows you to initialize the SDK, evaluate feature flags, and retrieve the appropriate flag variation for each customer.

**User Segment**: User segments are defined groups of users used for comparison and testing purposes. In Bucketeer, you can configure user segments in different ways, allowing you to offer different user experiences and test beta features with a subset of users.

**Variation**: Feature flags enable you to manage the visibility of specific features in your application, allowing you to activate them for specific users or roll them out to a percentage of your traffic. These feature flags consist of variations representing the different options or configurations available for the feature. In Bucketeer, a variation represents a possible value for a flag, defining the various options and states of a feature. Variations can be binary (e.g., true/false) for boolean flags or have multiple options for multivariate flags, each with a value of the same type, such as a string, and potentially accompanied by a name and description. 