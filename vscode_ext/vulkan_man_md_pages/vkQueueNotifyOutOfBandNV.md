# vkQueueNotifyOutOfBandNV(3) Manual Page

## Name

vkQueueNotifyOutOfBandNV - Notify out of band queue



## <a href="#_c_specification" class="anchor"></a>C Specification

An application can mark a queue as Out of Band to indicate that all
`vkQueueSubmit` calls on this queue are ignored for latency evaluation
by calling:

``` c
// Provided by VK_NV_low_latency2
void vkQueueNotifyOutOfBandNV(
    VkQueue                                     queue,
    const VkOutOfBandQueueTypeInfoNV*           pQueueTypeInfo);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `queue` is the VkQueue to be marked as out of band.

- `pQueueTypeInfo` is a pointer to a
  [VkOutOfBandQueueTypeInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkOutOfBandQueueTypeInfoNV.html)
  structure specifying the queue type.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a href="#VUID-vkQueueNotifyOutOfBandNV-queue-parameter"
  id="VUID-vkQueueNotifyOutOfBandNV-queue-parameter"></a>
  VUID-vkQueueNotifyOutOfBandNV-queue-parameter  
  `queue` **must** be a valid [VkQueue](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueue.html) handle

- <a href="#VUID-vkQueueNotifyOutOfBandNV-pQueueTypeInfo-parameter"
  id="VUID-vkQueueNotifyOutOfBandNV-pQueueTypeInfo-parameter"></a>
  VUID-vkQueueNotifyOutOfBandNV-pQueueTypeInfo-parameter  
  `pQueueTypeInfo` **must** be a valid pointer to a valid
  [VkOutOfBandQueueTypeInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkOutOfBandQueueTypeInfoNV.html)
  structure

Command Properties

| [Command Buffer Levels](#VkCommandBufferLevel) | [Render Pass Scope](#vkCmdBeginRenderPass) | [Video Coding Scope](#vkCmdBeginVideoCodingKHR) | [Supported Queue Types](#VkQueueFlagBits) | [Command Type](#fundamentals-queueoperation-command-types) |
|------------------------------------------------|--------------------------------------------|-------------------------------------------------|-------------------------------------------|------------------------------------------------------------|
| \-                                             | \-                                         | \-                                              | Any                                       | \-                                                         |

## <a href="#_see_also" class="anchor"></a>See Also

[VK_NV_low_latency2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_low_latency2.html),
[VkOutOfBandQueueTypeInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkOutOfBandQueueTypeInfoNV.html),
[VkQueue](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueue.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkQueueNotifyOutOfBandNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
