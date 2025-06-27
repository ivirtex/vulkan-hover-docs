# VkPhysicalDeviceImageViewImageFormatInfoEXT(3) Manual Page

## Name

VkPhysicalDeviceImageViewImageFormatInfoEXT - Structure for providing
image view type



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPhysicalDeviceImageViewImageFormatInfoEXT` structure is defined
as:

``` c
// Provided by VK_EXT_filter_cubic
typedef struct VkPhysicalDeviceImageViewImageFormatInfoEXT {
    VkStructureType    sType;
    void*              pNext;
    VkImageViewType    imageViewType;
} VkPhysicalDeviceImageViewImageFormatInfoEXT;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `imageViewType` is a [VkImageViewType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageViewType.html) value
  specifying the type of the image view.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a href="#VUID-VkPhysicalDeviceImageViewImageFormatInfoEXT-sType-sType"
  id="VUID-VkPhysicalDeviceImageViewImageFormatInfoEXT-sType-sType"></a>
  VUID-VkPhysicalDeviceImageViewImageFormatInfoEXT-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_IMAGE_VIEW_IMAGE_FORMAT_INFO_EXT`

- <a
  href="#VUID-VkPhysicalDeviceImageViewImageFormatInfoEXT-imageViewType-parameter"
  id="VUID-VkPhysicalDeviceImageViewImageFormatInfoEXT-imageViewType-parameter"></a>
  VUID-VkPhysicalDeviceImageViewImageFormatInfoEXT-imageViewType-parameter  
  `imageViewType` **must** be a valid
  [VkImageViewType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageViewType.html) value

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_filter_cubic](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_filter_cubic.html),
[VkImageViewType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageViewType.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDeviceImageViewImageFormatInfoEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
