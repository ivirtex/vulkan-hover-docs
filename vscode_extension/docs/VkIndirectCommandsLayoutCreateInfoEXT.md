# VkIndirectCommandsLayoutCreateInfoEXT(3) Manual Page

## Name

VkIndirectCommandsLayoutCreateInfoEXT - Structure specifying the parameters of a newly created indirect commands layout object



## [](#_c_specification)C Specification

The `VkIndirectCommandsLayoutCreateInfoEXT` structure is defined as:

```c++
// Provided by VK_EXT_device_generated_commands
typedef struct VkIndirectCommandsLayoutCreateInfoEXT {
    VkStructureType                            sType;
    const void*                                pNext;
    VkIndirectCommandsLayoutUsageFlagsEXT      flags;
    VkShaderStageFlags                         shaderStages;
    uint32_t                                   indirectStride;
    VkPipelineLayout                           pipelineLayout;
    uint32_t                                   tokenCount;
    const VkIndirectCommandsLayoutTokenEXT*    pTokens;
} VkIndirectCommandsLayoutCreateInfoEXT;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `flags` is a bitmask of [VkIndirectCommandsLayoutUsageFlagBitsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndirectCommandsLayoutUsageFlagBitsEXT.html) specifying usage rules for this layout.
- `shaderStages` is the [VkShaderStageFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderStageFlags.html) that this layout supports.
- `indirectStride` is the distance in bytes between sequences in the indirect buffer
- `pipelineLayout` is the optional [VkPipelineLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineLayout.html) that tokens in this layout use. If the [`dynamicGeneratedPipelineLayout`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-dynamicGeneratedPipelineLayout) feature is enabled, `pipelineLayout` **can** be [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html) and the layout **must** be specified by chaining the [VkPipelineLayoutCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineLayoutCreateInfo.html) structure off the `pNext`
- `tokenCount` is the length of the individual command sequence.
- `pTokens` is a pointer to an array of [VkIndirectCommandsLayoutTokenEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndirectCommandsLayoutTokenEXT.html) describing each command token in detail.

## [](#_description)Description

The following code illustrates some of the flags:

```c
void cmdProcessAllSequences(cmd, indirectExecutionSet, indirectCommandsLayout, indirectAddress, sequencesCount)
{
  for (s = 0; s < sequencesCount; s++)
  {
    sUsed = s;

    if (indirectCommandsLayout.flags & VK_INDIRECT_COMMANDS_LAYOUT_USAGE_UNORDERED_SEQUENCES_BIT_EXT) {
      sUsed = incoherent_implementation_dependent_permutation[ sUsed ];
    }

    cmdProcessSequence( cmd, indirectExecutionSet, indirectCommandsLayout, indirectAddress, sUsed );
  }
}
```

When tokens are consumed, an offset is computed based on token offset and stream stride. The resulting offset is required to be aligned. The alignment for a specific token is equal to the scalar alignment of the data type as defined in [Alignment Requirements](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#interfaces-alignment-requirements), or `4`, whichever is lower.

Valid Usage

- [](#VUID-VkIndirectCommandsLayoutCreateInfoEXT-indirectStride-11090)VUID-VkIndirectCommandsLayoutCreateInfoEXT-indirectStride-11090  
  `indirectStride` **must** be less than or equal to [VkPhysicalDeviceDeviceGeneratedCommandsPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceDeviceGeneratedCommandsPropertiesEXT.html)::`maxIndirectCommandsIndirectStride`
- [](#VUID-VkIndirectCommandsLayoutCreateInfoEXT-shaderStages-11091)VUID-VkIndirectCommandsLayoutCreateInfoEXT-shaderStages-11091  
  `shaderStages` **must** only contain stages supported by [](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-supportedIndirectCommandsShaderStages)[VkPhysicalDeviceDeviceGeneratedCommandsPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceDeviceGeneratedCommandsPropertiesEXT.html)::`supportedIndirectCommandsShaderStages`
- [](#VUID-VkIndirectCommandsLayoutCreateInfoEXT-tokenCount-11092)VUID-VkIndirectCommandsLayoutCreateInfoEXT-tokenCount-11092  
  `tokenCount` **must** be less than or equal to [VkPhysicalDeviceDeviceGeneratedCommandsPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceDeviceGeneratedCommandsPropertiesEXT.html)::`maxIndirectCommandsTokenCount`
- [](#VUID-VkIndirectCommandsLayoutCreateInfoEXT-pTokens-11093)VUID-VkIndirectCommandsLayoutCreateInfoEXT-pTokens-11093  
  The number of tokens in the `pTokens` array with `type` equal to `VK_INDIRECT_COMMANDS_TOKEN_TYPE_EXECUTION_SET_EXT` **must** be less than or equal to `1`
- [](#VUID-VkIndirectCommandsLayoutCreateInfoEXT-pTokens-11145)VUID-VkIndirectCommandsLayoutCreateInfoEXT-pTokens-11145  
  The number of tokens in the `pTokens` array with `type` equal to `VK_INDIRECT_COMMANDS_TOKEN_TYPE_SEQUENCE_INDEX_EXT` **must** be less than or equal to `1`
- [](#VUID-VkIndirectCommandsLayoutCreateInfoEXT-pTokens-11094)VUID-VkIndirectCommandsLayoutCreateInfoEXT-pTokens-11094  
  The number of tokens in the `pTokens` array with `type` equal to `VK_INDIRECT_COMMANDS_TOKEN_TYPE_INDEX_BUFFER_EXT` **must** be less than or equal to `1`
- [](#VUID-VkIndirectCommandsLayoutCreateInfoEXT-pTokens-11095)VUID-VkIndirectCommandsLayoutCreateInfoEXT-pTokens-11095  
  If the action command token in the `pTokens` array is not an indexed draw token, then `pTokens` **must** not contain a member with `type` set to `VK_INDIRECT_COMMANDS_TOKEN_TYPE_INDEX_BUFFER_EXT`
- [](#VUID-VkIndirectCommandsLayoutCreateInfoEXT-pTokens-11096)VUID-VkIndirectCommandsLayoutCreateInfoEXT-pTokens-11096  
  If the action command token in the `pTokens` array is not a non-mesh draw token, then `pTokens` **must** not contain a member with `type` set to `VK_INDIRECT_COMMANDS_TOKEN_TYPE_VERTEX_BUFFER_EXT`
- [](#VUID-VkIndirectCommandsLayoutCreateInfoEXT-pTokens-11097)VUID-VkIndirectCommandsLayoutCreateInfoEXT-pTokens-11097  
  If the `pTokens` array contains multiple tokens with `type` equal to `VK_INDIRECT_COMMANDS_TOKEN_TYPE_VERTEX_BUFFER_EXT`, then there **must** be no duplicate [VkIndirectCommandsVertexBufferTokenEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndirectCommandsVertexBufferTokenEXT.html)::`vertexBindingUnit` values
- [](#VUID-VkIndirectCommandsLayoutCreateInfoEXT-pTokens-11099)VUID-VkIndirectCommandsLayoutCreateInfoEXT-pTokens-11099  
  For all `VK_INDIRECT_COMMANDS_TOKEN_TYPE_PUSH_CONSTANT_EXT` and `VK_INDIRECT_COMMANDS_TOKEN_TYPE_SEQUENCE_INDEX_EXT` type tokens in `pTokens`, there **must** be no overlapping ranges between any specified push constant ranges
- [](#VUID-VkIndirectCommandsLayoutCreateInfoEXT-pTokens-11100)VUID-VkIndirectCommandsLayoutCreateInfoEXT-pTokens-11100  
  The action command token **must** be the last token in the `pTokens` array
- [](#VUID-VkIndirectCommandsLayoutCreateInfoEXT-pTokens-11139)VUID-VkIndirectCommandsLayoutCreateInfoEXT-pTokens-11139  
  If the `pTokens` array contains a `VK_INDIRECT_COMMANDS_TOKEN_TYPE_EXECUTION_SET_EXT` token, then this token **must** be the first token in the array
- [](#VUID-VkIndirectCommandsLayoutCreateInfoEXT-pTokens-11101)VUID-VkIndirectCommandsLayoutCreateInfoEXT-pTokens-11101  
  For any element of `pTokens`, if `type` is `VK_INDIRECT_COMMANDS_TOKEN_TYPE_PUSH_CONSTANT_EXT` or `VK_INDIRECT_COMMANDS_TOKEN_TYPE_SEQUENCE_INDEX_EXT` and the [`dynamicGeneratedPipelineLayout`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-dynamicGeneratedPipelineLayout) feature is not enabled, then the `pipelineLayout` **must** not be [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html)
- [](#VUID-VkIndirectCommandsLayoutCreateInfoEXT-pTokens-11102)VUID-VkIndirectCommandsLayoutCreateInfoEXT-pTokens-11102  
  For any element of `pTokens`, if `type` is either `VK_INDIRECT_COMMANDS_TOKEN_TYPE_PUSH_CONSTANT_EXT` or `VK_INDIRECT_COMMANDS_TOKEN_TYPE_SEQUENCE_INDEX_EXT` and `pipelineLayout` is [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), then the `pNext` chain **must** include a [VkPipelineLayoutCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineLayoutCreateInfo.html) struct
- [](#VUID-VkIndirectCommandsLayoutCreateInfoEXT-pTokens-11103)VUID-VkIndirectCommandsLayoutCreateInfoEXT-pTokens-11103  
  For any element of `pTokens`, the `offset` **must** be greater than or equal to the `offset` member of the previous tokens
- [](#VUID-VkIndirectCommandsLayoutCreateInfoEXT-pTokens-11104)VUID-VkIndirectCommandsLayoutCreateInfoEXT-pTokens-11104  
  For any element of `pTokens`, if `type` is `VK_INDIRECT_COMMANDS_TOKEN_TYPE_INDEX_BUFFER_EXT`, `VK_INDIRECT_COMMANDS_TOKEN_TYPE_VERTEX_BUFFER_EXT`, `VK_INDIRECT_COMMANDS_TOKEN_TYPE_DRAW_MESH_TASKS_EXT`, `VK_INDIRECT_COMMANDS_TOKEN_TYPE_DRAW_MESH_TASKS_COUNT_EXT`, `VK_INDIRECT_COMMANDS_TOKEN_TYPE_DRAW_MESH_TASKS_NV_EXT`, `VK_INDIRECT_COMMANDS_TOKEN_TYPE_DRAW_MESH_TASKS_COUNT_NV_EXT`, `VK_INDIRECT_COMMANDS_TOKEN_TYPE_DRAW_INDEXED_COUNT_EXT`, `VK_INDIRECT_COMMANDS_TOKEN_TYPE_DRAW_INDEXED_EXT`, or `VK_INDIRECT_COMMANDS_TOKEN_TYPE_DRAW_EXT`, then `shaderStages` **must** contain graphics stages
- [](#VUID-VkIndirectCommandsLayoutCreateInfoEXT-pTokens-11105)VUID-VkIndirectCommandsLayoutCreateInfoEXT-pTokens-11105  
  For any element of `pTokens`, if `type` is `VK_INDIRECT_COMMANDS_TOKEN_TYPE_DISPATCH_EXT`, then `shaderStages` **must** be `VK_SHADER_STAGE_COMPUTE_BIT`
- [](#VUID-VkIndirectCommandsLayoutCreateInfoEXT-pTokens-11106)VUID-VkIndirectCommandsLayoutCreateInfoEXT-pTokens-11106  
  For any element of `pTokens`, if `type` is `VK_INDIRECT_COMMANDS_TOKEN_TYPE_DRAW_MESH_TASKS_EXT` or `VK_INDIRECT_COMMANDS_TOKEN_TYPE_DRAW_MESH_TASKS_COUNT_EXT`, then `shaderStages` **must** contain `VK_SHADER_STAGE_MESH_BIT_EXT`
- [](#VUID-VkIndirectCommandsLayoutCreateInfoEXT-pTokens-11107)VUID-VkIndirectCommandsLayoutCreateInfoEXT-pTokens-11107  
  For any element of `pTokens`, if `type` is `VK_INDIRECT_COMMANDS_TOKEN_TYPE_DRAW_MESH_TASKS_NV_EXT` or `VK_INDIRECT_COMMANDS_TOKEN_TYPE_DRAW_MESH_TASKS_COUNT_NV_EXT`, then the `shaderStages` **must** contain `VK_SHADER_STAGE_MESH_BIT_NV`
- [](#VUID-VkIndirectCommandsLayoutCreateInfoEXT-pTokens-11108)VUID-VkIndirectCommandsLayoutCreateInfoEXT-pTokens-11108  
  For any element of `pTokens`, if `type` is `VK_INDIRECT_COMMANDS_TOKEN_TYPE_TRACE_RAYS2_EXT`, then `shaderStages` **must** contain ray tracing stages
- [](#VUID-VkIndirectCommandsLayoutCreateInfoEXT-shaderStages-11109)VUID-VkIndirectCommandsLayoutCreateInfoEXT-shaderStages-11109  
  If `shaderStages` contains graphics stages then the state tokens in `pTokens` **must** not include `VK_INDIRECT_COMMANDS_TOKEN_TYPE_TRACE_RAYS2_EXT`, `VK_INDIRECT_COMMANDS_TOKEN_TYPE_DISPATCH_EXT`
- [](#VUID-VkIndirectCommandsLayoutCreateInfoEXT-shaderStages-11110)VUID-VkIndirectCommandsLayoutCreateInfoEXT-shaderStages-11110  
  If `shaderStages` is `VK_SHADER_STAGE_COMPUTE_BIT` then the state tokens in `pTokens` **must** only include `VK_INDIRECT_COMMANDS_TOKEN_TYPE_DISPATCH_EXT`, `VK_INDIRECT_COMMANDS_TOKEN_TYPE_EXECUTION_SET_EXT`, `VK_INDIRECT_COMMANDS_TOKEN_TYPE_PUSH_CONSTANT_EXT`, or `VK_INDIRECT_COMMANDS_TOKEN_TYPE_SEQUENCE_INDEX_EXT`
- [](#VUID-VkIndirectCommandsLayoutCreateInfoEXT-shaderStages-11111)VUID-VkIndirectCommandsLayoutCreateInfoEXT-shaderStages-11111  
  If `shaderStages` contains ray tracing stages then the state tokens in `pTokens` **must** only include `VK_INDIRECT_COMMANDS_TOKEN_TYPE_TRACE_RAYS2_EXT`, `VK_INDIRECT_COMMANDS_TOKEN_TYPE_EXECUTION_SET_EXT`, `VK_INDIRECT_COMMANDS_TOKEN_TYPE_PUSH_CONSTANT_EXT`, or `VK_INDIRECT_COMMANDS_TOKEN_TYPE_SEQUENCE_INDEX_EXT`
- [](#VUID-VkIndirectCommandsLayoutCreateInfoEXT-shaderStages-11112)VUID-VkIndirectCommandsLayoutCreateInfoEXT-shaderStages-11112  
  The `shaderStages` **must** only contain stages from one of the following:
  
  - graphics stages
  - `VK_SHADER_STAGE_COMPUTE_BIT`
  - mesh stages and `VK_SHADER_STAGE_FRAGMENT_BIT`
  - ray tracing stages
- [](#VUID-VkIndirectCommandsLayoutCreateInfoEXT-shaderStages-11113)VUID-VkIndirectCommandsLayoutCreateInfoEXT-shaderStages-11113  
  If `shaderStages` contains `VK_SHADER_STAGE_FRAGMENT_BIT`, then `shaderStages` **must** also contain `VK_SHADER_STAGE_VERTEX_BIT` or `VK_SHADER_STAGE_MESH_BIT_EXT`

Valid Usage (Implicit)

- [](#VUID-VkIndirectCommandsLayoutCreateInfoEXT-sType-sType)VUID-VkIndirectCommandsLayoutCreateInfoEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_INDIRECT_COMMANDS_LAYOUT_CREATE_INFO_EXT`
- [](#VUID-VkIndirectCommandsLayoutCreateInfoEXT-pNext-pNext)VUID-VkIndirectCommandsLayoutCreateInfoEXT-pNext-pNext  
  `pNext` **must** be `NULL` or a pointer to a valid instance of [VkPipelineLayoutCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineLayoutCreateInfo.html)
- [](#VUID-VkIndirectCommandsLayoutCreateInfoEXT-sType-unique)VUID-VkIndirectCommandsLayoutCreateInfoEXT-sType-unique  
  The `sType` value of each structure in the `pNext` chain **must** be unique
- [](#VUID-VkIndirectCommandsLayoutCreateInfoEXT-flags-parameter)VUID-VkIndirectCommandsLayoutCreateInfoEXT-flags-parameter  
  `flags` **must** be a valid combination of [VkIndirectCommandsLayoutUsageFlagBitsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndirectCommandsLayoutUsageFlagBitsEXT.html) values
- [](#VUID-VkIndirectCommandsLayoutCreateInfoEXT-shaderStages-parameter)VUID-VkIndirectCommandsLayoutCreateInfoEXT-shaderStages-parameter  
  `shaderStages` **must** be a valid combination of [VkShaderStageFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderStageFlagBits.html) values
- [](#VUID-VkIndirectCommandsLayoutCreateInfoEXT-shaderStages-requiredbitmask)VUID-VkIndirectCommandsLayoutCreateInfoEXT-shaderStages-requiredbitmask  
  `shaderStages` **must** not be `0`
- [](#VUID-VkIndirectCommandsLayoutCreateInfoEXT-pipelineLayout-parameter)VUID-VkIndirectCommandsLayoutCreateInfoEXT-pipelineLayout-parameter  
  If `pipelineLayout` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), `pipelineLayout` **must** be a valid [VkPipelineLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineLayout.html) handle
- [](#VUID-VkIndirectCommandsLayoutCreateInfoEXT-pTokens-parameter)VUID-VkIndirectCommandsLayoutCreateInfoEXT-pTokens-parameter  
  `pTokens` **must** be a valid pointer to an array of `tokenCount` valid [VkIndirectCommandsLayoutTokenEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndirectCommandsLayoutTokenEXT.html) structures
- [](#VUID-VkIndirectCommandsLayoutCreateInfoEXT-tokenCount-arraylength)VUID-VkIndirectCommandsLayoutCreateInfoEXT-tokenCount-arraylength  
  `tokenCount` **must** be greater than `0`

## [](#_see_also)See Also

[VK\_EXT\_device\_generated\_commands](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_device_generated_commands.html), [VkIndirectCommandsLayoutTokenEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndirectCommandsLayoutTokenEXT.html), [VkIndirectCommandsLayoutUsageFlagsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndirectCommandsLayoutUsageFlagsEXT.html), [VkPipelineLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineLayout.html), [VkShaderStageFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderStageFlags.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkCreateIndirectCommandsLayoutEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateIndirectCommandsLayoutEXT.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkIndirectCommandsLayoutCreateInfoEXT).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0