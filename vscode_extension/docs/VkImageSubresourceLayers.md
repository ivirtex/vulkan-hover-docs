# VkImageSubresourceLayers(3) Manual Page

## Name

VkImageSubresourceLayers - Structure specifying an image subresource layers



## [](#_c_specification)C Specification

The `VkImageSubresourceLayers` structure is defined as:

```c++
// Provided by VK_VERSION_1_0
typedef struct VkImageSubresourceLayers {
    VkImageAspectFlags    aspectMask;
    uint32_t              mipLevel;
    uint32_t              baseArrayLayer;
    uint32_t              layerCount;
} VkImageSubresourceLayers;
```

## [](#_members)Members

- `aspectMask` is a combination of [VkImageAspectFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageAspectFlagBits.html), selecting the color, depth and/or stencil aspects to be copied.
- `mipLevel` is the mipmap level to copy
- `baseArrayLayer` and `layerCount` are the starting layer and number of layers to copy.

## [](#_description)Description

Valid Usage

- [](#VUID-VkImageSubresourceLayers-aspectMask-00167)VUID-VkImageSubresourceLayers-aspectMask-00167  
  If `aspectMask` contains `VK_IMAGE_ASPECT_COLOR_BIT`, it **must** not contain either of `VK_IMAGE_ASPECT_DEPTH_BIT` or `VK_IMAGE_ASPECT_STENCIL_BIT`
- [](#VUID-VkImageSubresourceLayers-aspectMask-00168)VUID-VkImageSubresourceLayers-aspectMask-00168  
  `aspectMask` **must** not contain `VK_IMAGE_ASPECT_METADATA_BIT`
- [](#VUID-VkImageSubresourceLayers-aspectMask-02247)VUID-VkImageSubresourceLayers-aspectMask-02247  
  `aspectMask` **must** not include `VK_IMAGE_ASPECT_MEMORY_PLANE_i_BIT_EXT` for any index *i*
- [](#VUID-VkImageSubresourceLayers-layerCount-09243)VUID-VkImageSubresourceLayers-layerCount-09243  
  If the [`maintenance5`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-maintenance5) feature is not enabled, `layerCount` **must** not be `VK_REMAINING_ARRAY_LAYERS`
- [](#VUID-VkImageSubresourceLayers-layerCount-01700)VUID-VkImageSubresourceLayers-layerCount-01700  
  If `layerCount` is not `VK_REMAINING_ARRAY_LAYERS`, it **must** be greater than 0

Valid Usage (Implicit)

- [](#VUID-VkImageSubresourceLayers-aspectMask-parameter)VUID-VkImageSubresourceLayers-aspectMask-parameter  
  `aspectMask` **must** be a valid combination of [VkImageAspectFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageAspectFlagBits.html) values
- [](#VUID-VkImageSubresourceLayers-aspectMask-requiredbitmask)VUID-VkImageSubresourceLayers-aspectMask-requiredbitmask  
  `aspectMask` **must** not be `0`

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkBufferImageCopy](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferImageCopy.html), [VkBufferImageCopy2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferImageCopy2.html), [VkCopyMemoryToImageIndirectCommandNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyMemoryToImageIndirectCommandNV.html), [VkImageAspectFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageAspectFlags.html), [VkImageBlit](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageBlit.html), [VkImageBlit2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageBlit2.html), [VkImageCopy](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCopy.html), [VkImageCopy2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCopy2.html), [VkImageResolve](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageResolve.html), [VkImageResolve2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageResolve2.html), [VkImageToMemoryCopy](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageToMemoryCopy.html), [VkMemoryToImageCopy](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryToImageCopy.html), [vkCmdCopyMemoryToImageIndirectNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdCopyMemoryToImageIndirectNV.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkImageSubresourceLayers).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0