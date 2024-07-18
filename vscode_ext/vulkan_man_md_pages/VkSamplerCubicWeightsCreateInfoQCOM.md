# VkSamplerCubicWeightsCreateInfoQCOM(3) Manual Page

## Name

VkSamplerCubicWeightsCreateInfoQCOM - Structure specifying sampler cubic
weights



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkSamplerCubicWeightsCreateInfoQCOM` structure is defined as:

``` c
// Provided by VK_QCOM_filter_cubic_weights
typedef struct VkSamplerCubicWeightsCreateInfoQCOM {
    VkStructureType             sType;
    const void*                 pNext;
    VkCubicFilterWeightsQCOM    cubicWeights;
} VkSamplerCubicWeightsCreateInfoQCOM;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `cubicWeights` is a
  [VkCubicFilterWeightsQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCubicFilterWeightsQCOM.html) value
  controlling which cubic weights are used.

## <a href="#_description" class="anchor"></a>Description

If the `pNext` chain of [VkSamplerCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerCreateInfo.html)
includes a `VkSamplerCubicWeightsCreateInfoQCOM` structure, then that
structure specifies which cubic weights are used.

If that structure is not present, `cubicWeights` is considered to be
`VK_CUBIC_FILTER_WEIGHTS_CATMULL_ROM_QCOM`.

Valid Usage (Implicit)

- <a href="#VUID-VkSamplerCubicWeightsCreateInfoQCOM-sType-sType"
  id="VUID-VkSamplerCubicWeightsCreateInfoQCOM-sType-sType"></a>
  VUID-VkSamplerCubicWeightsCreateInfoQCOM-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_SAMPLER_CUBIC_WEIGHTS_CREATE_INFO_QCOM`

- <a
  href="#VUID-VkSamplerCubicWeightsCreateInfoQCOM-cubicWeights-parameter"
  id="VUID-VkSamplerCubicWeightsCreateInfoQCOM-cubicWeights-parameter"></a>
  VUID-VkSamplerCubicWeightsCreateInfoQCOM-cubicWeights-parameter  
  `cubicWeights` **must** be a valid
  [VkCubicFilterWeightsQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCubicFilterWeightsQCOM.html) value

## <a href="#_see_also" class="anchor"></a>See Also

[VK_QCOM_filter_cubic_weights](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_QCOM_filter_cubic_weights.html),
[VkCubicFilterWeightsQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCubicFilterWeightsQCOM.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkSamplerCubicWeightsCreateInfoQCOM"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
