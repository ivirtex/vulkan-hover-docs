# VkSubresourceLayout2KHR(3) Manual Page

## Name

VkSubresourceLayout2KHR - Structure specifying subresource layout



## <a href="#_c_specification" class="anchor"></a>C Specification

Information about the layout of the image subresource is returned in a
`VkSubresourceLayout2KHR` structure:

``` c
// Provided by VK_KHR_maintenance5
typedef struct VkSubresourceLayout2KHR {
    VkStructureType        sType;
    void*                  pNext;
    VkSubresourceLayout    subresourceLayout;
} VkSubresourceLayout2KHR;
```

or the equivalent

``` c
// Provided by VK_EXT_host_image_copy, VK_EXT_image_compression_control
typedef VkSubresourceLayout2KHR VkSubresourceLayout2EXT;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `subresourceLayout` is a
  [VkSubresourceLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubresourceLayout.html) structure.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a href="#VUID-VkSubresourceLayout2KHR-sType-sType"
  id="VUID-VkSubresourceLayout2KHR-sType-sType"></a>
  VUID-VkSubresourceLayout2KHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_SUBRESOURCE_LAYOUT_2_KHR`

- <a href="#VUID-VkSubresourceLayout2KHR-pNext-pNext"
  id="VUID-VkSubresourceLayout2KHR-pNext-pNext"></a>
  VUID-VkSubresourceLayout2KHR-pNext-pNext  
  Each `pNext` member of any structure (including this one) in the
  `pNext` chain **must** be either `NULL` or a pointer to a valid
  instance of
  [VkImageCompressionPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCompressionPropertiesEXT.html)
  or
  [VkSubresourceHostMemcpySizeEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubresourceHostMemcpySizeEXT.html)

- <a href="#VUID-VkSubresourceLayout2KHR-sType-unique"
  id="VUID-VkSubresourceLayout2KHR-sType-unique"></a>
  VUID-VkSubresourceLayout2KHR-sType-unique  
  The `sType` value of each struct in the `pNext` chain **must** be
  unique

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_host_image_copy](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_host_image_copy.html),
[VK_EXT_image_compression_control](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_image_compression_control.html),
[VK_KHR_maintenance5](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_maintenance5.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[VkSubresourceLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubresourceLayout.html),
[vkGetDeviceImageSubresourceLayoutKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetDeviceImageSubresourceLayoutKHR.html),
[vkGetImageSubresourceLayout2EXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetImageSubresourceLayout2EXT.html),
[vkGetImageSubresourceLayout2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetImageSubresourceLayout2KHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkSubresourceLayout2KHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
