# VkShaderResourceUsageAMD(3) Manual Page

## Name

VkShaderResourceUsageAMD - Resource usage information about a particular
shader within a pipeline



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkShaderResourceUsageAMD` structure is defined as:

``` c
// Provided by VK_AMD_shader_info
typedef struct VkShaderResourceUsageAMD {
    uint32_t    numUsedVgprs;
    uint32_t    numUsedSgprs;
    uint32_t    ldsSizePerLocalWorkGroup;
    size_t      ldsUsageSizeInBytes;
    size_t      scratchMemUsageInBytes;
} VkShaderResourceUsageAMD;
```

## <a href="#_members" class="anchor"></a>Members

- `numUsedVgprs` is the number of vector instruction general-purpose
  registers used by this shader.

- `numUsedSgprs` is the number of scalar instruction general-purpose
  registers used by this shader.

- `ldsSizePerLocalWorkGroup` is the maximum local data store size per
  work group in bytes.

- `ldsUsageSizeInBytes` is the LDS usage size in bytes per work group by
  this shader.

- `scratchMemUsageInBytes` is the scratch memory usage in bytes by this
  shader.

## <a href="#_description" class="anchor"></a>Description

## <a href="#_see_also" class="anchor"></a>See Also

[VK_AMD_shader_info](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_AMD_shader_info.html),
[VkShaderStatisticsInfoAMD](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderStatisticsInfoAMD.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkShaderResourceUsageAMD"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
