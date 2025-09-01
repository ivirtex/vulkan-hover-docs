# VkPhysicalDeviceMemoryDecompressionPropertiesNV(3) Manual Page

## Name

VkPhysicalDeviceMemoryDecompressionPropertiesNV - Structure describing supported memory decompression methods by an implementation



## [](#_c_specification)C Specification

The `VkPhysicalDeviceMemoryDecompressionPropertiesNV` structure is defined as:

```c++
// Provided by VK_NV_memory_decompression
typedef struct VkPhysicalDeviceMemoryDecompressionPropertiesNV {
    VkStructureType                       sType;
    void*                                 pNext;
    VkMemoryDecompressionMethodFlagsNV    decompressionMethods;
    uint64_t                              maxDecompressionIndirectCount;
} VkPhysicalDeviceMemoryDecompressionPropertiesNV;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `decompressionMethods` is a bitmask of [VkMemoryDecompressionMethodFlagBitsNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryDecompressionMethodFlagBitsNV.html) specifying memory decompression methods supported by the implementation.
- `maxDecompressionIndirectCount` specifies the maximum supported count value in the `countBuffer` of [vkCmdDecompressMemoryIndirectCountNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdDecompressMemoryIndirectCountNV.html)

## [](#_description)Description

If the `VkPhysicalDeviceMemoryDecompressionPropertiesNV` structure is included in the `pNext` chain of the [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties2.html) structure passed to [vkGetPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceProperties2.html), it is filled in with each corresponding implementation-dependent property.

Valid Usage (Implicit)

- [](#VUID-VkPhysicalDeviceMemoryDecompressionPropertiesNV-sType-sType)VUID-VkPhysicalDeviceMemoryDecompressionPropertiesNV-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MEMORY_DECOMPRESSION_PROPERTIES_NV`

## [](#_see_also)See Also

[VK\_NV\_memory\_decompression](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_memory_decompression.html), [VkMemoryDecompressionMethodFlagsNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryDecompressionMethodFlagsNV.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPhysicalDeviceMemoryDecompressionPropertiesNV).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0