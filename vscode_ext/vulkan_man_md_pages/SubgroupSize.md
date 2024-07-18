# SubgroupSize(3) Manual Page

## Name

SubgroupSize - Size of a subgroup



## <a href="#_description" class="anchor"></a>Description

`SubgroupSize`  
Decorating a variable with the `SubgroupSize` builtin decoration will
make that variable contain the implementation-dependent <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-subgroupSize"
target="_blank" rel="noopener">number of invocations in a subgroup</a>.
This value **must** be a power-of-two integer.

If the pipeline was created with the
`VK_PIPELINE_SHADER_STAGE_CREATE_ALLOW_VARYING_SUBGROUP_SIZE_BIT` flag
set, or the shader object was created with the
`VK_SHADER_CREATE_ALLOW_VARYING_SUBGROUP_SIZE_BIT_EXT` flag set, or the
SPIR-V `module` is at least version 1.6, the `SubgroupSize` decorated
variable will contain the subgroup size for each subgroup that gets
dispatched. This value **must** be between <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-minSubgroupSize"
target="_blank" rel="noopener"><code>minSubgroupSize</code></a> and <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-maxSubgroupSize"
target="_blank" rel="noopener"><code>maxSubgroupSize</code></a> and
**must** be uniform with <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#shaders-scope-subgroup"
target="_blank" rel="noopener">subgroup scope</a>. The value **may**
vary across a single draw call, and for fragment shaders **may** vary
across a single primitive. In compute dispatches, `SubgroupSize`
**must** be uniform with <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#shaders-scope-command"
target="_blank" rel="noopener">command scope</a>.

If the pipeline was created with a chained
[VkPipelineShaderStageRequiredSubgroupSizeCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineShaderStageRequiredSubgroupSizeCreateInfo.html)
structure, or the shader object was created with a chained
[VkShaderRequiredSubgroupSizeCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderRequiredSubgroupSizeCreateInfoEXT.html)
structure, the `SubgroupSize` decorated variable will match <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-required-subgroup-size"
target="_blank" rel="noopener"><code>requiredSubgroupSize</code></a>.

If SPIR-V `module` is less than version 1.6 and the pipeline was not
created with the
`VK_PIPELINE_SHADER_STAGE_CREATE_ALLOW_VARYING_SUBGROUP_SIZE_BIT` flag
set and no
[VkPipelineShaderStageRequiredSubgroupSizeCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineShaderStageRequiredSubgroupSizeCreateInfo.html)
structure was chained, and the shader was not created with the
`VK_SHADER_CREATE_ALLOW_VARYING_SUBGROUP_SIZE_BIT_EXT` flag set and no
[VkShaderRequiredSubgroupSizeCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderRequiredSubgroupSizeCreateInfoEXT.html)
structure was chained, the variable decorated with `SubgroupSize` will
match <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-subgroupSize"
target="_blank" rel="noopener"><code>subgroupSize</code></a>.

The maximum number of invocations that an implementation can support per
subgroup is 128.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr>
<td class="icon"><em></em></td>
<td class="content">Note
<p>The old behavior for <code>SubgroupSize</code> is considered
deprecated as certain compute algorithms cannot be easily implemented
without the guarantees of
<code>VK_PIPELINE_SHADER_STAGE_CREATE_ALLOW_VARYING_SUBGROUP_SIZE_BIT</code>
and
<code>VK_PIPELINE_SHADER_STAGE_CREATE_REQUIRE_FULL_SUBGROUPS_BIT</code>.</p></td>
</tr>
</tbody>
</table>

Valid Usage

- <a href="#VUID-SubgroupSize-SubgroupSize-04382"
  id="VUID-SubgroupSize-SubgroupSize-04382"></a>
  VUID-SubgroupSize-SubgroupSize-04382  
  The variable decorated with `SubgroupSize` **must** be declared using
  the `Input` `Storage` `Class`

- <a href="#VUID-SubgroupSize-SubgroupSize-04383"
  id="VUID-SubgroupSize-SubgroupSize-04383"></a>
  VUID-SubgroupSize-SubgroupSize-04383  
  The variable decorated with `SubgroupSize` **must** be declared as a
  scalar 32-bit integer value

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#SubgroupSize"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
