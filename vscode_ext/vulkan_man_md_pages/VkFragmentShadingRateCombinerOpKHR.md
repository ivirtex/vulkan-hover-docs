# VkFragmentShadingRateCombinerOpKHR(3) Manual Page

## Name

VkFragmentShadingRateCombinerOpKHR - Control how fragment shading rates
are combined



## <a href="#_c_specification" class="anchor"></a>C Specification

The equation used for each combiner operation is defined by
`VkFragmentShadingRateCombinerOpKHR`:

``` c
// Provided by VK_KHR_fragment_shading_rate
typedef enum VkFragmentShadingRateCombinerOpKHR {
    VK_FRAGMENT_SHADING_RATE_COMBINER_OP_KEEP_KHR = 0,
    VK_FRAGMENT_SHADING_RATE_COMBINER_OP_REPLACE_KHR = 1,
    VK_FRAGMENT_SHADING_RATE_COMBINER_OP_MIN_KHR = 2,
    VK_FRAGMENT_SHADING_RATE_COMBINER_OP_MAX_KHR = 3,
    VK_FRAGMENT_SHADING_RATE_COMBINER_OP_MUL_KHR = 4,
} VkFragmentShadingRateCombinerOpKHR;
```

## <a href="#_description" class="anchor"></a>Description

- `VK_FRAGMENT_SHADING_RATE_COMBINER_OP_KEEP_KHR` specifies a combiner
  operation of combine(A<sub>xy</sub>,B<sub>xy</sub>) = A<sub>xy</sub>.

- `VK_FRAGMENT_SHADING_RATE_COMBINER_OP_REPLACE_KHR` specifies a
  combiner operation of combine(A<sub>xy</sub>,B<sub>xy</sub>) =
  B<sub>xy</sub>.

- `VK_FRAGMENT_SHADING_RATE_COMBINER_OP_MIN_KHR` specifies a combiner
  operation of combine(A<sub>xy</sub>,B<sub>xy</sub>) =
  min(A<sub>xy</sub>,B<sub>xy</sub>).

- `VK_FRAGMENT_SHADING_RATE_COMBINER_OP_MAX_KHR` specifies a combiner
  operation of combine(A<sub>xy</sub>,B<sub>xy</sub>) =
  max(A<sub>xy</sub>,B<sub>xy</sub>).

- `VK_FRAGMENT_SHADING_RATE_COMBINER_OP_MUL_KHR` specifies a combiner
  operation of combine(A<sub>xy</sub>,B<sub>xy</sub>) =
  A<sub>xy</sub>\*B<sub>xy</sub>.

where combine(A<sub>xy</sub>,B<sub>xy</sub>) is the combine operation,
and A<sub>xy</sub> and B<sub>xy</sub> are the inputs to the operation.

If <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-fragmentShadingRateStrictMultiplyCombiner"
target="_blank"
rel="noopener"><code>fragmentShadingRateStrictMultiplyCombiner</code></a>
is `VK_FALSE`, using `VK_FRAGMENT_SHADING_RATE_COMBINER_OP_MUL_KHR` with
values of 1 for both A and B in the same dimension results in the value
2 being produced for that dimension. See the definition of <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-fragmentShadingRateStrictMultiplyCombiner"
target="_blank"
rel="noopener"><code>fragmentShadingRateStrictMultiplyCombiner</code></a>
for more information.

These operations are performed in a component-wise fashion.

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_fragment_shading_rate](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_fragment_shading_rate.html),
[VkPipelineFragmentShadingRateEnumStateCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineFragmentShadingRateEnumStateCreateInfoNV.html),
[VkPipelineFragmentShadingRateStateCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineFragmentShadingRateStateCreateInfoKHR.html),
[vkCmdSetFragmentShadingRateEnumNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetFragmentShadingRateEnumNV.html),
[vkCmdSetFragmentShadingRateKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetFragmentShadingRateKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkFragmentShadingRateCombinerOpKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
