# VkPipelineFragmentShadingRateStateCreateInfoKHR(3) Manual Page

## Name

VkPipelineFragmentShadingRateStateCreateInfoKHR - Structure specifying parameters controlling the fragment shading rate



## [](#_c_specification)C Specification

The `VkPipelineFragmentShadingRateStateCreateInfoKHR` structure is defined as:

```c++
// Provided by VK_KHR_fragment_shading_rate
typedef struct VkPipelineFragmentShadingRateStateCreateInfoKHR {
    VkStructureType                       sType;
    const void*                           pNext;
    VkExtent2D                            fragmentSize;
    VkFragmentShadingRateCombinerOpKHR    combinerOps[2];
} VkPipelineFragmentShadingRateStateCreateInfoKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `fragmentSize` specifies a [VkExtent2D](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExtent2D.html) structure containing the fragment size used to define the pipeline fragment shading rate for drawing commands using this pipeline.
- `combinerOps` specifies a [VkFragmentShadingRateCombinerOpKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFragmentShadingRateCombinerOpKHR.html) value determining how the [pipeline](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#primsrast-fragment-shading-rate-pipeline), [primitive](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#primsrast-fragment-shading-rate-primitive), and [attachment shading rates](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#primsrast-fragment-shading-rate-attachment) are [combined](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#primsrast-fragment-shading-rate-combining) for fragments generated by drawing commands using the created pipeline.

## [](#_description)Description

If the `pNext` chain of [VkGraphicsPipelineCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkGraphicsPipelineCreateInfo.html) includes a `VkPipelineFragmentShadingRateStateCreateInfoKHR` structure, then that structure includes parameters controlling the pipeline fragment shading rate.

If this structure is not present, `fragmentSize` is considered to be equal to (1,1), and both elements of `combinerOps` are considered to be equal to `VK_FRAGMENT_SHADING_RATE_COMBINER_OP_KEEP_KHR`.

Valid Usage (Implicit)

- [](#VUID-VkPipelineFragmentShadingRateStateCreateInfoKHR-sType-sType)VUID-VkPipelineFragmentShadingRateStateCreateInfoKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PIPELINE_FRAGMENT_SHADING_RATE_STATE_CREATE_INFO_KHR`

## [](#_see_also)See Also

[VK\_KHR\_fragment\_shading\_rate](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_fragment_shading_rate.html), [VkExtent2D](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExtent2D.html), [VkFragmentShadingRateCombinerOpKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFragmentShadingRateCombinerOpKHR.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPipelineFragmentShadingRateStateCreateInfoKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0