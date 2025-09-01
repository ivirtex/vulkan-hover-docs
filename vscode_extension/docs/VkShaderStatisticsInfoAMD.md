# VkShaderStatisticsInfoAMD(3) Manual Page

## Name

VkShaderStatisticsInfoAMD - Statistical information about a particular shader within a pipeline



## [](#_c_specification)C Specification

The `VkShaderStatisticsInfoAMD` structure is defined as:

```c++
// Provided by VK_AMD_shader_info
typedef struct VkShaderStatisticsInfoAMD {
    VkShaderStageFlags          shaderStageMask;
    VkShaderResourceUsageAMD    resourceUsage;
    uint32_t                    numPhysicalVgprs;
    uint32_t                    numPhysicalSgprs;
    uint32_t                    numAvailableVgprs;
    uint32_t                    numAvailableSgprs;
    uint32_t                    computeWorkGroupSize[3];
} VkShaderStatisticsInfoAMD;
```

## [](#_members)Members

- `shaderStageMask` are the combination of logical shader stages contained within this shader.
- `resourceUsage` is a [VkShaderResourceUsageAMD](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderResourceUsageAMD.html) structure describing internal physical device resources used by this shader.
- `numPhysicalVgprs` is the maximum number of vector instruction general-purpose registers (VGPRs) available to the physical device.
- `numPhysicalSgprs` is the maximum number of scalar instruction general-purpose registers (SGPRs) available to the physical device.
- `numAvailableVgprs` is the maximum limit of VGPRs made available to the shader compiler.
- `numAvailableSgprs` is the maximum limit of SGPRs made available to the shader compiler.
- `computeWorkGroupSize` is the local workgroup size of this shader in { X, Y, Z } dimensions.

## [](#_description)Description

Some implementations may merge multiple logical shader stages together in a single shader. In such cases, `shaderStageMask` will contain a bitmask of all of the stages that are active within that shader. Consequently, if specifying those stages as input to [vkGetShaderInfoAMD](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetShaderInfoAMD.html), the same output information **may** be returned for all such shader stage queries.

The number of available VGPRs and SGPRs (`numAvailableVgprs` and `numAvailableSgprs` respectively) are the shader-addressable subset of physical registers that is given as a limit to the compiler for register assignment. These values **may** further be limited by implementations due to performance optimizations where register pressure is a bottleneck.

## [](#_see_also)See Also

[VK\_AMD\_shader\_info](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_AMD_shader_info.html), [VkShaderResourceUsageAMD](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderResourceUsageAMD.html), [VkShaderStageFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderStageFlags.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkShaderStatisticsInfoAMD).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0