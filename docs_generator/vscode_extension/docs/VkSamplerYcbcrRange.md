# VkSamplerYcbcrRange(3) Manual Page

## Name

VkSamplerYcbcrRange - Range of encoded values in a color space



## [](#_c_specification)C Specification

The [VkSamplerYcbcrRange](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerYcbcrRange.html) enum describes whether color components are encoded using the full range of numerical values or whether values are reserved for headroom and foot room. [VkSamplerYcbcrRange](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerYcbcrRange.html) is defined as:

```c++
// Provided by VK_VERSION_1_1
typedef enum VkSamplerYcbcrRange {
    VK_SAMPLER_YCBCR_RANGE_ITU_FULL = 0,
    VK_SAMPLER_YCBCR_RANGE_ITU_NARROW = 1,
  // Provided by VK_KHR_sampler_ycbcr_conversion
    VK_SAMPLER_YCBCR_RANGE_ITU_FULL_KHR = VK_SAMPLER_YCBCR_RANGE_ITU_FULL,
  // Provided by VK_KHR_sampler_ycbcr_conversion
    VK_SAMPLER_YCBCR_RANGE_ITU_NARROW_KHR = VK_SAMPLER_YCBCR_RANGE_ITU_NARROW,
} VkSamplerYcbcrRange;
```

or the equivalent

```c++
// Provided by VK_KHR_sampler_ycbcr_conversion
typedef VkSamplerYcbcrRange VkSamplerYcbcrRangeKHR;
```

## [](#_description)Description

- `VK_SAMPLER_YCBCR_RANGE_ITU_FULL` specifies that the full range of the encoded values are valid and interpreted according to the ITU “full range” quantization rules.
- `VK_SAMPLER_YCBCR_RANGE_ITU_NARROW` specifies that headroom and foot room are reserved in the numerical range of encoded values, and the remaining values are expanded according to the ITU “narrow range” quantization rules.

The formulae for these conversions is described in the [Sampler Y′CBCR Range Expansion](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#textures-sampler-YCbCr-conversion-rangeexpand) section of the [Image Operations](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#textures) chapter.

No range modification takes place if `ycbcrModel` is `VK_SAMPLER_YCBCR_MODEL_CONVERSION_RGB_IDENTITY`; the `ycbcrRange` field of [VkSamplerYcbcrConversionCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerYcbcrConversionCreateInfo.html) is ignored in this case.

## [](#_see_also)See Also

[VK\_VERSION\_1\_1](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_1.html), [VkAndroidHardwareBufferFormatProperties2ANDROID](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAndroidHardwareBufferFormatProperties2ANDROID.html), [VkAndroidHardwareBufferFormatPropertiesANDROID](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAndroidHardwareBufferFormatPropertiesANDROID.html), [VkBufferCollectionPropertiesFUCHSIA](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferCollectionPropertiesFUCHSIA.html), [VkSamplerYcbcrConversionCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerYcbcrConversionCreateInfo.html), [VkScreenBufferFormatPropertiesQNX](https://registry.khronos.org/vulkan/specs/latest/man/html/VkScreenBufferFormatPropertiesQNX.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkSamplerYcbcrRange)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0