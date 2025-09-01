# VkMicromapVersionInfoEXT(3) Manual Page

## Name

VkMicromapVersionInfoEXT - Micromap version information



## [](#_c_specification)C Specification

The `VkMicromapVersionInfoEXT` structure is defined as:

```c++
// Provided by VK_EXT_opacity_micromap
typedef struct VkMicromapVersionInfoEXT {
    VkStructureType    sType;
    const void*        pNext;
    const uint8_t*     pVersionData;
} VkMicromapVersionInfoEXT;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `pVersionData` is a pointer to the version header of a micromap as defined in [vkCmdCopyMicromapToMemoryEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdCopyMicromapToMemoryEXT.html)

## [](#_description)Description

Note

`pVersionData` is a *pointer* to an array of 2×`VK_UUID_SIZE` `uint8_t` values instead of two `VK_UUID_SIZE` arrays as the expected use case for this member is to be pointed at the header of a previously serialized micromap (via [vkCmdCopyMicromapToMemoryEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdCopyMicromapToMemoryEXT.html) or [vkCopyMicromapToMemoryEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCopyMicromapToMemoryEXT.html)) that is loaded in memory. Using arrays would necessitate extra memory copies of the UUIDs.

Valid Usage (Implicit)

- [](#VUID-VkMicromapVersionInfoEXT-sType-sType)VUID-VkMicromapVersionInfoEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_MICROMAP_VERSION_INFO_EXT`
- [](#VUID-VkMicromapVersionInfoEXT-pNext-pNext)VUID-VkMicromapVersionInfoEXT-pNext-pNext  
  `pNext` **must** be `NULL`
- [](#VUID-VkMicromapVersionInfoEXT-pVersionData-parameter)VUID-VkMicromapVersionInfoEXT-pVersionData-parameter  
  `pVersionData` **must** be a valid pointer to an array of 2×VK\_UUID\_SIZE `uint8_t` values

## [](#_see_also)See Also

[VK\_EXT\_opacity\_micromap](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_opacity_micromap.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkGetDeviceMicromapCompatibilityEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetDeviceMicromapCompatibilityEXT.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkMicromapVersionInfoEXT).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0