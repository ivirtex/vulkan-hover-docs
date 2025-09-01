# VkFilterCubicImageViewImageFormatPropertiesEXT(3) Manual Page

## Name

VkFilterCubicImageViewImageFormatPropertiesEXT - Structure for querying cubic filtering capabilities of an image view type



## [](#_c_specification)C Specification

The `VkFilterCubicImageViewImageFormatPropertiesEXT` structure is defined as:

```c++
// Provided by VK_EXT_filter_cubic
typedef struct VkFilterCubicImageViewImageFormatPropertiesEXT {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           filterCubic;
    VkBool32           filterCubicMinmax;
} VkFilterCubicImageViewImageFormatPropertiesEXT;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `filterCubic` tells if image format, image type and image view type **can** be used with cubic filtering. This field is set by the implementation. An application-specified value is ignored.
- `filterCubicMinmax` tells if image format, image type and image view type **can** be used with cubic filtering and minmax filtering. This field is set by the implementation. An application-specified value is ignored.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkFilterCubicImageViewImageFormatPropertiesEXT-sType-sType)VUID-VkFilterCubicImageViewImageFormatPropertiesEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_FILTER_CUBIC_IMAGE_VIEW_IMAGE_FORMAT_PROPERTIES_EXT`

Valid Usage

- [](#VUID-VkFilterCubicImageViewImageFormatPropertiesEXT-pNext-02627)VUID-VkFilterCubicImageViewImageFormatPropertiesEXT-pNext-02627  
  If the `pNext` chain of the [VkImageFormatProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageFormatProperties2.html) structure includes a [VkFilterCubicImageViewImageFormatPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFilterCubicImageViewImageFormatPropertiesEXT.html) structure, the `pNext` chain of the [VkPhysicalDeviceImageFormatInfo2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceImageFormatInfo2.html) structure **must** include a [VkPhysicalDeviceImageViewImageFormatInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceImageViewImageFormatInfoEXT.html) structure with an `imageViewType` that is compatible with `imageType`

## [](#_see_also)See Also

[VK\_EXT\_filter\_cubic](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_filter_cubic.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkFilterCubicImageViewImageFormatPropertiesEXT).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0