# VkPhysicalDeviceImageViewImageFormatInfoEXT(3) Manual Page

## Name

VkPhysicalDeviceImageViewImageFormatInfoEXT - Structure for providing image view type



## [](#_c_specification)C Specification

The `VkPhysicalDeviceImageViewImageFormatInfoEXT` structure is defined as:

```c++
// Provided by VK_EXT_filter_cubic
typedef struct VkPhysicalDeviceImageViewImageFormatInfoEXT {
    VkStructureType    sType;
    void*              pNext;
    VkImageViewType    imageViewType;
} VkPhysicalDeviceImageViewImageFormatInfoEXT;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `imageViewType` is a [VkImageViewType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageViewType.html) value specifying the type of the image view.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkPhysicalDeviceImageViewImageFormatInfoEXT-sType-sType)VUID-VkPhysicalDeviceImageViewImageFormatInfoEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_IMAGE_VIEW_IMAGE_FORMAT_INFO_EXT`
- [](#VUID-VkPhysicalDeviceImageViewImageFormatInfoEXT-imageViewType-parameter)VUID-VkPhysicalDeviceImageViewImageFormatInfoEXT-imageViewType-parameter  
  `imageViewType` **must** be a valid [VkImageViewType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageViewType.html) value

## [](#_see_also)See Also

[VK\_EXT\_filter\_cubic](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_filter_cubic.html), [VkImageViewType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageViewType.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPhysicalDeviceImageViewImageFormatInfoEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0