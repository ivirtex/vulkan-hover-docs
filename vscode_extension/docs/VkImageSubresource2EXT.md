# VkImageSubresource2(3) Manual Page

## Name

VkImageSubresource2 - Structure specifying an image subresource



## [](#_c_specification)C Specification

The `VkImageSubresource2` structure is defined as:

```c++
// Provided by VK_VERSION_1_4
typedef struct VkImageSubresource2 {
    VkStructureType       sType;
    void*                 pNext;
    VkImageSubresource    imageSubresource;
} VkImageSubresource2;
```

or the equivalent

```c++
// Provided by VK_KHR_maintenance5
typedef VkImageSubresource2 VkImageSubresource2KHR;
```

or the equivalent

```c++
// Provided by VK_EXT_host_image_copy, VK_EXT_image_compression_control
typedef VkImageSubresource2 VkImageSubresource2EXT;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `imageSubresource` is a [VkImageSubresource](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageSubresource.html) structure.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkImageSubresource2-sType-sType)VUID-VkImageSubresource2-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_IMAGE_SUBRESOURCE_2`
- [](#VUID-VkImageSubresource2-pNext-pNext)VUID-VkImageSubresource2-pNext-pNext  
  `pNext` **must** be `NULL`
- [](#VUID-VkImageSubresource2-imageSubresource-parameter)VUID-VkImageSubresource2-imageSubresource-parameter  
  `imageSubresource` **must** be a valid [VkImageSubresource](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageSubresource.html) structure

## [](#_see_also)See Also

[VK\_EXT\_host\_image\_copy](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_host_image_copy.html), [VK\_EXT\_image\_compression\_control](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_image_compression_control.html), [VK\_KHR\_maintenance5](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_maintenance5.html), [VK\_VERSION\_1\_4](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_4.html), [VkDeviceImageSubresourceInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceImageSubresourceInfo.html), [VkImageSubresource](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageSubresource.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkGetImageSubresourceLayout2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetImageSubresourceLayout2.html), [vkGetImageSubresourceLayout2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetImageSubresourceLayout2.html), [vkGetImageSubresourceLayout2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetImageSubresourceLayout2.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkImageSubresource2).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0