# VkFragmentShadingRateTypeNV(3) Manual Page

## Name

VkFragmentShadingRateTypeNV - Enumeration with fragment shading rate types



## [](#_c_specification)C Specification

The [VkFragmentShadingRateTypeNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFragmentShadingRateTypeNV.html) enumerated type specifies whether a graphics pipeline gets its pipeline fragment shading rates and combiners from the [VkPipelineFragmentShadingRateEnumStateCreateInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineFragmentShadingRateEnumStateCreateInfoNV.html) structure or the [VkPipelineFragmentShadingRateStateCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineFragmentShadingRateStateCreateInfoKHR.html) structure.

```c++
// Provided by VK_NV_fragment_shading_rate_enums
typedef enum VkFragmentShadingRateTypeNV {
    VK_FRAGMENT_SHADING_RATE_TYPE_FRAGMENT_SIZE_NV = 0,
    VK_FRAGMENT_SHADING_RATE_TYPE_ENUMS_NV = 1,
} VkFragmentShadingRateTypeNV;
```

## [](#_description)Description

- `VK_FRAGMENT_SHADING_RATE_TYPE_FRAGMENT_SIZE_NV` specifies that a graphics pipeline should obtain its pipeline fragment shading rate and shading rate combiner state from the [VkPipelineFragmentShadingRateStateCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineFragmentShadingRateStateCreateInfoKHR.html) structure and that any state specified by the [VkPipelineFragmentShadingRateEnumStateCreateInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineFragmentShadingRateEnumStateCreateInfoNV.html) structure should be ignored.
- `VK_FRAGMENT_SHADING_RATE_TYPE_ENUMS_NV` specifies that a graphics pipeline should obtain its pipeline fragment shading rate and shading rate combiner state from the [VkPipelineFragmentShadingRateEnumStateCreateInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineFragmentShadingRateEnumStateCreateInfoNV.html) structure and that any state specified by the [VkPipelineFragmentShadingRateStateCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineFragmentShadingRateStateCreateInfoKHR.html) structure should be ignored.

## [](#_see_also)See Also

[VK\_NV\_fragment\_shading\_rate\_enums](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_fragment_shading_rate_enums.html), [VkPipelineFragmentShadingRateEnumStateCreateInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineFragmentShadingRateEnumStateCreateInfoNV.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkFragmentShadingRateTypeNV).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0