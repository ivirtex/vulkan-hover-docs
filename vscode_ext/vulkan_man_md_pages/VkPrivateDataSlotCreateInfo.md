# VkPrivateDataSlotCreateInfo(3) Manual Page

## Name

VkPrivateDataSlotCreateInfo - Structure specifying the parameters of
private data slot construction



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPrivateDataSlotCreateInfo` structure is defined as:

``` c
// Provided by VK_VERSION_1_3
typedef struct VkPrivateDataSlotCreateInfo {
    VkStructureType                 sType;
    const void*                     pNext;
    VkPrivateDataSlotCreateFlags    flags;
} VkPrivateDataSlotCreateInfo;
```

or the equivalent

``` c
// Provided by VK_EXT_private_data
typedef VkPrivateDataSlotCreateInfo VkPrivateDataSlotCreateInfoEXT;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `flags` is reserved for future use.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a href="#VUID-VkPrivateDataSlotCreateInfo-sType-sType"
  id="VUID-VkPrivateDataSlotCreateInfo-sType-sType"></a>
  VUID-VkPrivateDataSlotCreateInfo-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PRIVATE_DATA_SLOT_CREATE_INFO`

- <a href="#VUID-VkPrivateDataSlotCreateInfo-pNext-pNext"
  id="VUID-VkPrivateDataSlotCreateInfo-pNext-pNext"></a>
  VUID-VkPrivateDataSlotCreateInfo-pNext-pNext  
  `pNext` **must** be `NULL`

- <a href="#VUID-VkPrivateDataSlotCreateInfo-flags-zerobitmask"
  id="VUID-VkPrivateDataSlotCreateInfo-flags-zerobitmask"></a>
  VUID-VkPrivateDataSlotCreateInfo-flags-zerobitmask  
  `flags` **must** be `0`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_private_data](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_private_data.html),
[VK_VERSION_1_3](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_3.html),
[VkPrivateDataSlotCreateFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPrivateDataSlotCreateFlags.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkCreatePrivateDataSlot](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreatePrivateDataSlot.html),
[vkCreatePrivateDataSlotEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreatePrivateDataSlotEXT.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPrivateDataSlotCreateInfo"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
