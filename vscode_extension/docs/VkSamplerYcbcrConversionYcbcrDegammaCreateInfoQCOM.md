# VkSamplerYcbcrConversionYcbcrDegammaCreateInfoQCOM(3) Manual Page

## Name

VkSamplerYcbcrConversionYcbcrDegammaCreateInfoQCOM - Structure
specifying Y′C<sub>B</sub>C<sub>R</sub> degamma parameters



## <a href="#_c_specification" class="anchor"></a>C Specification

Applications **can** enable sRGB to linear conversion for the R, G, and
B components of a Y′C<sub>B</sub>C<sub>R</sub> image during <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#textures-ycbcr-degamma"
target="_blank" rel="noopener">format conversion</a> by including
`VkSamplerYcbcrConversionYcbcrDegammaCreateInfoQCOM` structure in the
`pNext` chain of
[VkSamplerYcbcrConversionCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerYcbcrConversionCreateInfo.html).

The `VkSamplerYcbcrConversionYcbcrDegammaCreateInfoQCOM` structure is
defined as:

``` c
// Provided by VK_QCOM_ycbcr_degamma
typedef struct VkSamplerYcbcrConversionYcbcrDegammaCreateInfoQCOM {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           enableYDegamma;
    VkBool32           enableCbCrDegamma;
} VkSamplerYcbcrConversionYcbcrDegammaCreateInfoQCOM;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `enableYDegamma` indicates <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#textures-ycbcr-degamma"
  target="_blank" rel="noopener">sRGB to linear</a> conversion is
  enabled for the G component.

- `enableCbCrDegamma` indicates <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#textures-ycbcr-degamma"
  target="_blank" rel="noopener">sRGB to linear</a> conversion is
  enabled for the R and B components.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a
  href="#VUID-VkSamplerYcbcrConversionYcbcrDegammaCreateInfoQCOM-sType-sType"
  id="VUID-VkSamplerYcbcrConversionYcbcrDegammaCreateInfoQCOM-sType-sType"></a>
  VUID-VkSamplerYcbcrConversionYcbcrDegammaCreateInfoQCOM-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_SAMPLER_YCBCR_CONVERSION_YCBCR_DEGAMMA_CREATE_INFO_QCOM`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_QCOM_ycbcr_degamma](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_QCOM_ycbcr_degamma.html),
[VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkSamplerYcbcrConversionYcbcrDegammaCreateInfoQCOM"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
