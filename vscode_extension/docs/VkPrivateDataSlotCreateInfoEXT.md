# VkPrivateDataSlotCreateInfo(3) Manual Page

## Name

VkPrivateDataSlotCreateInfo - Structure specifying the parameters of private data slot construction



## [](#_c_specification)C Specification

The `VkPrivateDataSlotCreateInfo` structure is defined as:

```c++
// Provided by VK_VERSION_1_3
typedef struct VkPrivateDataSlotCreateInfo {
    VkStructureType                 sType;
    const void*                     pNext;
    VkPrivateDataSlotCreateFlags    flags;
} VkPrivateDataSlotCreateInfo;
```

or the equivalent

```c++
// Provided by VK_EXT_private_data
typedef VkPrivateDataSlotCreateInfo VkPrivateDataSlotCreateInfoEXT;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `flags` is reserved for future use.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkPrivateDataSlotCreateInfo-sType-sType)VUID-VkPrivateDataSlotCreateInfo-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PRIVATE_DATA_SLOT_CREATE_INFO`
- [](#VUID-VkPrivateDataSlotCreateInfo-pNext-pNext)VUID-VkPrivateDataSlotCreateInfo-pNext-pNext  
  `pNext` **must** be `NULL`
- [](#VUID-VkPrivateDataSlotCreateInfo-flags-zerobitmask)VUID-VkPrivateDataSlotCreateInfo-flags-zerobitmask  
  `flags` **must** be `0`

## [](#_see_also)See Also

[VK\_EXT\_private\_data](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_private_data.html), [VK\_VERSION\_1\_3](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_3.html), [VkPrivateDataSlotCreateFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPrivateDataSlotCreateFlags.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkCreatePrivateDataSlot](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreatePrivateDataSlot.html), [vkCreatePrivateDataSlotEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreatePrivateDataSlotEXT.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPrivateDataSlotCreateInfo)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0