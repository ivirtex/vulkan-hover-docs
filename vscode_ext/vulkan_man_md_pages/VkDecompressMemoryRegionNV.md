# VkDecompressMemoryRegionNV(3) Manual Page

## Name

VkDecompressMemoryRegionNV - Structure specifying decompression
parameters



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkDecompressMemoryRegionNV` structure is defined as:

``` c
// Provided by VK_NV_memory_decompression
typedef struct VkDecompressMemoryRegionNV {
    VkDeviceAddress                       srcAddress;
    VkDeviceAddress                       dstAddress;
    VkDeviceSize                          compressedSize;
    VkDeviceSize                          decompressedSize;
    VkMemoryDecompressionMethodFlagsNV    decompressionMethod;
} VkDecompressMemoryRegionNV;
```

## <a href="#_members" class="anchor"></a>Members

- `srcAddress` is the address where compressed data is stored.

- `dstAddress` is the destination address where decompressed data will
  be written.

- `compressedSize` is the size of compressed data in bytes.

- `decompressedSize` is the size of decompressed data in bytes.

- `decompressionMethod` is a bitmask of
  `VkMemoryDecompressionMethodFlagBitsNV` with a single bit set
  specifying the method used to decompress data.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-VkDecompressMemoryRegionNV-srcAddress-07685"
  id="VUID-VkDecompressMemoryRegionNV-srcAddress-07685"></a>
  VUID-VkDecompressMemoryRegionNV-srcAddress-07685  
  The `srcAddress` **must** be 4 byte aligned

- <a href="#VUID-VkDecompressMemoryRegionNV-srcAddress-07686"
  id="VUID-VkDecompressMemoryRegionNV-srcAddress-07686"></a>
  VUID-VkDecompressMemoryRegionNV-srcAddress-07686  
  The memory in range `srcAddress` and `srcAddress` + `compressedSize`
  **must** be valid and bound to a `VkDeviceMemory` object

- <a href="#VUID-VkDecompressMemoryRegionNV-dstAddress-07687"
  id="VUID-VkDecompressMemoryRegionNV-dstAddress-07687"></a>
  VUID-VkDecompressMemoryRegionNV-dstAddress-07687  
  The `dstAddress` **must** be 4 byte aligned

- <a href="#VUID-VkDecompressMemoryRegionNV-decompressionMethod-09395"
  id="VUID-VkDecompressMemoryRegionNV-decompressionMethod-09395"></a>
  VUID-VkDecompressMemoryRegionNV-decompressionMethod-09395  
  If `decompressionMethod` is
  `VK_MEMORY_DECOMPRESSION_METHOD_GDEFLATE_1_0_BIT_NV`, then
  `decompressedSize` **must** be less than or equal to 65536 bytes

- <a href="#VUID-VkDecompressMemoryRegionNV-dstAddress-07688"
  id="VUID-VkDecompressMemoryRegionNV-dstAddress-07688"></a>
  VUID-VkDecompressMemoryRegionNV-dstAddress-07688  
  The memory in range `dstAddress` and `dstAddress` + `decompressedSize`
  **must** be valid and bound to a `VkDeviceMemory` object

- <a href="#VUID-VkDecompressMemoryRegionNV-decompressedSize-07689"
  id="VUID-VkDecompressMemoryRegionNV-decompressedSize-07689"></a>
  VUID-VkDecompressMemoryRegionNV-decompressedSize-07689  
  The `decompressedSize` **must** be large enough to hold the
  decompressed data based on the `decompressionMethod`

- <a href="#VUID-VkDecompressMemoryRegionNV-decompressionMethod-07690"
  id="VUID-VkDecompressMemoryRegionNV-decompressionMethod-07690"></a>
  VUID-VkDecompressMemoryRegionNV-decompressionMethod-07690  
  The `decompressionMethod` **must** have a single bit set

- <a href="#VUID-VkDecompressMemoryRegionNV-srcAddress-07691"
  id="VUID-VkDecompressMemoryRegionNV-srcAddress-07691"></a>
  VUID-VkDecompressMemoryRegionNV-srcAddress-07691  
  The `srcAddress` to `srcAddress` + `compressedSize` region **must**
  not overlap with the `dstAddress` and `dstAddress` +
  `decompressedSize` region

Valid Usage (Implicit)

- <a href="#VUID-VkDecompressMemoryRegionNV-decompressionMethod-parameter"
  id="VUID-VkDecompressMemoryRegionNV-decompressionMethod-parameter"></a>
  VUID-VkDecompressMemoryRegionNV-decompressionMethod-parameter  
  `decompressionMethod` **must** be a valid combination of
  [VkMemoryDecompressionMethodFlagBitsNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryDecompressionMethodFlagBitsNV.html)
  values

- <a
  href="#VUID-VkDecompressMemoryRegionNV-decompressionMethod-requiredbitmask"
  id="VUID-VkDecompressMemoryRegionNV-decompressionMethod-requiredbitmask"></a>
  VUID-VkDecompressMemoryRegionNV-decompressionMethod-requiredbitmask  
  `decompressionMethod` **must** not be `0`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_NV_memory_decompression](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_memory_decompression.html),
[VkDeviceAddress](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceAddress.html),
[VkDeviceSize](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceSize.html),
[VkMemoryDecompressionMethodFlagsNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryDecompressionMethodFlagsNV.html),
[vkCmdDecompressMemoryNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdDecompressMemoryNV.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkDecompressMemoryRegionNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
