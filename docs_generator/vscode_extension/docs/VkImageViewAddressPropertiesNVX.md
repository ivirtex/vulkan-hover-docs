# VkImageViewAddressPropertiesNVX(3) Manual Page

## Name

VkImageViewAddressPropertiesNVX - Structure specifying the image view for handle queries



## [](#_c_specification)C Specification

The `VkImageViewAddressPropertiesNVX` structure is defined as:

```c++
// Provided by VK_NVX_image_view_handle
typedef struct VkImageViewAddressPropertiesNVX {
    VkStructureType    sType;
    void*              pNext;
    VkDeviceAddress    deviceAddress;
    VkDeviceSize       size;
} VkImageViewAddressPropertiesNVX;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `deviceAddress` is the device address of the image view.
- `size` is the size in bytes of the image view device memory.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkImageViewAddressPropertiesNVX-sType-sType)VUID-VkImageViewAddressPropertiesNVX-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_IMAGE_VIEW_ADDRESS_PROPERTIES_NVX`
- [](#VUID-VkImageViewAddressPropertiesNVX-pNext-pNext)VUID-VkImageViewAddressPropertiesNVX-pNext-pNext  
  `pNext` **must** be `NULL`

## [](#_see_also)See Also

[VK\_NVX\_image\_view\_handle](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NVX_image_view_handle.html), [VkDeviceAddress](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceAddress.html), [VkDeviceSize](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceSize.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkGetImageViewAddressNVX](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetImageViewAddressNVX.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkImageViewAddressPropertiesNVX)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0