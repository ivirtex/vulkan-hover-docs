# VkSubresourceLayout2(3) Manual Page

## Name

VkSubresourceLayout2 - Structure specifying subresource layout



## [](#_c_specification)C Specification

Information about the layout of the image subresource is returned in a `VkSubresourceLayout2` structure:

```c++
// Provided by VK_VERSION_1_4
typedef struct VkSubresourceLayout2 {
    VkStructureType        sType;
    void*                  pNext;
    VkSubresourceLayout    subresourceLayout;
} VkSubresourceLayout2;
```

or the equivalent

```c++
// Provided by VK_KHR_maintenance5
typedef VkSubresourceLayout2 VkSubresourceLayout2KHR;
```

or the equivalent

```c++
// Provided by VK_EXT_host_image_copy, VK_EXT_image_compression_control
typedef VkSubresourceLayout2 VkSubresourceLayout2EXT;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `subresourceLayout` is a [VkSubresourceLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubresourceLayout.html) structure.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkSubresourceLayout2-sType-sType)VUID-VkSubresourceLayout2-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_SUBRESOURCE_LAYOUT_2`
- [](#VUID-VkSubresourceLayout2-pNext-pNext)VUID-VkSubresourceLayout2-pNext-pNext  
  Each `pNext` member of any structure (including this one) in the `pNext` chain **must** be either `NULL` or a pointer to a valid instance of [VkImageCompressionPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCompressionPropertiesEXT.html) or [VkSubresourceHostMemcpySize](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubresourceHostMemcpySize.html)
- [](#VUID-VkSubresourceLayout2-sType-unique)VUID-VkSubresourceLayout2-sType-unique  
  The `sType` value of each structure in the `pNext` chain **must** be unique

## [](#_see_also)See Also

[VK\_EXT\_host\_image\_copy](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_host_image_copy.html), [VK\_EXT\_image\_compression\_control](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_image_compression_control.html), [VK\_KHR\_maintenance5](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_maintenance5.html), [VK\_VERSION\_1\_4](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_4.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [VkSubresourceLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubresourceLayout.html), [vkGetDeviceImageSubresourceLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetDeviceImageSubresourceLayout.html), [vkGetDeviceImageSubresourceLayoutKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetDeviceImageSubresourceLayoutKHR.html), [vkGetImageSubresourceLayout2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetImageSubresourceLayout2.html), [vkGetImageSubresourceLayout2EXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetImageSubresourceLayout2EXT.html), [vkGetImageSubresourceLayout2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetImageSubresourceLayout2KHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkSubresourceLayout2)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0