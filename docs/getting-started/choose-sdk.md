---
title: Choose an SDK
# sidebar_position: 
slug: /getting-started/choose-sdk
description: Lists and briefly describes each SKD available to use in the integration process.
tags: ['guide','sdk','select']
---

In this section you find information describing what are SDKs and listing the Bucketeer SDKs available to use in the integration process.

## What are SDKs

SDK stands for Software Development Kit, a set of resources provided to developers to facilitate and streamline the creation and integration of applications for a specific platform, operating system, or programming language. SDKs combine everything professionals need to develop and run software according to their goals and requirements. 

SDKs make it easier for developers to create and implement applications by providing libraries with pre-configured functionalities, components, requests, guides, and instructions.

With SDKs, developers don't have to start the development process from scratch. As a result, there is a reduction in the time and resources required for application implementation, leading to increased efficiency and productivity for the programming team.

## Choose a Bucketeer SDK

When integrating a Bucketeer SDK with you code, it is important to consider two factors. First, to choose the right SDK you need to asses  your specific requirements to determine whether you need a server-side or client-side SDK. Typically, you will only need one SDK per application or service. However, if your product consists of applications or services written in multiple languages, you can use multiple SDKs. After you need to select the programming language, which normally will be the one used on your system.

Currently, Bucketeer SDK supports the following programming languages for client-side:

- [Android](/sdk/client-side/android)
- [iOS](/sdk/client-side/ios)
- [JavaScript](/sdk/client-side/javascript)
- [Flutter](/sdk/client-side/flutter)

Regarding the server-side, Buckteers SDK supports:

- [Go](/sdk/server-side/go)
- [NodeJS](/sdk/server-side/node-js)

To understand the SDK integration process, access the [SDKs](/sdk) page. 

:::tip

We strongly recommend you check the [Integrate Bucketeers](./integrate-bucketeers.md) and [Using Feature Flags](../using-feature-flags/) guides before using Bucketeer SDKs in your application.

:::

### Client-side SDK

Client-side SDKs are designed for single-user desktop, mobile, and embedded applications. They are intended for use in potentially less secure environments, such as personal computers or mobile devices, including mobile SDKs. These SDKs typically run on end-user devices, which makes them vulnerable to compromise by users who unpack a mobile app to examine the SDK bytecode or use developer tools in their browser to inspect internal site data. Therefore, never using a server-side SDK key in a client-side or mobile application is essential.

### Server-side SDK

Server-side SDKs are designed for multi-user systems and intended for use in trusted environments, like corporate networks or web servers. They operate within server-architected applications running on your infrastructure or trusted cloud-based infrastructure. These locations are not directly accessible by end-users. The restricted access of server-based applications allows server-side SDKs to safely receive flag data and rulesets without the need to filter out sensitive information.
