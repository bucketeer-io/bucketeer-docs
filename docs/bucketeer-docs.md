---
title: Bucketeer Docs
sidebar_position: 1
slug: /
description: Describes what the Bucketeer is and its solution. In addition, the page also provides an overview of the main sections covered in the documentation.
tags: ['home', 'guide', 'presentation', 'overview', 'contact']
---

import Button from '@site/src/components/button/Button';
import ButtonShelf from '@site/src/components/button-shelf/ButtonShelf';

# Welcome to Bucketeer docs

Welcome to the Bucketeer documentation site, your comprehensive resource for everything related to the Bucketeer platform, integrations, and SDKs. You will find all the information you need to utilize the Bucketeer solution effectively here. Utilize the convenient navigation buttons below to access the desired sections quickly. Furthermore, this page will provide an in-depth overview of Bucketeer, highlighting its key features and capabilities.

<ButtonShelf>
  <Button
    redirect="../getting-started"
    title="Getting Started"
    info="Provides a quickstart guide on how to use the Bucketeer solution, covering flag creating and integration to your system."
  />
  <Button
    redirect="../feature-flags"
    title="Feature Flgs"
    info="Explain how to create and manage feature flags and execute tests using them."
  />
  <Button
    redirect="../sdk"
    title="SDK"
    info="Describes in detail how to integrate the Bucketeer SDK into your system, including server and client applications."
  />
  <Button
    redirect="../contribution-guide/contributing"
    title="Contribution Guide"
    info="Presents the process of contributing to the Bucketeer system and includes a style guide for those creating the documentation."
  />
</ButtonShelf>

:::info

Due to the renewal of the admin console, we are also renewing the documentation website.

We will update it weekly. Stay tuned!

:::

## What is Bucketeer

Bucketeer is a feature management platform that enables organizations to optimize software development by providing powerful tools. It offers feature flags for controlled feature releases, allowing easy experimentation, A/B testing, and targeted user segmentation. The platform supports beta testing programs, facilitates dynamic configuration changes, and enables dark launches and rolling out releases to specific user groups. With Bucketeer, teams can efficiently manage feature lifecycles, gather valuable user feedback, and make data-driven decisions to enhance the overall user experience.

## What you can do with the Bucketeer solution

### A/B Testing

Bucketeer offers a powerful A/B testing solution beyond superficial changes, allowing teams to test substantial functionality. With the feature management platform, organizations can establish goals and effectively measure the impact of different features.

Check how to create [A/B tests](./feature-flags/testing-with-flags/experiments) on Bucketeer.

### User Targeting

Unlock the potential of highly customized metadata to achieve precise and effective user targeting. By utilizing attributes such as region, age, or email, you can create segments or groups tailored to specific criteria. This level of granularity provides complete control over determining who sees what, empowering you to deliver personalized experiences.

Explore how [targeting works](./feature-flags/creating-feature-flags/targeting) within Bucketeer.

### Prerequisites

Use prerequisites to establish dependencies between flags and improve user targeting. Define the conditions or dependencies between flags, which will determine whether a flag should be evaluated or which variation should be provided to the end user.

Learn more about [Bucketeer prerequisites functionality](./feature-flags/creating-feature-flags/targeting#prerequisites).

### Kill Switch

No need to revert code to undo a feature release—Bucketeer allows anyone to deactivate any feature at any time instantly. Leveraging feature flags as a kill switch is a common practice for mitigating risks associated with feature releases. With this functionality, product and marketing teams can participate in feature testing and releases without heavy reliance on engineering support. Whether you're conducting tests in production, implementing a canary launch, or preparing to retire a feature, rest easy knowing that it's as simple as flipping a switch.

Check how [kill switch](./feature-flags/creating-feature-flags/auto-operation#how-auto-operation-works) works on Bucketeer.

### Beta Testing

Bucketeer's feature management platform is widely used for efficiently managing beta testing programs at scale. Creating specific groups based on targeting rules allows you to include entire groups or segments in tests as needed. The feedback from beta groups is invaluable, enabling you to validate new features and identify and address bugs before rolling them out to your entire user base.

Explore how to use [user groups](./feature-flags/creating-feature-flags/targeting#targeting) when targeting with Bucketeer.

### Dynamic Configuration

Changing configurations for released products can be an intimidating task. However, by employing feature flags, you can easily configure applications through a user-friendly interface, modifying values effortlessly. This capability is not limited to engineers alone. It also extends to business teams, allowing them to make changes according to their requirements.

### Dark Launch

With Bucketeer, you can introduce new features to the production environment without impacting anything by keeping the feature flags off. When the time is right, you can easily enable these features, ensuring a smooth transition.

### Rolling Out Releases

Effortlessly roll out new features to a subset of users, such as a dedicated beta tester group, and gather valuable feedback and bug reports from real-world usage scenarios. This allows you to refine and optimize features based on user experiences before launching them to a broader audience.

Learn how to create [roll out releases](./feature-flags/creating-feature-flags/targeting#rollout-percentage) on Bucketeer.

### Trunk Base Development

By turning off feature flags and merging feature branches more frequently, teams can reduce the number of work-in-progress branches. This approach alleviates reviewers' stress and apprehension regarding "big bang releases."

### Sunset Features

Over time, older features may conflict with new ones or become obsolete. Bucketeer's platform provides visibility into which features are still in use and who is using them, enabling teams to effectively manage their code base by determining what should be retained.

### Reactive Monitoring

Bucketeer's feature management platform aids teams in troubleshooting and resolving issues in real-time. With feature flags implemented within their applications, teams can utilize Bucketeer's audit log to identify changes that led to incidents, allowing for prompt resolution and continuous improvement.

## Contact the Bucketeer team

If you don't find your answer here, feel free to [contact us](https://app.slack.com/client/T08PSQ7BQ/C043026BME1).
