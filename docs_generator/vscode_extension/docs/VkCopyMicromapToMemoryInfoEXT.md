# VkCopyMicromapToMemoryInfoEXT(3) Manual Page

## Name

VkCopyMicromapToMemoryInfoEXT - Parameters for serializing a micromap



## [](#_c_specification)C Specification

```c++
// Provided by VK_EXT_opacity_micromap
typedef struct VkCopyMicromapToMemoryInfoEXT {
    VkStructureType             sType;
    const void*                 pNext;
    VkMicromapEXT               src;
    VkDeviceOrHostAddressKHR    dst;
    VkCopyMicromapModeEXT       mode;
} VkCopyMicromapToMemoryInfoEXT;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `src` is the source micromap for the copy
- `dst` is the device or host address to memory which is the target for the copy
- `mode` is a [VkCopyMicromapModeEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyMicromapModeEXT.html) value specifying additional operations to perform during the copy.

## [](#_description)Description

Valid Usage

- [](#VUID-VkCopyMicromapToMemoryInfoEXT-src-07540)VUID-VkCopyMicromapToMemoryInfoEXT-src-07540  
  The source micromap `src` **must** have been constructed prior to the execution of this command
- [](#VUID-VkCopyMicromapToMemoryInfoEXT-dst-07541)VUID-VkCopyMicromapToMemoryInfoEXT-dst-07541  
  The memory pointed to by `dst` **must** be at least as large as the serialization size of `src`, as reported by [vkWriteMicromapsPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkWriteMicromapsPropertiesEXT.html) or [vkCmdWriteMicromapsPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdWriteMicromapsPropertiesEXT.html) with a query type of `VK_QUERY_TYPE_MICROMAP_SERIALIZATION_SIZE_EXT`
- [](#VUID-VkCopyMicromapToMemoryInfoEXT-mode-07542)VUID-VkCopyMicromapToMemoryInfoEXT-mode-07542  
  `mode` **must** be `VK_COPY_MICROMAP_MODE_SERIALIZE_EXT`

Valid Usage (Implicit)

- [](#VUID-VkCopyMicromapToMemoryInfoEXT-sType-sType)VUID-VkCopyMicromapToMemoryInfoEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_COPY_MICROMAP_TO_MEMORY_INFO_EXT`
- [](#VUID-VkCopyMicromapToMemoryInfoEXT-pNext-pNext)VUID-VkCopyMicromapToMemoryInfoEXT-pNext-pNext  
  `pNext` **must** be `NULL`
- [](#VUID-VkCopyMicromapToMemoryInfoEXT-src-parameter)VUID-VkCopyMicromapToMemoryInfoEXT-src-parameter  
  `src` **must** be a valid [VkMicromapEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMicromapEXT.html) handle
- [](#VUID-VkCopyMicromapToMemoryInfoEXT-mode-parameter)VUID-VkCopyMicromapToMemoryInfoEXT-mode-parameter  
  `mode` **must** be a valid [VkCopyMicromapModeEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyMicromapModeEXT.html) value

## [](#_see_also)See Also

[VK\_EXT\_opacity\_micromap](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_opacity_micromap.html), [VkCopyMicromapModeEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyMicromapModeEXT.html), [VkDeviceOrHostAddressKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceOrHostAddressKHR.html), [VkMicromapEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMicromapEXT.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkCmdCopyMicromapToMemoryEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdCopyMicromapToMemoryEXT.html), [vkCopyMicromapToMemoryEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCopyMicromapToMemoryEXT.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkCopyMicromapToMemoryInfoEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0