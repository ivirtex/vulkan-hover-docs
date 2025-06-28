# VkSubresourceHostMemcpySize(3) Manual Page

## Name

VkSubresourceHostMemcpySize - Memory size needed to copy to or from an image on the host with VK\_HOST\_IMAGE\_COPY\_MEMCPY\_BIT



## [](#_c_specification)C Specification

To query the memory size needed to copy to or from an image using [vkCopyMemoryToImage](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCopyMemoryToImage.html) or [vkCopyImageToMemory](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCopyImageToMemory.html) when the `VK_HOST_IMAGE_COPY_MEMCPY_BIT` flag is specified, add a [VkSubresourceHostMemcpySize](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubresourceHostMemcpySize.html) structure to the `pNext` chain of the [VkSubresourceLayout2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubresourceLayout2.html) structure in a call to [vkGetImageSubresourceLayout2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetImageSubresourceLayout2.html).

The `VkSubresourceHostMemcpySize` structure is defined as:

```c++
// Provided by VK_VERSION_1_4
typedef struct VkSubresourceHostMemcpySize {
    VkStructureType    sType;
    void*              pNext;
    VkDeviceSize       size;
} VkSubresourceHostMemcpySize;
```

or the equivalent

```c++
// Provided by VK_EXT_host_image_copy
typedef VkSubresourceHostMemcpySize VkSubresourceHostMemcpySizeEXT;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `size` is the size in bytes of the image subresource.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkSubresourceHostMemcpySize-sType-sType)VUID-VkSubresourceHostMemcpySize-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_SUBRESOURCE_HOST_MEMCPY_SIZE`

## [](#_see_also)See Also

[VK\_EXT\_host\_image\_copy](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_host_image_copy.html), [VK\_VERSION\_1\_4](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_4.html), [VkDeviceSize](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceSize.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkSubresourceHostMemcpySize)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0