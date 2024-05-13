# VK_MAKE_API_VERSION(3) Manual Page

## Name

VK_MAKE_API_VERSION - Construct an API version number



## <a href="#_c_specification" class="anchor"></a>C Specification

`VK_MAKE_API_VERSION` constructs an API version number.

``` c
// Provided by VK_VERSION_1_0
#define VK_MAKE_API_VERSION(variant, major, minor, patch) \
    ((((uint32_t)(variant)) << 29U) | (((uint32_t)(major)) << 22U) | (((uint32_t)(minor)) << 12U) | ((uint32_t)(patch)))
```

## <a href="#_description" class="anchor"></a>Description

- `variant` is the variant number.

- `major` is the major version number.

- `minor` is the minor version number.

- `patch` is the patch version number.

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkApplicationInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkApplicationInfo.html),
[vkCreateInstance](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateInstance.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_MAKE_API_VERSION"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
