# VkPhysicalDeviceSparseImageFormatInfo2(3) Manual Page

## Name

VkPhysicalDeviceSparseImageFormatInfo2 - Structure specifying sparse image format inputs



## [](#_c_specification)C Specification

The `VkPhysicalDeviceSparseImageFormatInfo2` structure is defined as:

```c++
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

```c++
// Provided by VK_KHR_get_physical_device_properties2
typedef VkPhysicalDeviceSparseImageFormatInfo2 VkPhysicalDeviceSparseImageFormatInfo2KHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `format` is the image format.
- `type` is the dimensionality of the image.
- `samples` is a [VkSampleCountFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSampleCountFlagBits.html) value specifying the number of samples per texel.
- `usage` is a bitmask describing the intended usage of the image.
- `tiling` is the tiling arrangement of the texel blocks in memory.

## [](#_description)Description

Valid Usage

- [](#VUID-VkPhysicalDeviceSparseImageFormatInfo2-samples-01095)VUID-VkPhysicalDeviceSparseImageFormatInfo2-samples-01095  
  `samples` **must** be a valid [VkSampleCountFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSampleCountFlagBits.html) value that is set in `VkImageFormatProperties`::`sampleCounts` returned by `vkGetPhysicalDeviceImageFormatProperties` with `format`, `type`, `tiling`, and `usage` equal to those in this command

Valid Usage (Implicit)

- [](#VUID-VkPhysicalDeviceSparseImageFormatInfo2-sType-sType)VUID-VkPhysicalDeviceSparseImageFormatInfo2-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SPARSE_IMAGE_FORMAT_INFO_2`
- [](#VUID-VkPhysicalDeviceSparseImageFormatInfo2-pNext-pNext)VUID-VkPhysicalDeviceSparseImageFormatInfo2-pNext-pNext  
  `pNext` **must** be `NULL`
- [](#VUID-VkPhysicalDeviceSparseImageFormatInfo2-format-parameter)VUID-VkPhysicalDeviceSparseImageFormatInfo2-format-parameter  
  `format` **must** be a valid [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) value
- [](#VUID-VkPhysicalDeviceSparseImageFormatInfo2-type-parameter)VUID-VkPhysicalDeviceSparseImageFormatInfo2-type-parameter  
  `type` **must** be a valid [VkImageType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageType.html) value
- [](#VUID-VkPhysicalDeviceSparseImageFormatInfo2-samples-parameter)VUID-VkPhysicalDeviceSparseImageFormatInfo2-samples-parameter  
  `samples` **must** be a valid [VkSampleCountFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSampleCountFlagBits.html) value
- [](#VUID-VkPhysicalDeviceSparseImageFormatInfo2-usage-parameter)VUID-VkPhysicalDeviceSparseImageFormatInfo2-usage-parameter  
  `usage` **must** be a valid combination of [VkImageUsageFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageUsageFlagBits.html) values
- [](#VUID-VkPhysicalDeviceSparseImageFormatInfo2-usage-requiredbitmask)VUID-VkPhysicalDeviceSparseImageFormatInfo2-usage-requiredbitmask  
  `usage` **must** not be `0`
- [](#VUID-VkPhysicalDeviceSparseImageFormatInfo2-tiling-parameter)VUID-VkPhysicalDeviceSparseImageFormatInfo2-tiling-parameter  
  `tiling` **must** be a valid [VkImageTiling](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageTiling.html) value

## [](#_see_also)See Also

[VK\_VERSION\_1\_1](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_1.html), [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html), [VkImageTiling](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageTiling.html), [VkImageType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageType.html), [VkImageUsageFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageUsageFlags.html), [VkSampleCountFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSampleCountFlagBits.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkGetPhysicalDeviceSparseImageFormatProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceSparseImageFormatProperties2.html), [vkGetPhysicalDeviceSparseImageFormatProperties2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceSparseImageFormatProperties2KHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPhysicalDeviceSparseImageFormatInfo2)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0