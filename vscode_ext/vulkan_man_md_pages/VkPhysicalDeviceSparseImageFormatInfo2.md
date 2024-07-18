# VkPhysicalDeviceSparseImageFormatInfo2(3) Manual Page

## Name

VkPhysicalDeviceSparseImageFormatInfo2 - Structure specifying sparse
image format inputs



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPhysicalDeviceSparseImageFormatInfo2` structure is defined as:

``` c
// Provided by VK_VERSION_1_1
typedef struct VkPhysicalDeviceSparseImageFormatInfo2 {
    VkStructureType          sType;
    const void*              pNext;
    VkFormat                 format;
    VkImageType              type;
    VkSampleCountFlagBits    samples;
    VkImageUsageFlags        usage;
    VkImageTiling            tiling;
} VkPhysicalDeviceSparseImageFormatInfo2;
```

or the equivalent

``` c
// Provided by VK_KHR_get_physical_device_properties2
typedef VkPhysicalDeviceSparseImageFormatInfo2 VkPhysicalDeviceSparseImageFormatInfo2KHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `format` is the image format.

- `type` is the dimensionality of the image.

- `samples` is a [VkSampleCountFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSampleCountFlagBits.html)
  value specifying the number of samples per texel.

- `usage` is a bitmask describing the intended usage of the image.

- `tiling` is the tiling arrangement of the texel blocks in memory.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-VkPhysicalDeviceSparseImageFormatInfo2-samples-01095"
  id="VUID-VkPhysicalDeviceSparseImageFormatInfo2-samples-01095"></a>
  VUID-VkPhysicalDeviceSparseImageFormatInfo2-samples-01095  
  `samples` **must** be a valid
  [VkSampleCountFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSampleCountFlagBits.html) value that is set
  in `VkImageFormatProperties`::`sampleCounts` returned by
  `vkGetPhysicalDeviceImageFormatProperties` with `format`, `type`,
  `tiling`, and `usage` equal to those in this command

Valid Usage (Implicit)

- <a href="#VUID-VkPhysicalDeviceSparseImageFormatInfo2-sType-sType"
  id="VUID-VkPhysicalDeviceSparseImageFormatInfo2-sType-sType"></a>
  VUID-VkPhysicalDeviceSparseImageFormatInfo2-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SPARSE_IMAGE_FORMAT_INFO_2`

- <a href="#VUID-VkPhysicalDeviceSparseImageFormatInfo2-pNext-pNext"
  id="VUID-VkPhysicalDeviceSparseImageFormatInfo2-pNext-pNext"></a>
  VUID-VkPhysicalDeviceSparseImageFormatInfo2-pNext-pNext  
  `pNext` **must** be `NULL`

- <a href="#VUID-VkPhysicalDeviceSparseImageFormatInfo2-format-parameter"
  id="VUID-VkPhysicalDeviceSparseImageFormatInfo2-format-parameter"></a>
  VUID-VkPhysicalDeviceSparseImageFormatInfo2-format-parameter  
  `format` **must** be a valid [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) value

- <a href="#VUID-VkPhysicalDeviceSparseImageFormatInfo2-type-parameter"
  id="VUID-VkPhysicalDeviceSparseImageFormatInfo2-type-parameter"></a>
  VUID-VkPhysicalDeviceSparseImageFormatInfo2-type-parameter  
  `type` **must** be a valid [VkImageType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageType.html) value

- <a href="#VUID-VkPhysicalDeviceSparseImageFormatInfo2-samples-parameter"
  id="VUID-VkPhysicalDeviceSparseImageFormatInfo2-samples-parameter"></a>
  VUID-VkPhysicalDeviceSparseImageFormatInfo2-samples-parameter  
  `samples` **must** be a valid
  [VkSampleCountFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSampleCountFlagBits.html) value

- <a href="#VUID-VkPhysicalDeviceSparseImageFormatInfo2-usage-parameter"
  id="VUID-VkPhysicalDeviceSparseImageFormatInfo2-usage-parameter"></a>
  VUID-VkPhysicalDeviceSparseImageFormatInfo2-usage-parameter  
  `usage` **must** be a valid combination of
  [VkImageUsageFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageUsageFlagBits.html) values

- <a
  href="#VUID-VkPhysicalDeviceSparseImageFormatInfo2-usage-requiredbitmask"
  id="VUID-VkPhysicalDeviceSparseImageFormatInfo2-usage-requiredbitmask"></a>
  VUID-VkPhysicalDeviceSparseImageFormatInfo2-usage-requiredbitmask  
  `usage` **must** not be `0`

- <a href="#VUID-VkPhysicalDeviceSparseImageFormatInfo2-tiling-parameter"
  id="VUID-VkPhysicalDeviceSparseImageFormatInfo2-tiling-parameter"></a>
  VUID-VkPhysicalDeviceSparseImageFormatInfo2-tiling-parameter  
  `tiling` **must** be a valid [VkImageTiling](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageTiling.html) value

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_1](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_1.html), [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html),
[VkImageTiling](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageTiling.html), [VkImageType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageType.html),
[VkImageUsageFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageUsageFlags.html),
[VkSampleCountFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSampleCountFlagBits.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkGetPhysicalDeviceSparseImageFormatProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceSparseImageFormatProperties2.html),
[vkGetPhysicalDeviceSparseImageFormatProperties2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceSparseImageFormatProperties2KHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDeviceSparseImageFormatInfo2"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
