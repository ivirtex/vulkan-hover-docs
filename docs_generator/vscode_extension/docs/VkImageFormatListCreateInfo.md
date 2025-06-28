# VkImageFormatListCreateInfo(3) Manual Page

## Name

VkImageFormatListCreateInfo - Specify that an image can: be used with a particular set of formats



## [](#_c_specification)C Specification

If the `pNext` chain of [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateInfo.html) includes a `VkImageFormatListCreateInfo` structure, then that structure contains a list of all formats that **can** be used when creating views of this image.

The `VkImageFormatListCreateInfo` structure is defined as:

```c++
// Provided by VK_VERSION_1_2
typedef struct VkImageFormatListCreateInfo {
    VkStructureType    sType;
    const void*        pNext;
    uint32_t           viewFormatCount;
    const VkFormat*    pViewFormats;
} VkImageFormatListCreateInfo;
```

or the equivalent

```c++
// Provided by VK_KHR_image_format_list
typedef VkImageFormatListCreateInfo VkImageFormatListCreateInfoKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `viewFormatCount` is the number of entries in the `pViewFormats` array.
- `pViewFormats` is a pointer to an array of [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) values specifying all formats which **can** be used when creating views of this image.

## [](#_description)Description

If `viewFormatCount` is zero, `pViewFormats` is ignored and the image is created as if the `VkImageFormatListCreateInfo` structure were not included in the `pNext` chain of [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateInfo.html).

Valid Usage

- [](#VUID-VkImageFormatListCreateInfo-viewFormatCount-09540)VUID-VkImageFormatListCreateInfo-viewFormatCount-09540  
  If `viewFormatCount` is not 0, each element of `pViewFormats` **must** not be `VK_FORMAT_UNDEFINED`

Valid Usage (Implicit)

- [](#VUID-VkImageFormatListCreateInfo-sType-sType)VUID-VkImageFormatListCreateInfo-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_IMAGE_FORMAT_LIST_CREATE_INFO`
- [](#VUID-VkImageFormatListCreateInfo-pViewFormats-parameter)VUID-VkImageFormatListCreateInfo-pViewFormats-parameter  
  If `viewFormatCount` is not `0`, `pViewFormats` **must** be a valid pointer to an array of `viewFormatCount` valid [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) values

## [](#_see_also)See Also

[VK\_KHR\_image\_format\_list](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_image_format_list.html), [VK\_VERSION\_1\_2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_2.html), [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkImageFormatListCreateInfo)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0