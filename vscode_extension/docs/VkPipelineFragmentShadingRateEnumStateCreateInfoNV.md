# VkPipelineFragmentShadingRateEnumStateCreateInfoNV(3) Manual Page

## Name

VkPipelineFragmentShadingRateEnumStateCreateInfoNV - Structure specifying parameters controlling the fragment shading rate using rate enums



## [](#_c_specification)C Specification

The `VkPipelineFragmentShadingRateEnumStateCreateInfoNV` structure is defined as:

```c++
// Provided by VK_NV_fragment_shading_rate_enums
typedef struct VkPipelineFragmentShadingRateEnumStateCreateInfoNV {
    VkStructureType                       sType;
    const void*                           pNext;
    VkFragmentShadingRateTypeNV           shadingRateType;
    VkFragmentShadingRateNV               shadingRate;
    VkFragmentShadingRateCombinerOpKHR    combinerOps[2];
} VkPipelineFragmentShadingRateEnumStateCreateInfoNV;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `shadingRateType` specifies a [VkFragmentShadingRateTypeNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFragmentShadingRateTypeNV.html) value indicating whether fragment shading rates are specified using fragment sizes or [VkFragmentShadingRateNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFragmentShadingRateNV.html) enums.
- `shadingRate` specifies a [VkFragmentShadingRateNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFragmentShadingRateNV.html) value indicating the pipeline fragment shading rate.
- `combinerOps` specifies [VkFragmentShadingRateCombinerOpKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFragmentShadingRateCombinerOpKHR.html) values determining how the [pipeline](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#primsrast-fragment-shading-rate-pipeline), [primitive](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#primsrast-fragment-shading-rate-primitive), and [attachment shading rates](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#primsrast-fragment-shading-rate-attachment) are [combined](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#primsrast-fragment-shading-rate-combining) for fragments generated by drawing commands using the created pipeline.

## [](#_description)Description

If the `pNext` chain of [VkGraphicsPipelineCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkGraphicsPipelineCreateInfo.html) includes a `VkPipelineFragmentShadingRateEnumStateCreateInfoNV` structure, then that structure includes parameters controlling the pipeline fragment shading rate.

If this structure is not present, `shadingRateType` is considered to be equal to `VK_FRAGMENT_SHADING_RATE_TYPE_FRAGMENT_SIZE_NV`, `shadingRate` is considered to be equal to `VK_FRAGMENT_SHADING_RATE_1_INVOCATION_PER_PIXEL_NV`, and both elements of `combinerOps` are considered to be equal to `VK_FRAGMENT_SHADING_RATE_COMBINER_OP_KEEP_KHR`.

Valid Usage (Implicit)

- [](#VUID-VkPipelineFragmentShadingRateEnumStateCreateInfoNV-sType-sType)VUID-VkPipelineFragmentShadingRateEnumStateCreateInfoNV-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PIPELINE_FRAGMENT_SHADING_RATE_ENUM_STATE_CREATE_INFO_NV`

## [](#_see_also)See Also

[VK\_NV\_fragment\_shading\_rate\_enums](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_fragment_shading_rate_enums.html), [VkFragmentShadingRateCombinerOpKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFragmentShadingRateCombinerOpKHR.html), [VkFragmentShadingRateNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFragmentShadingRateNV.html), [VkFragmentShadingRateTypeNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFragmentShadingRateTypeNV.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPipelineFragmentShadingRateEnumStateCreateInfoNV)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0