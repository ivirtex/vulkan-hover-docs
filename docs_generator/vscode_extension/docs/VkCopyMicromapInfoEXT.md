# VkCopyMicromapInfoEXT(3) Manual Page

## Name

VkCopyMicromapInfoEXT - Parameters for copying a micromap



## [](#_c_specification)C Specification

The `VkCopyMicromapInfoEXT` structure is defined as:

```c++
// Provided by VK_EXT_opacity_micromap
typedef struct VkCopyMicromapInfoEXT {
    VkStructureType          sType;
    const void*              pNext;
    VkMicromapEXT            src;
    VkMicromapEXT            dst;
    VkCopyMicromapModeEXT    mode;
} VkCopyMicromapInfoEXT;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `src` is the source micromap for the copy.
- `dst` is the target micromap for the copy.
- `mode` is a [VkCopyMicromapModeEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyMicromapModeEXT.html) value specifying additional operations to perform during the copy.

## [](#_description)Description

Valid Usage

- [](#VUID-VkCopyMicromapInfoEXT-mode-07531)VUID-VkCopyMicromapInfoEXT-mode-07531  
  `mode` **must** be `VK_COPY_MICROMAP_MODE_COMPACT_EXT` or `VK_COPY_MICROMAP_MODE_CLONE_EXT`
- [](#VUID-VkCopyMicromapInfoEXT-src-07532)VUID-VkCopyMicromapInfoEXT-src-07532  
  The source acceleration structure `src` **must** have been constructed prior to the execution of this command
- [](#VUID-VkCopyMicromapInfoEXT-mode-07533)VUID-VkCopyMicromapInfoEXT-mode-07533  
  If `mode` is `VK_COPY_MICROMAP_MODE_COMPACT_EXT`, `src` **must** have been constructed with `VK_BUILD_MICROMAP_ALLOW_COMPACTION_BIT_EXT` in the build
- [](#VUID-VkCopyMicromapInfoEXT-buffer-07534)VUID-VkCopyMicromapInfoEXT-buffer-07534  
  The `buffer` used to create `src` **must** be bound to device memory
- [](#VUID-VkCopyMicromapInfoEXT-buffer-07535)VUID-VkCopyMicromapInfoEXT-buffer-07535  
  The `buffer` used to create `dst` **must** be bound to device memory

Valid Usage (Implicit)

- [](#VUID-VkCopyMicromapInfoEXT-sType-sType)VUID-VkCopyMicromapInfoEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_COPY_MICROMAP_INFO_EXT`
- [](#VUID-VkCopyMicromapInfoEXT-pNext-pNext)VUID-VkCopyMicromapInfoEXT-pNext-pNext  
  `pNext` **must** be `NULL`
- [](#VUID-VkCopyMicromapInfoEXT-src-parameter)VUID-VkCopyMicromapInfoEXT-src-parameter  
  `src` **must** be a valid [VkMicromapEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMicromapEXT.html) handle
- [](#VUID-VkCopyMicromapInfoEXT-dst-parameter)VUID-VkCopyMicromapInfoEXT-dst-parameter  
  `dst` **must** be a valid [VkMicromapEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMicromapEXT.html) handle
- [](#VUID-VkCopyMicromapInfoEXT-mode-parameter)VUID-VkCopyMicromapInfoEXT-mode-parameter  
  `mode` **must** be a valid [VkCopyMicromapModeEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyMicromapModeEXT.html) value
- [](#VUID-VkCopyMicromapInfoEXT-commonparent)VUID-VkCopyMicromapInfoEXT-commonparent  
  Both of `dst`, and `src` **must** have been created, allocated, or retrieved from the same [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html)

## [](#_see_also)See Also

[VK\_EXT\_opacity\_micromap](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_opacity_micromap.html), [VkCopyMicromapModeEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyMicromapModeEXT.html), [VkMicromapEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMicromapEXT.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkCmdCopyMicromapEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdCopyMicromapEXT.html), [vkCopyMicromapEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCopyMicromapEXT.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkCopyMicromapInfoEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0