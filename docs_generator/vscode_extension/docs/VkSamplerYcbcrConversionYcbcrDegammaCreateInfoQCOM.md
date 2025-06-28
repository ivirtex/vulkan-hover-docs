# VkSamplerYcbcrConversionYcbcrDegammaCreateInfoQCOM(3) Manual Page

## Name

VkSamplerYcbcrConversionYcbcrDegammaCreateInfoQCOM - Structure specifying Y′CBCR degamma parameters



## [](#_c_specification)C Specification

Applications **can** enable sRGB to linear conversion for the R, G, and B components of a Y′CBCR image during [format conversion](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#textures-ycbcr-degamma) by including `VkSamplerYcbcrConversionYcbcrDegammaCreateInfoQCOM` structure in the `pNext` chain of [VkSamplerYcbcrConversionCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerYcbcrConversionCreateInfo.html).

The `VkSamplerYcbcrConversionYcbcrDegammaCreateInfoQCOM` structure is defined as:

```c++
// Provided by VK_QCOM_ycbcr_degamma
typedef struct VkSamplerYcbcrConversionYcbcrDegammaCreateInfoQCOM {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           enableYDegamma;
    VkBool32           enableCbCrDegamma;
} VkSamplerYcbcrConversionYcbcrDegammaCreateInfoQCOM;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `enableYDegamma` indicates [sRGB to linear](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#textures-ycbcr-degamma) conversion is enabled for the G component.
- `enableCbCrDegamma` indicates [sRGB to linear](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#textures-ycbcr-degamma) conversion is enabled for the R and B components.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkSamplerYcbcrConversionYcbcrDegammaCreateInfoQCOM-sType-sType)VUID-VkSamplerYcbcrConversionYcbcrDegammaCreateInfoQCOM-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_SAMPLER_YCBCR_CONVERSION_YCBCR_DEGAMMA_CREATE_INFO_QCOM`

## [](#_see_also)See Also

[VK\_QCOM\_ycbcr\_degamma](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_QCOM_ycbcr_degamma.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkSamplerYcbcrConversionYcbcrDegammaCreateInfoQCOM)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0