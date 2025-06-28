# VkFilter(3) Manual Page

## Name

VkFilter - Specify filters used for texture lookups



## [](#_c_specification)C Specification

Possible values of the [VkSamplerCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerCreateInfo.html)::`magFilter` and `minFilter` parameters, specifying filters used for texture lookups, are:

```c++
// Provided by VK_VERSION_1_0
typedef enum VkFilter {
    VK_FILTER_NEAREST = 0,
    VK_FILTER_LINEAR = 1,
  // Provided by VK_EXT_filter_cubic
    VK_FILTER_CUBIC_EXT = 1000015000,
  // Provided by VK_IMG_filter_cubic
    VK_FILTER_CUBIC_IMG = VK_FILTER_CUBIC_EXT,
} VkFilter;
```

## [](#_description)Description

- `VK_FILTER_NEAREST` specifies nearest filtering.
- `VK_FILTER_LINEAR` specifies linear filtering.
- `VK_FILTER_CUBIC_EXT` specifies cubic filtering.

These filters are described in detail in [Texel Filtering](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#textures-texel-filtering).

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkBlitImageInfo2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBlitImageInfo2.html), [VkSamplerCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerCreateInfo.html), [VkSamplerYcbcrConversionCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerYcbcrConversionCreateInfo.html), [vkCmdBlitImage](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdBlitImage.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkFilter)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0