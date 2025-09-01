# VkCubicFilterWeightsQCOM(3) Manual Page

## Name

VkCubicFilterWeightsQCOM - Specify cubic weights for texture filtering



## [](#_c_specification)C Specification

Possible values of the [VkSamplerCubicWeightsCreateInfoQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerCubicWeightsCreateInfoQCOM.html)::`cubicWeights`, specifying cubic weights used in [Texel Cubic Filtering](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#textures-texel-cubic-filtering) are:

```c++
// Provided by VK_QCOM_filter_cubic_weights
typedef enum VkCubicFilterWeightsQCOM {
    VK_CUBIC_FILTER_WEIGHTS_CATMULL_ROM_QCOM = 0,
    VK_CUBIC_FILTER_WEIGHTS_ZERO_TANGENT_CARDINAL_QCOM = 1,
    VK_CUBIC_FILTER_WEIGHTS_B_SPLINE_QCOM = 2,
    VK_CUBIC_FILTER_WEIGHTS_MITCHELL_NETRAVALI_QCOM = 3,
} VkCubicFilterWeightsQCOM;
```

## [](#_description)Description

- `VK_CUBIC_FILTER_WEIGHTS_CATMULL_ROM_QCOM` specifies Catmull-Rom weights.
- `VK_CUBIC_FILTER_WEIGHTS_ZERO_TANGENT_CARDINAL_QCOM` specifies Zero Tangent Cardinal weights.
- `VK_CUBIC_FILTER_WEIGHTS_B_SPLINE_QCOM` specifies B-Spline weights.
- `VK_CUBIC_FILTER_WEIGHTS_MITCHELL_NETRAVALI_QCOM` specifies Mitchell-Netravali weights.

## [](#_see_also)See Also

[VK\_QCOM\_filter\_cubic\_weights](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_QCOM_filter_cubic_weights.html), [VkBlitImageCubicWeightsInfoQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBlitImageCubicWeightsInfoQCOM.html), [VkSamplerCubicWeightsCreateInfoQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerCubicWeightsCreateInfoQCOM.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkCubicFilterWeightsQCOM).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0