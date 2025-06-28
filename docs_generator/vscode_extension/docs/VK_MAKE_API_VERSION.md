# VK\_MAKE\_API\_VERSION(3) Manual Page

## Name

VK\_MAKE\_API\_VERSION - Construct an API version number



## [](#_c_specification)C Specification

`VK_MAKE_API_VERSION` constructs an API version number.

```c++
// Provided by VK_VERSION_1_0
#define VK_MAKE_API_VERSION(variant, major, minor, patch) \
    ((((uint32_t)(variant)) << 29U) | (((uint32_t)(major)) << 22U) | (((uint32_t)(minor)) << 12U) | ((uint32_t)(patch)))
```

## [](#_description)Description

- `variant` is the variant number.
- `major` is the major version number.
- `minor` is the minor version number.
- `patch` is the patch version number.

## [](#_see_also)See Also

[VK\_API\_VERSION](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_API_VERSION.html), [VK\_API\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_API_VERSION_1_0.html), [VK\_API\_VERSION\_1\_1](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_API_VERSION_1_1.html), [VK\_API\_VERSION\_1\_2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_API_VERSION_1_2.html), [VK\_API\_VERSION\_1\_3](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_API_VERSION_1_3.html), [VK\_API\_VERSION\_1\_4](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_API_VERSION_1_4.html), [VK\_HEADER\_VERSION\_COMPLETE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_HEADER_VERSION_COMPLETE.html), [VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkApplicationInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkApplicationInfo.html), [vkCreateInstance](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateInstance.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_MAKE_API_VERSION)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0