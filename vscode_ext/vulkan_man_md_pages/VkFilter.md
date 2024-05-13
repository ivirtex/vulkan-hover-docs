# VkFilter(3) Manual Page

## Name

VkFilter - Specify filters used for texture lookups



## <a href="#_c_specification" class="anchor"></a>C Specification

Possible values of the
[VkSamplerCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerCreateInfo.html)::`magFilter` and
`minFilter` parameters, specifying filters used for texture lookups,
are:

``` c
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

## <a href="#_description" class="anchor"></a>Description

- `VK_FILTER_NEAREST` specifies nearest filtering.

- `VK_FILTER_LINEAR` specifies linear filtering.

- `VK_FILTER_CUBIC_EXT` specifies cubic filtering.

These filters are described in detail in <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#textures-texel-filtering"
target="_blank" rel="noopener">Texel Filtering</a>.

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkBlitImageInfo2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBlitImageInfo2.html),
[VkSamplerCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerCreateInfo.html),
[VkSamplerYcbcrConversionCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerYcbcrConversionCreateInfo.html),
[vkCmdBlitImage](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBlitImage.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkFilter"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
