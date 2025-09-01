# vkQueueNotifyOutOfBandNV(3) Manual Page

## Name

vkQueueNotifyOutOfBandNV - Notify out of band queue



## [](#_c_specification)C Specification

To mark a queue as *out of band*, so that all `vkQueueSubmit` calls on the queue are ignored for latency evaluation, call:

```c++
// Provided by VK_NV_low_latency2
void vkQueueNotifyOutOfBandNV(
    VkQueue                                     queue,
    const VkOutOfBandQueueTypeInfoNV*           pQueueTypeInfo);
```

## [](#_parameters)Parameters

- `queue` is the VkQueue to be marked as out of band.
- `pQueueTypeInfo` is a pointer to a [VkOutOfBandQueueTypeInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkOutOfBandQueueTypeInfoNV.html) structure specifying the queue type.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-vkQueueNotifyOutOfBandNV-queue-parameter)VUID-vkQueueNotifyOutOfBandNV-queue-parameter  
  `queue` **must** be a valid [VkQueue](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueue.html) handle
- [](#VUID-vkQueueNotifyOutOfBandNV-pQueueTypeInfo-parameter)VUID-vkQueueNotifyOutOfBandNV-pQueueTypeInfo-parameter  
  `pQueueTypeInfo` **must** be a valid pointer to a valid [VkOutOfBandQueueTypeInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkOutOfBandQueueTypeInfoNV.html) structure

Command Properties

     [Command Buffer Levels](#VkCommandBufferLevel) [Render Pass Scope](#vkCmdBeginRenderPass) [Video Coding Scope](#vkCmdBeginVideoCodingKHR) [Supported Queue Types](#VkQueueFlagBits) [Command Type](#fundamentals-queueoperation-command-types)

\-

\-

\-

Any

\-

## [](#_see_also)See Also

[VK\_NV\_low\_latency2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_low_latency2.html), [VkOutOfBandQueueTypeInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkOutOfBandQueueTypeInfoNV.html), [VkQueue](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueue.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkQueueNotifyOutOfBandNV).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0