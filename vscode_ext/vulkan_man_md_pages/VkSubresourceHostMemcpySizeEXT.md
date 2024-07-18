# VkSubresourceHostMemcpySizeEXT(3) Manual Page

## Name

VkSubresourceHostMemcpySizeEXT - Memory size needed to copy to or from
an image on the host with VK_HOST_IMAGE_COPY_MEMCPY_EXT



## <a href="#_c_specification" class="anchor"></a>C Specification

To query the memory size needed to copy to or from an image using
[vkCopyMemoryToImageEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCopyMemoryToImageEXT.html) or
[vkCopyImageToMemoryEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCopyImageToMemoryEXT.html) when the
`VK_HOST_IMAGE_COPY_MEMCPY_EXT` flag is specified, add a
[VkSubresourceHostMemcpySizeEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubresourceHostMemcpySizeEXT.html)
structure to the `pNext` chain of the
[VkSubresourceLayout2EXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubresourceLayout2EXT.html) structure in a
call to
[vkGetImageSubresourceLayout2EXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetImageSubresourceLayout2EXT.html).

The `VkSubresourceHostMemcpySizeEXT` structure is defined as:

``` c
// Provided by VK_EXT_host_image_copy
typedef struct VkSubresourceHostMemcpySizeEXT {
    VkStructureType    sType;
    void*              pNext;
    VkDeviceSize       size;
} VkSubresourceHostMemcpySizeEXT;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `size` is the size in bytes of the image subresource.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a href="#VUID-VkSubresourceHostMemcpySizeEXT-sType-sType"
  id="VUID-VkSubresourceHostMemcpySizeEXT-sType-sType"></a>
  VUID-VkSubresourceHostMemcpySizeEXT-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_SUBRESOURCE_HOST_MEMCPY_SIZE_EXT`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_host_image_copy](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_host_image_copy.html),
[VkDeviceSize](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceSize.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkSubresourceHostMemcpySizeEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
