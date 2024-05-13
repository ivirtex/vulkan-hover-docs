# VK_MAKE_VERSION(3) Manual Page

## Name

VK_MAKE_VERSION - Construct an API version number



## <a href="#_c_specification" class="anchor"></a>C Specification

`VK_MAKE_VERSION` constructs an API version number.

``` c
// Provided by VK_VERSION_1_0
// DEPRECATED: This define is deprecated. VK_MAKE_API_VERSION should be used instead.
#define VK_MAKE_VERSION(major, minor, patch) \
    ((((uint32_t)(major)) << 22U) | (((uint32_t)(minor)) << 12U) | ((uint32_t)(patch)))
```

## <a href="#_description" class="anchor"></a>Description

- `major` is the major version number.

- `minor` is the minor version number.

- `patch` is the patch version number.

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkApplicationInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkApplicationInfo.html),
[vkCreateInstance](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateInstance.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_MAKE_VERSION"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
