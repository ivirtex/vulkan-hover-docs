# VkSamplerCubicWeightsCreateInfoQCOM(3) Manual Page

## Name

VkSamplerCubicWeightsCreateInfoQCOM - Structure specifying sampler cubic weights



## [](#_c_specification)C Specification

The `VkSamplerCubicWeightsCreateInfoQCOM` structure is defined as:

```c++
// Provided by VK_QCOM_filter_cubic_weights
typedef struct VkSamplerCubicWeightsCreateInfoQCOM {
    VkStructureType             sType;
    const void*                 pNext;
    VkCubicFilterWeightsQCOM    cubicWeights;
} VkSamplerCubicWeightsCreateInfoQCOM;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `cubicWeights` is a [VkCubicFilterWeightsQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCubicFilterWeightsQCOM.html) value controlling which cubic weights are used.

## [](#_description)Description

If the `pNext` chain of [VkSamplerCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerCreateInfo.html) includes a `VkSamplerCubicWeightsCreateInfoQCOM` structure, then that structure specifies which cubic weights are used.

If that structure is not present, `cubicWeights` is considered to be `VK_CUBIC_FILTER_WEIGHTS_CATMULL_ROM_QCOM`.

Valid Usage (Implicit)

- [](#VUID-VkSamplerCubicWeightsCreateInfoQCOM-sType-sType)VUID-VkSamplerCubicWeightsCreateInfoQCOM-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_SAMPLER_CUBIC_WEIGHTS_CREATE_INFO_QCOM`
- [](#VUID-VkSamplerCubicWeightsCreateInfoQCOM-cubicWeights-parameter)VUID-VkSamplerCubicWeightsCreateInfoQCOM-cubicWeights-parameter  
  `cubicWeights` **must** be a valid [VkCubicFilterWeightsQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCubicFilterWeightsQCOM.html) value

## [](#_see_also)See Also

[VK\_QCOM\_filter\_cubic\_weights](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_QCOM_filter_cubic_weights.html), [VkCubicFilterWeightsQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCubicFilterWeightsQCOM.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkSamplerCubicWeightsCreateInfoQCOM).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0