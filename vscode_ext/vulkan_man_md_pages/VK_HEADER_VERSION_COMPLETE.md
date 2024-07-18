# VK_HEADER_VERSION_COMPLETE(3) Manual Page

## Name

VK_HEADER_VERSION_COMPLETE - Vulkan header file complete version number



## <a href="#_c_specification" class="anchor"></a>C Specification

`VK_HEADER_VERSION_COMPLETE` is the complete version number of the
`vulkan_core.h` header, comprising the major, minor, and patch versions.
The major/minor values are kept synchronized with the complete version
of the released Specification. This value is intended for use by
automated tools to identify exactly which version of the header was used
during their generation.

Applications should not use this value as their
[VkApplicationInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkApplicationInfo.html)::`apiVersion`. Instead
applications should explicitly select a specific fixed major/minor API
version using, for example, one of the `VK_API_VERSION_`\*\_\* values.

``` c
// Provided by VK_VERSION_1_0
// Complete version of this file
#define VK_HEADER_VERSION_COMPLETE VK_MAKE_API_VERSION(0, 1, 3, VK_HEADER_VERSION)
```

## <a href="#_see_also" class="anchor"></a>See Also

[VK_HEADER_VERSION](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_HEADER_VERSION.html),
[VK_MAKE_API_VERSION](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_MAKE_API_VERSION.html),
[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_HEADER_VERSION_COMPLETE"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
