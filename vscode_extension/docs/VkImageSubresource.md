# VkImageSubresource(3) Manual Page

## Name

VkImageSubresource - Structure specifying an image subresource



## [](#_c_specification)C Specification

The `VkImageSubresource` structure is defined as:

```c++
// Provided by VK_VERSION_1_0
typedef struct VkImageSubresource {
    VkImageAspectFlags    aspectMask;
    uint32_t              mipLevel;
    uint32_t              arrayLayer;
} VkImageSubresource;
```

## [](#_members)Members

- `aspectMask` is a [VkImageAspectFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageAspectFlags.html) value selecting the image *aspect*.
- `mipLevel` selects the mipmap level.
- `arrayLayer` selects the array layer.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkImageSubresource-aspectMask-parameter)VUID-VkImageSubresource-aspectMask-parameter  
  `aspectMask` **must** be a valid combination of [VkImageAspectFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageAspectFlagBits.html) values
- [](#VUID-VkImageSubresource-aspectMask-requiredbitmask)VUID-VkImageSubresource-aspectMask-requiredbitmask  
  `aspectMask` **must** not be `0`

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkImageAspectFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageAspectFlags.html), [VkImageSubresource2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageSubresource2.html), [VkSparseImageMemoryBind](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSparseImageMemoryBind.html), [vkGetImageSubresourceLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetImageSubresourceLayout.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkImageSubresource).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0