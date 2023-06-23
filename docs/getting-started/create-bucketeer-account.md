---
title: Create a Bucketeer account
# sidebar_position:
slug: /getting-started/create-bucketeer-account
description: This page guides the user on how to create the Bucketeer account.
tags: ['guide', 'account', 'create']
---

import CenteredImg from '@site/src/components/centered-img/CenteredImg';

Bucketeer allows you to manage and track data related to different environments. To begin using Bucketeer, you need an account. Ask the administrator to create an account if you are a new team member.

The following steps show how to create an account as an Admin.

## Account types

There are two types of accounts available in Buckteers, each with different access levels:

- **Admin account**: it provides full access and control over all existing projects. Members with admin accounts can access all existing environments and manage other users' accounts.

- **Environment account**: it is associated with a specific environment on the platform. Users with Environment accounts can only access tags, reports, and data related to that particular environment.

### Environment account roles

In addition to being limited to a specific environment, Environment accounts also have predefined roles assigned during the account creation process. The defined role will change the account access level. The available roles for Environment accounts are described below:

- **Viewer**: viewers have read-only access to their assigned environments. They can view all data but cannot make any modifications. This role is suitable for individuals in your organization who need visibility into feature flags without the ability to modify rollout rules or administer the system.
- **Editor**: editors can modify feature flags, goals, experiments, and more within their assigned environments. However, they do not have the authority to add new team members to the account.
- **Owner**: owners possess complete control over their assigned environments. They can perform various actions and make changes across the environment, including adding and removing team members.

:::note

If you are using the Bucketeer system, to create an Admin account, you need to contact the Bucketeer team.

As an open-source project, you can host your Buckteers version and create an Admin account.

:::

## Creating an Environment account

To add a new member to a Buckteer organization, an Admin needs to create an Environment account.

:::note

Only Admins are capable of creating Environment accounts.

:::

As an Admin, you can follow the steps below to create an Environment account for a new user:

1. Access the [Bucketeer Dashboard](https://dev.bucketeer.jp/) using your Admin account credentials.
2. Select the desired project environment from the available options.
3. Navigate to the **Accounts** page within the Dashboard.
4. Click on the **+ Add** button to add a new account.

<CenteredImg
  imgURL='img/getting-started/create-bucketeer-account-1.png'
  alt='Account dashboard tab'
  wSize='100%'
/>

5. On the left panel, enter the new user's email and define their role.
6. Click **Submit** to create the account.

<CenteredImg 
  imgURL='img/getting-started/create-bucketeer-account-2.png'
  wSize='400px'
  alt='Create an account'
  borderWidth='1px'/>

The newly added member will receive an email invitation to join the project. If the new user does not receive the invitation email, the Admin can resend it by going to the Accounts page, locating the user's name in the Accounts list, and clicking on it.
