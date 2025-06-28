# VK\_HEADER\_VERSION\_COMPLETE(3) Manual Page

## Name

VK\_HEADER\_VERSION\_COMPLETE - Vulkan header file complete version number



## [](#_c_specification)C Specification

`VK_HEADER_VERSION_COMPLETE` is the complete version number of the `vulkan_core.h` header, comprising the major, minor, and patch versions. The major/minor values are kept synchronized with the complete version of the released Specification. This value is intended for use by automated tools to identify exactly which version of the header was used during their generation.

Applications should not use this value as their [VkApplicationInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkApplicationInfo.html)::`apiVersion`. Instead applications should explicitly select a specific fixed major/minor API version using, for example, one of the `VK_API_VERSION_`\*\_* values.

```c++
// Provided by VK_VERSION_1_0
// Complete version of this file
#define VK_HEADER_VERSION_COMPLETE VK_MAKE_API_VERSION(0, 1, 4, VK_HEADER_VERSION)
```

## [](#_see_also)See Also

[VK\_HEADER\_VERSION](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_HEADER_VERSION.html), [VK\_MAKE\_API\_VERSION](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_MAKE_API_VERSION.html), [VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_HEADER_VERSION_COMPLETE)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0