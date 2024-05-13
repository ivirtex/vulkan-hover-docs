# VkSamplerYcbcrRange(3) Manual Page

## Name

VkSamplerYcbcrRange - Range of encoded values in a color space



## <a href="#_c_specification" class="anchor"></a>C Specification

The [VkSamplerYcbcrRange](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerYcbcrRange.html) enum describes
whether color components are encoded using the full range of numerical
values or whether values are reserved for headroom and foot room.
[VkSamplerYcbcrRange](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerYcbcrRange.html) is defined as:

``` c
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

``` c
// Provided by VK_KHR_sampler_ycbcr_conversion
typedef VkSamplerYcbcrRange VkSamplerYcbcrRangeKHR;
```

## <a href="#_description" class="anchor"></a>Description

- `VK_SAMPLER_YCBCR_RANGE_ITU_FULL` specifies that the full range of the
  encoded values are valid and interpreted according to the ITU “full
  range” quantization rules.

- `VK_SAMPLER_YCBCR_RANGE_ITU_NARROW` specifies that headroom and foot
  room are reserved in the numerical range of encoded values, and the
  remaining values are expanded according to the ITU “narrow range”
  quantization rules.

The formulae for these conversions is described in the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#textures-sampler-YCbCr-conversion-rangeexpand"
target="_blank" rel="noopener">Sampler Y′C<sub>B</sub>C<sub>R</sub>
Range Expansion</a> section of the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#textures"
target="_blank" rel="noopener">Image Operations</a> chapter.

No range modification takes place if `ycbcrModel` is
`VK_SAMPLER_YCBCR_MODEL_CONVERSION_RGB_IDENTITY`; the `ycbcrRange` field
of
[VkSamplerYcbcrConversionCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerYcbcrConversionCreateInfo.html)
is ignored in this case.

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_1](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_1.html),
[VkAndroidHardwareBufferFormatProperties2ANDROID](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAndroidHardwareBufferFormatProperties2ANDROID.html),
[VkAndroidHardwareBufferFormatPropertiesANDROID](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAndroidHardwareBufferFormatPropertiesANDROID.html),
[VkBufferCollectionPropertiesFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferCollectionPropertiesFUCHSIA.html),
[VkSamplerYcbcrConversionCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerYcbcrConversionCreateInfo.html),
[VkScreenBufferFormatPropertiesQNX](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkScreenBufferFormatPropertiesQNX.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkSamplerYcbcrRange"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
