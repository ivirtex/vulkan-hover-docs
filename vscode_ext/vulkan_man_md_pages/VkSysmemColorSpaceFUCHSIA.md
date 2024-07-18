# VkSysmemColorSpaceFUCHSIA(3) Manual Page

## Name

VkSysmemColorSpaceFUCHSIA - Structure describing the buffer collections
color space



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkSysmemColorSpaceFUCHSIA` structure is defined as:

``` c
// Provided by VK_FUCHSIA_buffer_collection
typedef struct VkSysmemColorSpaceFUCHSIA {
    VkStructureType    sType;
    const void*        pNext;
    uint32_t           colorSpace;
} VkSysmemColorSpaceFUCHSIA;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this structure

- `colorSpace` value of the Sysmem `ColorSpaceType`

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-VkSysmemColorSpaceFUCHSIA-colorSpace-06402"
  id="VUID-VkSysmemColorSpaceFUCHSIA-colorSpace-06402"></a>
  VUID-VkSysmemColorSpaceFUCHSIA-colorSpace-06402  
  `colorSpace` **must** be a `ColorSpaceType` as defined in
  `fuchsia.sysmem/image_formats.fidl`

Valid Usage (Implicit)

- <a href="#VUID-VkSysmemColorSpaceFUCHSIA-sType-sType"
  id="VUID-VkSysmemColorSpaceFUCHSIA-sType-sType"></a>
  VUID-VkSysmemColorSpaceFUCHSIA-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_SYSMEM_COLOR_SPACE_FUCHSIA`

- <a href="#VUID-VkSysmemColorSpaceFUCHSIA-pNext-pNext"
  id="VUID-VkSysmemColorSpaceFUCHSIA-pNext-pNext"></a>
  VUID-VkSysmemColorSpaceFUCHSIA-pNext-pNext  
  `pNext` **must** be `NULL`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_FUCHSIA_buffer_collection](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_FUCHSIA_buffer_collection.html),
[VkBufferCollectionPropertiesFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferCollectionPropertiesFUCHSIA.html),
[VkImageFormatConstraintsInfoFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageFormatConstraintsInfoFUCHSIA.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkSysmemColorSpaceFUCHSIA"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
