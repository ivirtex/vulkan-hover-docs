# VkExtent3D(3) Manual Page

## Name

VkExtent3D - Structure specifying a three-dimensional extent



## [](#_c_specification)C Specification

A three-dimensional extent is defined by the structure:

```c++
// Provided by VK_VERSION_1_0
typedef struct VkExtent3D {
    uint32_t    width;
    uint32_t    height;
    uint32_t    depth;
} VkExtent3D;
```

## [](#_members)Members

- `width` is the width of the extent.
- `height` is the height of the extent.
- `depth` is the depth of the extent.

## [](#_description)Description

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkBufferImageCopy](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferImageCopy.html), [VkBufferImageCopy2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferImageCopy2.html), [VkCopyMemoryToImageIndirectCommandKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyMemoryToImageIndirectCommandKHR.html), [VkImageCopy](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCopy.html), [VkImageCopy2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCopy2.html), [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateInfo.html), [VkImageFormatProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageFormatProperties.html), [VkImageResolve](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageResolve.html), [VkImageResolve2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageResolve2.html), [VkImageToMemoryCopy](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageToMemoryCopy.html), [VkMemoryToImageCopy](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryToImageCopy.html), [VkQueueFamilyProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueueFamilyProperties.html), [VkSparseImageFormatProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSparseImageFormatProperties.html), [VkSparseImageMemoryBind](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSparseImageMemoryBind.html), [VkTilePropertiesQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTilePropertiesQCOM.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkExtent3D).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0