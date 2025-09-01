# VkSamplerYcbcrConversionInfo(3) Manual Page

## Name

VkSamplerYcbcrConversionInfo - Structure specifying Y′CBCR conversion to a sampler or image view



## [](#_c_specification)C Specification

To create a sampler with Y′CBCR conversion enabled, add a [VkSamplerYcbcrConversionInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerYcbcrConversionInfo.html) structure to the `pNext` chain of the [VkSamplerCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerCreateInfo.html) structure. To create a sampler Y′CBCR conversion, the [`samplerYcbcrConversion`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-samplerYcbcrConversion) feature **must** be enabled. Conversion **must** be fixed at pipeline creation time, through use of a combined image sampler with an immutable sampler in `VkDescriptorSetLayoutBinding`.

A [VkSamplerYcbcrConversionInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerYcbcrConversionInfo.html) **must** be provided for samplers to be used with image views that access `VK_IMAGE_ASPECT_COLOR_BIT` if the format is one of the [formats that require a sampler Y′CBCR conversion](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#formats-requiring-sampler-ycbcr-conversion) , or if the image view has an [external format](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#memory-external-android-hardware-buffer-external-formats) .

The `VkSamplerYcbcrConversionInfo` structure is defined as:

```c++
// Provided by VK_VERSION_1_1
typedef struct VkSamplerYcbcrConversionInfo {
    VkStructureType             sType;
    const void*                 pNext;
    VkSamplerYcbcrConversion    conversion;
} VkSamplerYcbcrConversionInfo;
```

or the equivalent

```c++
// Provided by VK_KHR_sampler_ycbcr_conversion
typedef VkSamplerYcbcrConversionInfo VkSamplerYcbcrConversionInfoKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `conversion` is a [VkSamplerYcbcrConversion](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerYcbcrConversion.html) handle created with [vkCreateSamplerYcbcrConversion](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateSamplerYcbcrConversion.html).

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkSamplerYcbcrConversionInfo-sType-sType)VUID-VkSamplerYcbcrConversionInfo-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_SAMPLER_YCBCR_CONVERSION_INFO`
- [](#VUID-VkSamplerYcbcrConversionInfo-conversion-parameter)VUID-VkSamplerYcbcrConversionInfo-conversion-parameter  
  `conversion` **must** be a valid [VkSamplerYcbcrConversion](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerYcbcrConversion.html) handle

## [](#_see_also)See Also

[VK\_VERSION\_1\_1](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_1.html), [VkSamplerYcbcrConversion](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerYcbcrConversion.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkSamplerYcbcrConversionInfo).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0