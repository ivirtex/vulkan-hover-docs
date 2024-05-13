# vkQueueSetPerformanceConfigurationINTEL(3) Manual Page

## Name

vkQueueSetPerformanceConfigurationINTEL - Set a performance query



## <a href="#_c_specification" class="anchor"></a>C Specification

To set a performance configuration, call:

``` c
// Provided by VK_INTEL_performance_query
VkResult vkQueueSetPerformanceConfigurationINTEL(
    VkQueue                                     queue,
    VkPerformanceConfigurationINTEL             configuration);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `queue` is the queue on which the configuration will be used.

- `configuration` is the configuration to use.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a href="#VUID-vkQueueSetPerformanceConfigurationINTEL-queue-parameter"
  id="VUID-vkQueueSetPerformanceConfigurationINTEL-queue-parameter"></a>
  VUID-vkQueueSetPerformanceConfigurationINTEL-queue-parameter  
  `queue` **must** be a valid [VkQueue](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueue.html) handle

- <a
  href="#VUID-vkQueueSetPerformanceConfigurationINTEL-configuration-parameter"
  id="VUID-vkQueueSetPerformanceConfigurationINTEL-configuration-parameter"></a>
  VUID-vkQueueSetPerformanceConfigurationINTEL-configuration-parameter  
  `configuration` **must** be a valid
  [VkPerformanceConfigurationINTEL](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPerformanceConfigurationINTEL.html)
  handle

- <a href="#VUID-vkQueueSetPerformanceConfigurationINTEL-commonparent"
  id="VUID-vkQueueSetPerformanceConfigurationINTEL-commonparent"></a>
  VUID-vkQueueSetPerformanceConfigurationINTEL-commonparent  
  Both of `configuration`, and `queue` **must** have been created,
  allocated, or retrieved from the same [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html)

Command Properties

| [Command Buffer Levels](#VkCommandBufferLevel) | [Render Pass Scope](#vkCmdBeginRenderPass) | [Video Coding Scope](#vkCmdBeginVideoCodingKHR) | [Supported Queue Types](#VkQueueFlagBits) | [Command Type](#fundamentals-queueoperation-command-types) |
|------------------------------------------------|--------------------------------------------|-------------------------------------------------|-------------------------------------------|------------------------------------------------------------|
| \-                                             | \-                                         | \-                                              | Any                                       | \-                                                         |

Return Codes

On success, this command returns  
- `VK_SUCCESS`

On failure, this command returns  
- `VK_ERROR_TOO_MANY_OBJECTS`

- `VK_ERROR_OUT_OF_HOST_MEMORY`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_INTEL_performance_query](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_INTEL_performance_query.html),
[VkPerformanceConfigurationINTEL](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPerformanceConfigurationINTEL.html),
[VkQueue](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueue.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkQueueSetPerformanceConfigurationINTEL"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
