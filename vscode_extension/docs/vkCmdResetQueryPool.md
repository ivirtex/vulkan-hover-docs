# vkCmdResetQueryPool(3) Manual Page

## Name

vkCmdResetQueryPool - Reset queries in a query pool



## [](#_c_specification)C Specification

To reset a range of queries in a query pool on a queue, call:

```c++
// Provided by VK_VERSION_1_0
void vkCmdResetQueryPool(
    VkCommandBuffer                             commandBuffer,
    VkQueryPool                                 queryPool,
    uint32_t                                    firstQuery,
    uint32_t                                    queryCount);
```

## [](#_parameters)Parameters

- `commandBuffer` is the command buffer into which this command will be recorded.
- `queryPool` is the handle of the query pool managing the queries being reset.
- `firstQuery` is the initial query index to reset.
- `queryCount` is the number of queries to reset.

## [](#_description)Description

When executed on a queue, this command sets the status of query indices \[`firstQuery`, `firstQuery` + `queryCount` - 1] to unavailable.

This command defines an execution dependency between other query commands that reference the same query.

The first [synchronization scope](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-dependencies-scopes) includes all commands which reference the queries in `queryPool` indicated by `firstQuery` and `queryCount` that occur earlier in [submission order](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-submission-order).

The second [synchronization scope](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-dependencies-scopes) includes all commands which reference the queries in `queryPool` indicated by `firstQuery` and `queryCount` that occur later in [submission order](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-submission-order).

The operation of this command happens after the first scope and happens before the second scope.

If the `queryType` used to create `queryPool` was `VK_QUERY_TYPE_PERFORMANCE_QUERY_KHR`, this command sets the status of query indices \[`firstQuery`, `firstQuery` + `queryCount` - 1] to unavailable for each pass of `queryPool`, as indicated by a call to [vkGetPhysicalDeviceQueueFamilyPerformanceQueryPassesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceQueueFamilyPerformanceQueryPassesKHR.html).

Note

Because `vkCmdResetQueryPool` resets all the passes of the indicated queries, applications must not record a `vkCmdResetQueryPool` command for a `queryPool` created with `VK_QUERY_TYPE_PERFORMANCE_QUERY_KHR` in a command buffer that needs to be submitted multiple times as indicated by a call to [vkGetPhysicalDeviceQueueFamilyPerformanceQueryPassesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceQueueFamilyPerformanceQueryPassesKHR.html). Otherwise applications will never be able to complete the recorded queries.

Valid Usage

- [](#VUID-vkCmdResetQueryPool-firstQuery-09436)VUID-vkCmdResetQueryPool-firstQuery-09436  
  `firstQuery` **must** be less than the number of queries in `queryPool`
- [](#VUID-vkCmdResetQueryPool-firstQuery-09437)VUID-vkCmdResetQueryPool-firstQuery-09437  
  The sum of `firstQuery` and `queryCount` **must** be less than or equal to the number of queries in `queryPool`

<!--THE END-->

- [](#VUID-vkCmdResetQueryPool-None-02841)VUID-vkCmdResetQueryPool-None-02841  
  All queries used by the command **must** not be active
- [](#VUID-vkCmdResetQueryPool-firstQuery-02862)VUID-vkCmdResetQueryPool-firstQuery-02862  
  If `queryPool` was created with `VK_QUERY_TYPE_PERFORMANCE_QUERY_KHR`, this command **must** not be recorded in a command buffer that, either directly or through secondary command buffers, also contains begin commands for a query from the set of queries \[`firstQuery`, `firstQuery` + `queryCount` - 1]

Valid Usage (Implicit)

- [](#VUID-vkCmdResetQueryPool-commandBuffer-parameter)VUID-vkCmdResetQueryPool-commandBuffer-parameter  
  `commandBuffer` **must** be a valid [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html) handle
- [](#VUID-vkCmdResetQueryPool-queryPool-parameter)VUID-vkCmdResetQueryPool-queryPool-parameter  
  `queryPool` **must** be a valid [VkQueryPool](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueryPool.html) handle
- [](#VUID-vkCmdResetQueryPool-commandBuffer-recording)VUID-vkCmdResetQueryPool-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording state](#commandbuffers-lifecycle)
- [](#VUID-vkCmdResetQueryPool-commandBuffer-cmdpool)VUID-vkCmdResetQueryPool-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must** support graphics, compute, decode, encode, or optical flow operations
- [](#VUID-vkCmdResetQueryPool-renderpass)VUID-vkCmdResetQueryPool-renderpass  
  This command **must** only be called outside of a render pass instance
- [](#VUID-vkCmdResetQueryPool-videocoding)VUID-vkCmdResetQueryPool-videocoding  
  This command **must** only be called outside of a video coding scope
- [](#VUID-vkCmdResetQueryPool-commonparent)VUID-vkCmdResetQueryPool-commonparent  
  Both of `commandBuffer`, and `queryPool` **must** have been created, allocated, or retrieved from the same [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html)

Host Synchronization

- Host access to `commandBuffer` **must** be externally synchronized
- Host access to the `VkCommandPool` that `commandBuffer` was allocated from **must** be externally synchronized

Command Properties

     [Command Buffer Levels](#VkCommandBufferLevel) [Render Pass Scope](#vkCmdBeginRenderPass) [Video Coding Scope](#vkCmdBeginVideoCodingKHR) [Supported Queue Types](#VkQueueFlagBits) [Command Type](#fundamentals-queueoperation-command-types)

Primary  
Secondary

Outside

Outside

Graphics  
Compute  
Decode  
Encode  
Opticalflow

Action

Conditional Rendering

vkCmdResetQueryPool is not affected by [conditional rendering](#drawing-conditional-rendering)

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html), [VkQueryPool](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueryPool.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCmdResetQueryPool).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0