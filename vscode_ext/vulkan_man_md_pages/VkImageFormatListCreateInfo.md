# VkImageFormatListCreateInfo(3) Manual Page

## Name

VkImageFormatListCreateInfo - Specify that an image can: be used with a
particular set of formats



## <a href="#_c_specification" class="anchor"></a>C Specification

If the `pNext` chain of [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html)
includes a `VkImageFormatListCreateInfo` structure, then that structure
contains a list of all formats that **can** be used when creating views
of this image.

The `VkImageFormatListCreateInfo` structure is defined as:

``` c
// Provided by VK_VERSION_1_2
typedef struct VkImageFormatListCreateInfo {
    VkStructureType    sType;
    const void*        pNext;
    uint32_t           viewFormatCount;
    const VkFormat*    pViewFormats;
} VkImageFormatListCreateInfo;
```

or the equivalent

``` c
// Provided by VK_KHR_image_format_list
typedef VkImageFormatListCreateInfo VkImageFormatListCreateInfoKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `viewFormatCount` is the number of entries in the `pViewFormats`
  array.

- `pViewFormats` is a pointer to an array of [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html)
  values specifying all formats which **can** be used when creating
  views of this image.

## <a href="#_description" class="anchor"></a>Description

If `viewFormatCount` is zero, `pViewFormats` is ignored and the image is
created as if the `VkImageFormatListCreateInfo` structure were not
included in the `pNext` chain of
[VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html).

Valid Usage

- <a href="#VUID-VkImageFormatListCreateInfo-viewFormatCount-09540"
  id="VUID-VkImageFormatListCreateInfo-viewFormatCount-09540"></a>
  VUID-VkImageFormatListCreateInfo-viewFormatCount-09540  
  If `viewFormatCount` is not 0, each element of `pViewFormats` **must**
  not be `VK_FORMAT_UNDEFINED`

Valid Usage (Implicit)

- <a href="#VUID-VkImageFormatListCreateInfo-sType-sType"
  id="VUID-VkImageFormatListCreateInfo-sType-sType"></a>
  VUID-VkImageFormatListCreateInfo-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_IMAGE_FORMAT_LIST_CREATE_INFO`

- <a href="#VUID-VkImageFormatListCreateInfo-pViewFormats-parameter"
  id="VUID-VkImageFormatListCreateInfo-pViewFormats-parameter"></a>
  VUID-VkImageFormatListCreateInfo-pViewFormats-parameter  
  If `viewFormatCount` is not `0`, `pViewFormats` **must** be a valid
  pointer to an array of `viewFormatCount` valid
  [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) values

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_image_format_list](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_image_format_list.html),
[VK_VERSION_1_2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_2.html), [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkImageFormatListCreateInfo"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
