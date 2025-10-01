# vkCmdEndQuery(3) Manual Page

## Name

vkCmdEndQuery - Ends a query



## [](#_c_specification)C Specification

To end a query after the set of desired drawing or dispatching commands is executed, call:

```c++
// Provided by VK_VERSION_1_0
void vkCmdEndQuery(
    VkCommandBuffer                             commandBuffer,
    VkQueryPool                                 queryPool,
    uint32_t                                    query);
```

## [](#_parameters)Parameters

- `commandBuffer` is the command buffer into which this command will be recorded.
- `queryPool` is the query pool that is managing the results of the query.
- `query` is the query index within the query pool where the result is stored.

## [](#_description)Description

The command completes the query in `queryPool` identified by `query`, and marks it as available.

This command defines an execution dependency between other query commands that reference the same query.

The first [synchronization scope](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-dependencies-scopes) includes all commands which reference the queries in `queryPool` indicated by `query` that occur earlier in [submission order](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-submission-order).

The second [synchronization scope](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-dependencies-scopes) includes only the operation of this command.

Calling `vkCmdEndQuery` is equivalent to calling [vkCmdEndQueryIndexedEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdEndQueryIndexedEXT.html) with the `index` parameter set to zero.

Valid Usage

- [](#VUID-vkCmdEndQuery-None-01923)VUID-vkCmdEndQuery-None-01923  
  All queries used by the command **must** be [active](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#queries-operation-active)
- [](#VUID-vkCmdEndQuery-query-00810)VUID-vkCmdEndQuery-query-00810  
  `query` **must** be less than the number of queries in `queryPool`
- [](#VUID-vkCmdEndQuery-commandBuffer-01886)VUID-vkCmdEndQuery-commandBuffer-01886  
  `commandBuffer` **must** not be a protected command buffer
- [](#VUID-vkCmdEndQuery-query-00812)VUID-vkCmdEndQuery-query-00812  
  If `vkCmdEndQuery` is called within a render pass instance, the sum of `query` and the number of bits set in the current subpass’s view mask **must** be less than or equal to the number of queries in `queryPool`
- [](#VUID-vkCmdEndQuery-queryPool-03227)VUID-vkCmdEndQuery-queryPool-03227  
  If `queryPool` was created with a `queryType` of `VK_QUERY_TYPE_PERFORMANCE_QUERY_KHR` and one or more of the counters used to create `queryPool` was `VK_PERFORMANCE_COUNTER_SCOPE_COMMAND_BUFFER_KHR`, the [vkCmdEndQuery](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdEndQuery.html) **must** be the last recorded command in `commandBuffer`
- [](#VUID-vkCmdEndQuery-queryPool-03228)VUID-vkCmdEndQuery-queryPool-03228  
  If `queryPool` was created with a `queryType` of `VK_QUERY_TYPE_PERFORMANCE_QUERY_KHR` and one or more of the counters used to create `queryPool` was `VK_PERFORMANCE_COUNTER_SCOPE_RENDER_PASS_KHR`, the [vkCmdEndQuery](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdEndQuery.html) **must** not be recorded within a render pass instance

<!--THE END-->

- [](#VUID-vkCmdEndQuery-None-07007)VUID-vkCmdEndQuery-None-07007  
  If called within a subpass of a render pass instance, the corresponding `vkCmdBeginQuery`* command **must** have been called previously within the same subpass
- [](#VUID-vkCmdEndQuery-None-10682)VUID-vkCmdEndQuery-None-10682  
  This command **must** not be recorded when [per-tile execution model](#renderpass-per-tile-execution-model) is enabled

Valid Usage (Implicit)

- [](#VUID-vkCmdEndQuery-commandBuffer-parameter)VUID-vkCmdEndQuery-commandBuffer-parameter  
  `commandBuffer` **must** be a valid [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html) handle
- [](#VUID-vkCmdEndQuery-queryPool-parameter)VUID-vkCmdEndQuery-queryPool-parameter  
  `queryPool` **must** be a valid [VkQueryPool](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueryPool.html) handle
- [](#VUID-vkCmdEndQuery-commandBuffer-recording)VUID-vkCmdEndQuery-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording state](#commandbuffers-lifecycle)
- [](#VUID-vkCmdEndQuery-commandBuffer-cmdpool)VUID-vkCmdEndQuery-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must** support VK\_QUEUE\_COMPUTE\_BIT, VK\_QUEUE\_GRAPHICS\_BIT, VK\_QUEUE\_VIDEO\_DECODE\_BIT\_KHR, or VK\_QUEUE\_VIDEO\_ENCODE\_BIT\_KHR operations
- [](#VUID-vkCmdEndQuery-commonparent)VUID-vkCmdEndQuery-commonparent  
  Both of `commandBuffer`, and `queryPool` **must** have been created, allocated, or retrieved from the same [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html)

Host Synchronization

- Host access to `commandBuffer` **must** be externally synchronized
- Host access to the `VkCommandPool` that `commandBuffer` was allocated from **must** be externally synchronized

Command Properties

     [Command Buffer Levels](#VkCommandBufferLevel) [Render Pass Scope](#vkCmdBeginRenderPass) [Video Coding Scope](#vkCmdBeginVideoCodingKHR) [Supported Queue Types](#VkQueueFlagBits) [Command Type](#fundamentals-queueoperation-command-types)

Primary  
Secondary

Both

Both

VK\_QUEUE\_COMPUTE\_BIT  
VK\_QUEUE\_GRAPHICS\_BIT  
VK\_QUEUE\_VIDEO\_DECODE\_BIT\_KHR  
VK\_QUEUE\_VIDEO\_ENCODE\_BIT\_KHR

Action  
State

Conditional Rendering

vkCmdEndQuery is not affected by [conditional rendering](#drawing-conditional-rendering)

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html), [VkQueryPool](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueryPool.html), [vkCmdBeginQuery](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdBeginQuery.html), [vkCmdBeginQueryIndexedEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdBeginQueryIndexedEXT.html), [vkCmdEndQueryIndexedEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdEndQueryIndexedEXT.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCmdEndQuery).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0