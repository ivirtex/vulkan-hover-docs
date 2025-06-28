# VkSamplerYcbcrConversionImageFormatProperties(3) Manual Page

## Name

VkSamplerYcbcrConversionImageFormatProperties - Structure specifying combined image sampler descriptor count for multi-planar images



## [](#_c_specification)C Specification

To determine the number of combined image samplers required to support a [multi-planar format](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#formats-multiplanar), add [VkSamplerYcbcrConversionImageFormatProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerYcbcrConversionImageFormatProperties.html) to the `pNext` chain of the [VkImageFormatProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageFormatProperties2.html) structure in a call to `vkGetPhysicalDeviceImageFormatProperties2`.

The `VkSamplerYcbcrConversionImageFormatProperties` structure is defined as:

```c++
// Provided by VK_VERSION_1_1
typedef struct VkSamplerYcbcrConversionImageFormatProperties {
    VkStructureType    sType;
    void*              pNext;
    uint32_t           combinedImageSamplerDescriptorCount;
} VkSamplerYcbcrConversionImageFormatProperties;
```

or the equivalent

```c++
// Provided by VK_KHR_sampler_ycbcr_conversion
typedef VkSamplerYcbcrConversionImageFormatProperties VkSamplerYcbcrConversionImageFormatPropertiesKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `combinedImageSamplerDescriptorCount` is the number of combined image sampler descriptors that the implementation uses to access the format.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkSamplerYcbcrConversionImageFormatProperties-sType-sType)VUID-VkSamplerYcbcrConversionImageFormatProperties-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_SAMPLER_YCBCR_CONVERSION_IMAGE_FORMAT_PROPERTIES`

## [](#_see_also)See Also

[VK\_VERSION\_1\_1](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_1.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkSamplerYcbcrConversionImageFormatProperties)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0