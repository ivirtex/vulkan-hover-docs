# VK\_MAKE\_VERSION(3) Manual Page

## Name

VK\_MAKE\_VERSION - Construct an API version number



## [](#_c_specification)C Specification

`VK_MAKE_VERSION` constructs an API version number.

Warning

This functionality is deprecated by [Vulkan Version 1.0](#versions-1.0). See [Deprecated Functionality](#deprecation-version-macros) for more information.

```c++
// Provided by VK_VERSION_1_0
#define VK_MAKE_VERSION(major, minor, patch) \
    ((((uint32_t)(major)) << 22U) | (((uint32_t)(minor)) << 12U) | ((uint32_t)(patch)))
```

## [](#_description)Description

- `major` is the major version number.
- `minor` is the minor version number.
- `patch` is the patch version number.

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkApplicationInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkApplicationInfo.html), [vkCreateInstance](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateInstance.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_MAKE_VERSION)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0