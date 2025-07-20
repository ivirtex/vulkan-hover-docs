# vkCmdEndQueryIndexedEXT(3) Manual Page

## Name

vkCmdEndQueryIndexedEXT - Ends a query



## [](#_c_specification)C Specification

To end an indexed query after the set of desired drawing or dispatching commands is recorded, call:

```c++
// Provided by VK_EXT_transform_feedback
void vkCmdEndQueryIndexedEXT(
    VkCommandBuffer                             commandBuffer,
    VkQueryPool                                 queryPool,
    uint32_t                                    query,
    uint32_t                                    index);
```

## [](#_parameters)Parameters

- `commandBuffer` is the command buffer into which this command will be recorded.
- `queryPool` is the query pool that is managing the results of the query.
- `query` is the query index within the query pool where the result is stored.
- `index` is the query type specific index.

## [](#_description)Description

The command completes the query in `queryPool` identified by `query` and `index`, and marks it as available.

The `vkCmdEndQueryIndexedEXT` command operates the same as the [vkCmdEndQuery](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdEndQuery.html) command, except that it also accepts a query type specific `index` parameter.

This command defines an execution dependency between other query commands that reference the same query index.

The first [synchronization scope](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-dependencies-scopes) includes all commands which reference the queries in `queryPool` indicated by `query` that occur earlier in [submission order](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-submission-order).

The second [synchronization scope](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-dependencies-scopes) includes only the operation of this command.

Valid Usage

- [](#VUID-vkCmdEndQueryIndexedEXT-None-02342)VUID-vkCmdEndQueryIndexedEXT-None-02342  
  All queries used by the command **must** be [active](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#queries-operation-active)
- [](#VUID-vkCmdEndQueryIndexedEXT-query-02343)VUID-vkCmdEndQueryIndexedEXT-query-02343  
  `query` **must** be less than the number of queries in `queryPool`
- [](#VUID-vkCmdEndQueryIndexedEXT-commandBuffer-02344)VUID-vkCmdEndQueryIndexedEXT-commandBuffer-02344  
  `commandBuffer` **must** not be a protected command buffer
- [](#VUID-vkCmdEndQueryIndexedEXT-query-02345)VUID-vkCmdEndQueryIndexedEXT-query-02345  
  If `vkCmdEndQueryIndexedEXT` is called within a render pass instance, the sum of `query` and the number of bits set in the current subpass’s view mask **must** be less than or equal to the number of queries in `queryPool`
- [](#VUID-vkCmdEndQueryIndexedEXT-queryType-06694)VUID-vkCmdEndQueryIndexedEXT-queryType-06694  
  If the `queryType` used to create `queryPool` was `VK_QUERY_TYPE_TRANSFORM_FEEDBACK_STREAM_EXT` or `VK_QUERY_TYPE_PRIMITIVES_GENERATED_EXT`, the `index` parameter **must** be less than `VkPhysicalDeviceTransformFeedbackPropertiesEXT`::`maxTransformFeedbackStreams`
- [](#VUID-vkCmdEndQueryIndexedEXT-queryType-06695)VUID-vkCmdEndQueryIndexedEXT-queryType-06695  
  If the `queryType` used to create `queryPool` was not `VK_QUERY_TYPE_TRANSFORM_FEEDBACK_STREAM_EXT` and not `VK_QUERY_TYPE_PRIMITIVES_GENERATED_EXT`, the `index` **must** be zero
- [](#VUID-vkCmdEndQueryIndexedEXT-queryType-06696)VUID-vkCmdEndQueryIndexedEXT-queryType-06696  
  If the `queryType` used to create `queryPool` was `VK_QUERY_TYPE_TRANSFORM_FEEDBACK_STREAM_EXT` or `VK_QUERY_TYPE_PRIMITIVES_GENERATED_EXT`, `index` **must** equal the `index` used to begin the query

<!--THE END-->

- [](#VUID-vkCmdEndQueryIndexedEXT-None-07007)VUID-vkCmdEndQueryIndexedEXT-None-07007  
  If called within a subpass of a render pass instance, the corresponding `vkCmdBeginQuery`* command **must** have been called previously within the same subpass
- [](#VUID-vkCmdEndQueryIndexedEXT-None-10682)VUID-vkCmdEndQueryIndexedEXT-None-10682  
  This command **must** not be recorded when [per-tile execution model](#renderpass-per-tile-execution-model) is enabled

Valid Usage (Implicit)

- [](#VUID-vkCmdEndQueryIndexedEXT-commandBuffer-parameter)VUID-vkCmdEndQueryIndexedEXT-commandBuffer-parameter  
  `commandBuffer` **must** be a valid [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html) handle
- [](#VUID-vkCmdEndQueryIndexedEXT-queryPool-parameter)VUID-vkCmdEndQueryIndexedEXT-queryPool-parameter  
  `queryPool` **must** be a valid [VkQueryPool](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueryPool.html) handle
- [](#VUID-vkCmdEndQueryIndexedEXT-commandBuffer-recording)VUID-vkCmdEndQueryIndexedEXT-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording state](#commandbuffers-lifecycle)
- [](#VUID-vkCmdEndQueryIndexedEXT-commandBuffer-cmdpool)VUID-vkCmdEndQueryIndexedEXT-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must** support graphics, compute, decode, or encode operations
- [](#VUID-vkCmdEndQueryIndexedEXT-videocoding)VUID-vkCmdEndQueryIndexedEXT-videocoding  
  This command **must** only be called outside of a video coding scope
- [](#VUID-vkCmdEndQueryIndexedEXT-commonparent)VUID-vkCmdEndQueryIndexedEXT-commonparent  
  Both of `commandBuffer`, and `queryPool` **must** have been created, allocated, or retrieved from the same [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html)

Host Synchronization

- Host access to `commandBuffer` **must** be externally synchronized
- Host access to the `VkCommandPool` that `commandBuffer` was allocated from **must** be externally synchronized

Command Properties

     [Command Buffer Levels](#VkCommandBufferLevel) [Render Pass Scope](#vkCmdBeginRenderPass) [Video Coding Scope](#vkCmdBeginVideoCodingKHR) [Supported Queue Types](#VkQueueFlagBits) [Command Type](#fundamentals-queueoperation-command-types)

Primary  
Secondary

Both

Outside

Graphics  
Compute  
Decode  
Encode

Action  
State

Conditional Rendering

vkCmdEndQueryIndexedEXT is not affected by [conditional rendering](#drawing-conditional-rendering)

## [](#_see_also)See Also

[VK\_EXT\_transform\_feedback](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_transform_feedback.html), [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html), [VkQueryPool](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueryPool.html), [vkCmdBeginQuery](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdBeginQuery.html), [vkCmdBeginQueryIndexedEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdBeginQueryIndexedEXT.html), [vkCmdEndQuery](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdEndQuery.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCmdEndQueryIndexedEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0