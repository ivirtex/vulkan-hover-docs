# vkQueueSetPerformanceConfigurationINTEL(3) Manual Page

## Name

vkQueueSetPerformanceConfigurationINTEL - Set a performance query



## [](#_c_specification)C Specification

To set a performance configuration, call:

```c++
// Provided by VK_INTEL_performance_query
VkResult vkQueueSetPerformanceConfigurationINTEL(
    VkQueue                                     queue,
    VkPerformanceConfigurationINTEL             configuration);
```

## [](#_parameters)Parameters

- `queue` is the queue on which the configuration will be used.
- `configuration` is the configuration to use.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-vkQueueSetPerformanceConfigurationINTEL-queue-parameter)VUID-vkQueueSetPerformanceConfigurationINTEL-queue-parameter  
  `queue` **must** be a valid [VkQueue](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueue.html) handle
- [](#VUID-vkQueueSetPerformanceConfigurationINTEL-configuration-parameter)VUID-vkQueueSetPerformanceConfigurationINTEL-configuration-parameter  
  `configuration` **must** be a valid [VkPerformanceConfigurationINTEL](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPerformanceConfigurationINTEL.html) handle
- [](#VUID-vkQueueSetPerformanceConfigurationINTEL-commonparent)VUID-vkQueueSetPerformanceConfigurationINTEL-commonparent  
  Both of `configuration`, and `queue` **must** have been created, allocated, or retrieved from the same [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html)

Host Synchronization

- Host access to `queue` **must** be externally synchronized

Command Properties

     [Command Buffer Levels](#VkCommandBufferLevel) [Render Pass Scope](#vkCmdBeginRenderPass) [Video Coding Scope](#vkCmdBeginVideoCodingKHR) [Supported Queue Types](#VkQueueFlagBits) [Command Type](#fundamentals-queueoperation-command-types)

\-

\-

\-

Any

\-

Return Codes

On success, this command returns

- `VK_SUCCESS`

On failure, this command returns

- `VK_ERROR_OUT_OF_HOST_MEMORY`
- `VK_ERROR_TOO_MANY_OBJECTS`
- `VK_ERROR_UNKNOWN`
- `VK_ERROR_VALIDATION_FAILED`

## [](#_see_also)See Also

[VK\_INTEL\_performance\_query](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_INTEL_performance_query.html), [VkPerformanceConfigurationINTEL](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPerformanceConfigurationINTEL.html), [VkQueue](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueue.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkQueueSetPerformanceConfigurationINTEL).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0