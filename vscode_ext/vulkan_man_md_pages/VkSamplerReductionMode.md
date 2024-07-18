# VkSamplerReductionMode(3) Manual Page

## Name

VkSamplerReductionMode - Specify reduction mode for texture filtering



## <a href="#_c_specification" class="anchor"></a>C Specification

Reduction modes are specified by
[VkSamplerReductionMode](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerReductionMode.html), which takes
values:

``` c
// Provided by VK_VERSION_1_2
typedef enum VkSamplerReductionMode {
    VK_SAMPLER_REDUCTION_MODE_WEIGHTED_AVERAGE = 0,
    VK_SAMPLER_REDUCTION_MODE_MIN = 1,
    VK_SAMPLER_REDUCTION_MODE_MAX = 2,
  // Provided by VK_QCOM_filter_cubic_clamp
    VK_SAMPLER_REDUCTION_MODE_WEIGHTED_AVERAGE_RANGECLAMP_QCOM = 1000521000,
  // Provided by VK_EXT_sampler_filter_minmax
    VK_SAMPLER_REDUCTION_MODE_WEIGHTED_AVERAGE_EXT = VK_SAMPLER_REDUCTION_MODE_WEIGHTED_AVERAGE,
  // Provided by VK_EXT_sampler_filter_minmax
    VK_SAMPLER_REDUCTION_MODE_MIN_EXT = VK_SAMPLER_REDUCTION_MODE_MIN,
  // Provided by VK_EXT_sampler_filter_minmax
    VK_SAMPLER_REDUCTION_MODE_MAX_EXT = VK_SAMPLER_REDUCTION_MODE_MAX,
} VkSamplerReductionMode;
```

or the equivalent

``` c
// Provided by VK_EXT_sampler_filter_minmax
typedef VkSamplerReductionMode VkSamplerReductionModeEXT;
```

## <a href="#_description" class="anchor"></a>Description

- `VK_SAMPLER_REDUCTION_MODE_WEIGHTED_AVERAGE` specifies that texel
  values are combined by computing a weighted average of values in the
  footprint, using weights as specified in <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#textures-unnormalized-to-integer"
  target="_blank" rel="noopener">the image operations chapter</a>.

- `VK_SAMPLER_REDUCTION_MODE_MIN` specifies that texel values are
  combined by taking the component-wise minimum of values in the
  footprint with non-zero weights.

- `VK_SAMPLER_REDUCTION_MODE_MAX` specifies that texel values are
  combined by taking the component-wise maximum of values in the
  footprint with non-zero weights.

- `VK_SAMPLER_REDUCTION_MODE_WEIGHTED_AVERAGE_RANGECLAMP_QCOM` specifies
  values are combined as described by
  `VK_SAMPLER_REDUCTION_MODE_WEIGHTED_AVERAGE`, followed by a <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#textures-texel-range-clamp"
  target="_blank" rel="noopener">texel range clamp</a>.

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_sampler_filter_minmax](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_sampler_filter_minmax.html),
[VK_VERSION_1_2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_2.html),
[VkSamplerReductionModeCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerReductionModeCreateInfo.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkSamplerReductionMode"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
