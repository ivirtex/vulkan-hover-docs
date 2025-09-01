# VkPipelineShaderStageCreateFlagBits(3) Manual Page

## Name

VkPipelineShaderStageCreateFlagBits - Bitmask controlling how a pipeline shader stage is created



## [](#_c_specification)C Specification

Possible values of the `flags` member of [VkPipelineShaderStageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineShaderStageCreateInfo.html) specifying how a pipeline shader stage is created, are:

```c++
// Provided by VK_VERSION_1_0
typedef enum VkPipelineShaderStageCreateFlagBits {
  // Provided by VK_VERSION_1_3
    VK_PIPELINE_SHADER_STAGE_CREATE_ALLOW_VARYING_SUBGROUP_SIZE_BIT = 0x00000001,
  // Provided by VK_VERSION_1_3
    VK_PIPELINE_SHADER_STAGE_CREATE_REQUIRE_FULL_SUBGROUPS_BIT = 0x00000002,
  // Provided by VK_EXT_subgroup_size_control
    VK_PIPELINE_SHADER_STAGE_CREATE_ALLOW_VARYING_SUBGROUP_SIZE_BIT_EXT = VK_PIPELINE_SHADER_STAGE_CREATE_ALLOW_VARYING_SUBGROUP_SIZE_BIT,
  // Provided by VK_EXT_subgroup_size_control
    VK_PIPELINE_SHADER_STAGE_CREATE_REQUIRE_FULL_SUBGROUPS_BIT_EXT = VK_PIPELINE_SHADER_STAGE_CREATE_REQUIRE_FULL_SUBGROUPS_BIT,
} VkPipelineShaderStageCreateFlagBits;
```

## [](#_description)Description

- `VK_PIPELINE_SHADER_STAGE_CREATE_ALLOW_VARYING_SUBGROUP_SIZE_BIT` specifies that the [`SubgroupSize`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#interfaces-builtin-variables-sgs) **may** vary in the shader stage.
- `VK_PIPELINE_SHADER_STAGE_CREATE_REQUIRE_FULL_SUBGROUPS_BIT` specifies that the subgroup sizes **must** be launched with all invocations active in the task, mesh, or compute stage.

Note

If `VK_PIPELINE_SHADER_STAGE_CREATE_ALLOW_VARYING_SUBGROUP_SIZE_BIT_EXT` and `VK_PIPELINE_SHADER_STAGE_CREATE_REQUIRE_FULL_SUBGROUPS_BIT_EXT` are specified and [`minSubgroupSize`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-minSubgroupSize) does not equal [`maxSubgroupSize`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-maxSubgroupSize) and no [required subgroup size](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#pipelines-required-subgroup-size) is specified, then the only way to guarantee that the 'X' dimension of the local workgroup size is a multiple of [`SubgroupSize`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#interfaces-builtin-variables-sgs) is to make it a multiple of `maxSubgroupSize`. Under these conditions, you are guaranteed full subgroups but not any particular subgroup size.

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkPipelineShaderStageCreateFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineShaderStageCreateFlags.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPipelineShaderStageCreateFlagBits).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0