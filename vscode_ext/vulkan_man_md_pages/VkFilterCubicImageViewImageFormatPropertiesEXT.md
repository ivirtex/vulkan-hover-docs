# VkFilterCubicImageViewImageFormatPropertiesEXT(3) Manual Page

## Name

VkFilterCubicImageViewImageFormatPropertiesEXT - Structure for querying
cubic filtering capabilities of an image view type



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkFilterCubicImageViewImageFormatPropertiesEXT` structure is
defined as:

``` c
// Provided by VK_EXT_filter_cubic
typedef struct VkFilterCubicImageViewImageFormatPropertiesEXT {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           filterCubic;
    VkBool32           filterCubicMinmax;
} VkFilterCubicImageViewImageFormatPropertiesEXT;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `filterCubic` tells if image format, image type and image view type
  **can** be used with cubic filtering. This field is set by the
  implementation. User-specified value is ignored.

- `filterCubicMinmax` tells if image format, image type and image view
  type **can** be used with cubic filtering and minmax filtering. This
  field is set by the implementation. User-specified value is ignored.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a
  href="#VUID-VkFilterCubicImageViewImageFormatPropertiesEXT-sType-sType"
  id="VUID-VkFilterCubicImageViewImageFormatPropertiesEXT-sType-sType"></a>
  VUID-VkFilterCubicImageViewImageFormatPropertiesEXT-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_FILTER_CUBIC_IMAGE_VIEW_IMAGE_FORMAT_PROPERTIES_EXT`

Valid Usage

- <a
  href="#VUID-VkFilterCubicImageViewImageFormatPropertiesEXT-pNext-02627"
  id="VUID-VkFilterCubicImageViewImageFormatPropertiesEXT-pNext-02627"></a>
  VUID-VkFilterCubicImageViewImageFormatPropertiesEXT-pNext-02627  
  If the `pNext` chain of the
  [VkImageFormatProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageFormatProperties2.html) structure
  includes a
  [VkFilterCubicImageViewImageFormatPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFilterCubicImageViewImageFormatPropertiesEXT.html)
  structure, the `pNext` chain of the
  [VkPhysicalDeviceImageFormatInfo2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceImageFormatInfo2.html)
  structure **must** include a
  [VkPhysicalDeviceImageViewImageFormatInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceImageViewImageFormatInfoEXT.html)
  structure with an `imageViewType` that is compatible with `imageType`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_filter_cubic](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_filter_cubic.html),
[VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkFilterCubicImageViewImageFormatPropertiesEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
