# VkCubicFilterWeightsQCOM(3) Manual Page

## Name

VkCubicFilterWeightsQCOM - Specify cubic weights for texture filtering



## <a href="#_c_specification" class="anchor"></a>C Specification

Possible values of the
[VkSamplerCubicWeightsCreateInfoQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerCubicWeightsCreateInfoQCOM.html)::`cubicWeights`,
specifying cubic weights used in <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#textures-texel-cubic-filtering"
target="_blank" rel="noopener">Texel Cubic Filtering</a> are:

``` c
// Provided by VK_QCOM_filter_cubic_weights
typedef enum VkCubicFilterWeightsQCOM {
    VK_CUBIC_FILTER_WEIGHTS_CATMULL_ROM_QCOM = 0,
    VK_CUBIC_FILTER_WEIGHTS_ZERO_TANGENT_CARDINAL_QCOM = 1,
    VK_CUBIC_FILTER_WEIGHTS_B_SPLINE_QCOM = 2,
    VK_CUBIC_FILTER_WEIGHTS_MITCHELL_NETRAVALI_QCOM = 3,
} VkCubicFilterWeightsQCOM;
```

## <a href="#_description" class="anchor"></a>Description

- `VK_CUBIC_FILTER_WEIGHTS_CATMULL_ROM_QCOM` specifies Catmull-Rom
  weights.

- `VK_CUBIC_FILTER_WEIGHTS_ZERO_TANGENT_CARDINAL_QCOM` specifies Zero
  Tangent Cardinal weights.

- `VK_CUBIC_FILTER_WEIGHTS_B_SPLINE_QCOM` specifies B-Spline weights.

- `VK_CUBIC_FILTER_WEIGHTS_MITCHELL_NETRAVALI_QCOM` specifies
  Mitchell-Netravali weights.

## <a href="#_see_also" class="anchor"></a>See Also

[VK_QCOM_filter_cubic_weights](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_QCOM_filter_cubic_weights.html),
[VkBlitImageCubicWeightsInfoQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBlitImageCubicWeightsInfoQCOM.html),
[VkSamplerCubicWeightsCreateInfoQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerCubicWeightsCreateInfoQCOM.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkCubicFilterWeightsQCOM"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
