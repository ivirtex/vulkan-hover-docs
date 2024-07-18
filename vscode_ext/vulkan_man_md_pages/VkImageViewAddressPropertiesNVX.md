# VkImageViewAddressPropertiesNVX(3) Manual Page

## Name

VkImageViewAddressPropertiesNVX - Structure specifying the image view
for handle queries



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkImageViewAddressPropertiesNVX` structure is defined as:

``` c
// Provided by VK_NVX_image_view_handle
typedef struct VkImageViewAddressPropertiesNVX {
    VkStructureType    sType;
    void*              pNext;
    VkDeviceAddress    deviceAddress;
    VkDeviceSize       size;
} VkImageViewAddressPropertiesNVX;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `deviceAddress` is the device address of the image view.

- `size` is the size in bytes of the image view device memory.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a href="#VUID-VkImageViewAddressPropertiesNVX-sType-sType"
  id="VUID-VkImageViewAddressPropertiesNVX-sType-sType"></a>
  VUID-VkImageViewAddressPropertiesNVX-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_IMAGE_VIEW_ADDRESS_PROPERTIES_NVX`

- <a href="#VUID-VkImageViewAddressPropertiesNVX-pNext-pNext"
  id="VUID-VkImageViewAddressPropertiesNVX-pNext-pNext"></a>
  VUID-VkImageViewAddressPropertiesNVX-pNext-pNext  
  `pNext` **must** be `NULL`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_NVX_image_view_handle](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NVX_image_view_handle.html),
[VkDeviceAddress](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceAddress.html),
[VkDeviceSize](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceSize.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkGetImageViewAddressNVX](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetImageViewAddressNVX.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkImageViewAddressPropertiesNVX"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
