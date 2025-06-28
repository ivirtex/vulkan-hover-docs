# VK\_VERSION\_PATCH(3) Manual Page

## Name

VK\_VERSION\_PATCH - Extract API patch version number



## [](#_c_specification)C Specification

`VK_VERSION_PATCH` extracts the API patch version number from a packed version number:

Warning

This functionality is deprecated by [Vulkan Version 1.0](#versions-1.0). See [Deprecated Functionality](#deprecation-version-macros) for more information.

```c++
// Provided by VK_VERSION_1_0
#define VK_VERSION_PATCH(version) ((uint32_t)(version) & 0xFFFU)
```

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_VERSION_PATCH)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0