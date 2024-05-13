# VkFragmentShadingRateTypeNV(3) Manual Page

## Name

VkFragmentShadingRateTypeNV - Enumeration with fragment shading rate
types



## <a href="#_c_specification" class="anchor"></a>C Specification

The [VkFragmentShadingRateTypeNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFragmentShadingRateTypeNV.html)
enumerated type specifies whether a graphics pipeline gets its pipeline
fragment shading rates and combiners from the
[VkPipelineFragmentShadingRateEnumStateCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineFragmentShadingRateEnumStateCreateInfoNV.html)
structure or the
[VkPipelineFragmentShadingRateStateCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineFragmentShadingRateStateCreateInfoKHR.html)
structure.

``` c
// Provided by VK_NV_fragment_shading_rate_enums
typedef enum VkFragmentShadingRateTypeNV {
    VK_FRAGMENT_SHADING_RATE_TYPE_FRAGMENT_SIZE_NV = 0,
    VK_FRAGMENT_SHADING_RATE_TYPE_ENUMS_NV = 1,
} VkFragmentShadingRateTypeNV;
```

## <a href="#_description" class="anchor"></a>Description

- `VK_FRAGMENT_SHADING_RATE_TYPE_FRAGMENT_SIZE_NV` specifies that a
  graphics pipeline should obtain its pipeline fragment shading rate and
  shading rate combiner state from the
  [VkPipelineFragmentShadingRateStateCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineFragmentShadingRateStateCreateInfoKHR.html)
  structure and that any state specified by the
  [VkPipelineFragmentShadingRateEnumStateCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineFragmentShadingRateEnumStateCreateInfoNV.html)
  structure should be ignored.

- `VK_FRAGMENT_SHADING_RATE_TYPE_ENUMS_NV` specifies that a graphics
  pipeline should obtain its pipeline fragment shading rate and shading
  rate combiner state from the
  [VkPipelineFragmentShadingRateEnumStateCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineFragmentShadingRateEnumStateCreateInfoNV.html)
  structure and that any state specified by the
  [VkPipelineFragmentShadingRateStateCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineFragmentShadingRateStateCreateInfoKHR.html)
  structure should be ignored.

## <a href="#_see_also" class="anchor"></a>See Also

[VK_NV_fragment_shading_rate_enums](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_fragment_shading_rate_enums.html),
[VkPipelineFragmentShadingRateEnumStateCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineFragmentShadingRateEnumStateCreateInfoNV.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkFragmentShadingRateTypeNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
