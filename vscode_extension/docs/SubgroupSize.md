# SubgroupSize(3) Manual Page

## Name

SubgroupSize - Size of a subgroup



## [](#_description)Description

`SubgroupSize`

Decorating a variable with the `SubgroupSize` builtin decoration will make that variable contain the implementation-dependent [number of invocations in a subgroup](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-subgroupSize). This value **must** be a power-of-two integer.

If the pipeline was created with the `VK_PIPELINE_SHADER_STAGE_CREATE_ALLOW_VARYING_SUBGROUP_SIZE_BIT` flag set, or the shader object was created with the `VK_SHADER_CREATE_ALLOW_VARYING_SUBGROUP_SIZE_BIT_EXT` flag set, or the SPIR-V `module` is at least version 1.6, the `SubgroupSize` decorated variable will contain the subgroup size for each subgroup that gets dispatched. This value **must** be between [`minSubgroupSize`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-minSubgroupSize) and [`maxSubgroupSize`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-maxSubgroupSize) and **must** be uniform with [subgroup scope](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#shaders-scope-subgroup). The value **may** vary across a single draw call, and for fragment shaders **may** vary across a single primitive. In compute dispatches, `SubgroupSize` **must** be uniform with [command scope](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#shaders-scope-command).

If the pipeline was created with a chained [VkPipelineShaderStageRequiredSubgroupSizeCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineShaderStageRequiredSubgroupSizeCreateInfo.html) structure, or the shader object was created with a chained [VkShaderRequiredSubgroupSizeCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderRequiredSubgroupSizeCreateInfoEXT.html) structure, the `SubgroupSize` decorated variable will match [`requiredSubgroupSize`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#pipelines-required-subgroup-size).

If SPIR-V `module` is less than version 1.6 and the pipeline was not created with the `VK_PIPELINE_SHADER_STAGE_CREATE_ALLOW_VARYING_SUBGROUP_SIZE_BIT` flag set and no [VkPipelineShaderStageRequiredSubgroupSizeCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineShaderStageRequiredSubgroupSizeCreateInfo.html) structure was chained, and the shader was not created with the `VK_SHADER_CREATE_ALLOW_VARYING_SUBGROUP_SIZE_BIT_EXT` flag set and no [VkShaderRequiredSubgroupSizeCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderRequiredSubgroupSizeCreateInfoEXT.html) structure was chained, the variable decorated with `SubgroupSize` will match [`subgroupSize`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-subgroupSize).

The maximum number of invocations that an implementation can support per subgroup is 128.

Note

The old behavior for `SubgroupSize` is considered deprecated as certain compute algorithms cannot be easily implemented without the guarantees of `VK_PIPELINE_SHADER_STAGE_CREATE_ALLOW_VARYING_SUBGROUP_SIZE_BIT` and `VK_PIPELINE_SHADER_STAGE_CREATE_REQUIRE_FULL_SUBGROUPS_BIT`.

Valid Usage

- [](#VUID-SubgroupSize-SubgroupSize-04382)VUID-SubgroupSize-SubgroupSize-04382  
  The variable decorated with `SubgroupSize` **must** be declared using the `Input` `Storage` `Class`
- [](#VUID-SubgroupSize-SubgroupSize-04383)VUID-SubgroupSize-SubgroupSize-04383  
  The variable decorated with `SubgroupSize` **must** be declared as a scalar 32-bit integer value

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#SubgroupSize).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0