# VkDecompressMemoryRegionNV(3) Manual Page

## Name

VkDecompressMemoryRegionNV - Structure specifying decompression parameters



## [](#_c_specification)C Specification

The `VkDecompressMemoryRegionNV` structure is defined as:

```c++
// Provided by VK_NV_memory_decompression
typedef struct VkDecompressMemoryRegionNV {
    VkDeviceAddress                       srcAddress;
    VkDeviceAddress                       dstAddress;
    VkDeviceSize                          compressedSize;
    VkDeviceSize                          decompressedSize;
    VkMemoryDecompressionMethodFlagsNV    decompressionMethod;
} VkDecompressMemoryRegionNV;
```

## [](#_members)Members

- `srcAddress` is the address where compressed data is stored.
- `dstAddress` is the destination address where decompressed data will be written.
- `compressedSize` is the size of compressed data in bytes.
- `decompressedSize` is the size of decompressed data in bytes.
- `decompressionMethod` is a bitmask of `VkMemoryDecompressionMethodFlagBitsNV` with a single bit set specifying the method used to decompress data.

## [](#_description)Description

Valid Usage

- [](#VUID-VkDecompressMemoryRegionNV-srcAddress-07685)VUID-VkDecompressMemoryRegionNV-srcAddress-07685  
  The `srcAddress` **must** be 4 byte aligned
- [](#VUID-VkDecompressMemoryRegionNV-srcAddress-07686)VUID-VkDecompressMemoryRegionNV-srcAddress-07686  
  The memory in range `srcAddress` and `srcAddress` + `compressedSize` **must** be valid and bound to a `VkDeviceMemory` object
- [](#VUID-VkDecompressMemoryRegionNV-dstAddress-07687)VUID-VkDecompressMemoryRegionNV-dstAddress-07687  
  The `dstAddress` **must** be 4 byte aligned
- [](#VUID-VkDecompressMemoryRegionNV-decompressionMethod-09395)VUID-VkDecompressMemoryRegionNV-decompressionMethod-09395  
  If `decompressionMethod` is `VK_MEMORY_DECOMPRESSION_METHOD_GDEFLATE_1_0_BIT_NV`, then `decompressedSize` **must** be less than or equal to 65536 bytes
- [](#VUID-VkDecompressMemoryRegionNV-dstAddress-07688)VUID-VkDecompressMemoryRegionNV-dstAddress-07688  
  The memory in range `dstAddress` and `dstAddress` + `decompressedSize` **must** be valid and bound to a `VkDeviceMemory` object
- [](#VUID-VkDecompressMemoryRegionNV-decompressedSize-07689)VUID-VkDecompressMemoryRegionNV-decompressedSize-07689  
  The `decompressedSize` **must** be large enough to hold the decompressed data based on the `decompressionMethod`
- [](#VUID-VkDecompressMemoryRegionNV-decompressionMethod-07690)VUID-VkDecompressMemoryRegionNV-decompressionMethod-07690  
  The `decompressionMethod` **must** have a single bit set
- [](#VUID-VkDecompressMemoryRegionNV-srcAddress-07691)VUID-VkDecompressMemoryRegionNV-srcAddress-07691  
  The `srcAddress` to `srcAddress` + `compressedSize` region **must** not overlap with the `dstAddress` and `dstAddress` + `decompressedSize` region

Valid Usage (Implicit)

- [](#VUID-VkDecompressMemoryRegionNV-srcAddress-parameter)VUID-VkDecompressMemoryRegionNV-srcAddress-parameter  
  `srcAddress` **must** be a valid [VkDeviceAddress](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceAddress.html) value
- [](#VUID-VkDecompressMemoryRegionNV-dstAddress-parameter)VUID-VkDecompressMemoryRegionNV-dstAddress-parameter  
  `dstAddress` **must** be a valid [VkDeviceAddress](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceAddress.html) value
- [](#VUID-VkDecompressMemoryRegionNV-decompressionMethod-parameter)VUID-VkDecompressMemoryRegionNV-decompressionMethod-parameter  
  `decompressionMethod` **must** be a valid combination of [VkMemoryDecompressionMethodFlagBitsNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryDecompressionMethodFlagBitsNV.html) values
- [](#VUID-VkDecompressMemoryRegionNV-decompressionMethod-requiredbitmask)VUID-VkDecompressMemoryRegionNV-decompressionMethod-requiredbitmask  
  `decompressionMethod` **must** not be `0`

## [](#_see_also)See Also

[VK\_NV\_memory\_decompression](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_memory_decompression.html), [VkDeviceAddress](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceAddress.html), [VkDeviceSize](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceSize.html), [VkMemoryDecompressionMethodFlagsNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryDecompressionMethodFlagsNV.html), [vkCmdDecompressMemoryNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdDecompressMemoryNV.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkDecompressMemoryRegionNV)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0