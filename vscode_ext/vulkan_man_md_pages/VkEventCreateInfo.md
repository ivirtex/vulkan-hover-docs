# VkEventCreateInfo(3) Manual Page

## Name

VkEventCreateInfo - Structure specifying parameters of a newly created
event



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkEventCreateInfo` structure is defined as:

``` c
// Provided by VK_VERSION_1_0
typedef struct VkEventCreateInfo {
    VkStructureType       sType;
    const void*           pNext;
    VkEventCreateFlags    flags;
} VkEventCreateInfo;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `flags` is a bitmask of
  [VkEventCreateFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkEventCreateFlagBits.html) defining
  additional creation parameters.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-VkEventCreateInfo-pNext-06790"
  id="VUID-VkEventCreateInfo-pNext-06790"></a>
  VUID-VkEventCreateInfo-pNext-06790  
  If the `pNext` chain includes a
  [VkExportMetalObjectCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExportMetalObjectCreateInfoEXT.html)
  structure, its `exportObjectType` member **must** be
  `VK_EXPORT_METAL_OBJECT_TYPE_METAL_SHARED_EVENT_BIT_EXT`

Valid Usage (Implicit)

- <a href="#VUID-VkEventCreateInfo-sType-sType"
  id="VUID-VkEventCreateInfo-sType-sType"></a>
  VUID-VkEventCreateInfo-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_EVENT_CREATE_INFO`

- <a href="#VUID-VkEventCreateInfo-pNext-pNext"
  id="VUID-VkEventCreateInfo-pNext-pNext"></a>
  VUID-VkEventCreateInfo-pNext-pNext  
  Each `pNext` member of any structure (including this one) in the
  `pNext` chain **must** be either `NULL` or a pointer to a valid
  instance of
  [VkExportMetalObjectCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExportMetalObjectCreateInfoEXT.html)
  or
  [VkImportMetalSharedEventInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImportMetalSharedEventInfoEXT.html)

- <a href="#VUID-VkEventCreateInfo-sType-unique"
  id="VUID-VkEventCreateInfo-sType-unique"></a>
  VUID-VkEventCreateInfo-sType-unique  
  The `sType` value of each struct in the `pNext` chain **must** be
  unique, with the exception of structures of type
  [VkExportMetalObjectCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExportMetalObjectCreateInfoEXT.html)

- <a href="#VUID-VkEventCreateInfo-flags-parameter"
  id="VUID-VkEventCreateInfo-flags-parameter"></a>
  VUID-VkEventCreateInfo-flags-parameter  
  `flags` **must** be a valid combination of
  [VkEventCreateFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkEventCreateFlagBits.html) values

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkEventCreateFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkEventCreateFlags.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkCreateEvent](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateEvent.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkEventCreateInfo"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
