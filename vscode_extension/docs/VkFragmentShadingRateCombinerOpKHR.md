# VkFragmentShadingRateCombinerOpKHR(3) Manual Page

## Name

VkFragmentShadingRateCombinerOpKHR - Control how fragment shading rates are combined



## [](#_c_specification)C Specification

The equation used for each combiner operation is defined by `VkFragmentShadingRateCombinerOpKHR`:

```c++
// Provided by VK_KHR_fragment_shading_rate
typedef enum VkFragmentShadingRateCombinerOpKHR {
    VK_FRAGMENT_SHADING_RATE_COMBINER_OP_KEEP_KHR = 0,
    VK_FRAGMENT_SHADING_RATE_COMBINER_OP_REPLACE_KHR = 1,
    VK_FRAGMENT_SHADING_RATE_COMBINER_OP_MIN_KHR = 2,
    VK_FRAGMENT_SHADING_RATE_COMBINER_OP_MAX_KHR = 3,
    VK_FRAGMENT_SHADING_RATE_COMBINER_OP_MUL_KHR = 4,
} VkFragmentShadingRateCombinerOpKHR;
```

## [](#_description)Description

- `VK_FRAGMENT_SHADING_RATE_COMBINER_OP_KEEP_KHR` specifies a combiner operation of combine(Axy,Bxy) = Axy.
- `VK_FRAGMENT_SHADING_RATE_COMBINER_OP_REPLACE_KHR` specifies a combiner operation of combine(Axy,Bxy) = Bxy.
- `VK_FRAGMENT_SHADING_RATE_COMBINER_OP_MIN_KHR` specifies a combiner operation of combine(Axy,Bxy) = min(Axy,Bxy).
- `VK_FRAGMENT_SHADING_RATE_COMBINER_OP_MAX_KHR` specifies a combiner operation of combine(Axy,Bxy) = max(Axy,Bxy).
- `VK_FRAGMENT_SHADING_RATE_COMBINER_OP_MUL_KHR` specifies a combiner operation of combine(Axy,Bxy) = Axy\*Bxy.

where combine(Axy,Bxy) is the combine operation, and Axy and Bxy are the inputs to the operation.

If [`fragmentShadingRateStrictMultiplyCombiner`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-fragmentShadingRateStrictMultiplyCombiner) is `VK_FALSE`, using `VK_FRAGMENT_SHADING_RATE_COMBINER_OP_MUL_KHR` with values of 1 for both A and B in the same dimension results in the value 2 being produced for that dimension. See the definition of [`fragmentShadingRateStrictMultiplyCombiner`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-fragmentShadingRateStrictMultiplyCombiner) for more information.

These operations are performed in a component-wise fashion.

## [](#_see_also)See Also

[VK\_KHR\_fragment\_shading\_rate](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_fragment_shading_rate.html), [VkPipelineFragmentShadingRateEnumStateCreateInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineFragmentShadingRateEnumStateCreateInfoNV.html), [VkPipelineFragmentShadingRateStateCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineFragmentShadingRateStateCreateInfoKHR.html), [vkCmdSetFragmentShadingRateEnumNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetFragmentShadingRateEnumNV.html), [vkCmdSetFragmentShadingRateKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetFragmentShadingRateKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkFragmentShadingRateCombinerOpKHR).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0