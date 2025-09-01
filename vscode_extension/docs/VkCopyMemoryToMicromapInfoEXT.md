# VkCopyMemoryToMicromapInfoEXT(3) Manual Page

## Name

VkCopyMemoryToMicromapInfoEXT - Parameters for deserializing a micromap



## [](#_c_specification)C Specification

The `VkCopyMemoryToMicromapInfoEXT` structure is defined as:

```c++
// Provided by VK_EXT_opacity_micromap
typedef struct VkCopyMemoryToMicromapInfoEXT {
    VkStructureType                  sType;
    const void*                      pNext;
    VkDeviceOrHostAddressConstKHR    src;
    VkMicromapEXT                    dst;
    VkCopyMicromapModeEXT            mode;
} VkCopyMemoryToMicromapInfoEXT;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `src` is the device or host address to memory containing the source data for the copy.
- `dst` is the target micromap for the copy.
- `mode` is a [VkCopyMicromapModeEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyMicromapModeEXT.html) value specifying additional operations to perform during the copy.

## [](#_description)Description

Valid Usage

- [](#VUID-VkCopyMemoryToMicromapInfoEXT-src-07547)VUID-VkCopyMemoryToMicromapInfoEXT-src-07547  
  The source memory pointed to by `src` **must** contain data previously serialized using [vkCmdCopyMicromapToMemoryEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdCopyMicromapToMemoryEXT.html)
- [](#VUID-VkCopyMemoryToMicromapInfoEXT-mode-07548)VUID-VkCopyMemoryToMicromapInfoEXT-mode-07548  
  `mode` **must** be `VK_COPY_MICROMAP_MODE_DESERIALIZE_EXT`
- [](#VUID-VkCopyMemoryToMicromapInfoEXT-src-07549)VUID-VkCopyMemoryToMicromapInfoEXT-src-07549  
  The data in `src` **must** have a format compatible with the destination physical device as returned by [vkGetDeviceMicromapCompatibilityEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetDeviceMicromapCompatibilityEXT.html)
- [](#VUID-VkCopyMemoryToMicromapInfoEXT-dst-07550)VUID-VkCopyMemoryToMicromapInfoEXT-dst-07550  
  `dst` **must** have been created with a `size` greater than or equal to that used to serialize the data in `src`

Valid Usage (Implicit)

- [](#VUID-VkCopyMemoryToMicromapInfoEXT-sType-sType)VUID-VkCopyMemoryToMicromapInfoEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_COPY_MEMORY_TO_MICROMAP_INFO_EXT`
- [](#VUID-VkCopyMemoryToMicromapInfoEXT-pNext-pNext)VUID-VkCopyMemoryToMicromapInfoEXT-pNext-pNext  
  `pNext` **must** be `NULL`
- [](#VUID-VkCopyMemoryToMicromapInfoEXT-dst-parameter)VUID-VkCopyMemoryToMicromapInfoEXT-dst-parameter  
  `dst` **must** be a valid [VkMicromapEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMicromapEXT.html) handle
- [](#VUID-VkCopyMemoryToMicromapInfoEXT-mode-parameter)VUID-VkCopyMemoryToMicromapInfoEXT-mode-parameter  
  `mode` **must** be a valid [VkCopyMicromapModeEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyMicromapModeEXT.html) value

## [](#_see_also)See Also

[VK\_EXT\_opacity\_micromap](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_opacity_micromap.html), [VkCopyMicromapModeEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyMicromapModeEXT.html), [VkDeviceOrHostAddressConstKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceOrHostAddressConstKHR.html), [VkMicromapEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMicromapEXT.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkCmdCopyMemoryToMicromapEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdCopyMemoryToMicromapEXT.html), [vkCopyMemoryToMicromapEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCopyMemoryToMicromapEXT.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkCopyMemoryToMicromapInfoEXT).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0