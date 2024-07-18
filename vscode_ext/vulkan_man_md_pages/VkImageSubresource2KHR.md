# VkImageSubresource2KHR(3) Manual Page

## Name

VkImageSubresource2KHR - Structure specifying an image subresource



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkImageSubresource2KHR` structure is defined as:

``` c
// Provided by VK_KHR_maintenance5
typedef struct VkImageSubresource2KHR {
    VkStructureType       sType;
    void*                 pNext;
    VkImageSubresource    imageSubresource;
} VkImageSubresource2KHR;
```

or the equivalent

``` c
// Provided by VK_EXT_host_image_copy, VK_EXT_image_compression_control
typedef VkImageSubresource2KHR VkImageSubresource2EXT;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `imageSubresource` is a [VkImageSubresource](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageSubresource.html)
  structure.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a href="#VUID-VkImageSubresource2KHR-sType-sType"
  id="VUID-VkImageSubresource2KHR-sType-sType"></a>
  VUID-VkImageSubresource2KHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_IMAGE_SUBRESOURCE_2_KHR`

- <a href="#VUID-VkImageSubresource2KHR-pNext-pNext"
  id="VUID-VkImageSubresource2KHR-pNext-pNext"></a>
  VUID-VkImageSubresource2KHR-pNext-pNext  
  `pNext` **must** be `NULL`

- <a href="#VUID-VkImageSubresource2KHR-imageSubresource-parameter"
  id="VUID-VkImageSubresource2KHR-imageSubresource-parameter"></a>
  VUID-VkImageSubresource2KHR-imageSubresource-parameter  
  `imageSubresource` **must** be a valid
  [VkImageSubresource](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageSubresource.html) structure

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_host_image_copy](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_host_image_copy.html),
[VK_EXT_image_compression_control](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_image_compression_control.html),
[VK_KHR_maintenance5](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_maintenance5.html),
[VkDeviceImageSubresourceInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceImageSubresourceInfoKHR.html),
[VkImageSubresource](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageSubresource.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkGetImageSubresourceLayout2EXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetImageSubresourceLayout2EXT.html),
[vkGetImageSubresourceLayout2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetImageSubresourceLayout2KHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkImageSubresource2KHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
