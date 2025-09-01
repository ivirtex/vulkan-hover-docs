# VkOffset3D(3) Manual Page

## Name

VkOffset3D - Structure specifying a three-dimensional offset



## [](#_c_specification)C Specification

A three-dimensional offset is defined by the structure:

```c++
// Provided by VK_VERSION_1_0
typedef struct VkOffset3D {
    int32_t    x;
    int32_t    y;
    int32_t    z;
} VkOffset3D;
```

## [](#_members)Members

- `x` is the x offset.
- `y` is the y offset.
- `z` is the z offset.

## [](#_description)Description

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkBufferImageCopy](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferImageCopy.html), [VkBufferImageCopy2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferImageCopy2.html), [VkCopyMemoryToImageIndirectCommandNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyMemoryToImageIndirectCommandNV.html), [VkImageBlit](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageBlit.html), [VkImageBlit2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageBlit2.html), [VkImageCopy](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCopy.html), [VkImageCopy2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCopy2.html), [VkImageResolve](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageResolve.html), [VkImageResolve2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageResolve2.html), [VkImageToMemoryCopy](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageToMemoryCopy.html), [VkMemoryToImageCopy](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryToImageCopy.html), [VkSparseImageMemoryBind](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSparseImageMemoryBind.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkOffset3D).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0