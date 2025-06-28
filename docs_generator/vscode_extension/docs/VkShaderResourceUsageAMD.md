# VkShaderResourceUsageAMD(3) Manual Page

## Name

VkShaderResourceUsageAMD - Resource usage information about a particular shader within a pipeline



## [](#_c_specification)C Specification

The `VkShaderResourceUsageAMD` structure is defined as:

```c++
// Provided by VK_AMD_shader_info
typedef struct VkShaderResourceUsageAMD {
    uint32_t    numUsedVgprs;
    uint32_t    numUsedSgprs;
    uint32_t    ldsSizePerLocalWorkGroup;
    size_t      ldsUsageSizeInBytes;
    size_t      scratchMemUsageInBytes;
} VkShaderResourceUsageAMD;
```

## [](#_members)Members

- `numUsedVgprs` is the number of vector instruction general-purpose registers used by this shader.
- `numUsedSgprs` is the number of scalar instruction general-purpose registers used by this shader.
- `ldsSizePerLocalWorkGroup` is the maximum local data store size per work group in bytes.
- `ldsUsageSizeInBytes` is the LDS usage size in bytes per work group by this shader.
- `scratchMemUsageInBytes` is the scratch memory usage in bytes by this shader.

## [](#_description)Description

## [](#_see_also)See Also

[VK\_AMD\_shader\_info](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_AMD_shader_info.html), [VkShaderStatisticsInfoAMD](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderStatisticsInfoAMD.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkShaderResourceUsageAMD)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0