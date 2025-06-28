# VkMemoryDecompressionMethodFlagBitsNV(3) Manual Page

## Name

VkMemoryDecompressionMethodFlagBitsNV - List the supported memory decompression methods



## [](#_c_specification)C Specification

Bits which **can** be set in `VkDecompressMemoryRegionNV`::`decompressionMethod` specifying the decompression method to select, or returned in `VkPhysicalDeviceMemoryDecompressionPropertiesNV`::`decompressionMethods` specifying the available decompression methods are:

```c++
// Provided by VK_NV_memory_decompression
// Flag bits for VkMemoryDecompressionMethodFlagBitsNV
typedef VkFlags64 VkMemoryDecompressionMethodFlagBitsNV;
static const VkMemoryDecompressionMethodFlagBitsNV VK_MEMORY_DECOMPRESSION_METHOD_GDEFLATE_1_0_BIT_NV = 0x00000001ULL;
```

## [](#_description)Description

- `VK_MEMORY_DECOMPRESSION_METHOD_GDEFLATE_1_0_BIT_NV` specifies that the GDeflate 1.0 algorithm is used to decompress data.

## [](#_see_also)See Also

[VK\_NV\_memory\_decompression](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_memory_decompression.html), [VkMemoryDecompressionMethodFlagsNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryDecompressionMethodFlagsNV.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkMemoryDecompressionMethodFlagBitsNV)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0