# VkBorderColor(3) Manual Page

## Name

VkBorderColor - Specify border color used for texture lookups



## <a href="#_c_specification" class="anchor"></a>C Specification

Possible values of
[VkSamplerCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerCreateInfo.html)::`borderColor`,
specifying the border color used for texture lookups, are:

``` c
// Provided by VK_VERSION_1_0
typedef enum VkBorderColor {
    VK_BORDER_COLOR_FLOAT_TRANSPARENT_BLACK = 0,
    VK_BORDER_COLOR_INT_TRANSPARENT_BLACK = 1,
    VK_BORDER_COLOR_FLOAT_OPAQUE_BLACK = 2,
    VK_BORDER_COLOR_INT_OPAQUE_BLACK = 3,
    VK_BORDER_COLOR_FLOAT_OPAQUE_WHITE = 4,
    VK_BORDER_COLOR_INT_OPAQUE_WHITE = 5,
  // Provided by VK_EXT_custom_border_color
    VK_BORDER_COLOR_FLOAT_CUSTOM_EXT = 1000287003,
  // Provided by VK_EXT_custom_border_color
    VK_BORDER_COLOR_INT_CUSTOM_EXT = 1000287004,
} VkBorderColor;
```

## <a href="#_description" class="anchor"></a>Description

- `VK_BORDER_COLOR_FLOAT_TRANSPARENT_BLACK` specifies a transparent,
  floating-point format, black color.

- `VK_BORDER_COLOR_INT_TRANSPARENT_BLACK` specifies a transparent,
  integer format, black color.

- `VK_BORDER_COLOR_FLOAT_OPAQUE_BLACK` specifies an opaque,
  floating-point format, black color.

- `VK_BORDER_COLOR_INT_OPAQUE_BLACK` specifies an opaque, integer
  format, black color.

- `VK_BORDER_COLOR_FLOAT_OPAQUE_WHITE` specifies an opaque,
  floating-point format, white color.

- `VK_BORDER_COLOR_INT_OPAQUE_WHITE` specifies an opaque, integer
  format, white color.

- `VK_BORDER_COLOR_FLOAT_CUSTOM_EXT` indicates that a
  [VkSamplerCustomBorderColorCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerCustomBorderColorCreateInfoEXT.html)
  structure is included in the
  [VkSamplerCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerCreateInfo.html)::`pNext` chain
  containing the color data in floating-point format.

- `VK_BORDER_COLOR_INT_CUSTOM_EXT` indicates that a
  [VkSamplerCustomBorderColorCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerCustomBorderColorCreateInfoEXT.html)
  structure is included in the
  [VkSamplerCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerCreateInfo.html)::`pNext` chain
  containing the color data in integer format.

These colors are described in detail in <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#textures-texel-replacement"
target="_blank" rel="noopener">Texel Replacement</a>.

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkSamplerCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerCreateInfo.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkBorderColor"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
