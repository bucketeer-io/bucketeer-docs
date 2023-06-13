---
title: Choose an SDK
# sidebar_position: 
slug: /getting-started/choose-sdk
description: Lists and briefly describes each SKD available to use in the integration process.
tags: ['guide','sdk','select']
---

Lists and briefly describes each SKD available to use in the integration process.

## important

before using starting using the SDK you need to have your `apiKey`, `apiURL`, and `featureTag`. To get this information you need to create a feature flag. Therefore, we recomend you to go thougth the [Using Feature Flags](test) guides before trying to use any SDK.

The featureTag é definido ao criar a feature flag. Ele tem a função de facilitar a busca, que é um caso de uso muito comum em todos os sistemas. No entanto, ele tem outra função importantíssima para otimização do funcionamenot do sistema Bucketeer. As tags servem de limitador ao avaliar usuários. Assim, se a chamada utiliza uma tag muito abrangente, utilizada por várias feature flags, ou mesmo não utiliza tags, muitas feature flags serão retornadas para avaliar o usuário. Isso, pode não ser um problema para sistemas pequenos, com poucas featureflags. No entanto, ao passo que o sistema cresce, a não utilização de tags para identificar e otimizar as chamadas de avaliação pode consumir muito processamento do seu servidor e elevar seus custos. Além disso, o tamanho e o tempo de resposta crescem, prejudicando a experiência do usuário. Por esse motivo, as tags tem tanta relevância no sistema Bucketeer.

Portanto, a correta utilização de tags irá:

- ascelerar o processo de avaliação, uma vez que somente os aspectos necessários serão analisados.
- reduzir o trafego de informações entre servidor e a aplicação local.

Apesar dessa relevância e possíveis impactos na performance e custos, a utilização de tags é opcional.

Exemplo: Se você pretende opearar várias feature flags em vários plataformas, nós recomentadamos você a usar as tags para distinguir as plataformas e separar os conteúdos. Assim você irá reduzir o tempo de resposta e gastos operacionais.
