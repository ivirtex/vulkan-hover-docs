# VkPhysicalDeviceMemoryDecompressionPropertiesNV(3) Manual Page

## Name

VkPhysicalDeviceMemoryDecompressionPropertiesNV - Structure describing
supported memory decompression methods by an implementation



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPhysicalDeviceMemoryDecompressionPropertiesNV` structure is
defined as:

``` c
// Provided by VK_NV_memory_decompression
typedef struct VkPhysicalDeviceMemoryDecompressionPropertiesNV {
    VkStructureType                       sType;
    void*                                 pNext;
    VkMemoryDecompressionMethodFlagsNV    decompressionMethods;
    uint64_t                              maxDecompressionIndirectCount;
} VkPhysicalDeviceMemoryDecompressionPropertiesNV;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `decompressionMethods` is a bitmask of
  [VkMemoryDecompressionMethodFlagBitsNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryDecompressionMethodFlagBitsNV.html)
  specifying memory decompression methods supported by the
  implementation.

- `maxDecompressionIndirectCount` specifies the maximum supported count
  value in the `countBuffer` of
  [vkCmdDecompressMemoryIndirectCountNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdDecompressMemoryIndirectCountNV.html)

## <a href="#_description" class="anchor"></a>Description

If the `VkPhysicalDeviceMemoryDecompressionPropertiesNV` structure is
included in the `pNext` chain of the
[VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProperties2.html)
structure passed to
[vkGetPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceProperties2.html),
it is filled in with each corresponding implementation-dependent
property.

Valid Usage (Implicit)

- <a
  href="#VUID-VkPhysicalDeviceMemoryDecompressionPropertiesNV-sType-sType"
  id="VUID-VkPhysicalDeviceMemoryDecompressionPropertiesNV-sType-sType"></a>
  VUID-VkPhysicalDeviceMemoryDecompressionPropertiesNV-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MEMORY_DECOMPRESSION_PROPERTIES_NV`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_NV_memory_decompression](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_memory_decompression.html),
[VkMemoryDecompressionMethodFlagsNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryDecompressionMethodFlagsNV.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDeviceMemoryDecompressionPropertiesNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
