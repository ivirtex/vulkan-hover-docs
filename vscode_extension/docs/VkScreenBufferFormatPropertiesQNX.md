# VkScreenBufferFormatPropertiesQNX(3) Manual Page

## Name

VkScreenBufferFormatPropertiesQNX - Structure describing the image format properties of a QNX Screen buffer



## [](#_c_specification)C Specification

To obtain format properties of a QNX Screen buffer, include a `VkScreenBufferFormatPropertiesQNX` structure in the `pNext` chain of the [VkScreenBufferPropertiesQNX](https://registry.khronos.org/vulkan/specs/latest/man/html/VkScreenBufferPropertiesQNX.html) structure passed to [vkGetScreenBufferPropertiesQNX](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetScreenBufferPropertiesQNX.html). This structure is defined as:

```c++
// Provided by VK_QNX_external_memory_screen_buffer
typedef struct VkScreenBufferFormatPropertiesQNX {
    VkStructureType                  sType;
    void*                            pNext;
    VkFormat                         format;
    uint64_t                         externalFormat;
    uint64_t                         screenUsage;
    VkFormatFeatureFlags             formatFeatures;
    VkComponentMapping               samplerYcbcrConversionComponents;
    VkSamplerYcbcrModelConversion    suggestedYcbcrModel;
    VkSamplerYcbcrRange              suggestedYcbcrRange;
    VkChromaLocation                 suggestedXChromaOffset;
    VkChromaLocation                 suggestedYChromaOffset;
} VkScreenBufferFormatPropertiesQNX;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `format` is the Vulkan format corresponding to the Screen buffer’s format or `VK_FORMAT_UNDEFINED` if there is not an equivalent Vulkan format.
- `externalFormat` is an implementation-defined external format identifier for use with [VkExternalFormatQNX](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalFormatQNX.html). It **must** not be zero.
- `screenUsage` is an implementation-defined external usage identifier for the QNX Screen buffer.
- `formatFeatures` describes the capabilities of this external format when used with an image bound to memory imported from `buffer`.
- `samplerYcbcrConversionComponents` is the component swizzle that **should** be used in [VkSamplerYcbcrConversionCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerYcbcrConversionCreateInfo.html).
- `suggestedYcbcrModel` is a suggested color model to use in the [VkSamplerYcbcrConversionCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerYcbcrConversionCreateInfo.html).
- `suggestedYcbcrRange` is a suggested numerical value range to use in [VkSamplerYcbcrConversionCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerYcbcrConversionCreateInfo.html).
- `suggestedXChromaOffset` is a suggested X chroma offset to use in [VkSamplerYcbcrConversionCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerYcbcrConversionCreateInfo.html).
- `suggestedYChromaOffset` is a suggested Y chroma offset to use in [VkSamplerYcbcrConversionCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerYcbcrConversionCreateInfo.html).

## [](#_description)Description

If the QNX Screen buffer has one of the formats listed in the [QNX Screen Format Equivalence table](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#memory-external-qnx-screen-buffer-formats), then `format` **must** have the equivalent Vulkan format listed in the table. Otherwise, `format` **may** be `VK_FORMAT_UNDEFINED`, indicating the QNX Screen buffer **can** only be used with an external format. The `formatFeatures` member **must** include `VK_FORMAT_FEATURE_SAMPLED_IMAGE_BIT` and **should** include `VK_FORMAT_FEATURE_SAMPLED_IMAGE_FILTER_LINEAR_BIT` and `VK_FORMAT_FEATURE_SAMPLED_IMAGE_YCBCR_CONVERSION_LINEAR_FILTER_BIT`.

Valid Usage (Implicit)

- [](#VUID-VkScreenBufferFormatPropertiesQNX-sType-sType)VUID-VkScreenBufferFormatPropertiesQNX-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_SCREEN_BUFFER_FORMAT_PROPERTIES_QNX`

## [](#_see_also)See Also

[VK\_QNX\_external\_memory\_screen\_buffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_QNX_external_memory_screen_buffer.html), [VkChromaLocation](https://registry.khronos.org/vulkan/specs/latest/man/html/VkChromaLocation.html), [VkComponentMapping](https://registry.khronos.org/vulkan/specs/latest/man/html/VkComponentMapping.html), [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html), [VkFormatFeatureFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormatFeatureFlags.html), [VkSamplerYcbcrModelConversion](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerYcbcrModelConversion.html), [VkSamplerYcbcrRange](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerYcbcrRange.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkScreenBufferFormatPropertiesQNX)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0