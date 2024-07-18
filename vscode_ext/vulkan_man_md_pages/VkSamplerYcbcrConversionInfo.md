# VkSamplerYcbcrConversionInfo(3) Manual Page

## Name

VkSamplerYcbcrConversionInfo - Structure specifying
Y′C<sub>B</sub>C<sub>R</sub> conversion to a sampler or image view



## <a href="#_c_specification" class="anchor"></a>C Specification

To create a sampler with Y′C<sub>B</sub>C<sub>R</sub> conversion
enabled, add a
[VkSamplerYcbcrConversionInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerYcbcrConversionInfo.html)
structure to the `pNext` chain of the
[VkSamplerCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerCreateInfo.html) structure. To create a
sampler Y′C<sub>B</sub>C<sub>R</sub> conversion, the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-samplerYcbcrConversion"
target="_blank" rel="noopener"><code>samplerYcbcrConversion</code></a>
feature **must** be enabled. Conversion **must** be fixed at pipeline
creation time, through use of a combined image sampler with an immutable
sampler in `VkDescriptorSetLayoutBinding`.

A [VkSamplerYcbcrConversionInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerYcbcrConversionInfo.html)
**must** be provided for samplers to be used with image views that
access `VK_IMAGE_ASPECT_COLOR_BIT` if the format is one of the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#formats-requiring-sampler-ycbcr-conversion"
target="_blank" rel="noopener">formats that require a sampler
Y′C<sub>B</sub>C<sub>R</sub> conversion</a> , or if the image view has
an <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#memory-external-android-hardware-buffer-external-formats"
target="_blank" rel="noopener">external format</a> .

The `VkSamplerYcbcrConversionInfo` structure is defined as:

``` c
// Provided by VK_VERSION_1_1
typedef struct VkSamplerYcbcrConversionInfo {
    VkStructureType             sType;
    const void*                 pNext;
    VkSamplerYcbcrConversion    conversion;
} VkSamplerYcbcrConversionInfo;
```

or the equivalent

``` c
// Provided by VK_KHR_sampler_ycbcr_conversion
typedef VkSamplerYcbcrConversionInfo VkSamplerYcbcrConversionInfoKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `conversion` is a
  [VkSamplerYcbcrConversion](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerYcbcrConversion.html) handle
  created with
  [vkCreateSamplerYcbcrConversion](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateSamplerYcbcrConversion.html).

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a href="#VUID-VkSamplerYcbcrConversionInfo-sType-sType"
  id="VUID-VkSamplerYcbcrConversionInfo-sType-sType"></a>
  VUID-VkSamplerYcbcrConversionInfo-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_SAMPLER_YCBCR_CONVERSION_INFO`

- <a href="#VUID-VkSamplerYcbcrConversionInfo-conversion-parameter"
  id="VUID-VkSamplerYcbcrConversionInfo-conversion-parameter"></a>
  VUID-VkSamplerYcbcrConversionInfo-conversion-parameter  
  `conversion` **must** be a valid
  [VkSamplerYcbcrConversion](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerYcbcrConversion.html) handle

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_1](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_1.html),
[VkSamplerYcbcrConversion](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerYcbcrConversion.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkSamplerYcbcrConversionInfo"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
