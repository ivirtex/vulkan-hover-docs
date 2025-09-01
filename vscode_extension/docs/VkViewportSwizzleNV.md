# VkViewportSwizzleNV(3) Manual Page

## Name

VkViewportSwizzleNV - Structure specifying a viewport swizzle



## [](#_c_specification)C Specification

The `VkViewportSwizzleNV` structure is defined as:

```c++
// Provided by VK_NV_viewport_swizzle
typedef struct VkViewportSwizzleNV {
    VkViewportCoordinateSwizzleNV    x;
    VkViewportCoordinateSwizzleNV    y;
    VkViewportCoordinateSwizzleNV    z;
    VkViewportCoordinateSwizzleNV    w;
} VkViewportSwizzleNV;
```

## [](#_members)Members

- `x` is a [VkViewportCoordinateSwizzleNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkViewportCoordinateSwizzleNV.html) value specifying the swizzle operation to apply to the x component of the primitive
- `y` is a [VkViewportCoordinateSwizzleNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkViewportCoordinateSwizzleNV.html) value specifying the swizzle operation to apply to the y component of the primitive
- `z` is a [VkViewportCoordinateSwizzleNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkViewportCoordinateSwizzleNV.html) value specifying the swizzle operation to apply to the z component of the primitive
- `w` is a [VkViewportCoordinateSwizzleNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkViewportCoordinateSwizzleNV.html) value specifying the swizzle operation to apply to the w component of the primitive

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkViewportSwizzleNV-x-parameter)VUID-VkViewportSwizzleNV-x-parameter  
  `x` **must** be a valid [VkViewportCoordinateSwizzleNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkViewportCoordinateSwizzleNV.html) value
- [](#VUID-VkViewportSwizzleNV-y-parameter)VUID-VkViewportSwizzleNV-y-parameter  
  `y` **must** be a valid [VkViewportCoordinateSwizzleNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkViewportCoordinateSwizzleNV.html) value
- [](#VUID-VkViewportSwizzleNV-z-parameter)VUID-VkViewportSwizzleNV-z-parameter  
  `z` **must** be a valid [VkViewportCoordinateSwizzleNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkViewportCoordinateSwizzleNV.html) value
- [](#VUID-VkViewportSwizzleNV-w-parameter)VUID-VkViewportSwizzleNV-w-parameter  
  `w` **must** be a valid [VkViewportCoordinateSwizzleNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkViewportCoordinateSwizzleNV.html) value

## [](#_see_also)See Also

[VK\_NV\_viewport\_swizzle](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_viewport_swizzle.html), [VkPipelineViewportSwizzleStateCreateInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineViewportSwizzleStateCreateInfoNV.html), [VkViewportCoordinateSwizzleNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkViewportCoordinateSwizzleNV.html), [vkCmdSetViewportSwizzleNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetViewportSwizzleNV.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkViewportSwizzleNV).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0