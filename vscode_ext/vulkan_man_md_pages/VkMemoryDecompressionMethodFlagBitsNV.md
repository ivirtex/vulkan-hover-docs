# VkMemoryDecompressionMethodFlagBitsNV(3) Manual Page

## Name

VkMemoryDecompressionMethodFlagBitsNV - List the supported memory
decompression methods



## <a href="#_c_specification" class="anchor"></a>C Specification

Bits which **can** be set in
`VkDecompressMemoryRegionNV`::`decompressionMethod` specifying the
decompression method to select, or returned in
`VkPhysicalDeviceMemoryDecompressionPropertiesNV`::`decompressionMethods`
specifying the available decompression methods are:

``` c
// Provided by VK_NV_memory_decompression
// Flag bits for VkMemoryDecompressionMethodFlagBitsNV
typedef VkFlags64 VkMemoryDecompressionMethodFlagBitsNV;
static const VkMemoryDecompressionMethodFlagBitsNV VK_MEMORY_DECOMPRESSION_METHOD_GDEFLATE_1_0_BIT_NV = 0x00000001ULL;
```

## <a href="#_description" class="anchor"></a>Description

- `VK_MEMORY_DECOMPRESSION_METHOD_GDEFLATE_1_0_BIT_NV` specifies that
  the GDeflate 1.0 algorithm is used to decompress data.

## <a href="#_see_also" class="anchor"></a>See Also

[VK_NV_memory_decompression](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_memory_decompression.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkMemoryDecompressionMethodFlagBitsNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
