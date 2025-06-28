# VkIndirectCommandsLayoutCreateInfoNV(3) Manual Page

## Name

VkIndirectCommandsLayoutCreateInfoNV - Structure specifying the parameters of a newly created indirect commands layout object



## [](#_c_specification)C Specification

The `VkIndirectCommandsLayoutCreateInfoNV` structure is defined as:

```c++
// Provided by VK_NV_device_generated_commands
typedef struct VkIndirectCommandsLayoutCreateInfoNV {
    VkStructureType                           sType;
    const void*                               pNext;
    VkIndirectCommandsLayoutUsageFlagsNV      flags;
    VkPipelineBindPoint                       pipelineBindPoint;
    uint32_t                                  tokenCount;
    const VkIndirectCommandsLayoutTokenNV*    pTokens;
    uint32_t                                  streamCount;
    const uint32_t*                           pStreamStrides;
} VkIndirectCommandsLayoutCreateInfoNV;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `pipelineBindPoint` is the [VkPipelineBindPoint](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineBindPoint.html) that this layout targets.
- `flags` is a bitmask of [VkIndirectCommandsLayoutUsageFlagBitsNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndirectCommandsLayoutUsageFlagBitsNV.html) specifying usage hints of this layout.
- `tokenCount` is the length of the individual command sequence.
- `pTokens` is an array describing each command token in detail. See [VkIndirectCommandsTokenTypeNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndirectCommandsTokenTypeNV.html) and [VkIndirectCommandsLayoutTokenNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndirectCommandsLayoutTokenNV.html) below for details.
- `streamCount` is the number of streams used to provide the token inputs.
- `pStreamStrides` is an array defining the byte stride for each input stream.

## [](#_description)Description

The following code illustrates some of the flags:

```c
void cmdProcessAllSequences(cmd, pipeline, indirectCommandsLayout, pIndirectCommandsTokens, sequencesCount, indexbuffer, indexbufferOffset)
{
  for (s = 0; s < sequencesCount; s++)
  {
    sUsed = s;

    if (indirectCommandsLayout.flags & VK_INDIRECT_COMMANDS_LAYOUT_USAGE_INDEXED_SEQUENCES_BIT_NV) {
      sUsed = indexbuffer.load_uint32( sUsed * sizeof(uint32_t) + indexbufferOffset);
    }

    if (indirectCommandsLayout.flags & VK_INDIRECT_COMMANDS_LAYOUT_USAGE_UNORDERED_SEQUENCES_BIT_NV) {
      sUsed = incoherent_implementation_dependent_permutation[ sUsed ];
    }

    cmdProcessSequence( cmd, pipeline, indirectCommandsLayout, pIndirectCommandsTokens, sUsed );
  }
}
```

When tokens are consumed, an offset is computed based on token offset and stream stride. The resulting offset is required to be aligned. The alignment for a specific token is equal to the scalar alignment of the data type as defined in [Alignment Requirements](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#interfaces-alignment-requirements), or `VkPhysicalDeviceDeviceGeneratedCommandsPropertiesNV`::`minIndirectCommandsBufferOffsetAlignment`, whichever is lower.

Note

A `minIndirectCommandsBufferOffsetAlignment` of 4 allows [VkDeviceAddress](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceAddress.html) to be packed as `uvec2` with scalar layout instead of `uint64_t` with 8 byte alignment. This enables direct compatibility with D3D12 command signature layouts.

Valid Usage

- [](#VUID-VkIndirectCommandsLayoutCreateInfoNV-pipelineBindPoint-02930)VUID-VkIndirectCommandsLayoutCreateInfoNV-pipelineBindPoint-02930  
  The `pipelineBindPoint` **must** be `VK_PIPELINE_BIND_POINT_GRAPHICS` or `VK_PIPELINE_BIND_POINT_COMPUTE`
- [](#VUID-VkIndirectCommandsLayoutCreateInfoNV-tokenCount-02931)VUID-VkIndirectCommandsLayoutCreateInfoNV-tokenCount-02931  
  `tokenCount` **must** be greater than `0` and less than or equal to `VkPhysicalDeviceDeviceGeneratedCommandsPropertiesNV`::`maxIndirectCommandsTokenCount`
- [](#VUID-VkIndirectCommandsLayoutCreateInfoNV-pTokens-02932)VUID-VkIndirectCommandsLayoutCreateInfoNV-pTokens-02932  
  If `pTokens` contains an entry of `VK_INDIRECT_COMMANDS_TOKEN_TYPE_SHADER_GROUP_NV` it **must** be the first element of the array and there **must** be only a single element of such token type
- [](#VUID-VkIndirectCommandsLayoutCreateInfoNV-pTokens-09585)VUID-VkIndirectCommandsLayoutCreateInfoNV-pTokens-09585  
  If `pTokens` contains an entry of `VK_INDIRECT_COMMANDS_TOKEN_TYPE_PIPELINE_NV` it **must** be the first element of the array and there **must** be only a single element of such token type
- [](#VUID-VkIndirectCommandsLayoutCreateInfoNV-pTokens-02933)VUID-VkIndirectCommandsLayoutCreateInfoNV-pTokens-02933  
  If `pTokens` contains an entry of `VK_INDIRECT_COMMANDS_TOKEN_TYPE_STATE_FLAGS_NV` there **must** be only a single element of such token type
- [](#VUID-VkIndirectCommandsLayoutCreateInfoNV-pTokens-02934)VUID-VkIndirectCommandsLayoutCreateInfoNV-pTokens-02934  
  All state tokens in `pTokens` **must** occur before any action command tokens (`VK_INDIRECT_COMMANDS_TOKEN_TYPE_DRAW_NV`, `VK_INDIRECT_COMMANDS_TOKEN_TYPE_DRAW_INDEXED_NV`, `VK_INDIRECT_COMMANDS_TOKEN_TYPE_DRAW_TASKS_NV`, `VK_INDIRECT_COMMANDS_TOKEN_TYPE_DRAW_MESH_TASKS_NV` , `VK_INDIRECT_COMMANDS_TOKEN_TYPE_DISPATCH_NV` )
- [](#VUID-VkIndirectCommandsLayoutCreateInfoNV-pTokens-02935)VUID-VkIndirectCommandsLayoutCreateInfoNV-pTokens-02935  
  The content of `pTokens` **must** include one single action command token that is compatible with the `pipelineBindPoint`
- [](#VUID-VkIndirectCommandsLayoutCreateInfoNV-streamCount-02936)VUID-VkIndirectCommandsLayoutCreateInfoNV-streamCount-02936  
  `streamCount` **must** be greater than `0` and less or equal to `VkPhysicalDeviceDeviceGeneratedCommandsPropertiesNV`::`maxIndirectCommandsStreamCount`
- [](#VUID-VkIndirectCommandsLayoutCreateInfoNV-pStreamStrides-02937)VUID-VkIndirectCommandsLayoutCreateInfoNV-pStreamStrides-02937  
  each element of `pStreamStrides` **must** be greater than `0` and less than or equal to `VkPhysicalDeviceDeviceGeneratedCommandsPropertiesNV`::`maxIndirectCommandsStreamStride`. Furthermore the alignment of each token input **must** be ensured
- [](#VUID-VkIndirectCommandsLayoutCreateInfoNV-pipelineBindPoint-09088)VUID-VkIndirectCommandsLayoutCreateInfoNV-pipelineBindPoint-09088  
  If `pipelineBindPoint` is `VK_PIPELINE_BIND_POINT_COMPUTE` then the [`VkPhysicalDeviceDeviceGeneratedCommandsComputeFeaturesNV`::`deviceGeneratedCompute`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-deviceGeneratedCompute) feature **must** be enabled
- [](#VUID-VkIndirectCommandsLayoutCreateInfoNV-pipelineBindPoint-09089)VUID-VkIndirectCommandsLayoutCreateInfoNV-pipelineBindPoint-09089  
  If `pipelineBindPoint` is `VK_PIPELINE_BIND_POINT_COMPUTE` then the state tokens in `pTokens` **must** only include `VK_INDIRECT_COMMANDS_TOKEN_TYPE_DISPATCH_NV`, `VK_INDIRECT_COMMANDS_TOKEN_TYPE_PIPELINE_NV`, or `VK_INDIRECT_COMMANDS_TOKEN_TYPE_PUSH_CONSTANT_NV`
- [](#VUID-VkIndirectCommandsLayoutCreateInfoNV-pipelineBindPoint-09090)VUID-VkIndirectCommandsLayoutCreateInfoNV-pipelineBindPoint-09090  
  If `pipelineBindPoint` is `VK_PIPELINE_BIND_POINT_COMPUTE` and `pTokens` includes `VK_INDIRECT_COMMANDS_TOKEN_TYPE_PIPELINE_NV`, then the [`VkPhysicalDeviceDeviceGeneratedCommandsComputeFeaturesNV`::`deviceGeneratedComputePipelines`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-deviceGeneratedComputePipelines) feature **must** be enabled

Valid Usage (Implicit)

- [](#VUID-VkIndirectCommandsLayoutCreateInfoNV-sType-sType)VUID-VkIndirectCommandsLayoutCreateInfoNV-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_INDIRECT_COMMANDS_LAYOUT_CREATE_INFO_NV`
- [](#VUID-VkIndirectCommandsLayoutCreateInfoNV-pNext-pNext)VUID-VkIndirectCommandsLayoutCreateInfoNV-pNext-pNext  
  `pNext` **must** be `NULL`
- [](#VUID-VkIndirectCommandsLayoutCreateInfoNV-flags-parameter)VUID-VkIndirectCommandsLayoutCreateInfoNV-flags-parameter  
  `flags` **must** be a valid combination of [VkIndirectCommandsLayoutUsageFlagBitsNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndirectCommandsLayoutUsageFlagBitsNV.html) values
- [](#VUID-VkIndirectCommandsLayoutCreateInfoNV-pipelineBindPoint-parameter)VUID-VkIndirectCommandsLayoutCreateInfoNV-pipelineBindPoint-parameter  
  `pipelineBindPoint` **must** be a valid [VkPipelineBindPoint](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineBindPoint.html) value
- [](#VUID-VkIndirectCommandsLayoutCreateInfoNV-pTokens-parameter)VUID-VkIndirectCommandsLayoutCreateInfoNV-pTokens-parameter  
  `pTokens` **must** be a valid pointer to an array of `tokenCount` valid [VkIndirectCommandsLayoutTokenNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndirectCommandsLayoutTokenNV.html) structures
- [](#VUID-VkIndirectCommandsLayoutCreateInfoNV-pStreamStrides-parameter)VUID-VkIndirectCommandsLayoutCreateInfoNV-pStreamStrides-parameter  
  `pStreamStrides` **must** be a valid pointer to an array of `streamCount` `uint32_t` values
- [](#VUID-VkIndirectCommandsLayoutCreateInfoNV-tokenCount-arraylength)VUID-VkIndirectCommandsLayoutCreateInfoNV-tokenCount-arraylength  
  `tokenCount` **must** be greater than `0`
- [](#VUID-VkIndirectCommandsLayoutCreateInfoNV-streamCount-arraylength)VUID-VkIndirectCommandsLayoutCreateInfoNV-streamCount-arraylength  
  `streamCount` **must** be greater than `0`

## [](#_see_also)See Also

[VK\_NV\_device\_generated\_commands](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_device_generated_commands.html), [VkIndirectCommandsLayoutTokenNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndirectCommandsLayoutTokenNV.html), [VkIndirectCommandsLayoutUsageFlagsNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndirectCommandsLayoutUsageFlagsNV.html), [VkPipelineBindPoint](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineBindPoint.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkCreateIndirectCommandsLayoutNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateIndirectCommandsLayoutNV.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkIndirectCommandsLayoutCreateInfoNV)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0