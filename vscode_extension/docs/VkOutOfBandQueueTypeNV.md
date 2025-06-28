# VkOutOfBandQueueTypeNV(3) Manual Page

## Name

VkOutOfBandQueueTypeNV - Type of out of band queue



## [](#_c_specification)C Specification

The [VkOutOfBandQueueTypeNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkOutOfBandQueueTypeNV.html) enum is defined as:

```c++
// Provided by VK_NV_low_latency2
typedef enum VkOutOfBandQueueTypeNV {
    VK_OUT_OF_BAND_QUEUE_TYPE_RENDER_NV = 0,
    VK_OUT_OF_BAND_QUEUE_TYPE_PRESENT_NV = 1,
} VkOutOfBandQueueTypeNV;
```

## [](#_description)Description

The members of the [VkOutOfBandQueueTypeNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkOutOfBandQueueTypeNV.html) are used to describe the queue type in [VkOutOfBandQueueTypeInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkOutOfBandQueueTypeInfoNV.html) as described below:

- `VK_OUT_OF_BAND_QUEUE_TYPE_RENDER_NV` specifies that work will be submitted to this queue.
- `VK_OUT_OF_BAND_QUEUE_TYPE_PRESENT_NV` specifies that this queue will be presented from.

## [](#_see_also)See Also

[VK\_NV\_low\_latency2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_low_latency2.html), [VkOutOfBandQueueTypeInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkOutOfBandQueueTypeInfoNV.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkOutOfBandQueueTypeNV)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0