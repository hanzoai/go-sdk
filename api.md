# hanzoai

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#GetHomeResponse">GetHomeResponse</a>

Methods:

- <code title="get /">client.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#HanzoaiService.GetHome">GetHome</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#GetHomeResponse">GetHomeResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Models

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#ModelListResponse">ModelListResponse</a>

Methods:

- <code title="get /v1/models">client.Models.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#ModelService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#ModelListParams">ModelListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#ModelListResponse">ModelListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# OpenAI

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#OpenAINewResponse">OpenAINewResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#OpenAIGetResponse">OpenAIGetResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#OpenAIUpdateResponse">OpenAIUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#OpenAIDeleteResponse">OpenAIDeleteResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#OpenAIPatchResponse">OpenAIPatchResponse</a>

Methods:

- <code title="post /openai/{endpoint}">client.OpenAI.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#OpenAIService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, endpoint <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#OpenAINewResponse">OpenAINewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /openai/{endpoint}">client.OpenAI.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#OpenAIService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, endpoint <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#OpenAIGetResponse">OpenAIGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /openai/{endpoint}">client.OpenAI.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#OpenAIService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, endpoint <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#OpenAIUpdateResponse">OpenAIUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /openai/{endpoint}">client.OpenAI.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#OpenAIService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, endpoint <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#OpenAIDeleteResponse">OpenAIDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /openai/{endpoint}">client.OpenAI.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#OpenAIService.Patch">Patch</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, endpoint <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#OpenAIPatchResponse">OpenAIPatchResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Deployments

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#OpenAIDeploymentCompleteResponse">OpenAIDeploymentCompleteResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#OpenAIDeploymentEmbedResponse">OpenAIDeploymentEmbedResponse</a>

Methods:

- <code title="post /openai/deployments/{model}/completions">client.OpenAI.Deployments.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#OpenAIDeploymentService.Complete">Complete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, model <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#OpenAIDeploymentCompleteResponse">OpenAIDeploymentCompleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /openai/deployments/{model}/embeddings">client.OpenAI.Deployments.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#OpenAIDeploymentService.Embed">Embed</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, model <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#OpenAIDeploymentEmbedResponse">OpenAIDeploymentEmbedResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Chat

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#OpenAIDeploymentChatCompleteResponse">OpenAIDeploymentChatCompleteResponse</a>

Methods:

- <code title="post /openai/deployments/{model}/chat/completions">client.OpenAI.Deployments.Chat.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#OpenAIDeploymentChatService.Complete">Complete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, model <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#OpenAIDeploymentChatCompleteResponse">OpenAIDeploymentChatCompleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Engines

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#EngineCompleteResponse">EngineCompleteResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#EngineEmbedResponse">EngineEmbedResponse</a>

Methods:

- <code title="post /engines/{model}/completions">client.Engines.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#EngineService.Complete">Complete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, model <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#EngineCompleteResponse">EngineCompleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /engines/{model}/embeddings">client.Engines.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#EngineService.Embed">Embed</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, model <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#EngineEmbedResponse">EngineEmbedResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Chat

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#EngineChatCompleteResponse">EngineChatCompleteResponse</a>

Methods:

- <code title="post /engines/{model}/chat/completions">client.Engines.Chat.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#EngineChatService.Complete">Complete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, model <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#EngineChatCompleteResponse">EngineChatCompleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Chat

## Completions

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#ChatCompletionNewResponse">ChatCompletionNewResponse</a>

Methods:

- <code title="post /v1/chat/completions">client.Chat.Completions.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#ChatCompletionService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#ChatCompletionNewParams">ChatCompletionNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#ChatCompletionNewResponse">ChatCompletionNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Completions

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#CompletionNewResponse">CompletionNewResponse</a>

Methods:

- <code title="post /completions">client.Completions.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#CompletionService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#CompletionNewParams">CompletionNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#CompletionNewResponse">CompletionNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Embeddings

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#EmbeddingNewResponse">EmbeddingNewResponse</a>

Methods:

- <code title="post /embeddings">client.Embeddings.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#EmbeddingService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#EmbeddingNewParams">EmbeddingNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#EmbeddingNewResponse">EmbeddingNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Images

## Generations

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#ImageGenerationNewResponse">ImageGenerationNewResponse</a>

Methods:

- <code title="post /v1/images/generations">client.Images.Generations.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#ImageGenerationService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#ImageGenerationNewResponse">ImageGenerationNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Audio

## Speech

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#AudioSpeechNewResponse">AudioSpeechNewResponse</a>

Methods:

- <code title="post /v1/audio/speech">client.Audio.Speech.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#AudioSpeechService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#AudioSpeechNewResponse">AudioSpeechNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Transcriptions

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#AudioTranscriptionNewResponse">AudioTranscriptionNewResponse</a>

Methods:

- <code title="post /v1/audio/transcriptions">client.Audio.Transcriptions.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#AudioTranscriptionService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#AudioTranscriptionNewParams">AudioTranscriptionNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#AudioTranscriptionNewResponse">AudioTranscriptionNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Assistants

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#AssistantNewResponse">AssistantNewResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#AssistantListResponse">AssistantListResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#AssistantDeleteResponse">AssistantDeleteResponse</a>

Methods:

- <code title="post /v1/assistants">client.Assistants.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#AssistantService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#AssistantNewResponse">AssistantNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/assistants">client.Assistants.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#AssistantService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#AssistantListResponse">AssistantListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/assistants/{assistant_id}">client.Assistants.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#AssistantService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, assistantID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#AssistantDeleteResponse">AssistantDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Threads

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#ThreadNewResponse">ThreadNewResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#ThreadGetResponse">ThreadGetResponse</a>

Methods:

- <code title="post /v1/threads">client.Threads.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#ThreadService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#ThreadNewResponse">ThreadNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/threads/{thread_id}">client.Threads.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#ThreadService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, threadID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#ThreadGetResponse">ThreadGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Messages

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#ThreadMessageNewResponse">ThreadMessageNewResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#ThreadMessageListResponse">ThreadMessageListResponse</a>

Methods:

- <code title="post /v1/threads/{thread_id}/messages">client.Threads.Messages.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#ThreadMessageService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, threadID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#ThreadMessageNewResponse">ThreadMessageNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/threads/{thread_id}/messages">client.Threads.Messages.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#ThreadMessageService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, threadID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#ThreadMessageListResponse">ThreadMessageListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Runs

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#ThreadRunNewResponse">ThreadRunNewResponse</a>

Methods:

- <code title="post /v1/threads/{thread_id}/runs">client.Threads.Runs.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#ThreadRunService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, threadID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#ThreadRunNewResponse">ThreadRunNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Moderations

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#ModerationNewResponse">ModerationNewResponse</a>

Methods:

- <code title="post /v1/moderations">client.Moderations.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#ModerationService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#ModerationNewResponse">ModerationNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Utils

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#UtilGetSupportedOpenAIParamsResponse">UtilGetSupportedOpenAIParamsResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#UtilTokenCounterResponse">UtilTokenCounterResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#UtilTransformRequestResponse">UtilTransformRequestResponse</a>

Methods:

- <code title="get /utils/supported_openai_params">client.Utils.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#UtilService.GetSupportedOpenAIParams">GetSupportedOpenAIParams</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#UtilGetSupportedOpenAIParamsParams">UtilGetSupportedOpenAIParamsParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#UtilGetSupportedOpenAIParamsResponse">UtilGetSupportedOpenAIParamsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /utils/token_counter">client.Utils.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#UtilService.TokenCounter">TokenCounter</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#UtilTokenCounterParams">UtilTokenCounterParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#UtilTokenCounterResponse">UtilTokenCounterResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /utils/transform_request">client.Utils.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#UtilService.TransformRequest">TransformRequest</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#UtilTransformRequestParams">UtilTransformRequestParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#UtilTransformRequestResponse">UtilTransformRequestResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Model

Params Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#ConfigurableClientsideParamsCustomAuth">ConfigurableClientsideParamsCustomAuth</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#ModelInfoParam">ModelInfoParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#ModelNewResponse">ModelNewResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#ModelDeleteResponse">ModelDeleteResponse</a>

Methods:

- <code title="post /model/new">client.Model.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#ModelService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#ModelNewParams">ModelNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#ModelNewResponse">ModelNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /model/delete">client.Model.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#ModelService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#ModelDeleteParams">ModelDeleteParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#ModelDeleteResponse">ModelDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Info

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#ModelInfoListResponse">ModelInfoListResponse</a>

Methods:

- <code title="get /model/info">client.Model.Info.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#ModelInfoService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#ModelInfoListParams">ModelInfoListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#ModelInfoListResponse">ModelInfoListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Update

Params Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#UpdateDeploymentParam">UpdateDeploymentParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#ModelUpdateFullResponse">ModelUpdateFullResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#ModelUpdatePartialResponse">ModelUpdatePartialResponse</a>

Methods:

- <code title="post /model/update">client.Model.Update.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#ModelUpdateService.Full">Full</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#ModelUpdateFullParams">ModelUpdateFullParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#ModelUpdateFullResponse">ModelUpdateFullResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /model/{model_id}/update">client.Model.Update.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#ModelUpdateService.Partial">Partial</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, modelID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#ModelUpdatePartialParams">ModelUpdatePartialParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#ModelUpdatePartialResponse">ModelUpdatePartialResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# ModelGroup

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#ModelGroupGetInfoResponse">ModelGroupGetInfoResponse</a>

Methods:

- <code title="get /model_group/info">client.ModelGroup.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#ModelGroupService.GetInfo">GetInfo</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#ModelGroupGetInfoParams">ModelGroupGetInfoParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#ModelGroupGetInfoResponse">ModelGroupGetInfoResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Routes

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#RouteListResponse">RouteListResponse</a>

Methods:

- <code title="get /routes">client.Routes.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#RouteService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#RouteListResponse">RouteListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Responses

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#ResponseNewResponse">ResponseNewResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#ResponseGetResponse">ResponseGetResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#ResponseDeleteResponse">ResponseDeleteResponse</a>

Methods:

- <code title="post /v1/responses">client.Responses.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#ResponseService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#ResponseNewResponse">ResponseNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/responses/{response_id}">client.Responses.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#ResponseService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, responseID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#ResponseGetResponse">ResponseGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/responses/{response_id}">client.Responses.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#ResponseService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, responseID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#ResponseDeleteResponse">ResponseDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## InputItems

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#ResponseInputItemListResponse">ResponseInputItemListResponse</a>

Methods:

- <code title="get /v1/responses/{response_id}/input_items">client.Responses.InputItems.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#ResponseInputItemService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, responseID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#ResponseInputItemListResponse">ResponseInputItemListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Batches

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#BatchNewResponse">BatchNewResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#BatchGetResponse">BatchGetResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#BatchListResponse">BatchListResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#BatchCancelWithProviderResponse">BatchCancelWithProviderResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#BatchNewWithProviderResponse">BatchNewWithProviderResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#BatchListWithProviderResponse">BatchListWithProviderResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#BatchGetWithProviderResponse">BatchGetWithProviderResponse</a>

Methods:

- <code title="post /v1/batches">client.Batches.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#BatchService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#BatchNewParams">BatchNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#BatchNewResponse">BatchNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/batches/{batch_id}">client.Batches.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#BatchService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, batchID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#BatchGetParams">BatchGetParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#BatchGetResponse">BatchGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/batches">client.Batches.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#BatchService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#BatchListParams">BatchListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#BatchListResponse">BatchListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /{provider}/v1/batches/{batch_id}/cancel">client.Batches.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#BatchService.CancelWithProvider">CancelWithProvider</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, provider <a href="https://pkg.go.dev/builtin#string">string</a>, batchID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#BatchCancelWithProviderResponse">BatchCancelWithProviderResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /{provider}/v1/batches">client.Batches.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#BatchService.NewWithProvider">NewWithProvider</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, provider <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#BatchNewWithProviderResponse">BatchNewWithProviderResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /{provider}/v1/batches">client.Batches.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#BatchService.ListWithProvider">ListWithProvider</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, provider <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#BatchListWithProviderParams">BatchListWithProviderParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#BatchListWithProviderResponse">BatchListWithProviderResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /{provider}/v1/batches/{batch_id}">client.Batches.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#BatchService.GetWithProvider">GetWithProvider</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, provider <a href="https://pkg.go.dev/builtin#string">string</a>, batchID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#BatchGetWithProviderResponse">BatchGetWithProviderResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Cancel

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#BatchCancelCancelResponse">BatchCancelCancelResponse</a>

Methods:

- <code title="post /batches/{batch_id}/cancel">client.Batches.Cancel.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#BatchCancelService.Cancel">Cancel</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, batchID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#BatchCancelCancelParams">BatchCancelCancelParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#BatchCancelCancelResponse">BatchCancelCancelResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Rerank

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#RerankNewResponse">RerankNewResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#RerankNewV1Response">RerankNewV1Response</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#RerankNewV2Response">RerankNewV2Response</a>

Methods:

- <code title="post /rerank">client.Rerank.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#RerankService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#RerankNewResponse">RerankNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/rerank">client.Rerank.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#RerankService.NewV1">NewV1</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#RerankNewV1Response">RerankNewV1Response</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v2/rerank">client.Rerank.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#RerankService.NewV2">NewV2</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#RerankNewV2Response">RerankNewV2Response</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# FineTuning

## Jobs

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#FineTuningJobNewResponse">FineTuningJobNewResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#FineTuningJobGetResponse">FineTuningJobGetResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#FineTuningJobListResponse">FineTuningJobListResponse</a>

Methods:

- <code title="post /v1/fine_tuning/jobs">client.FineTuning.Jobs.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#FineTuningJobService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#FineTuningJobNewParams">FineTuningJobNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#FineTuningJobNewResponse">FineTuningJobNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/fine_tuning/jobs/{fine_tuning_job_id}">client.FineTuning.Jobs.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#FineTuningJobService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, fineTuningJobID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#FineTuningJobGetParams">FineTuningJobGetParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#FineTuningJobGetResponse">FineTuningJobGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/fine_tuning/jobs">client.FineTuning.Jobs.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#FineTuningJobService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#FineTuningJobListParams">FineTuningJobListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#FineTuningJobListResponse">FineTuningJobListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Cancel

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#FineTuningJobCancelNewResponse">FineTuningJobCancelNewResponse</a>

Methods:

- <code title="post /v1/fine_tuning/jobs/{fine_tuning_job_id}/cancel">client.FineTuning.Jobs.Cancel.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#FineTuningJobCancelService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, fineTuningJobID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#FineTuningJobCancelNewResponse">FineTuningJobCancelNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Credentials

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#CredentialNewResponse">CredentialNewResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#CredentialListResponse">CredentialListResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#CredentialDeleteResponse">CredentialDeleteResponse</a>

Methods:

- <code title="post /credentials">client.Credentials.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#CredentialService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#CredentialNewParams">CredentialNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#CredentialNewResponse">CredentialNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /credentials">client.Credentials.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#CredentialService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#CredentialListResponse">CredentialListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /credentials/{credential_name}">client.Credentials.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#CredentialService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, credentialName <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#CredentialDeleteResponse">CredentialDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# VertexAI

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#VertexAINewResponse">VertexAINewResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#VertexAIGetResponse">VertexAIGetResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#VertexAIUpdateResponse">VertexAIUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#VertexAIDeleteResponse">VertexAIDeleteResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#VertexAIPatchResponse">VertexAIPatchResponse</a>

Methods:

- <code title="post /vertex_ai/{endpoint}">client.VertexAI.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#VertexAIService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, endpoint <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#VertexAINewResponse">VertexAINewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /vertex_ai/{endpoint}">client.VertexAI.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#VertexAIService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, endpoint <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#VertexAIGetResponse">VertexAIGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /vertex_ai/{endpoint}">client.VertexAI.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#VertexAIService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, endpoint <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#VertexAIUpdateResponse">VertexAIUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /vertex_ai/{endpoint}">client.VertexAI.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#VertexAIService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, endpoint <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#VertexAIDeleteResponse">VertexAIDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /vertex_ai/{endpoint}">client.VertexAI.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#VertexAIService.Patch">Patch</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, endpoint <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#VertexAIPatchResponse">VertexAIPatchResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Gemini

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#GeminiNewResponse">GeminiNewResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#GeminiGetResponse">GeminiGetResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#GeminiUpdateResponse">GeminiUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#GeminiDeleteResponse">GeminiDeleteResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#GeminiPatchResponse">GeminiPatchResponse</a>

Methods:

- <code title="post /gemini/{endpoint}">client.Gemini.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#GeminiService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, endpoint <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#GeminiNewResponse">GeminiNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /gemini/{endpoint}">client.Gemini.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#GeminiService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, endpoint <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#GeminiGetResponse">GeminiGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /gemini/{endpoint}">client.Gemini.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#GeminiService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, endpoint <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#GeminiUpdateResponse">GeminiUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /gemini/{endpoint}">client.Gemini.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#GeminiService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, endpoint <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#GeminiDeleteResponse">GeminiDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /gemini/{endpoint}">client.Gemini.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#GeminiService.Patch">Patch</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, endpoint <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#GeminiPatchResponse">GeminiPatchResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Cohere

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#CohereNewResponse">CohereNewResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#CohereGetResponse">CohereGetResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#CohereUpdateResponse">CohereUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#CohereDeleteResponse">CohereDeleteResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#CohereModifyResponse">CohereModifyResponse</a>

Methods:

- <code title="post /cohere/{endpoint}">client.Cohere.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#CohereService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, endpoint <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#CohereNewResponse">CohereNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /cohere/{endpoint}">client.Cohere.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#CohereService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, endpoint <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#CohereGetResponse">CohereGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /cohere/{endpoint}">client.Cohere.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#CohereService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, endpoint <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#CohereUpdateResponse">CohereUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /cohere/{endpoint}">client.Cohere.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#CohereService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, endpoint <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#CohereDeleteResponse">CohereDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /cohere/{endpoint}">client.Cohere.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#CohereService.Modify">Modify</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, endpoint <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#CohereModifyResponse">CohereModifyResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Anthropic

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#AnthropicNewResponse">AnthropicNewResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#AnthropicGetResponse">AnthropicGetResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#AnthropicUpdateResponse">AnthropicUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#AnthropicDeleteResponse">AnthropicDeleteResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#AnthropicModifyResponse">AnthropicModifyResponse</a>

Methods:

- <code title="post /anthropic/{endpoint}">client.Anthropic.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#AnthropicService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, endpoint <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#AnthropicNewResponse">AnthropicNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /anthropic/{endpoint}">client.Anthropic.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#AnthropicService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, endpoint <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#AnthropicGetResponse">AnthropicGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /anthropic/{endpoint}">client.Anthropic.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#AnthropicService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, endpoint <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#AnthropicUpdateResponse">AnthropicUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /anthropic/{endpoint}">client.Anthropic.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#AnthropicService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, endpoint <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#AnthropicDeleteResponse">AnthropicDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /anthropic/{endpoint}">client.Anthropic.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#AnthropicService.Modify">Modify</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, endpoint <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#AnthropicModifyResponse">AnthropicModifyResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Bedrock

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#BedrockNewResponse">BedrockNewResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#BedrockGetResponse">BedrockGetResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#BedrockUpdateResponse">BedrockUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#BedrockDeleteResponse">BedrockDeleteResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#BedrockPatchResponse">BedrockPatchResponse</a>

Methods:

- <code title="post /bedrock/{endpoint}">client.Bedrock.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#BedrockService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, endpoint <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#BedrockNewResponse">BedrockNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /bedrock/{endpoint}">client.Bedrock.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#BedrockService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, endpoint <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#BedrockGetResponse">BedrockGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /bedrock/{endpoint}">client.Bedrock.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#BedrockService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, endpoint <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#BedrockUpdateResponse">BedrockUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /bedrock/{endpoint}">client.Bedrock.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#BedrockService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, endpoint <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#BedrockDeleteResponse">BedrockDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /bedrock/{endpoint}">client.Bedrock.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#BedrockService.Patch">Patch</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, endpoint <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#BedrockPatchResponse">BedrockPatchResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# EuAssemblyai

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#EuAssemblyaiNewResponse">EuAssemblyaiNewResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#EuAssemblyaiGetResponse">EuAssemblyaiGetResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#EuAssemblyaiUpdateResponse">EuAssemblyaiUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#EuAssemblyaiDeleteResponse">EuAssemblyaiDeleteResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#EuAssemblyaiPatchResponse">EuAssemblyaiPatchResponse</a>

Methods:

- <code title="post /eu.assemblyai/{endpoint}">client.EuAssemblyai.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#EuAssemblyaiService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, endpoint <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#EuAssemblyaiNewResponse">EuAssemblyaiNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /eu.assemblyai/{endpoint}">client.EuAssemblyai.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#EuAssemblyaiService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, endpoint <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#EuAssemblyaiGetResponse">EuAssemblyaiGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /eu.assemblyai/{endpoint}">client.EuAssemblyai.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#EuAssemblyaiService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, endpoint <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#EuAssemblyaiUpdateResponse">EuAssemblyaiUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /eu.assemblyai/{endpoint}">client.EuAssemblyai.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#EuAssemblyaiService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, endpoint <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#EuAssemblyaiDeleteResponse">EuAssemblyaiDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /eu.assemblyai/{endpoint}">client.EuAssemblyai.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#EuAssemblyaiService.Patch">Patch</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, endpoint <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#EuAssemblyaiPatchResponse">EuAssemblyaiPatchResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Assemblyai

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#AssemblyaiNewResponse">AssemblyaiNewResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#AssemblyaiGetResponse">AssemblyaiGetResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#AssemblyaiUpdateResponse">AssemblyaiUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#AssemblyaiDeleteResponse">AssemblyaiDeleteResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#AssemblyaiPatchResponse">AssemblyaiPatchResponse</a>

Methods:

- <code title="post /assemblyai/{endpoint}">client.Assemblyai.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#AssemblyaiService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, endpoint <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#AssemblyaiNewResponse">AssemblyaiNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /assemblyai/{endpoint}">client.Assemblyai.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#AssemblyaiService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, endpoint <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#AssemblyaiGetResponse">AssemblyaiGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /assemblyai/{endpoint}">client.Assemblyai.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#AssemblyaiService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, endpoint <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#AssemblyaiUpdateResponse">AssemblyaiUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /assemblyai/{endpoint}">client.Assemblyai.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#AssemblyaiService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, endpoint <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#AssemblyaiDeleteResponse">AssemblyaiDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /assemblyai/{endpoint}">client.Assemblyai.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#AssemblyaiService.Patch">Patch</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, endpoint <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#AssemblyaiPatchResponse">AssemblyaiPatchResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Azure

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#AzureNewResponse">AzureNewResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#AzureUpdateResponse">AzureUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#AzureDeleteResponse">AzureDeleteResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#AzureCallResponse">AzureCallResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#AzurePatchResponse">AzurePatchResponse</a>

Methods:

- <code title="post /azure/{endpoint}">client.Azure.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#AzureService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, endpoint <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#AzureNewResponse">AzureNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /azure/{endpoint}">client.Azure.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#AzureService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, endpoint <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#AzureUpdateResponse">AzureUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /azure/{endpoint}">client.Azure.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#AzureService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, endpoint <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#AzureDeleteResponse">AzureDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /azure/{endpoint}">client.Azure.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#AzureService.Call">Call</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, endpoint <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#AzureCallResponse">AzureCallResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /azure/{endpoint}">client.Azure.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#AzureService.Patch">Patch</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, endpoint <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#AzurePatchResponse">AzurePatchResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Langfuse

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#LangfuseNewResponse">LangfuseNewResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#LangfuseGetResponse">LangfuseGetResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#LangfuseUpdateResponse">LangfuseUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#LangfuseDeleteResponse">LangfuseDeleteResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#LangfusePatchResponse">LangfusePatchResponse</a>

Methods:

- <code title="post /langfuse/{endpoint}">client.Langfuse.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#LangfuseService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, endpoint <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#LangfuseNewResponse">LangfuseNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /langfuse/{endpoint}">client.Langfuse.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#LangfuseService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, endpoint <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#LangfuseGetResponse">LangfuseGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /langfuse/{endpoint}">client.Langfuse.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#LangfuseService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, endpoint <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#LangfuseUpdateResponse">LangfuseUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /langfuse/{endpoint}">client.Langfuse.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#LangfuseService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, endpoint <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#LangfuseDeleteResponse">LangfuseDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /langfuse/{endpoint}">client.Langfuse.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#LangfuseService.Patch">Patch</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, endpoint <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#LangfusePatchResponse">LangfusePatchResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Config

## PassThroughEndpoint

Params Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#PassThroughGenericEndpointParam">PassThroughGenericEndpointParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#PassThroughEndpointResponse">PassThroughEndpointResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#PassThroughGenericEndpoint">PassThroughGenericEndpoint</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#ConfigPassThroughEndpointNewResponse">ConfigPassThroughEndpointNewResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#ConfigPassThroughEndpointUpdateResponse">ConfigPassThroughEndpointUpdateResponse</a>

Methods:

- <code title="post /config/pass_through_endpoint">client.Config.PassThroughEndpoint.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#ConfigPassThroughEndpointService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#ConfigPassThroughEndpointNewParams">ConfigPassThroughEndpointNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#ConfigPassThroughEndpointNewResponse">ConfigPassThroughEndpointNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /config/pass_through_endpoint/{endpoint_id}">client.Config.PassThroughEndpoint.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#ConfigPassThroughEndpointService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, endpointID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#ConfigPassThroughEndpointUpdateResponse">ConfigPassThroughEndpointUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /config/pass_through_endpoint">client.Config.PassThroughEndpoint.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#ConfigPassThroughEndpointService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#ConfigPassThroughEndpointListParams">ConfigPassThroughEndpointListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#PassThroughEndpointResponse">PassThroughEndpointResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /config/pass_through_endpoint">client.Config.PassThroughEndpoint.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#ConfigPassThroughEndpointService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#ConfigPassThroughEndpointDeleteParams">ConfigPassThroughEndpointDeleteParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#PassThroughEndpointResponse">PassThroughEndpointResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Test

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#TestPingResponse">TestPingResponse</a>

Methods:

- <code title="get /test">client.Test.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#TestService.Ping">Ping</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#TestPingResponse">TestPingResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Health

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#HealthCheckAllResponse">HealthCheckAllResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#HealthCheckLivelinessResponse">HealthCheckLivelinessResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#HealthCheckLivenessResponse">HealthCheckLivenessResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#HealthCheckReadinessResponse">HealthCheckReadinessResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#HealthCheckServicesResponse">HealthCheckServicesResponse</a>

Methods:

- <code title="get /health">client.Health.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#HealthService.CheckAll">CheckAll</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#HealthCheckAllParams">HealthCheckAllParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#HealthCheckAllResponse">HealthCheckAllResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /health/liveliness">client.Health.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#HealthService.CheckLiveliness">CheckLiveliness</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#HealthCheckLivelinessResponse">HealthCheckLivelinessResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /health/liveness">client.Health.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#HealthService.CheckLiveness">CheckLiveness</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#HealthCheckLivenessResponse">HealthCheckLivenessResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /health/readiness">client.Health.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#HealthService.CheckReadiness">CheckReadiness</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#HealthCheckReadinessResponse">HealthCheckReadinessResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /health/services">client.Health.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#HealthService.CheckServices">CheckServices</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#HealthCheckServicesParams">HealthCheckServicesParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#HealthCheckServicesResponse">HealthCheckServicesResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Active

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#ActiveListCallbacksResponse">ActiveListCallbacksResponse</a>

Methods:

- <code title="get /active/callbacks">client.Active.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#ActiveService.ListCallbacks">ListCallbacks</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#ActiveListCallbacksResponse">ActiveListCallbacksResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Settings

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#SettingGetResponse">SettingGetResponse</a>

Methods:

- <code title="get /settings">client.Settings.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#SettingService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#SettingGetResponse">SettingGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Key

Params Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#BlockKeyRequestParam">BlockKeyRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#GenerateKeyResponse">GenerateKeyResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#KeyUpdateResponse">KeyUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#KeyListResponse">KeyListResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#KeyDeleteResponse">KeyDeleteResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#KeyBlockResponse">KeyBlockResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#KeyCheckHealthResponse">KeyCheckHealthResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#KeyGetInfoResponse">KeyGetInfoResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#KeyUnblockResponse">KeyUnblockResponse</a>

Methods:

- <code title="post /key/update">client.Key.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#KeyService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#KeyUpdateParams">KeyUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#KeyUpdateResponse">KeyUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /key/list">client.Key.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#KeyService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#KeyListParams">KeyListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#KeyListResponse">KeyListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /key/delete">client.Key.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#KeyService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#KeyDeleteParams">KeyDeleteParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#KeyDeleteResponse">KeyDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /key/block">client.Key.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#KeyService.Block">Block</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#KeyBlockParams">KeyBlockParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#KeyBlockResponse">KeyBlockResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /key/health">client.Key.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#KeyService.CheckHealth">CheckHealth</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#KeyCheckHealthResponse">KeyCheckHealthResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /key/generate">client.Key.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#KeyService.Generate">Generate</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#KeyGenerateParams">KeyGenerateParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#GenerateKeyResponse">GenerateKeyResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /key/{key}/regenerate">client.Key.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#KeyService.RegenerateByKey">RegenerateByKey</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, key <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#KeyRegenerateByKeyParams">KeyRegenerateByKeyParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#GenerateKeyResponse">GenerateKeyResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /key/info">client.Key.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#KeyService.GetInfo">GetInfo</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#KeyGetInfoParams">KeyGetInfoParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#KeyGetInfoResponse">KeyGetInfoResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /key/unblock">client.Key.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#KeyService.Unblock">Unblock</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#KeyUnblockParams">KeyUnblockParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#KeyUnblockResponse">KeyUnblockResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Regenerate

Params Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#RegenerateKeyRequestParam">RegenerateKeyRequestParam</a>

# User

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#UserNewResponse">UserNewResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#UserUpdateResponse">UserUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#UserListResponse">UserListResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#UserDeleteResponse">UserDeleteResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#UserGetInfoResponse">UserGetInfoResponse</a>

Methods:

- <code title="post /user/new">client.User.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#UserService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#UserNewParams">UserNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#UserNewResponse">UserNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /user/update">client.User.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#UserService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#UserUpdateParams">UserUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#UserUpdateResponse">UserUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /user/get_users">client.User.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#UserService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#UserListParams">UserListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#UserListResponse">UserListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /user/delete">client.User.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#UserService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#UserDeleteParams">UserDeleteParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#UserDeleteResponse">UserDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /user/info">client.User.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#UserService.GetInfo">GetInfo</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#UserGetInfoParams">UserGetInfoParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#UserGetInfoResponse">UserGetInfoResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Team

Params Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#BlockTeamRequestParam">BlockTeamRequestParam</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#MemberParam">MemberParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#Member">Member</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#TeamNewResponse">TeamNewResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#TeamUpdateResponse">TeamUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#TeamListResponse">TeamListResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#TeamDeleteResponse">TeamDeleteResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#TeamAddMemberResponse">TeamAddMemberResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#TeamBlockResponse">TeamBlockResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#TeamDisableLoggingResponse">TeamDisableLoggingResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#TeamListAvailableResponse">TeamListAvailableResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#TeamRemoveMemberResponse">TeamRemoveMemberResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#TeamGetInfoResponse">TeamGetInfoResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#TeamUnblockResponse">TeamUnblockResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#TeamUpdateMemberResponse">TeamUpdateMemberResponse</a>

Methods:

- <code title="post /team/new">client.Team.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#TeamService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#TeamNewParams">TeamNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#TeamNewResponse">TeamNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /team/update">client.Team.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#TeamService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#TeamUpdateParams">TeamUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#TeamUpdateResponse">TeamUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /team/list">client.Team.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#TeamService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#TeamListParams">TeamListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#TeamListResponse">TeamListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /team/delete">client.Team.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#TeamService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#TeamDeleteParams">TeamDeleteParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#TeamDeleteResponse">TeamDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /team/member_add">client.Team.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#TeamService.AddMember">AddMember</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#TeamAddMemberParams">TeamAddMemberParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#TeamAddMemberResponse">TeamAddMemberResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /team/block">client.Team.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#TeamService.Block">Block</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#TeamBlockParams">TeamBlockParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#TeamBlockResponse">TeamBlockResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /team/{team_id}/disable_logging">client.Team.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#TeamService.DisableLogging">DisableLogging</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, teamID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#TeamDisableLoggingResponse">TeamDisableLoggingResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /team/available">client.Team.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#TeamService.ListAvailable">ListAvailable</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#TeamListAvailableParams">TeamListAvailableParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#TeamListAvailableResponse">TeamListAvailableResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /team/member_delete">client.Team.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#TeamService.RemoveMember">RemoveMember</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#TeamRemoveMemberParams">TeamRemoveMemberParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#TeamRemoveMemberResponse">TeamRemoveMemberResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /team/info">client.Team.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#TeamService.GetInfo">GetInfo</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#TeamGetInfoParams">TeamGetInfoParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#TeamGetInfoResponse">TeamGetInfoResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /team/unblock">client.Team.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#TeamService.Unblock">Unblock</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#TeamUnblockParams">TeamUnblockParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#TeamUnblockResponse">TeamUnblockResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /team/member_update">client.Team.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#TeamService.UpdateMember">UpdateMember</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#TeamUpdateMemberParams">TeamUpdateMemberParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#TeamUpdateMemberResponse">TeamUpdateMemberResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Model

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#TeamModelAddResponse">TeamModelAddResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#TeamModelRemoveResponse">TeamModelRemoveResponse</a>

Methods:

- <code title="post /team/model/add">client.Team.Model.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#TeamModelService.Add">Add</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#TeamModelAddParams">TeamModelAddParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#TeamModelAddResponse">TeamModelAddResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /team/model/delete">client.Team.Model.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#TeamModelService.Remove">Remove</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#TeamModelRemoveParams">TeamModelRemoveParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#TeamModelRemoveResponse">TeamModelRemoveResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Callback

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#TeamCallbackGetResponse">TeamCallbackGetResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#TeamCallbackAddResponse">TeamCallbackAddResponse</a>

Methods:

- <code title="get /team/{team_id}/callback">client.Team.Callback.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#TeamCallbackService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, teamID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#TeamCallbackGetResponse">TeamCallbackGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /team/{team_id}/callback">client.Team.Callback.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#TeamCallbackService.Add">Add</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, teamID <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#TeamCallbackAddParams">TeamCallbackAddParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#TeamCallbackAddResponse">TeamCallbackAddResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Organization

Params Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#OrgMemberParam">OrgMemberParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#OrganizationNewResponse">OrganizationNewResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#OrganizationUpdateResponse">OrganizationUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#OrganizationListResponse">OrganizationListResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#OrganizationDeleteResponse">OrganizationDeleteResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#OrganizationAddMemberResponse">OrganizationAddMemberResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#OrganizationDeleteMemberResponse">OrganizationDeleteMemberResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#OrganizationUpdateMemberResponse">OrganizationUpdateMemberResponse</a>

Methods:

- <code title="post /organization/new">client.Organization.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#OrganizationService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#OrganizationNewParams">OrganizationNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#OrganizationNewResponse">OrganizationNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /organization/update">client.Organization.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#OrganizationService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#OrganizationUpdateParams">OrganizationUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#OrganizationUpdateResponse">OrganizationUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /organization/list">client.Organization.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#OrganizationService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) ([]<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#OrganizationListResponse">OrganizationListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /organization/delete">client.Organization.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#OrganizationService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#OrganizationDeleteParams">OrganizationDeleteParams</a>) ([]<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#OrganizationDeleteResponse">OrganizationDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /organization/member_add">client.Organization.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#OrganizationService.AddMember">AddMember</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#OrganizationAddMemberParams">OrganizationAddMemberParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#OrganizationAddMemberResponse">OrganizationAddMemberResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /organization/member_delete">client.Organization.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#OrganizationService.DeleteMember">DeleteMember</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#OrganizationDeleteMemberParams">OrganizationDeleteMemberParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#OrganizationDeleteMemberResponse">OrganizationDeleteMemberResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /organization/member_update">client.Organization.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#OrganizationService.UpdateMember">UpdateMember</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#OrganizationUpdateMemberParams">OrganizationUpdateMemberParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#OrganizationUpdateMemberResponse">OrganizationUpdateMemberResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Info

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#OrganizationInfoGetResponse">OrganizationInfoGetResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#OrganizationInfoDeprecatedResponse">OrganizationInfoDeprecatedResponse</a>

Methods:

- <code title="get /organization/info">client.Organization.Info.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#OrganizationInfoService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#OrganizationInfoGetParams">OrganizationInfoGetParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#OrganizationInfoGetResponse">OrganizationInfoGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /organization/info">client.Organization.Info.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#OrganizationInfoService.Deprecated">Deprecated</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#OrganizationInfoDeprecatedParams">OrganizationInfoDeprecatedParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#OrganizationInfoDeprecatedResponse">OrganizationInfoDeprecatedResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Customer

Params Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#BlockUsersParam">BlockUsersParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#CustomerNewResponse">CustomerNewResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#CustomerUpdateResponse">CustomerUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#CustomerListResponse">CustomerListResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#CustomerDeleteResponse">CustomerDeleteResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#CustomerBlockResponse">CustomerBlockResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#CustomerGetInfoResponse">CustomerGetInfoResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#CustomerUnblockResponse">CustomerUnblockResponse</a>

Methods:

- <code title="post /customer/new">client.Customer.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#CustomerService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#CustomerNewParams">CustomerNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#CustomerNewResponse">CustomerNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /customer/update">client.Customer.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#CustomerService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#CustomerUpdateParams">CustomerUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#CustomerUpdateResponse">CustomerUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /customer/list">client.Customer.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#CustomerService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) ([]<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#CustomerListResponse">CustomerListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /customer/delete">client.Customer.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#CustomerService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#CustomerDeleteParams">CustomerDeleteParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#CustomerDeleteResponse">CustomerDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /customer/block">client.Customer.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#CustomerService.Block">Block</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#CustomerBlockParams">CustomerBlockParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#CustomerBlockResponse">CustomerBlockResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /customer/info">client.Customer.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#CustomerService.GetInfo">GetInfo</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#CustomerGetInfoParams">CustomerGetInfoParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#CustomerGetInfoResponse">CustomerGetInfoResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /customer/unblock">client.Customer.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#CustomerService.Unblock">Unblock</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#CustomerUnblockParams">CustomerUnblockParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#CustomerUnblockResponse">CustomerUnblockResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Spend

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#SpendCalculateSpendResponse">SpendCalculateSpendResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#SpendListLogsResponse">SpendListLogsResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#SpendListTagsResponse">SpendListTagsResponse</a>

Methods:

- <code title="post /spend/calculate">client.Spend.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#SpendService.CalculateSpend">CalculateSpend</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#SpendCalculateSpendParams">SpendCalculateSpendParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#SpendCalculateSpendResponse">SpendCalculateSpendResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /spend/logs">client.Spend.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#SpendService.ListLogs">ListLogs</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#SpendListLogsParams">SpendListLogsParams</a>) ([]<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#SpendListLogsResponse">SpendListLogsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /spend/tags">client.Spend.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#SpendService.ListTags">ListTags</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#SpendListTagsParams">SpendListTagsParams</a>) ([]<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#SpendListTagsResponse">SpendListTagsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Global

## Spend

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#GlobalSpendListTagsResponse">GlobalSpendListTagsResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#GlobalSpendResetResponse">GlobalSpendResetResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#GlobalSpendGetReportResponse">GlobalSpendGetReportResponse</a>

Methods:

- <code title="get /global/spend/tags">client.Global.Spend.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#GlobalSpendService.ListTags">ListTags</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#GlobalSpendListTagsParams">GlobalSpendListTagsParams</a>) ([]<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#GlobalSpendListTagsResponse">GlobalSpendListTagsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /global/spend/reset">client.Global.Spend.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#GlobalSpendService.Reset">Reset</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#GlobalSpendResetResponse">GlobalSpendResetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /global/spend/report">client.Global.Spend.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#GlobalSpendService.GetReport">GetReport</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#GlobalSpendGetReportParams">GlobalSpendGetReportParams</a>) ([]<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#GlobalSpendGetReportResponse">GlobalSpendGetReportResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Provider

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#ProviderListBudgetsResponse">ProviderListBudgetsResponse</a>

Methods:

- <code title="get /provider/budgets">client.Provider.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#ProviderService.ListBudgets">ListBudgets</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#ProviderListBudgetsResponse">ProviderListBudgetsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Cache

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#CacheDeleteResponse">CacheDeleteResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#CacheFlushAllResponse">CacheFlushAllResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#CachePingResponse">CachePingResponse</a>

Methods:

- <code title="post /cache/delete">client.Cache.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#CacheService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#CacheDeleteResponse">CacheDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /cache/flushall">client.Cache.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#CacheService.FlushAll">FlushAll</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#CacheFlushAllResponse">CacheFlushAllResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /cache/ping">client.Cache.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#CacheService.Ping">Ping</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#CachePingResponse">CachePingResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Redis

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#CacheRediGetInfoResponse">CacheRediGetInfoResponse</a>

Methods:

- <code title="get /cache/redis/info">client.Cache.Redis.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#CacheRediService.GetInfo">GetInfo</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#CacheRediGetInfoResponse">CacheRediGetInfoResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Guardrails

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#GuardrailListResponse">GuardrailListResponse</a>

Methods:

- <code title="get /guardrails/list">client.Guardrails.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#GuardrailService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#GuardrailListResponse">GuardrailListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Add

Params Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#IPAddressParam">IPAddressParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#AddAddAllowedIPResponse">AddAddAllowedIPResponse</a>

Methods:

- <code title="post /add/allowed_ip">client.Add.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#AddService.AddAllowedIP">AddAllowedIP</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#AddAddAllowedIPParams">AddAddAllowedIPParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#AddAddAllowedIPResponse">AddAddAllowedIPResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Delete

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#DeleteNewAllowedIPResponse">DeleteNewAllowedIPResponse</a>

Methods:

- <code title="post /delete/allowed_ip">client.Delete.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#DeleteService.NewAllowedIP">NewAllowedIP</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#DeleteNewAllowedIPParams">DeleteNewAllowedIPParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#DeleteNewAllowedIPResponse">DeleteNewAllowedIPResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Files

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#FileNewResponse">FileNewResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#FileGetResponse">FileGetResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#FileListResponse">FileListResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#FileDeleteResponse">FileDeleteResponse</a>

Methods:

- <code title="post /{provider}/v1/files">client.Files.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#FileService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, provider <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#FileNewParams">FileNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#FileNewResponse">FileNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /{provider}/v1/files/{file_id}">client.Files.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#FileService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, provider <a href="https://pkg.go.dev/builtin#string">string</a>, fileID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#FileGetResponse">FileGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /{provider}/v1/files">client.Files.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#FileService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, provider <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#FileListParams">FileListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#FileListResponse">FileListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /{provider}/v1/files/{file_id}">client.Files.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#FileService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, provider <a href="https://pkg.go.dev/builtin#string">string</a>, fileID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#FileDeleteResponse">FileDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Content

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#FileContentGetResponse">FileContentGetResponse</a>

Methods:

- <code title="get /{provider}/v1/files/{file_id}/content">client.Files.Content.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#FileContentService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, provider <a href="https://pkg.go.dev/builtin#string">string</a>, fileID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#FileContentGetResponse">FileContentGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Budget

Params Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#BudgetNewParam">BudgetNewParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#BudgetNewResponse">BudgetNewResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#BudgetUpdateResponse">BudgetUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#BudgetListResponse">BudgetListResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#BudgetDeleteResponse">BudgetDeleteResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#BudgetInfoResponse">BudgetInfoResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#BudgetSettingsResponse">BudgetSettingsResponse</a>

Methods:

- <code title="post /budget/new">client.Budget.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#BudgetService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#BudgetNewParams">BudgetNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#BudgetNewResponse">BudgetNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /budget/update">client.Budget.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#BudgetService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#BudgetUpdateParams">BudgetUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#BudgetUpdateResponse">BudgetUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /budget/list">client.Budget.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#BudgetService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#BudgetListResponse">BudgetListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /budget/delete">client.Budget.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#BudgetService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#BudgetDeleteParams">BudgetDeleteParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#BudgetDeleteResponse">BudgetDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /budget/info">client.Budget.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#BudgetService.Info">Info</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#BudgetInfoParams">BudgetInfoParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#BudgetInfoResponse">BudgetInfoResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /budget/settings">client.Budget.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#BudgetService.Settings">Settings</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#BudgetSettingsParams">BudgetSettingsParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go">hanzoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Hanzo-AI-go#BudgetSettingsResponse">BudgetSettingsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
