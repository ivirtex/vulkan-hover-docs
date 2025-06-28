# vkCmdSetPerformanceOverrideINTEL(3) Manual Page

## Name

vkCmdSetPerformanceOverrideINTEL - Performance override settings



## [](#_c_specification)C Specification

Some applications might want measure the effect of a set of commands with a different settings. It is possible to override a particular settings using :

```c++
// Provided by VK_INTEL_performance_query
VkResult vkCmdSetPerformanceOverrideINTEL(
    VkCommandBuffer                             commandBuffer,
    const VkPerformanceOverrideInfoINTEL*       pOverrideInfo);
```

## [](#_parameters)Parameters

- `commandBuffer` is the command buffer where the override takes place.
- `pOverrideInfo` is a pointer to a [VkPerformanceOverrideInfoINTEL](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPerformanceOverrideInfoINTEL.html) structure selecting the parameter to override.

## [](#_description)Description

Valid Usage

- [](#VUID-vkCmdSetPerformanceOverrideINTEL-pOverrideInfo-02736)VUID-vkCmdSetPerformanceOverrideINTEL-pOverrideInfo-02736  
  `pOverrideInfo` **must** not be used with a [VkPerformanceOverrideTypeINTEL](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPerformanceOverrideTypeINTEL.html) that is not reported available by `vkGetPerformanceParameterINTEL`

Valid Usage (Implicit)

- [](#VUID-vkCmdSetPerformanceOverrideINTEL-commandBuffer-parameter)VUID-vkCmdSetPerformanceOverrideINTEL-commandBuffer-parameter  
  `commandBuffer` **must** be a valid [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html) handle
- [](#VUID-vkCmdSetPerformanceOverrideINTEL-pOverrideInfo-parameter)VUID-vkCmdSetPerformanceOverrideINTEL-pOverrideInfo-parameter  
  `pOverrideInfo` **must** be a valid pointer to a valid [VkPerformanceOverrideInfoINTEL](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPerformanceOverrideInfoINTEL.html) structure
- [](#VUID-vkCmdSetPerformanceOverrideINTEL-commandBuffer-recording)VUID-vkCmdSetPerformanceOverrideINTEL-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording state](#commandbuffers-lifecycle)
- [](#VUID-vkCmdSetPerformanceOverrideINTEL-commandBuffer-cmdpool)VUID-vkCmdSetPerformanceOverrideINTEL-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must** support graphics, compute, or transfer operations
- [](#VUID-vkCmdSetPerformanceOverrideINTEL-videocoding)VUID-vkCmdSetPerformanceOverrideINTEL-videocoding  
  This command **must** only be called outside of a video coding scope

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
Transfer

State

Return Codes

On success, this command returns

- `VK_SUCCESS`

On failure, this command returns

- `VK_ERROR_TOO_MANY_OBJECTS`
- `VK_ERROR_OUT_OF_HOST_MEMORY`

## [](#_see_also)See Also

[VK\_INTEL\_performance\_query](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_INTEL_performance_query.html), [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html), [VkPerformanceOverrideInfoINTEL](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPerformanceOverrideInfoINTEL.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCmdSetPerformanceOverrideINTEL)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0