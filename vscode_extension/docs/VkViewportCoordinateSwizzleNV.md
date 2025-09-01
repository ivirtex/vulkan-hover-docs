# VkViewportCoordinateSwizzleNV(3) Manual Page

## Name

VkViewportCoordinateSwizzleNV - Specify how a viewport coordinate is swizzled



## [](#_c_specification)C Specification

Possible values of the [VkViewportSwizzleNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkViewportSwizzleNV.html)::`x`, `y`, `z`, and `w` members, specifying swizzling of the corresponding components of primitives, are:

```c++
// Provided by VK_NV_viewport_swizzle
typedef enum VkViewportCoordinateSwizzleNV {
    VK_VIEWPORT_COORDINATE_SWIZZLE_POSITIVE_X_NV = 0,
    VK_VIEWPORT_COORDINATE_SWIZZLE_NEGATIVE_X_NV = 1,
    VK_VIEWPORT_COORDINATE_SWIZZLE_POSITIVE_Y_NV = 2,
    VK_VIEWPORT_COORDINATE_SWIZZLE_NEGATIVE_Y_NV = 3,
    VK_VIEWPORT_COORDINATE_SWIZZLE_POSITIVE_Z_NV = 4,
    VK_VIEWPORT_COORDINATE_SWIZZLE_NEGATIVE_Z_NV = 5,
    VK_VIEWPORT_COORDINATE_SWIZZLE_POSITIVE_W_NV = 6,
    VK_VIEWPORT_COORDINATE_SWIZZLE_NEGATIVE_W_NV = 7,
} VkViewportCoordinateSwizzleNV;
```

## [](#_description)Description

These values are described in detail in [Viewport Swizzle](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vertexpostproc-viewport-swizzle).

## [](#_see_also)See Also

[VK\_NV\_viewport\_swizzle](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_viewport_swizzle.html), [VkViewportSwizzleNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkViewportSwizzleNV.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkViewportCoordinateSwizzleNV).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0