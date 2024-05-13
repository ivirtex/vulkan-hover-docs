# VkOutOfBandQueueTypeInfoNV(3) Manual Page

## Name

VkOutOfBandQueueTypeInfoNV - Structure used to describe the queue that
is being marked as Out of Band



## <a href="#_c_specification" class="anchor"></a>C Specification

The [VkOutOfBandQueueTypeInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkOutOfBandQueueTypeInfoNV.html)
structure is defined as:

``` c
// Provided by VK_NV_low_latency2
typedef struct VkOutOfBandQueueTypeInfoNV {
    VkStructureType           sType;
    const void*               pNext;
    VkOutOfBandQueueTypeNV    queueType;
} VkOutOfBandQueueTypeInfoNV;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `queueType` describes the usage of the queue to be marked as out of
  band.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a href="#VUID-VkOutOfBandQueueTypeInfoNV-sType-sType"
  id="VUID-VkOutOfBandQueueTypeInfoNV-sType-sType"></a>
  VUID-VkOutOfBandQueueTypeInfoNV-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_OUT_OF_BAND_QUEUE_TYPE_INFO_NV`

- <a href="#VUID-VkOutOfBandQueueTypeInfoNV-queueType-parameter"
  id="VUID-VkOutOfBandQueueTypeInfoNV-queueType-parameter"></a>
  VUID-VkOutOfBandQueueTypeInfoNV-queueType-parameter  
  `queueType` **must** be a valid
  [VkOutOfBandQueueTypeNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkOutOfBandQueueTypeNV.html) value

## <a href="#_see_also" class="anchor"></a>See Also

[VK_NV_low_latency2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_low_latency2.html),
[VkOutOfBandQueueTypeNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkOutOfBandQueueTypeNV.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkQueueNotifyOutOfBandNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkQueueNotifyOutOfBandNV.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkOutOfBandQueueTypeInfoNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
