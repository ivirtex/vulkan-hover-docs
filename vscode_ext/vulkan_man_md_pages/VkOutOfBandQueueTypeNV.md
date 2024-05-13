# VkOutOfBandQueueTypeNV(3) Manual Page

## Name

VkOutOfBandQueueTypeNV - Type of out of band queue



## <a href="#_c_specification" class="anchor"></a>C Specification

The [VkOutOfBandQueueTypeNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkOutOfBandQueueTypeNV.html) enum is
defined as:

``` c
// Provided by VK_NV_low_latency2
typedef enum VkOutOfBandQueueTypeNV {
    VK_OUT_OF_BAND_QUEUE_TYPE_RENDER_NV = 0,
    VK_OUT_OF_BAND_QUEUE_TYPE_PRESENT_NV = 1,
} VkOutOfBandQueueTypeNV;
```

## <a href="#_description" class="anchor"></a>Description

The members of the [VkOutOfBandQueueTypeNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkOutOfBandQueueTypeNV.html)
are used to describe the queue type in
[VkOutOfBandQueueTypeInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkOutOfBandQueueTypeInfoNV.html) as
described below:

- `VK_OUT_OF_BAND_QUEUE_TYPE_RENDER_NV` indicates that work will be
  submitted to this queue.

- `VK_OUT_OF_BAND_QUEUE_TYPE_PRESENT_NV` indicates that this queue will
  be presented from.

## <a href="#_see_also" class="anchor"></a>See Also

[VK_NV_low_latency2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_low_latency2.html),
[VkOutOfBandQueueTypeInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkOutOfBandQueueTypeInfoNV.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkOutOfBandQueueTypeNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
