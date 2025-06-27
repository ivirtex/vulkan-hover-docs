# VkPipelineShaderStageCreateFlagBits(3) Manual Page

## Name

VkPipelineShaderStageCreateFlagBits - Bitmask controlling how a pipeline
shader stage is created



## <a href="#_c_specification" class="anchor"></a>C Specification

Possible values of the `flags` member of
[VkPipelineShaderStageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineShaderStageCreateInfo.html)
specifying how a pipeline shader stage is created, are:

``` c
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

## <a href="#_description" class="anchor"></a>Description

- `VK_PIPELINE_SHADER_STAGE_CREATE_ALLOW_VARYING_SUBGROUP_SIZE_BIT`
  specifies that the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#interfaces-builtin-variables-sgs"
  target="_blank" rel="noopener"><code>SubgroupSize</code></a> **may**
  vary in the shader stage.

- `VK_PIPELINE_SHADER_STAGE_CREATE_REQUIRE_FULL_SUBGROUPS_BIT` specifies
  that the subgroup sizes **must** be launched with all invocations
  active in the task, mesh, or compute stage.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr>
<td class="icon"><em></em></td>
<td class="content">Note
<p>If
<code>VK_PIPELINE_SHADER_STAGE_CREATE_ALLOW_VARYING_SUBGROUP_SIZE_BIT_EXT</code>
and
<code>VK_PIPELINE_SHADER_STAGE_CREATE_REQUIRE_FULL_SUBGROUPS_BIT_EXT</code>
are specified and <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-minSubgroupSize"
target="_blank" rel="noopener"><code>minSubgroupSize</code></a> does not
equal <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-maxSubgroupSize"
target="_blank" rel="noopener"><code>maxSubgroupSize</code></a> and no
<a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-required-subgroup-size"
target="_blank" rel="noopener">required subgroup size</a> is specified,
then the only way to guarantee that the 'X' dimension of the local
workgroup size is a multiple of <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#interfaces-builtin-variables-sgs"
target="_blank" rel="noopener"><code>SubgroupSize</code></a> is to make
it a multiple of <code>maxSubgroupSize</code>. Under these conditions,
you are guaranteed full subgroups but not any particular subgroup
size.</p></td>
</tr>
</tbody>
</table>

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkPipelineShaderStageCreateFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineShaderStageCreateFlags.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPipelineShaderStageCreateFlagBits"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
