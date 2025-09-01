# VkOutOfBandQueueTypeInfoNV(3) Manual Page

## Name

VkOutOfBandQueueTypeInfoNV - Structure used to describe the queue that is being marked as Out of Band



## [](#_c_specification)C Specification

The [VkOutOfBandQueueTypeInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkOutOfBandQueueTypeInfoNV.html) structure is defined as:

```c++
// Provided by VK_NV_low_latency2
typedef struct VkOutOfBandQueueTypeInfoNV {
    VkStructureType           sType;
    const void*               pNext;
    VkOutOfBandQueueTypeNV    queueType;
} VkOutOfBandQueueTypeInfoNV;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `queueType` describes the usage of the queue to be marked as out of band.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkOutOfBandQueueTypeInfoNV-sType-sType)VUID-VkOutOfBandQueueTypeInfoNV-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_OUT_OF_BAND_QUEUE_TYPE_INFO_NV`
- [](#VUID-VkOutOfBandQueueTypeInfoNV-queueType-parameter)VUID-VkOutOfBandQueueTypeInfoNV-queueType-parameter  
  `queueType` **must** be a valid [VkOutOfBandQueueTypeNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkOutOfBandQueueTypeNV.html) value

## [](#_see_also)See Also

[VK\_NV\_low\_latency2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_low_latency2.html), [VkOutOfBandQueueTypeNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkOutOfBandQueueTypeNV.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkQueueNotifyOutOfBandNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkQueueNotifyOutOfBandNV.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkOutOfBandQueueTypeInfoNV).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0